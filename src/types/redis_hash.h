/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

#pragma once

#include <rocksdb/status.h>

#include <string>
#include <vector>

#include "encoding.h"
#include "storage/redis_db.h"
#include "storage/redis_metadata.h"

struct FieldValue {
  std::string field;
  std::string value;
};

enum class HashFetchType { kAll = 0, kOnlyKey = 1, kOnlyValue = 2 };

struct HashSpec {
  std::string min, max;
  bool minex, maxex; /* are min or max exclusive */
  bool max_infinite; /* are max infinite */
  int64_t offset, count;
  bool removed, reversed;
  HashSpec():minex(false),maxex(false),max_infinite(false),offset(-1),count(-1),removed(false),reversed(false){}
};

namespace Redis {
class Hash : public SubKeyScanner {
 public:
  Hash(Engine::Storage *storage, const std::string &ns) : SubKeyScanner(storage, ns) {}
  rocksdb::Status Size(const Slice &user_key, uint32_t *ret);
  rocksdb::Status Get(const Slice &user_key, const Slice &field, std::string *value);
  rocksdb::Status Set(const Slice &user_key, const Slice &field, const Slice &value, int *ret);
  rocksdb::Status SetNX(const Slice &user_key, const Slice &field, Slice value, int *ret);
  rocksdb::Status Delete(const Slice &user_key, const std::vector<Slice> &fields, int *ret);
  rocksdb::Status IncrBy(const Slice &user_key, const Slice &field, int64_t increment, int64_t *ret);
  rocksdb::Status IncrByFloat(const Slice &user_key, const Slice &field, double increment, double *ret);
  rocksdb::Status MSet(const Slice &user_key, const std::vector<FieldValue> &field_values, bool nx, int *ret);
  static Status ParseRangeLexSpec(const std::string &min, const std::string &max, HashSpec *spec);
  rocksdb::Status RangeByLex(const Slice &user_key, const Slice &start, const Slice &stop, HashSpec spec,
                        std::vector<FieldValue> *field_values);
  rocksdb::Status MGet(const Slice &user_key, const std::vector<Slice> &fields, std::vector<std::string> *values,
                       std::vector<rocksdb::Status> *statuses);
  rocksdb::Status GetAll(const Slice &user_key, std::vector<FieldValue> *field_values,
                         HashFetchType type = HashFetchType::kAll);
  rocksdb::Status Scan(const Slice &user_key, const std::string &cursor, uint64_t limit,
                       const std::string &field_prefix, std::vector<std::string> *fields,
                       std::vector<std::string> *values = nullptr);

 private:
  rocksdb::Status GetMetadata(const Slice &ns_key, HashMetadata *metadata);
};
}  // namespace Redis
