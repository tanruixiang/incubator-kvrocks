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
 */

package zset

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/apache/incubator-kvrocks/tests/gocase/util"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/require"
)

func createZset(rdb *redis.Client, ctx context.Context, key string, items []redis.Z) {
	rdb.Del(ctx, key)
	for _, it := range items {
		rdb.ZAdd(ctx, key, it)
	}
}

func createDefaultZset(rdb *redis.Client, ctx context.Context) {
	createZset(rdb, ctx, "zset", []redis.Z{
		{math.Inf(-1), "a"},
		{1, "b"},
		{2, "c"},
		{3, "d"},
		{4, "e"},
		{5, "f"},
		{math.Inf(1), "g"}})
}

func createDefaultLexZset(rdb *redis.Client, ctx context.Context) {
	createZset(rdb, ctx, "zset", []redis.Z{
		{0, "alpha"},
		{0, "bar"},
		{0, "cool"},
		{0, "down"},
		{0, "elephant"},
		{0, "foo"},
		{0, "great"},
		{0, "hill"},
		{0, "omega"}})
}

func basicTests(t *testing.T, rdb *redis.Client, ctx context.Context, encoding string) {
	t.Run(fmt.Sprintf("Check encoding - %s", encoding), func(t *testing.T) {
		rdb.Del(ctx, "ztmp")
		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 10, Member: "x"})
	})

	t.Run(fmt.Sprintf("ZSET basic ZADD and score update - %s", encoding), func(t *testing.T) {
		rdb.Del(ctx, "ztmp")
		// rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 10, Member: "x"})
		// rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 20, Member: "y"})
		// rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 30, Member: "z"})
		for i := 0; i < 1000; i++ {
			rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 10, Member: "x" + strconv.FormatInt(int64(i), 10)})
		}
		index := 1000
		clnts := make([]chan int, index)
		for i := 0; i < index; i++ {
			clnts[i] = make(chan int, 1)
		}
		for i := 0; i < index; i++ {
			// t.Log(i)
			go func(i int) {
				now := i
				// t.Log(now)
				if now%2 == 0 {
					require.Equal(t, []interface{}{"x999", "x998", "x997", "x996", "x995", "x994", "x993", "x992", "x991", "x990", "x99", "x989", "x988", "x987", "x986", "x985", "x984", "x983", "x982", "x981", "x980", "x98", "x979", "x978", "x977", "x976", "x975", "x974", "x973", "x972", "x971", "x970", "x97", "x969", "x968", "x967", "x966", "x965", "x964", "x963", "x962", "x961", "x960", "x96", "x959", "x958", "x957", "x956", "x955", "x954", "x953", "x952", "x951", "x950", "x95", "x949", "x948", "x947", "x946", "x945", "x944", "x943", "x942", "x941", "x940", "x94", "x939", "x938", "x937", "x936", "x935", "x934", "x933", "x932", "x931", "x930", "x93", "x929", "x928", "x927", "x926", "x925", "x924", "x923", "x922", "x921", "x920", "x92", "x919", "x918", "x917", "x916", "x915", "x914", "x913", "x912", "x911", "x910", "x91", "x909", "x908", "x907", "x906", "x905", "x904", "x903", "x902", "x901", "x900", "x90", "x9", "x899", "x898", "x897", "x896", "x895", "x894", "x893", "x892", "x891", "x890", "x89", "x889", "x888", "x887", "x886", "x885", "x884", "x883", "x882", "x881", "x880", "x88", "x879", "x878", "x877", "x876", "x875", "x874", "x873", "x872", "x871", "x870", "x87", "x869", "x868", "x867", "x866", "x865", "x864", "x863", "x862", "x861", "x860", "x86", "x859", "x858", "x857", "x856", "x855", "x854", "x853", "x852", "x851", "x850", "x85", "x849", "x848", "x847", "x846", "x845", "x844", "x843", "x842", "x841", "x840", "x84", "x839", "x838", "x837", "x836", "x835", "x834", "x833", "x832", "x831", "x830", "x83", "x829", "x828", "x827", "x826", "x825", "x824", "x823", "x822", "x821", "x820", "x82", "x819", "x818", "x817", "x816", "x815", "x814", "x813", "x812", "x811", "x810", "x81", "x809", "x808", "x807", "x806", "x805", "x804", "x803", "x802", "x801", "x800", "x80", "x8", "x799", "x798", "x797", "x796", "x795", "x794", "x793", "x792", "x791", "x790", "x79", "x789", "x788", "x787", "x786", "x785", "x784", "x783", "x782", "x781", "x780", "x78", "x779", "x778", "x777", "x776", "x775", "x774", "x773", "x772", "x771", "x770", "x77", "x769", "x768", "x767", "x766", "x765", "x764", "x763", "x762", "x761", "x760", "x76", "x759", "x758", "x757", "x756", "x755", "x754", "x753", "x752", "x751", "x750", "x75", "x749", "x748", "x747", "x746", "x745", "x744", "x743", "x742", "x741", "x740", "x74", "x739", "x738", "x737", "x736", "x735", "x734", "x733", "x732", "x731", "x730", "x73", "x729", "x728", "x727", "x726", "x725", "x724", "x723", "x722", "x721", "x720", "x72", "x719", "x718", "x717", "x716", "x715", "x714", "x713", "x712", "x711", "x710", "x71", "x709", "x708", "x707", "x706", "x705", "x704", "x703", "x702", "x701", "x700", "x70", "x7", "x699", "x698", "x697", "x696", "x695", "x694", "x693", "x692", "x691", "x690", "x69", "x689", "x688", "x687", "x686", "x685", "x684", "x683", "x682", "x681", "x680", "x68", "x679", "x678", "x677", "x676", "x675", "x674", "x673", "x672", "x671", "x670", "x67", "x669", "x668", "x667", "x666", "x665", "x664", "x663", "x662", "x661", "x660", "x66", "x659", "x658", "x657", "x656", "x655", "x654", "x653", "x652", "x651", "x650", "x65", "x649", "x648", "x647", "x646", "x645", "x644", "x643", "x642", "x641", "x640", "x64", "x639", "x638", "x637", "x636", "x635", "x634", "x633", "x632", "x631", "x630", "x63", "x629", "x628", "x627", "x626", "x625", "x624", "x623", "x622", "x621", "x620", "x62", "x619", "x618", "x617", "x616", "x615", "x614", "x613", "x612", "x611", "x610", "x61", "x609", "x608", "x607", "x606", "x605", "x604", "x603", "x602", "x601", "x600", "x60", "x6", "x599", "x598", "x597", "x596", "x595", "x594", "x593", "x592", "x591", "x590", "x59", "x589", "x588", "x587", "x586", "x585", "x584", "x583", "x582", "x581", "x580", "x58", "x579", "x578", "x577", "x576", "x575", "x574", "x573", "x572", "x571", "x570", "x57", "x569", "x568", "x567", "x566", "x565", "x564", "x563", "x562", "x561", "x560", "x56", "x559", "x558", "x557", "x556", "x555", "x554", "x553", "x552", "x551", "x550", "x55", "x549", "x548", "x547", "x546", "x545", "x544", "x543", "x542", "x541", "x540", "x54", "x539", "x538", "x537", "x536", "x535", "x534", "x533", "x532", "x531", "x530", "x53", "x529", "x528", "x527", "x526", "x525", "x524", "x523", "x522", "x521", "x520", "x52", "x519", "x518", "x517", "x516", "x515", "x514", "x513", "x512", "x511", "x510", "x51", "x509", "x508", "x507", "x506", "x505", "x504", "x503", "x502", "x501", "x500", "x50", "x5", "x499", "x498", "x497", "x496", "x495", "x494", "x493", "x492", "x491", "x490", "x49", "x489", "x488", "x487", "x486", "x485", "x484", "x483", "x482", "x481", "x480", "x48", "x479", "x478", "x477", "x476", "x475", "x474", "x473", "x472", "x471", "x470", "x47", "x469", "x468", "x467", "x466", "x465", "x464", "x463", "x462", "x461", "x460", "x46", "x459", "x458", "x457", "x456", "x455", "x454", "x453", "x452", "x451", "x450", "x45", "x449", "x448", "x447", "x446", "x445", "x444", "x443", "x442", "x441", "x440", "x44", "x439", "x438", "x437", "x436", "x435", "x434", "x433", "x432", "x431", "x430", "x43", "x429", "x428", "x427", "x426", "x425", "x424", "x423", "x422", "x421", "x420", "x42", "x419", "x418", "x417", "x416", "x415", "x414", "x413", "x412", "x411", "x410", "x41", "x409", "x408", "x407", "x406", "x405", "x404", "x403", "x402", "x401", "x400", "x40", "x4", "x399", "x398", "x397", "x396", "x395", "x394", "x393", "x392", "x391", "x390", "x39", "x389", "x388", "x387", "x386", "x385", "x384", "x383", "x382", "x381", "x380", "x38", "x379", "x378", "x377", "x376", "x375", "x374", "x373", "x372", "x371", "x370", "x37", "x369", "x368", "x367", "x366", "x365", "x364", "x363", "x362", "x361", "x360", "x36", "x359", "x358", "x357", "x356", "x355", "x354", "x353", "x352", "x351", "x350", "x35", "x349", "x348", "x347", "x346", "x345", "x344", "x343", "x342", "x341", "x340", "x34", "x339", "x338", "x337", "x336", "x335", "x334", "x333", "x332", "x331", "x330", "x33", "x329", "x328", "x327", "x326", "x325", "x324", "x323", "x322", "x321", "x320", "x32", "x319", "x318", "x317", "x316", "x315", "x314", "x313", "x312", "x311", "x310", "x31", "x309", "x308", "x307", "x306", "x305", "x304", "x303", "x302", "x301", "x300", "x30", "x3", "x299", "x298", "x297", "x296", "x295", "x294", "x293", "x292", "x291", "x290", "x29", "x289", "x288", "x287", "x286", "x285", "x284", "x283", "x282", "x281", "x280", "x28", "x279", "x278", "x277", "x276", "x275", "x274", "x273", "x272", "x271", "x270", "x27", "x269", "x268", "x267", "x266", "x265", "x264", "x263", "x262", "x261", "x260", "x26", "x259", "x258", "x257", "x256", "x255", "x254", "x253", "x252", "x251", "x250", "x25", "x249", "x248", "x247", "x246", "x245", "x244", "x243", "x242", "x241", "x240", "x24", "x239", "x238", "x237", "x236", "x235", "x234", "x233", "x232", "x231", "x230", "x23", "x229", "x228", "x227", "x226", "x225", "x224", "x223", "x222", "x221", "x220", "x22", "x219", "x218", "x217", "x216", "x215", "x214", "x213", "x212", "x211", "x210", "x21", "x209", "x208", "x207", "x206", "x205", "x204", "x203", "x202", "x201", "x200", "x20", "x2", "x199", "x198", "x197", "x196", "x195", "x194", "x193", "x192", "x191", "x190", "x19", "x189", "x188", "x187", "x186", "x185", "x184", "x183", "x182", "x181", "x180", "x18", "x179", "x178", "x177", "x176", "x175", "x174", "x173", "x172", "x171", "x170", "x17", "x169", "x168", "x167", "x166", "x165", "x164", "x163", "x162", "x161", "x160", "x16", "x159", "x158", "x157", "x156", "x155", "x154", "x153", "x152", "x151", "x150", "x15", "x149", "x148", "x147", "x146", "x145", "x144", "x143", "x142", "x141", "x140", "x14", "x139", "x138", "x137", "x136", "x135", "x134", "x133", "x132", "x131", "x130", "x13", "x129", "x128", "x127", "x126", "x125", "x124", "x123", "x122", "x121", "x120", "x12", "x119", "x118", "x117", "x116", "x115", "x114", "x113", "x112", "x111", "x110", "x11", "x109", "x108", "x107", "x106", "x105", "x104", "x103", "x102", "x101", "x100", "x10", "x1", "x0"}, rdb.Do(ctx, "zrange", "ztmp", 0, -1, "REV").Val())
				} else {
					require.Equal(t, []interface{}{"x0", "x1", "x10", "x100", "x101", "x102", "x103", "x104", "x105", "x106", "x107", "x108", "x109", "x11", "x110", "x111", "x112", "x113", "x114", "x115", "x116", "x117", "x118", "x119", "x12", "x120", "x121", "x122", "x123", "x124", "x125", "x126", "x127", "x128", "x129", "x13", "x130", "x131", "x132", "x133", "x134", "x135", "x136", "x137", "x138", "x139", "x14", "x140", "x141", "x142", "x143", "x144", "x145", "x146", "x147", "x148", "x149", "x15", "x150", "x151", "x152", "x153", "x154", "x155", "x156", "x157", "x158", "x159", "x16", "x160", "x161", "x162", "x163", "x164", "x165", "x166", "x167", "x168", "x169", "x17", "x170", "x171", "x172", "x173", "x174", "x175", "x176", "x177", "x178", "x179", "x18", "x180", "x181", "x182", "x183", "x184", "x185", "x186", "x187", "x188", "x189", "x19", "x190", "x191", "x192", "x193", "x194", "x195", "x196", "x197", "x198", "x199", "x2", "x20", "x200", "x201", "x202", "x203", "x204", "x205", "x206", "x207", "x208", "x209", "x21", "x210", "x211", "x212", "x213", "x214", "x215", "x216", "x217", "x218", "x219", "x22", "x220", "x221", "x222", "x223", "x224", "x225", "x226", "x227", "x228", "x229", "x23", "x230", "x231", "x232", "x233", "x234", "x235", "x236", "x237", "x238", "x239", "x24", "x240", "x241", "x242", "x243", "x244", "x245", "x246", "x247", "x248", "x249", "x25", "x250", "x251", "x252", "x253", "x254", "x255", "x256", "x257", "x258", "x259", "x26", "x260", "x261", "x262", "x263", "x264", "x265", "x266", "x267", "x268", "x269", "x27", "x270", "x271", "x272", "x273", "x274", "x275", "x276", "x277", "x278", "x279", "x28", "x280", "x281", "x282", "x283", "x284", "x285", "x286", "x287", "x288", "x289", "x29", "x290", "x291", "x292", "x293", "x294", "x295", "x296", "x297", "x298", "x299", "x3", "x30", "x300", "x301", "x302", "x303", "x304", "x305", "x306", "x307", "x308", "x309", "x31", "x310", "x311", "x312", "x313", "x314", "x315", "x316", "x317", "x318", "x319", "x32", "x320", "x321", "x322", "x323", "x324", "x325", "x326", "x327", "x328", "x329", "x33", "x330", "x331", "x332", "x333", "x334", "x335", "x336", "x337", "x338", "x339", "x34", "x340", "x341", "x342", "x343", "x344", "x345", "x346", "x347", "x348", "x349", "x35", "x350", "x351", "x352", "x353", "x354", "x355", "x356", "x357", "x358", "x359", "x36", "x360", "x361", "x362", "x363", "x364", "x365", "x366", "x367", "x368", "x369", "x37", "x370", "x371", "x372", "x373", "x374", "x375", "x376", "x377", "x378", "x379", "x38", "x380", "x381", "x382", "x383", "x384", "x385", "x386", "x387", "x388", "x389", "x39", "x390", "x391", "x392", "x393", "x394", "x395", "x396", "x397", "x398", "x399", "x4", "x40", "x400", "x401", "x402", "x403", "x404", "x405", "x406", "x407", "x408", "x409", "x41", "x410", "x411", "x412", "x413", "x414", "x415", "x416", "x417", "x418", "x419", "x42", "x420", "x421", "x422", "x423", "x424", "x425", "x426", "x427", "x428", "x429", "x43", "x430", "x431", "x432", "x433", "x434", "x435", "x436", "x437", "x438", "x439", "x44", "x440", "x441", "x442", "x443", "x444", "x445", "x446", "x447", "x448", "x449", "x45", "x450", "x451", "x452", "x453", "x454", "x455", "x456", "x457", "x458", "x459", "x46", "x460", "x461", "x462", "x463", "x464", "x465", "x466", "x467", "x468", "x469", "x47", "x470", "x471", "x472", "x473", "x474", "x475", "x476", "x477", "x478", "x479", "x48", "x480", "x481", "x482", "x483", "x484", "x485", "x486", "x487", "x488", "x489", "x49", "x490", "x491", "x492", "x493", "x494", "x495", "x496", "x497", "x498", "x499", "x5", "x50", "x500", "x501", "x502", "x503", "x504", "x505", "x506", "x507", "x508", "x509", "x51", "x510", "x511", "x512", "x513", "x514", "x515", "x516", "x517", "x518", "x519", "x52", "x520", "x521", "x522", "x523", "x524", "x525", "x526", "x527", "x528", "x529", "x53", "x530", "x531", "x532", "x533", "x534", "x535", "x536", "x537", "x538", "x539", "x54", "x540", "x541", "x542", "x543", "x544", "x545", "x546", "x547", "x548", "x549", "x55", "x550", "x551", "x552", "x553", "x554", "x555", "x556", "x557", "x558", "x559", "x56", "x560", "x561", "x562", "x563", "x564", "x565", "x566", "x567", "x568", "x569", "x57", "x570", "x571", "x572", "x573", "x574", "x575", "x576", "x577", "x578", "x579", "x58", "x580", "x581", "x582", "x583", "x584", "x585", "x586", "x587", "x588", "x589", "x59", "x590", "x591", "x592", "x593", "x594", "x595", "x596", "x597", "x598", "x599", "x6", "x60", "x600", "x601", "x602", "x603", "x604", "x605", "x606", "x607", "x608", "x609", "x61", "x610", "x611", "x612", "x613", "x614", "x615", "x616", "x617", "x618", "x619", "x62", "x620", "x621", "x622", "x623", "x624", "x625", "x626", "x627", "x628", "x629", "x63", "x630", "x631", "x632", "x633", "x634", "x635", "x636", "x637", "x638", "x639", "x64", "x640", "x641", "x642", "x643", "x644", "x645", "x646", "x647", "x648", "x649", "x65", "x650", "x651", "x652", "x653", "x654", "x655", "x656", "x657", "x658", "x659", "x66", "x660", "x661", "x662", "x663", "x664", "x665", "x666", "x667", "x668", "x669", "x67", "x670", "x671", "x672", "x673", "x674", "x675", "x676", "x677", "x678", "x679", "x68", "x680", "x681", "x682", "x683", "x684", "x685", "x686", "x687", "x688", "x689", "x69", "x690", "x691", "x692", "x693", "x694", "x695", "x696", "x697", "x698", "x699", "x7", "x70", "x700", "x701", "x702", "x703", "x704", "x705", "x706", "x707", "x708", "x709", "x71", "x710", "x711", "x712", "x713", "x714", "x715", "x716", "x717", "x718", "x719", "x72", "x720", "x721", "x722", "x723", "x724", "x725", "x726", "x727", "x728", "x729", "x73", "x730", "x731", "x732", "x733", "x734", "x735", "x736", "x737", "x738", "x739", "x74", "x740", "x741", "x742", "x743", "x744", "x745", "x746", "x747", "x748", "x749", "x75", "x750", "x751", "x752", "x753", "x754", "x755", "x756", "x757", "x758", "x759", "x76", "x760", "x761", "x762", "x763", "x764", "x765", "x766", "x767", "x768", "x769", "x77", "x770", "x771", "x772", "x773", "x774", "x775", "x776", "x777", "x778", "x779", "x78", "x780", "x781", "x782", "x783", "x784", "x785", "x786", "x787", "x788", "x789", "x79", "x790", "x791", "x792", "x793", "x794", "x795", "x796", "x797", "x798", "x799", "x8", "x80", "x800", "x801", "x802", "x803", "x804", "x805", "x806", "x807", "x808", "x809", "x81", "x810", "x811", "x812", "x813", "x814", "x815", "x816", "x817", "x818", "x819", "x82", "x820", "x821", "x822", "x823", "x824", "x825", "x826", "x827", "x828", "x829", "x83", "x830", "x831", "x832", "x833", "x834", "x835", "x836", "x837", "x838", "x839", "x84", "x840", "x841", "x842", "x843", "x844", "x845", "x846", "x847", "x848", "x849", "x85", "x850", "x851", "x852", "x853", "x854", "x855", "x856", "x857", "x858", "x859", "x86", "x860", "x861", "x862", "x863", "x864", "x865", "x866", "x867", "x868", "x869", "x87", "x870", "x871", "x872", "x873", "x874", "x875", "x876", "x877", "x878", "x879", "x88", "x880", "x881", "x882", "x883", "x884", "x885", "x886", "x887", "x888", "x889", "x89", "x890", "x891", "x892", "x893", "x894", "x895", "x896", "x897", "x898", "x899", "x9", "x90", "x900", "x901", "x902", "x903", "x904", "x905", "x906", "x907", "x908", "x909", "x91", "x910", "x911", "x912", "x913", "x914", "x915", "x916", "x917", "x918", "x919", "x92", "x920", "x921", "x922", "x923", "x924", "x925", "x926", "x927", "x928", "x929", "x93", "x930", "x931", "x932", "x933", "x934", "x935", "x936", "x937", "x938", "x939", "x94", "x940", "x941", "x942", "x943", "x944", "x945", "x946", "x947", "x948", "x949", "x95", "x950", "x951", "x952", "x953", "x954", "x955", "x956", "x957", "x958", "x959", "x96", "x960", "x961", "x962", "x963", "x964", "x965", "x966", "x967", "x968", "x969", "x97", "x970", "x971", "x972", "x973", "x974", "x975", "x976", "x977", "x978", "x979", "x98", "x980", "x981", "x982", "x983", "x984", "x985", "x986", "x987", "x988", "x989", "x99", "x990", "x991", "x992", "x993", "x994", "x995", "x996", "x997", "x998", "x999"}, rdb.Do(ctx, "zrange", "ztmp", 0, -1).Val())
				}
				clnts[now] <- now
			}(i)
		}

		for i := 0; i < index; i++ {
			<-clnts[i]
		}
		// rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 1, Member: "y"})
		// require.Equal(t, []string{"y", "x", "z"}, rdb.ZRange(ctx, "ztmp", 0, -1).Val())
	})

	// 	t.Run(fmt.Sprintf("ZSET basic ZADD the same member with different scores - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		require.Equal(t, int64(1), rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 10, Member: "x"}, redis.Z{Score: 20, Member: "x"}).Val())
	// 		require.Equal(t, []string{"x"}, rdb.ZRange(ctx, "ztmp", 0, -1).Val())
	// 		require.Equal(t, float64(20), rdb.ZScore(ctx, "ztmp", "x").Val())

	// 		require.Equal(t, int64(2), rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 30, Member: "x"}, redis.Z{Score: 40, Member: "y"}, redis.Z{Score: 50, Member: "z"}).Val())
	// 		require.Equal(t, []string{"x", "y", "z"}, rdb.ZRange(ctx, "ztmp", 0, -1).Val())
	// 		require.Equal(t, float64(30), rdb.ZScore(ctx, "ztmp", "x").Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZSET ZADD INCR option supports a single pair - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		require.Equal(t, 1.5, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.Contains(t, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{Members: []redis.Z{{Member: "abc", Score: 1.5}, {Member: "adc"}}}).Err(),
	// 			"INCR option supports a single increment-element pair")
	// 	})

	// 	t.Run(fmt.Sprintf("ZSET ZADD IncrMixedOtherOptions - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		require.Equal(t, "1.5", rdb.Do(ctx, "zadd", "ztmp", "nx", "nx", "nx", "nx", "incr", "1.5", "abc").Val())
	// 		require.Equal(t, redis.Nil, rdb.Do(ctx, "zadd", "ztmp", "nx", "nx", "nx", "nx", "incr", "1.5", "abc").Err())
	// 		require.Equal(t, "3", rdb.Do(ctx, "zadd", "ztmp", "xx", "xx", "xx", "xx", "incr", "1.5", "abc").Val())

	// 		rdb.Del(ctx, "ztmp")
	// 		require.Equal(t, 1.5, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.Equal(t, redis.Nil, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Err())

	// 		rdb.Del(ctx, "ztmp")
	// 		require.Equal(t, redis.Nil, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{XX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Err())
	// 		require.Equal(t, 1.5, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())

	// 		rdb.Del(ctx, "ztmp")
	// 		require.Equal(t, 1.5, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.Equal(t, 3.0, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{GT: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.Equal(t, 0.0, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{GT: true, Members: []redis.Z{{Member: "abc", Score: -1.5}}}).Val())
	// 		require.Equal(t, redis.Nil, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{GT: true, Members: []redis.Z{{Member: "abc", Score: -1.5}}}).Err())

	// 		rdb.Del(ctx, "ztmp")
	// 		require.Equal(t, 1.5, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.Equal(t, 0.0, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{LT: true, Members: []redis.Z{{Member: "abc", Score: -1.5}}}).Val())
	// 		require.Equal(t, 0.0, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{LT: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.Equal(t, redis.Nil, rdb.ZAddArgsIncr(ctx, "ztmp", redis.ZAddArgs{LT: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Err())
	// 	})

	// 	t.Run(fmt.Sprintf("ZSET ZADD LT/GT with other options - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		require.EqualValues(t, 1, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.EqualValues(t, 1, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{GT: true, Ch: true, Members: []redis.Z{{Member: "abc", Score: 2.5}}}).Val())
	// 		require.EqualValues(t, 0, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{GT: true, Ch: true, Members: []redis.Z{{Member: "abc", Score: 2.5}}}).Val())
	// 		require.EqualValues(t, 0, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{GT: true, Ch: false, Members: []redis.Z{{Member: "abc", Score: 2.5}}}).Val())
	// 		require.EqualValues(t, 0, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{GT: true, Ch: false, Members: []redis.Z{{Member: "abc", Score: 100}}}).Val())
	// 		require.Contains(t, rdb.Do(ctx, "zadd", "ztmp", "lt", "gt", "1", "m1", "2", "m2").Err(),
	// 			"GT, LT, and/or NX options at the same time are not compatible")

	// 		rdb.Del(ctx, "ztmp")
	// 		require.EqualValues(t, 1, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{LT: true, Ch: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.EqualValues(t, 1, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{LT: true, Ch: true, Members: []redis.Z{{Member: "abc", Score: 1.2}}}).Val())
	// 		require.EqualValues(t, 0, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{LT: true, Ch: false, Members: []redis.Z{{Member: "abc", Score: 0.5}}}).Val())

	// 		rdb.Del(ctx, "newAbc1", "newAbc2")
	// 		require.EqualValues(t, 2, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{Ch: true, Members: []redis.Z{{Member: "abc", Score: 0.5}, {Member: "newAbc1", Score: 10}, {Member: "newAbc2"}}}).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZSET ZADD NX/XX option supports a single pair - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		require.EqualValues(t, 2, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "a", Score: 1}, {Member: "b", Score: 2}}}).Val())
	// 		require.EqualValues(t, 1, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "c", Score: 3}}}).Val())

	// 		rdb.Del(ctx, "ztmp")
	// 		require.EqualValues(t, 1, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.EqualValues(t, 0, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.EqualValues(t, 1, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{XX: true, Ch: true, Members: []redis.Z{{Member: "abc", Score: 2.5}}}).Val())
	// 		require.EqualValues(t, 0, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{XX: true, Ch: true, Members: []redis.Z{{Member: "abc", Score: 2.5}}}).Val())
	// 		require.Contains(t, rdb.Do(ctx, "zadd", "ztmp", "nx", "xx", "1", "m1", "2", "m2").Err(),
	// 			"XX and NX options at the same time are not compatible")

	// 		require.Contains(t, rdb.Do(ctx, "zadd", "ztmp", "lt", "nx", "1", "m1", "2", "m2").Err(),
	// 			"GT, LT, and/or NX options at the same time are not compatible")
	// 		require.Contains(t, rdb.Do(ctx, "zadd", "ztmp", "gt", "nx", "1", "m1", "2", "m2").Err(),
	// 			"GT, LT, and/or NX options at the same time are not compatible")

	// 		rdb.Del(ctx, "ztmp")
	// 		require.EqualValues(t, 1, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{NX: true, Ch: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 		require.EqualValues(t, 0, rdb.ZAddArgs(ctx, "ztmp", redis.ZAddArgs{NX: true, Members: []redis.Z{{Member: "abc", Score: 1.5}}}).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZSET element can't be set to NaN with ZADD - %s", encoding), func(t *testing.T) {
	// 		require.Contains(t, rdb.ZAdd(ctx, "myzset", redis.Z{Score: math.NaN(), Member: "abc"}).Err(), "float")
	// 	})

	// 	t.Run("ZSET element can't be set to NaN with ZINCRBY", func(t *testing.T) {
	// 		require.Contains(t, rdb.ZAdd(ctx, "myzset", redis.Z{Score: math.NaN(), Member: "abc"}).Err(), "float")
	// 	})

	// 	t.Run("ZINCRBY calls leading to NaN result in error", func(t *testing.T) {
	// 		rdb.ZIncrBy(ctx, "myzset", math.Inf(1), "abc")
	// 		util.ErrorRegexp(t, rdb.ZIncrBy(ctx, "myzset", math.Inf(-1), "abc").Err(), ".*NaN.*")
	// 	})

	// 	t.Run("ZADD - Variadic version base case", func(t *testing.T) {
	// 		rdb.Del(ctx, "myzset")
	// 		require.Equal(t, int64(3), rdb.ZAdd(ctx, "myzset", redis.Z{Score: 10, Member: "a"}, redis.Z{Score: 20, Member: "b"}, redis.Z{Score: 30, Member: "c"}).Val())
	// 		require.Equal(t, []redis.Z{{10, "a"}, {20, "b"}, {30, "c"}}, rdb.ZRangeWithScores(ctx, "myzset", 0, -1).Val())
	// 	})

	// 	t.Run("ZADD - Return value is the number of actually added items", func(t *testing.T) {
	// 		require.Equal(t, int64(1), rdb.ZAdd(ctx, "myzset", redis.Z{Score: 5, Member: "x"}, redis.Z{Score: 20, Member: "b"}, redis.Z{Score: 30, Member: "c"}).Val())
	// 		require.Equal(t, []redis.Z{{5, "x"}, {10, "a"}, {20, "b"}, {30, "c"}}, rdb.ZRangeWithScores(ctx, "myzset", 0, -1).Val())
	// 	})

	// 	t.Run("ZADD - Variadic version will raise error on missing arg", func(t *testing.T) {
	// 		rdb.Del(ctx, "myzset")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zadd", "myzset", 10, "a", 20, "b", 30, "c", 40).Err(), ".*syntax.*")
	// 	})

	// 	t.Run("ZINCRBY does not work variadic even if shares ZADD implementation", func(t *testing.T) {
	// 		rdb.Del(ctx, "myzset")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zincrby", "myzset", 10, "a", 20, "b", 30, "c").Err(), ".*ERR.*wrong.*number.*arg.*")
	// 	})

	// 	t.Run(fmt.Sprintf("ZCARD basics - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 10, Member: "a"}, redis.Z{Score: 20, Member: "b"}, redis.Z{Score: 30, Member: "c"})
	// 		require.Equal(t, int64(3), rdb.ZCard(ctx, "ztmp").Val())
	// 		require.Equal(t, int64(0), rdb.ZCard(ctx, "zdoesntexist").Val())
	// 	})

	// 	t.Run("ZREM removes key after last element is removed", func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 10, Member: "x"}, redis.Z{Score: 20, Member: "y"})
	// 		require.Equal(t, int64(1), rdb.Exists(ctx, "ztmp").Val())
	// 		require.Equal(t, int64(0), rdb.ZRem(ctx, "ztmp", "z").Val())
	// 		require.Equal(t, int64(1), rdb.ZRem(ctx, "ztmp", "y").Val())
	// 		require.Equal(t, int64(1), rdb.ZRem(ctx, "ztmp", "x").Val())
	// 		require.Equal(t, int64(0), rdb.Exists(ctx, "ztmp").Val())
	// 	})

	// 	t.Run("ZREM variadic version", func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 10, Member: "a"}, redis.Z{Score: 20, Member: "b"}, redis.Z{Score: 30, Member: "c"})
	// 		require.Equal(t, int64(2), rdb.ZRem(ctx, "ztmp", []string{"x", "y", "a", "b", "k"}).Val())
	// 		require.Equal(t, int64(0), rdb.ZRem(ctx, "ztmp", []string{"foo", "bar"}).Val())
	// 		require.Equal(t, int64(1), rdb.ZRem(ctx, "ztmp", []string{"c"}).Val())
	// 		require.Equal(t, int64(0), rdb.Exists(ctx, "ztmp").Val())
	// 	})

	// 	t.Run("ZREM variadic version -- remove elements after key deletion", func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 10, Member: "a"}, redis.Z{Score: 20, Member: "b"}, redis.Z{Score: 30, Member: "c"})
	// 		require.Equal(t, int64(3), rdb.ZRem(ctx, "ztmp", []string{"a", "b", "c", "d", "e", "f", "g"}).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZRANGE basics - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 1, Member: "a"})
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 2, Member: "b"})
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 3, Member: "c"})
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 4, Member: "d"})

	// 		require.Equal(t, []string{"a", "b", "c", "d"}, rdb.ZRange(ctx, "ztmp", 0, -1).Val())
	// 		require.Equal(t, []string{"a", "b", "c"}, rdb.ZRange(ctx, "ztmp", 0, -2).Val())
	// 		require.Equal(t, []string{"b", "c", "d"}, rdb.ZRange(ctx, "ztmp", 1, -1).Val())
	// 		require.Equal(t, []string{"b", "c"}, rdb.ZRange(ctx, "ztmp", 1, -2).Val())
	// 		require.Equal(t, []string{"c", "d"}, rdb.ZRange(ctx, "ztmp", -2, -1).Val())
	// 		require.Equal(t, []string{"c"}, rdb.ZRange(ctx, "ztmp", -2, -2).Val())

	// 		// out of range start index
	// 		require.Equal(t, []string{"a", "b", "c"}, rdb.ZRange(ctx, "ztmp", -5, 2).Val())
	// 		require.Equal(t, []string{"a", "b"}, rdb.ZRange(ctx, "ztmp", -5, 1).Val())
	// 		require.Equal(t, []string{}, rdb.ZRange(ctx, "ztmp", 5, -1).Val())
	// 		require.Equal(t, []string{}, rdb.ZRange(ctx, "ztmp", 5, -2).Val())

	// 		// out of range end index
	// 		require.Equal(t, []string{"a", "b", "c", "d"}, rdb.ZRange(ctx, "ztmp", 0, 5).Val())
	// 		require.Equal(t, []string{"b", "c", "d"}, rdb.ZRange(ctx, "ztmp", 1, 5).Val())
	// 		require.Equal(t, []string{}, rdb.ZRange(ctx, "ztmp", 0, -5).Val())
	// 		require.Equal(t, []string{}, rdb.ZRange(ctx, "ztmp", 1, -5).Val())

	// 		// withscores
	// 		require.Equal(t, []redis.Z{
	// 			{1, "a"},
	// 			{2, "b"},
	// 			{3, "c"},
	// 			{4, "d"},
	// 		}, rdb.ZRangeWithScores(ctx, "ztmp", 0, -1).Val())

	// 		// use limit and offset
	// 		require.Equal(t, []interface{}([]interface{}{"a", "b", "c", "d"}), rdb.Do(ctx, "zrange", "ztmp", 0, -1, "limit", 0, -1).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "ztmp", 0, -1, "limit", 0, 0).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"a", "b"}), rdb.Do(ctx, "zrange", "ztmp", 0, -1, "limit", 0, 2).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"b", "c"}), rdb.Do(ctx, "zrange", "ztmp", 0, -1, "limit", 1, 2).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "ztmp", 0, -1, "limit", 5, 5).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"c"}), rdb.Do(ctx, "zrange", "ztmp", 1, 2, "limit", 1, 1).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "ztmp", 0, -1, "limit", 0, 0).Val())

	// 		// use withscores
	// 		require.Equal(t, []interface{}{"a", "1", "b", "2", "c", "3", "d", "4"}, rdb.Do(ctx, "zrange", "ztmp", 0, -1, "limit", 0, -1, "withscores").Val())

	// 		// use rev
	// 		require.Equal(t, []interface{}{"d", "4", "c", "3", "b", "2", "a", "1"}, rdb.Do(ctx, "zrange", "ztmp", 0, -1, "limit", 0, -1, "withscores", "rev").Val())

	// 	})

	// 	t.Run(fmt.Sprintf("ZREVRANGE basics - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 1, Member: "a"})
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 2, Member: "b"})
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 3, Member: "c"})
	// 		rdb.ZAdd(ctx, "ztmp", redis.Z{Score: 4, Member: "d"})

	// 		require.Equal(t, []string{"d", "c", "b", "a"}, rdb.ZRevRange(ctx, "ztmp", 0, -1).Val())
	// 		require.Equal(t, []string{"d", "c", "b"}, rdb.ZRevRange(ctx, "ztmp", 0, -2).Val())
	// 		require.Equal(t, []string{"c", "b", "a"}, rdb.ZRevRange(ctx, "ztmp", 1, -1).Val())
	// 		require.Equal(t, []string{"c", "b"}, rdb.ZRevRange(ctx, "ztmp", 1, -2).Val())
	// 		require.Equal(t, []string{"b", "a"}, rdb.ZRevRange(ctx, "ztmp", -2, -1).Val())
	// 		require.Equal(t, []string{"b"}, rdb.ZRevRange(ctx, "ztmp", -2, -2).Val())

	// 		// out of range start index
	// 		require.Equal(t, []string{"d", "c", "b"}, rdb.ZRevRange(ctx, "ztmp", -5, 2).Val())
	// 		require.Equal(t, []string{"d", "c"}, rdb.ZRevRange(ctx, "ztmp", -5, 1).Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRange(ctx, "ztmp", 5, -1).Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRange(ctx, "ztmp", 5, -2).Val())

	// 		// out of range end index
	// 		require.Equal(t, []string{"d", "c", "b", "a"}, rdb.ZRevRange(ctx, "ztmp", 0, 5).Val())
	// 		require.Equal(t, []string{"c", "b", "a"}, rdb.ZRevRange(ctx, "ztmp", 1, 5).Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRange(ctx, "ztmp", 0, -5).Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRange(ctx, "ztmp", 1, -5).Val())

	// 		// withscores
	// 		require.Equal(t, []redis.Z{
	// 			{4, "d"},
	// 			{3, "c"},
	// 			{2, "b"},
	// 			{1, "a"},
	// 		}, rdb.ZRevRangeWithScores(ctx, "ztmp", 0, -1).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZRANK/ZREVRANK basics - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "zranktmp")
	// 		rdb.ZAdd(ctx, "zranktmp", redis.Z{Score: 10, Member: "x"})
	// 		rdb.ZAdd(ctx, "zranktmp", redis.Z{Score: 20, Member: "y"})
	// 		rdb.ZAdd(ctx, "zranktmp", redis.Z{Score: 30, Member: "z"})
	// 		require.Equal(t, int64(0), rdb.ZRank(ctx, "zranktmp", "x").Val())
	// 		require.Equal(t, int64(1), rdb.ZRank(ctx, "zranktmp", "y").Val())
	// 		require.Equal(t, int64(2), rdb.ZRank(ctx, "zranktmp", "z").Val())
	// 		require.Equal(t, int64(0), rdb.ZRank(ctx, "zranktmp", "foo").Val())
	// 		require.Equal(t, int64(2), rdb.ZRevRank(ctx, "zranktmp", "x").Val())
	// 		require.Equal(t, int64(1), rdb.ZRevRank(ctx, "zranktmp", "y").Val())
	// 		require.Equal(t, int64(0), rdb.ZRevRank(ctx, "zranktmp", "z").Val())
	// 		require.Equal(t, int64(0), rdb.ZRevRank(ctx, "zranktmp", "foo").Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZRANK - after deletion -%s", encoding), func(t *testing.T) {
	// 		rdb.ZRem(ctx, "zranktmp", "y")
	// 		require.Equal(t, int64(0), rdb.ZRank(ctx, "zranktmp", "x").Val())
	// 		require.Equal(t, int64(1), rdb.ZRank(ctx, "zranktmp", "z").Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZINCRBY - can create a new sorted set - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "zset")
	// 		rdb.ZIncrBy(ctx, "zset", 1, "foo")
	// 		require.Equal(t, []string{"foo"}, rdb.ZRange(ctx, "zset", 0, -1).Val())
	// 		require.Equal(t, float64(1), rdb.ZScore(ctx, "zset", "foo").Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZINCRBY - increment and decrement - %s", encoding), func(t *testing.T) {
	// 		rdb.ZIncrBy(ctx, "zset", 2, "foo")
	// 		rdb.ZIncrBy(ctx, "zset", 1, "bar")
	// 		require.Equal(t, []string{"bar", "foo"}, rdb.ZRange(ctx, "zset", 0, -1).Val())
	// 		rdb.ZIncrBy(ctx, "zset", 10, "bar")
	// 		rdb.ZIncrBy(ctx, "zset", -5, "foo")
	// 		rdb.ZIncrBy(ctx, "zset", -5, "bar")
	// 		require.Equal(t, []string{"foo", "bar"}, rdb.ZRange(ctx, "zset", 0, -1).Val())
	// 		require.Equal(t, float64(-2), rdb.ZScore(ctx, "zset", "foo").Val())
	// 		require.Equal(t, float64(6), rdb.ZScore(ctx, "zset", "bar").Val())
	// 	})

	// 	t.Run("ZINCRBY return value", func(t *testing.T) {
	// 		rdb.Del(ctx, "ztmp")
	// 		require.Equal(t, float64(1), rdb.ZIncrBy(ctx, "ztmp", 1.0, "x").Val())
	// 	})

	// 	t.Run("ZRANGEBYSCORE/ZREVRANGEBYSCORE/ZCOUNT basics", func(t *testing.T) {
	// 		createDefaultZset(rdb, ctx)

	// 		// inclusive range
	// 		require.Equal(t, []string{"a", "b", "c"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "-inf", Max: "2"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"a", "b", "c"}), rdb.Do(ctx, "zrange", "zset", "-inf", "2", "BYSCORE").Val())
	// 		require.Equal(t, []string{"b", "c", "d"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "3"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"b", "c", "d"}), rdb.Do(ctx, "zrange", "zset", "0", "3", "BYSCORE").Val())
	// 		require.Equal(t, []string{"d", "e", "f"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "3", Max: "6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"d", "e", "f"}), rdb.Do(ctx, "zrange", "zset", "3", "6", "BYSCORE").Val())
	// 		require.Equal(t, []string{"e", "f", "g"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "4", Max: "+inf"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"e", "f", "g"}), rdb.Do(ctx, "zrange", "zset", "4", "+inf", "BYSCORE").Val())
	// 		require.Equal(t, []string{"c", "b", "a"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "2", Min: "-inf"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"c", "b", "a"}), rdb.Do(ctx, "zrange", "zset", "2", "-inf", "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{"d", "c", "b"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "3", Min: "0"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"d", "c", "b"}), rdb.Do(ctx, "zrange", "zset", "3", "0", "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{"f", "e", "d"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "6", Min: "3"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"f", "e", "d"}), rdb.Do(ctx, "zrange", "zset", "6", "3", "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{"g", "f", "e"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "+inf", Min: "4"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"g", "f", "e"}), rdb.Do(ctx, "zrange", "zset", "+inf", "4", "BYSCORE", "REV").Val())
	// 		require.Equal(t, int64(3), rdb.ZCount(ctx, "zset", "0", "3").Val())

	// 		// exclusive range
	// 		require.Equal(t, []string{"b"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(-inf", Max: "(2"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"b"}), rdb.Do(ctx, "zrange", "zset", "(-inf", "(2", "BYSCORE").Val())
	// 		require.Equal(t, []string{"b", "c"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(0", Max: "(3"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"b", "c"}), rdb.Do(ctx, "zrange", "zset", "(0", "(3", "BYSCORE").Val())
	// 		require.Equal(t, []string{"e", "f"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(3", Max: "(6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"e", "f"}), rdb.Do(ctx, "zrange", "zset", "(3", "(6", "BYSCORE").Val())
	// 		require.Equal(t, []string{"f"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(4", Max: "(+inf"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"f"}), rdb.Do(ctx, "zrange", "zset", "(4", "(+inf", "BYSCORE").Val())
	// 		require.Equal(t, []string{"b"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "(2", Min: "(-inf"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"b"}), rdb.Do(ctx, "zrange", "zset", "(2", "(-inf", "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{"c", "b"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "(3", Min: "(0"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"c", "b"}), rdb.Do(ctx, "zrange", "zset", "(3", "(0", "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{"f", "e"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "(6", Min: "(3"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"f", "e"}), rdb.Do(ctx, "zrange", "zset", "(6", "(3", "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{"f"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "(+inf", Min: "(4"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"f"}), rdb.Do(ctx, "zrange", "zset", "(+inf", "(4", "BYSCORE", "REV").Val())
	// 		require.Equal(t, int64(2), rdb.ZCount(ctx, "zset", "(0", "(3").Val())

	// 		// test empty ranges
	// 		rdb.ZRem(ctx, "zset", "a")
	// 		rdb.ZRem(ctx, "zset", "g")

	// 		// inclusive range
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "4", Max: "2"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "4", "2", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "6", Max: "+inf"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "6", "+inf", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "-inf", Max: "-6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "-inf", "-6", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "+inf", Min: "6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "+inf", "6", "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "-6", Min: "-inf"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "-6", "-inf", "BYSCORE", "REV").Val())

	// 		// exclusive range
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(4", Max: "(2"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(4", "(2", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "2", Max: "(2"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "2", "(2", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(2", Max: "2"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(2", "2", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(6", Max: "(+inf"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(6", "(+inf", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(-inf", Max: "(-6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(-inf", "(-6", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "(+inf", Min: "(6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(+inf", "(6", "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Max: "(-6", Min: "(-inf"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(-6", "(-inf", "BYSCORE", "REV").Val())

	// 		// empty inner range
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "2.4", Max: "2.6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "2.4", "2.6", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(2.4", Max: "2.6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(2.4", "2.6", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "2.4", Max: "(2.6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "2.4", "(2.6", "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "(2.4", Max: "(2.6"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(2.4", "(2.6", "BYSCORE").Val())
	// 	})

	// 	t.Run("ZRANGEBYSCORE with WITHSCORES", func(t *testing.T) {
	// 		createDefaultZset(rdb, ctx)
	// 		require.Equal(t, []redis.Z{{1, "b"}, {2, "c"}, {3, "d"}}, rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "3"}).Val())
	// 		require.Equal(t, []interface{}{"b", "1", "c", "2", "d", "3"}, rdb.Do(ctx, "zrange", "zset", "0", "3", "BYSCORE", "withscores").Val())
	// 		require.Equal(t, []redis.Z{{3, "d"}, {2, "c"}, {1, "b"}}, rdb.ZRevRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "3"}).Val())
	// 		require.Equal(t, []interface{}{"d", "3", "c", "2", "b", "1"}, rdb.Do(ctx, "zrange", "zset", "3", "0", "BYSCORE", "withscores", "REV").Val())
	// 	})

	// 	t.Run("ZRANGEBYSCORE with LIMIT", func(t *testing.T) {
	// 		createDefaultZset(rdb, ctx)
	// 		require.Equal(t, []string{"b", "c"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 0, Count: 2}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"b", "c"}), rdb.Do(ctx, "zrange", "zset", "0", "10", "limit", 0, 2, "BYSCORE").Val())
	// 		require.Equal(t, []string{"d", "e", "f"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 2, Count: 3}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"d", "e", "f"}), rdb.Do(ctx, "zrange", "zset", "0", "10", "limit", 2, 3, "BYSCORE").Val())
	// 		require.Equal(t, []string{"d", "e", "f"}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 2, Count: 10}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"d", "e", "f"}), rdb.Do(ctx, "zrange", "zset", "0", "10", "limit", 2, 10, "BYSCORE").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 20, Count: 10}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "0", "10", "limit", 20, 10, "BYSCORE").Val())
	// 		require.Equal(t, []string{"f", "e"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 0, Count: 2}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"f", "e"}), rdb.Do(ctx, "zrange", "zset", "10", "0", "limit", 0, 2, "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{"d", "c", "b"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 2, Count: 3}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"d", "c", "b"}), rdb.Do(ctx, "zrange", "zset", "10", "0", "limit", 2, 3, "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{"d", "c", "b"}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 2, Count: 10}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"d", "c", "b"}), rdb.Do(ctx, "zrange", "zset", "10", "0", "limit", 2, 10, "BYSCORE", "REV").Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 20, Count: 10}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "10", "0", "limit", 20, 10, "BYSCORE", "REV").Val())
	// 	})

	// 	t.Run("ZRANGEBYSCORE with LIMIT and WITHSCORES", func(t *testing.T) {
	// 		createDefaultZset(rdb, ctx)
	// 		require.Equal(t, []redis.Z{{4, "e"}, {5, "f"}}, rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{Min: "2", Max: "5", Offset: 2, Count: 3}).Val())
	// 		require.Equal(t, []interface{}{"e", "4", "f", "5"}, rdb.Do(ctx, "zrange", "zset", "2", "5", "limit", 2, 3, "BYSCORE", "withscores").Val())
	// 		require.Equal(t, []redis.Z{{3, "d"}, {2, "c"}}, rdb.ZRevRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{Min: "2", Max: "5", Offset: 2, Count: 3}).Val())
	// 		require.Equal(t, []interface{}{"d", "3", "c", "2"}, rdb.Do(ctx, "zrange", "zset", "5", "2", "limit", 2, 3, "BYSCORE", "REV", "withscores").Val())
	// 	})

	// 	t.Run("ZRANGEBYSCORE with non-value min or max", func(t *testing.T) {
	// 		util.ErrorRegexp(t, rdb.ZRangeByScore(ctx, "fooz", &redis.ZRangeBy{Min: "str", Max: "1"}).Err(), ".*double.*")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zrange", "fooz", "str", "1", "BYSCORE").Err(), ".*double.*")
	// 		util.ErrorRegexp(t, rdb.ZRangeByScore(ctx, "fooz", &redis.ZRangeBy{Min: "1", Max: "str"}).Err(), ".*double.*")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zrange", "fooz", "1", "str", "BYSCORE").Err(), ".*double.*")
	// 		util.ErrorRegexp(t, rdb.ZRangeByScore(ctx, "fooz", &redis.ZRangeBy{Min: "1", Max: "NaN"}).Err(), ".*double.*")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zrange", "fooz", "1", "NaN", "BYSCORE").Err(), ".*double.*")
	// 	})

	// 	t.Run("ZRANGEBYSCORE for min/max score with multi member", func(t *testing.T) {
	// 		zsetInt := []redis.Z{
	// 			{math.Inf(-1), "a"},
	// 			{math.Inf(-1), "b"},
	// 			{-1, "c"},
	// 			{2, "d"},
	// 			{3, "e"},
	// 			{math.Inf(1), "f"},
	// 			{math.Inf(1), "g"}}
	// 		createZset(rdb, ctx, "mzset", zsetInt)
	// 		require.Equal(t, zsetInt, rdb.ZRangeByScoreWithScores(ctx, "mzset", &redis.ZRangeBy{Min: "-inf", Max: "+inf"}).Val())
	// 		require.Equal(t, []interface{}{"a", "-inf", "b", "-inf", "c", "-1", "d", "2", "e", "3", "f", "inf", "g", "inf"}, rdb.Do(ctx, "zrange", "mzset", "-inf", "+inf", "BYSCORE", "withscores").Val())
	// 		util.ReverseSlice(zsetInt)
	// 		require.Equal(t, zsetInt, rdb.ZRevRangeByScoreWithScores(ctx, "mzset", &redis.ZRangeBy{Min: "-inf", Max: "+inf"}).Val())
	// 		require.Equal(t, []interface{}{"g", "inf", "f", "inf", "e", "3", "d", "2", "c", "-1", "b", "-inf", "a", "-inf"}, rdb.Do(ctx, "zrange", "mzset", "+inf", "-inf", "BYSCORE", "withscores", "REV").Val())
	// 		zsetDouble := []redis.Z{
	// 			{-1.004, "a"},
	// 			{-1.004, "b"},
	// 			{-1.002, "c"},
	// 			{1.002, "d"},
	// 			{1.004, "e"},
	// 			{1.004, "f"}}
	// 		createZset(rdb, ctx, "mzset", zsetDouble)
	// 		require.Equal(t, zsetDouble, rdb.ZRangeByScoreWithScores(ctx, "mzset", &redis.ZRangeBy{Min: "-inf", Max: "+inf"}).Val())
	// 		require.Equal(t, []interface{}{"a", "-1.004", "b", "-1.004", "c", "-1.002", "d", "1.002", "e", "1.004", "f", "1.004"}, rdb.Do(ctx, "zrange", "mzset", "-inf", "+inf", "BYSCORE", "withscores").Val())
	// 		util.ReverseSlice(zsetDouble)
	// 		require.Equal(t, zsetDouble, rdb.ZRevRangeByScoreWithScores(ctx, "mzset", &redis.ZRangeBy{Min: "-inf", Max: "+inf"}).Val())
	// 		require.Equal(t, []interface{}{"f", "1.004", "e", "1.004", "d", "1.002", "c", "-1.002", "b", "-1.004", "a", "-1.004"}, rdb.Do(ctx, "zrange", "mzset", "+inf", "-inf", "BYSCORE", "withscores", "REV").Val())
	// 	})

	// 	t.Run("ZRANGEBYLEX/ZREVRANGEBYLEX/ZLEXCOUNT basics", func(t *testing.T) {
	// 		createDefaultLexZset(rdb, ctx)

	// 		// inclusive range
	// 		require.Equal(t, []string{"alpha", "bar", "cool"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "-", Max: "[cool"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"alpha", "bar", "cool"}), rdb.Do(ctx, "zrange", "zset", "-", "[cool", "BYLEX").Val())
	// 		require.Equal(t, []string{"bar", "cool", "down"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[bar", Max: "[down"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"bar", "cool", "down"}), rdb.Do(ctx, "zrange", "zset", "[bar", "[down", "BYLEX").Val())
	// 		require.Equal(t, []string{"great", "hill", "omega"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[g", Max: "+"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"great", "hill", "omega"}), rdb.Do(ctx, "zrange", "zset", "[g", "+", "BYLEX").Val())
	// 		require.Equal(t, []string{"cool", "bar", "alpha"}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "-", Max: "[cool"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"cool", "bar", "alpha"}), rdb.Do(ctx, "zrange", "zset", "[cool", "-", "BYLEX", "REV").Val())
	// 		require.Equal(t, []string{"down", "cool", "bar"}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[bar", Max: "[down"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"down", "cool", "bar"}), rdb.Do(ctx, "zrange", "zset", "[down", "[bar", "BYLEX", "REV").Val())
	// 		require.Equal(t, []string{"omega", "hill", "great", "foo", "elephant", "down"}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[d", Max: "+"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"omega", "hill", "great", "foo", "elephant", "down"}), rdb.Do(ctx, "zrange", "zset", "+", "[d", "BYLEX", "rev").Val())

	// 		// exclusive range
	// 		require.Equal(t, []string{"alpha", "bar"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "-", Max: "(cool"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"alpha", "bar"}), rdb.Do(ctx, "zrange", "zset", "-", "(cool", "BYLEX").Val())
	// 		require.Equal(t, []string{"cool"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "(bar", Max: "(down"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"cool"}), rdb.Do(ctx, "zrange", "zset", "(bar", "(down", "BYLEX").Val())
	// 		require.Equal(t, []string{"hill", "omega"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "(great", Max: "+"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"hill", "omega"}), rdb.Do(ctx, "zrange", "zset", "(great", "+", "BYLEX").Val())
	// 		require.Equal(t, []string{"bar", "alpha"}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "-", Max: "(cool"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"bar", "alpha"}), rdb.Do(ctx, "zrange", "zset", "(cool", "-", "BYLEX", "REV").Val())
	// 		require.Equal(t, []string{"cool"}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "(bar", Max: "(down"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"cool"}), rdb.Do(ctx, "zrange", "zset", "(down", "(bar", "BYLEX", "REV").Val())
	// 		require.Equal(t, []string{"omega", "hill"}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "(great", Max: "+"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"omega", "hill"}), rdb.Do(ctx, "zrange", "zset", "+", "(great", "BYLEX", "REV").Val())

	// 		// inclusive and exclusive
	// 		require.Equal(t, []string{}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "(az", Max: "(b"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(az", "(b", "BYLEX").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "(z", Max: "+"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(z", "+", "BYLEX").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "-", Max: "[aaaa"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "-", "[aaaa", "BYLEX").Val())
	// 		require.Equal(t, []string{}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[elez", Max: "[elex"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "[elex", "[elez", "BYLEX", "REV").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "(hill", Max: "(omega"}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "(hill", "(omega", "BYLEX").Val())
	// 	})

	// 	t.Run("ZRANGEBYLEX with LIMIT", func(t *testing.T) {
	// 		createDefaultLexZset(rdb, ctx)
	// 		require.Equal(t, []string{"alpha", "bar"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "-", Max: "[cool", Offset: 0, Count: 2}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"alpha", "bar"}), rdb.Do(ctx, "zrange", "zset", "-", "[cool", "BYLEX", "limit", 0, 2).Val())
	// 		require.Equal(t, []string{"bar", "cool"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "-", Max: "[cool", Offset: 1, Count: 2}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"bar", "cool"}), rdb.Do(ctx, "zrange", "zset", "-", "[cool", "BYLEX", "limit", 1, 2).Val())
	// 		require.Equal(t, []interface{}{}, rdb.Do(ctx, "zrangebylex", "zset", "[bar", "[down", "limit", "0", "0").Val())
	// 		require.Equal(t, []string{}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[bar", Max: "[down", Offset: 2, Count: 0}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{}), rdb.Do(ctx, "zrange", "zset", "[bar", "[down", "BYLEX", "limit", 2, 0).Val())
	// 		require.Equal(t, []string{"bar"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[bar", Max: "[down", Offset: 0, Count: 1}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"bar"}), rdb.Do(ctx, "zrange", "zset", "[bar", "[down", "BYLEX", "limit", 0, 1).Val())
	// 		require.Equal(t, []string{"cool"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[bar", Max: "[down", Offset: 1, Count: 1}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"cool"}), rdb.Do(ctx, "zrange", "zset", "[bar", "[down", "BYLEX", "limit", 1, 1).Val())
	// 		require.Equal(t, []string{"bar", "cool", "down"}, rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[bar", Max: "[down", Offset: 0, Count: 100}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"bar", "cool", "down"}), rdb.Do(ctx, "zrange", "zset", "[bar", "[down", "BYLEX", "limit", 0, 100).Val())
	// 		require.Equal(t, []string{"omega", "hill", "great", "foo", "elephant"}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[d", Max: "+", Offset: 0, Count: 5}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"omega", "hill", "great", "foo", "elephant"}), rdb.Do(ctx, "zrange", "zset", "+", "[d", "BYLEX", "limit", 0, 5, "REV").Val())
	// 		require.Equal(t, []string{"omega", "hill", "great", "foo"}, rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: "[d", Max: "+", Offset: 0, Count: 4}).Val())
	// 		require.Equal(t, []interface{}([]interface{}{"omega", "hill", "great", "foo"}), rdb.Do(ctx, "zrange", "zset", "+", "[d", "BYLEX", "limit", 0, 4, "REV").Val())
	// 	})

	// 	t.Run("ZRANGEBYLEX withscores", func(t *testing.T) {
	// 		createDefaultLexZset(rdb, ctx)
	// 		require.Equal(t, []interface{}{"alpha", "0", "bar", "0", "cool", "0"}, rdb.Do(ctx, "zrange", "zset", "-", "[cool", "BYLEX", "withscores").Val())
	// 		require.Equal(t, []interface{}{"cool", "0", "bar", "0", "alpha", "0"}, rdb.Do(ctx, "zrange", "zset", "[cool", "-", "BYLEX", "withscores", "REV").Val())
	// 		require.Equal(t, []interface{}{}, rdb.Do(ctx, "zrange", "zset", "(a", "(a", "BYLEX", "withscores").Val())
	// 		require.Equal(t, []interface{}{}, rdb.Do(ctx, "zrange", "zset", "(a", "(a", "BYLEX", "withscores", "REV").Val())
	// 	})

	// 	t.Run("ZRANGEBYLEX with invalid lex range specifiers", func(t *testing.T) {
	// 		util.ErrorRegexp(t, rdb.ZRangeByLex(ctx, "fooz", &redis.ZRangeBy{Min: "foo", Max: "bar"}).Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zrange", "fooz", "foo", "bar", "BYLEX").Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.ZRangeByLex(ctx, "fooz", &redis.ZRangeBy{Min: "[foo", Max: "bar"}).Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zrange", "fooz", "[foo", "bar", "BYLEX").Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.ZRangeByLex(ctx, "fooz", &redis.ZRangeBy{Min: "foo", Max: "[bar"}).Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zrange", "fooz", "foo", "[bar", "BYLEX").Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.ZRangeByLex(ctx, "fooz", &redis.ZRangeBy{Min: "+x", Max: "[bar"}).Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zrange", "fooz", "+x", "[bar", "BYLEX").Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.ZRangeByLex(ctx, "fooz", &redis.ZRangeBy{Min: "-x", Max: "[bar"}).Err(), ".*illegal.*")
	// 		util.ErrorRegexp(t, rdb.Do(ctx, "zrange", "fooz", "-x", "[bar", "BYLEX").Err(), ".*illegal.*")
	// 	})

	// 	t.Run("ZREMRANGEBYSCORE basics", func(t *testing.T) {
	// 		remrangebyscore := func(min, max string) int64 {
	// 			createZset(rdb, ctx, "zset", []redis.Z{{1, "a"}, {2, "b"}, {3, "c"},
	// 				{4, "d"}, {5, "e"}})
	// 			require.Equal(t, int64(1), rdb.Exists(ctx, "zset").Val())
	// 			return rdb.ZRemRangeByScore(ctx, "zset", min, max).Val()
	// 		}

	// 		// inner range
	// 		require.Equal(t, int64(3), remrangebyscore("2", "4"))
	// 		require.Equal(t, []string{"a", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// start underflow
	// 		require.Equal(t, int64(1), remrangebyscore("-10", "1"))
	// 		require.Equal(t, []string{"b", "c", "d", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// end overflow
	// 		require.Equal(t, int64(1), remrangebyscore("5", "10"))
	// 		require.Equal(t, []string{"a", "b", "c", "d"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// switch min and max
	// 		require.Equal(t, int64(0), remrangebyscore("4", "2"))
	// 		require.Equal(t, []string{"a", "b", "c", "d", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// -inf to mid
	// 		require.Equal(t, int64(3), remrangebyscore("-inf", "3"))
	// 		require.Equal(t, []string{"d", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// mid to +inf
	// 		require.Equal(t, int64(3), remrangebyscore("3", "+inf"))
	// 		require.Equal(t, []string{"a", "b"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// -inf to +inf
	// 		require.Equal(t, int64(5), remrangebyscore("-inf", "+inf"))
	// 		require.Equal(t, []string{}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// exclusive min
	// 		require.Equal(t, int64(4), remrangebyscore("(1", "5"))
	// 		require.Equal(t, []string{"a"}, rdb.ZRange(ctx, "zset", 0, -1).Val())
	// 		require.Equal(t, int64(3), remrangebyscore("(2", "5"))
	// 		require.Equal(t, []string{"a", "b"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// exclusive max
	// 		require.Equal(t, int64(4), remrangebyscore("1", "(5"))
	// 		require.Equal(t, []string{"e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())
	// 		require.Equal(t, int64(3), remrangebyscore("1", "(4"))
	// 		require.Equal(t, []string{"d", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// exclusive min and max
	// 		require.Equal(t, int64(3), remrangebyscore("(1", "(5"))
	// 		require.Equal(t, []string{"a", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// destroy when empty
	// 		require.Equal(t, int64(5), remrangebyscore("1", "5"))
	// 		require.Equal(t, int64(0), rdb.Exists(ctx, "zset").Val())
	// 	})

	// 	t.Run("ZREMRANGEBYSCORE with non-value min or max", func(t *testing.T) {
	// 		util.ErrorRegexp(t, rdb.ZRemRangeByScore(ctx, "fooz", "str", "1").Err(), ".*double.*")
	// 		util.ErrorRegexp(t, rdb.ZRemRangeByScore(ctx, "fooz", "1", "str").Err(), ".*double.*")
	// 		util.ErrorRegexp(t, rdb.ZRemRangeByScore(ctx, "fooz", "1", "NaN").Err(), ".*double.*")
	// 	})

	// 	t.Run("ZREMRANGEBYRANK basics", func(t *testing.T) {
	// 		remrangebyrank := func(min, max int64) int64 {
	// 			createZset(rdb, ctx, "zset", []redis.Z{{1, "a"}, {2, "b"}, {3, "c"},
	// 				{4, "d"}, {5, "e"}})
	// 			require.Equal(t, int64(1), rdb.Exists(ctx, "zset").Val())
	// 			return rdb.ZRemRangeByRank(ctx, "zset", min, max).Val()
	// 		}

	// 		// inner range
	// 		require.Equal(t, int64(3), remrangebyrank(1, 3))
	// 		require.Equal(t, []string{"a", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// start underflow
	// 		require.Equal(t, int64(1), remrangebyrank(-10, 0))
	// 		require.Equal(t, []string{"b", "c", "d", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// start overflow
	// 		require.Equal(t, int64(0), remrangebyrank(10, -1))
	// 		require.Equal(t, []string{"a", "b", "c", "d", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// end underflow
	// 		require.Equal(t, int64(0), remrangebyrank(0, -10))
	// 		require.Equal(t, []string{"a", "b", "c", "d", "e"}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// end overflow
	// 		require.Equal(t, int64(5), remrangebyrank(0, 10))
	// 		require.Equal(t, []string{}, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 		// destroy when empty
	// 		require.Equal(t, int64(5), remrangebyrank(0, 4))
	// 		require.Equal(t, int64(0), rdb.Exists(ctx, "zset").Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZUNIONSTORE against non-existing key doesn't set destination - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "zseta")
	// 		require.Equal(t, int64(0), rdb.ZUnionStore(ctx, "dst_key", &redis.ZStore{Keys: []string{"zseta"}}).Val())
	// 		require.Equal(t, int64(0), rdb.Exists(ctx, "dst_key").Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZUNIONSTORE with empty set - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "zseta", "zsetb")
	// 		rdb.ZAdd(ctx, "zseta", redis.Z{Score: 1, Member: "a"})
	// 		rdb.ZAdd(ctx, "zsetb", redis.Z{Score: 2, Member: "b"})
	// 		rdb.ZUnionStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}})
	// 		require.Equal(t, []redis.Z{{1, "a"}, {2, "b"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZUNIONSTORE basics - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "zseta", "zsetb", "zsetc")
	// 		rdb.ZAdd(ctx, "zseta", redis.Z{Score: 1, Member: "a"})
	// 		rdb.ZAdd(ctx, "zseta", redis.Z{Score: 2, Member: "b"})
	// 		rdb.ZAdd(ctx, "zseta", redis.Z{Score: 3, Member: "c"})
	// 		rdb.ZAdd(ctx, "zsetb", redis.Z{Score: 1, Member: "b"})
	// 		rdb.ZAdd(ctx, "zsetb", redis.Z{Score: 2, Member: "c"})
	// 		rdb.ZAdd(ctx, "zsetb", redis.Z{Score: 3, Member: "d"})
	// 		require.Equal(t, int64(4), rdb.ZUnionStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}}).Val())
	// 		require.Equal(t, []redis.Z{{1, "a"}, {3, "b"}, {3, "d"}, {5, "c"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZUNIONSTORE with weights - %s", encoding), func(t *testing.T) {
	// 		require.Equal(t, int64(4), rdb.ZUnionStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}, Weights: []float64{2, 3}}).Val())
	// 		require.Equal(t, []redis.Z{{2, "a"}, {7, "b"}, {9, "d"}, {12, "c"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZUNIONSTORE with AGGREGATE MIN - %s", encoding), func(t *testing.T) {
	// 		require.Equal(t, int64(4), rdb.ZUnionStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}, Aggregate: "min"}).Val())
	// 		require.Equal(t, []redis.Z{{1, "a"}, {1, "b"}, {2, "c"}, {3, "d"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())

	// 	})

	// 	t.Run(fmt.Sprintf("ZUNIONSTORE with AGGREGATE MAX - %s", encoding), func(t *testing.T) {
	// 		require.Equal(t, int64(4), rdb.ZUnionStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}, Aggregate: "max"}).Val())
	// 		require.Equal(t, []redis.Z{{1, "a"}, {2, "b"}, {3, "c"}, {3, "d"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZINTERSTORE basics - %s", encoding), func(t *testing.T) {
	// 		require.Equal(t, int64(2), rdb.ZInterStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}}).Val())
	// 		require.Equal(t, []redis.Z{{3, "b"}, {5, "c"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZINTERSTORE with weights - %s", encoding), func(t *testing.T) {
	// 		require.Equal(t, int64(2), rdb.ZInterStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}, Weights: []float64{2, 3}}).Val())
	// 		require.Equal(t, []redis.Z{{7, "b"}, {12, "c"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZINTERSTORE with AGGREGATE MIN - %s", encoding), func(t *testing.T) {
	// 		require.Equal(t, int64(2), rdb.ZInterStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}, Aggregate: "min"}).Val())
	// 		require.Equal(t, []redis.Z{{1, "b"}, {2, "c"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())
	// 	})

	// 	t.Run(fmt.Sprintf("ZINTERSTORE with AGGREGATE MAX - %s", encoding), func(t *testing.T) {
	// 		require.Equal(t, int64(2), rdb.ZInterStore(ctx, "zsetc", &redis.ZStore{Keys: []string{"zseta", "zsetb"}, Aggregate: "max"}).Val())
	// 		require.Equal(t, []redis.Z{{2, "b"}, {3, "c"}}, rdb.ZRangeWithScores(ctx, "zsetc", 0, -1).Val())
	// 	})

	// 	for i, cmd := range []func(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd{rdb.ZInterStore, rdb.ZUnionStore} {
	// 		var funcName string
	// 		switch i {
	// 		case 0:
	// 			funcName = "ZINTERSTORE"
	// 		case 1:
	// 			funcName = "ZUNIONSTORE"
	// 		}

	// 		t.Run(fmt.Sprintf("%s with +inf/-inf scores - %s", funcName, encoding), func(t *testing.T) {
	// 			rdb.Del(ctx, "zsetinf1", "zsetinf2")

	// 			rdb.ZAdd(ctx, "zsetinf1", redis.Z{Score: math.Inf(1), Member: "key"})
	// 			rdb.ZAdd(ctx, "zsetinf2", redis.Z{Score: math.Inf(1), Member: "key"})
	// 			cmd(ctx, "zsetinf3", &redis.ZStore{Keys: []string{"zsetinf1", "zsetinf2"}})
	// 			require.Equal(t, math.Inf(1), rdb.ZScore(ctx, "zsetinf3", "key").Val())

	// 			rdb.ZAdd(ctx, "zsetinf1", redis.Z{Score: math.Inf(-1), Member: "key"})
	// 			rdb.ZAdd(ctx, "zsetinf2", redis.Z{Score: math.Inf(1), Member: "key"})
	// 			cmd(ctx, "zsetinf3", &redis.ZStore{Keys: []string{"zsetinf1", "zsetinf2"}})
	// 			require.Equal(t, float64(0), rdb.ZScore(ctx, "zsetinf3", "key").Val())

	// 			rdb.ZAdd(ctx, "zsetinf1", redis.Z{Score: math.Inf(1), Member: "key"})
	// 			rdb.ZAdd(ctx, "zsetinf2", redis.Z{Score: math.Inf(-1), Member: "key"})
	// 			cmd(ctx, "zsetinf3", &redis.ZStore{Keys: []string{"zsetinf1", "zsetinf2"}})
	// 			require.Equal(t, float64(0), rdb.ZScore(ctx, "zsetinf3", "key").Val())

	// 			rdb.ZAdd(ctx, "zsetinf1", redis.Z{Score: math.Inf(-1), Member: "key"})
	// 			rdb.ZAdd(ctx, "zsetinf2", redis.Z{Score: math.Inf(-1), Member: "key"})
	// 			cmd(ctx, "zsetinf3", &redis.ZStore{Keys: []string{"zsetinf1", "zsetinf2"}})
	// 			require.Equal(t, math.Inf(-1), rdb.ZScore(ctx, "zsetinf3", "key").Val())
	// 		})

	// 		t.Run(fmt.Sprintf("%s with NaN weights - %s", funcName, encoding), func(t *testing.T) {
	// 			rdb.Del(ctx, "zsetinf1", "zsetinf2")
	// 			rdb.ZAdd(ctx, "zsetinf1", redis.Z{Score: 1.0, Member: "key"})
	// 			rdb.ZAdd(ctx, "zsetinf2", redis.Z{Score: 1.0, Member: "key"})
	// 			util.ErrorRegexp(t, cmd(ctx, "zsetinf3", &redis.ZStore{
	// 				Keys:    []string{"zsetinf1", "zsetinf2"},
	// 				Weights: []float64{math.NaN(), math.NaN()}},
	// 			).Err(), ".*weight.*not.*double.*")
	// 		})
	// 	}
	// }

	// func stressTests(t *testing.T, rdb *redis.Client, ctx context.Context, encoding string) {
	// 	var elements int
	// 	if encoding == "ziplist" {
	// 		elements = 128
	// 	} else if encoding == "skiplist" {
	// 		elements = 100
	// 	} else {
	// 		fmt.Println("Unknown sorted set encoding")
	// 		return
	// 	}
	// 	t.Run(fmt.Sprintf("ZSCORE - %s", encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "zscoretest")
	// 		aux := make([]float64, 0)
	// 		for i := 0; i < elements; i++ {
	// 			score := rand.Float64()
	// 			aux = append(aux, score)
	// 			rdb.ZAdd(ctx, "zscoretest", redis.Z{Score: score, Member: strconv.Itoa(i)})
	// 		}
	// 		for i := 0; i < elements; i++ {
	// 			require.Equal(t, aux[i], rdb.ZScore(ctx, "zscoretest", strconv.Itoa(i)).Val())
	// 		}
	// 	})

	// 	t.Run(fmt.Sprintf("ZSET sorting stresser - %s", encoding), func(t *testing.T) {
	// 		delta := 0
	// 		for test := 0; test < 2; test++ {
	// 			auxArray := make(map[string]float64)
	// 			auxList := make([]redis.Z, 0)
	// 			rdb.Del(ctx, "myzset")
	// 			var score float64
	// 			for i := 0; i < elements; i++ {
	// 				if test == 0 {
	// 					score = rand.Float64()
	// 				} else {
	// 					score = float64(rand.Intn(10))
	// 				}
	// 				auxArray[strconv.Itoa(i)] = score
	// 				rdb.ZAdd(ctx, "myzset", redis.Z{Score: score, Member: strconv.Itoa(i)})
	// 				if rand.Float64() < 0.2 {
	// 					j := rand.Intn(1000)
	// 					if test == 0 {
	// 						score = rand.Float64()
	// 					} else {
	// 						score = float64(rand.Intn(10))

	// 					}
	// 					auxArray[strconv.Itoa(j)] = score
	// 					rdb.ZAdd(ctx, "myzset", redis.Z{Score: score, Member: strconv.Itoa(j)})
	// 				}
	// 			}
	// 			for i, s := range auxArray {
	// 				auxList = append(auxList, redis.Z{Score: s, Member: i})
	// 			}
	// 			sort.Slice(auxList, func(i, j int) bool {
	// 				if auxList[i].Score < auxList[j].Score {
	// 					return true
	// 				} else if auxList[i].Score > auxList[j].Score {
	// 					return false
	// 				} else {
	// 					if strings.Compare(auxList[i].Member.(string), auxList[j].Member.(string)) == 1 {
	// 						return false
	// 					} else {
	// 						return true
	// 					}
	// 				}
	// 			})
	// 			var aux []string
	// 			for _, z := range auxList {
	// 				aux = append(aux, z.Member.(string))
	// 			}
	// 			fromRedis := rdb.ZRange(ctx, "myzset", 0, -1).Val()
	// 			for i := 0; i < len(fromRedis); i++ {
	// 				if aux[i] != fromRedis[i] {
	// 					delta++
	// 				}
	// 			}
	// 			require.Equal(t, 0, delta)
	// 		}
	// 	})

	// 	t.Run(fmt.Sprintf("ZRANGEBYSCORE fuzzy test, 100 ranges in %d element sorted set - %s", elements, encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "zset")
	// 		for i := 0; i < elements; i++ {
	// 			rdb.ZAdd(ctx, "zset", redis.Z{Score: rand.Float64(), Member: strconv.Itoa(i)})
	// 		}

	// 		for i := 0; i < 100; i++ {
	// 			min, max := rand.Float64(), rand.Float64()
	// 			min, max = math.Min(min, max), math.Max(min, max)
	// 			low := rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "-inf", Max: fmt.Sprintf("%f", min)}).Val()
	// 			ok := rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: fmt.Sprintf("%f", min), Max: fmt.Sprintf("%f", max)}).Val()
	// 			high := rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: fmt.Sprintf("%f", max), Max: "+inf"}).Val()
	// 			lowEx := rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: "-inf", Max: fmt.Sprintf("(%f", min)}).Val()
	// 			okEx := rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: fmt.Sprintf("(%f", min), Max: fmt.Sprintf("(%f", max)}).Val()
	// 			highEx := rdb.ZRangeByScore(ctx, "zset", &redis.ZRangeBy{Min: fmt.Sprintf("(%f", max), Max: "+inf"}).Val()

	// 			require.Len(t, low, int(rdb.ZCount(ctx, "zset", "-inf", fmt.Sprintf("%f", min)).Val()))
	// 			require.Len(t, ok, int(rdb.ZCount(ctx, "zset", fmt.Sprintf("%f", min), fmt.Sprintf("%f", max)).Val()))
	// 			require.Len(t, high, int(rdb.ZCount(ctx, "zset", fmt.Sprintf("%f", max), "+inf").Val()))
	// 			require.Len(t, lowEx, int(rdb.ZCount(ctx, "zset", "-inf", fmt.Sprintf("(%f", min)).Val()))
	// 			require.Len(t, okEx, int(rdb.ZCount(ctx, "zset", fmt.Sprintf("(%f", min), fmt.Sprintf("(%f", max)).Val()))
	// 			require.Len(t, highEx, int(rdb.ZCount(ctx, "zset", fmt.Sprintf("(%f", max), "+inf").Val()))

	// 			for _, x := range low {
	// 				require.LessOrEqual(t, rdb.ZScore(ctx, "zset", x).Val(), min)
	// 			}
	// 			for _, x := range lowEx {
	// 				require.Less(t, rdb.ZScore(ctx, "zset", x).Val(), min)
	// 			}
	// 			for _, x := range ok {
	// 				util.BetweenValues(t, rdb.ZScore(ctx, "zset", x).Val(), min, max)
	// 			}
	// 			for _, x := range okEx {
	// 				util.BetweenValuesEx(t, rdb.ZScore(ctx, "zset", x).Val(), min, max)
	// 			}
	// 			for _, x := range high {
	// 				require.GreaterOrEqual(t, rdb.ZScore(ctx, "zset", x).Val(), min)
	// 			}
	// 			for _, x := range highEx {
	// 				require.Greater(t, rdb.ZScore(ctx, "zset", x).Val(), min)
	// 			}
	// 		}
	// 	})

	// 	t.Run(fmt.Sprintf("ZRANGEBYLEX fuzzy test, 100 ranges in %d element sorted set - %s", elements, encoding), func(t *testing.T) {
	// 		rdb.Del(ctx, "zset")

	// 		var lexSet []string
	// 		for i := 0; i < elements; i++ {
	// 			e := util.RandString(0, 30, util.Alpha)
	// 			lexSet = append(lexSet, e)
	// 			rdb.ZAdd(ctx, "zset", redis.Z{Member: e})
	// 		}
	// 		sort.Strings(lexSet)
	// 		lexSet = slices.Compact(lexSet)

	// 		for i := 0; i < 100; i++ {
	// 			min, max := util.RandString(0, 30, util.Alpha), util.RandString(0, 30, util.Alpha)
	// 			minInc, maxInc := util.RandomBool(), util.RandomBool()
	// 			cMin, cMax := "("+min, "("+max
	// 			if minInc {
	// 				cMin = "[" + min
	// 			}
	// 			if maxInc {
	// 				cMax = "[" + max
	// 			}
	// 			rev := util.RandomBool()

	// 			// make sure data is the same in both sides
	// 			require.Equal(t, lexSet, rdb.ZRange(ctx, "zset", 0, -1).Val())

	// 			var output []string
	// 			var outLen int64
	// 			if rev {
	// 				output = rdb.ZRevRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: cMax, Max: cMin}).Val()
	// 				outLen = rdb.ZLexCount(ctx, "zset", cMax, cMin).Val()
	// 			} else {
	// 				output = rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: cMin, Max: cMax}).Val()
	// 				outLen = rdb.ZLexCount(ctx, "zset", cMin, cMax).Val()
	// 			}

	// 			// compute the same output by programming
	// 			o := make([]string, 0)
	// 			c := lexSet
	// 			if (!rev && min > max) || (rev && max > min) {
	// 				// empty output when ranges are inverted
	// 			} else {
	// 				if rev {
	// 					c = rdb.ZRevRange(ctx, "zset", 0, -1).Val()
	// 					min, max, minInc, maxInc = max, min, maxInc, minInc
	// 				}

	// 				for _, e := range c {
	// 					if (minInc && e >= min || !minInc && e > min) && (maxInc && e <= max || !maxInc && e < max) {
	// 						o = append(o, e)
	// 					}
	// 				}
	// 			}
	// 			require.Equal(t, o, output)
	// 			require.Len(t, output, int(outLen))
	// 		}
	// 	})

	// 	t.Run(fmt.Sprintf("ZREMRANGEBYLEX fuzzy test, 100 ranges in %d element sorted set - %s", elements, encoding), func(t *testing.T) {
	// 		var lexSet []string
	// 		rdb.Del(ctx, "zset", "zsetcopy")
	// 		for i := 0; i < elements; i++ {
	// 			e := util.RandString(0, 30, util.Alpha)
	// 			lexSet = append(lexSet, e)
	// 			rdb.ZAdd(ctx, "zset", redis.Z{Member: e})
	// 		}
	// 		sort.Strings(lexSet)
	// 		lexSet = slices.Compact(lexSet)
	// 		for i := 0; i < 100; i++ {
	// 			rdb.ZUnionStore(ctx, "zsetcopy", &redis.ZStore{Keys: []string{"zset"}})
	// 			var lexSetCopy []string
	// 			lexSetCopy = append(lexSetCopy, lexSet...)
	// 			min, max := util.RandString(0, 30, util.Alpha), util.RandString(0, 30, util.Alpha)
	// 			minInc, maxInc := util.RandomBool(), util.RandomBool()
	// 			cMin, cMax := "("+min, "("+max
	// 			if minInc {
	// 				cMin = "[" + min
	// 			}
	// 			if maxInc {
	// 				cMax = "[" + max
	// 			}
	// 			require.Equal(t, lexSet, rdb.ZRange(ctx, "zset", 0, -1).Val())
	// 			toRem := rdb.ZRangeByLex(ctx, "zset", &redis.ZRangeBy{Min: cMin, Max: cMax}).Val()
	// 			toRemLen := rdb.ZLexCount(ctx, "zset", cMin, cMax).Val()
	// 			rdb.ZRemRangeByLex(ctx, "zsetcopy", cMin, cMax)
	// 			output := rdb.ZRange(ctx, "zsetcopy", 0, -1).Val()
	// 			if toRemLen > 0 {
	// 				var first, last int64
	// 				for idx, v := range lexSetCopy {
	// 					if v == toRem[0] {
	// 						first = int64(idx)
	// 					}
	// 				}
	// 				last = first + toRemLen - 1
	// 				lexSetCopy = append(lexSetCopy[:first], lexSetCopy[last+1:]...)
	// 			}
	// 			require.Equal(t, lexSetCopy, output)
	// 		}
	// 	})

	// 	t.Run(fmt.Sprintf("ZSETs skiplist implementation backlink consistency test - %s", encoding), func(t *testing.T) {
	// 		diff := 0
	// 		for i := 0; i < elements; i++ {
	// 			rdb.ZAdd(ctx, "zset", redis.Z{Score: rand.Float64(), Member: fmt.Sprintf("Element-%d", i)})
	// 			rdb.ZRem(ctx, "myzset", fmt.Sprintf("Element-%d", rand.Intn(elements)))
	// 		}
	// 		l1 := rdb.ZRange(ctx, "myzset", 0, -1).Val()
	// 		l2 := rdb.ZRevRange(ctx, "myzset", 0, -1).Val()
	// 		for j := 0; j < len(l1); j++ {
	// 			if l1[j] != l2[len(l1)-j-1] {
	// 				diff++
	// 			}
	// 		}
	// 		require.Equal(t, 0, diff)
	// 	})

	// 	t.Run(fmt.Sprintf("ZSETs ZRANK augmented skip list stress testing - %s", encoding), func(t *testing.T) {
	// 		var err error
	// 		rdb.Del(ctx, "myzset")
	// 		for k := 0; k < 2000; k++ {
	// 			i := k % elements
	// 			if rand.Float64() < 0.2 {
	// 				rdb.ZRem(ctx, "myzset", strconv.Itoa(i))
	// 			} else {
	// 				score := rand.Float64()
	// 				rdb.ZAdd(ctx, "myzset", redis.Z{Score: score, Member: strconv.Itoa(i)})
	// 			}
	// 			card := rdb.ZCard(ctx, "myzset").Val()
	// 			if card > 0 {
	// 				index := util.RandomInt(card)
	// 				ele := rdb.ZRange(ctx, "myzset", index, index).Val()[0]
	// 				rank := rdb.ZRank(ctx, "myzset", ele).Val()
	// 				if rank != index {
	// 					err = fmt.Errorf("%s RANK is wrong! (%d != %d)", ele, rank, index)
	// 					break
	// 				}
	// 			}
	// 		}
	// 		require.NoError(t, err)
	// 	})
	// }
}
func TestZset(t *testing.T) {
	srv := util.StartServer(t, map[string]string{})
	defer srv.Close()
	ctx := context.Background()
	rdb := srv.NewClient()
	defer func() { require.NoError(t, rdb.Close()) }()

	basicTests(t, rdb, ctx, "skiplist")

	// t.Run("ZUNIONSTORE regression, should not create NaN in scores", func(t *testing.T) {
	// 	rdb.ZAdd(ctx, "z", redis.Z{Score: math.Inf(-1), Member: "neginf"})
	// 	rdb.ZUnionStore(ctx, "out", &redis.ZStore{Keys: []string{"z"}, Weights: []float64{0}})
	// 	require.Equal(t, []redis.Z{{0, "neginf"}}, rdb.ZRangeWithScores(ctx, "out", 0, -1).Val())
	// })

	// t.Run("ZUNIONSTORE result is sorted", func(t *testing.T) {
	// 	rdb.Del(ctx, "one", "two", "dest")
	// 	var zset1 []redis.Z
	// 	var zset2 []redis.Z
	// 	for j := 0; j < 1000; j++ {
	// 		zset1 = append(zset1, redis.Z{Score: float64(util.RandomInt(1000)), Member: util.RandomValue()})
	// 		zset2 = append(zset2, redis.Z{Score: float64(util.RandomInt(1000)), Member: util.RandomValue()})
	// 	}
	// 	rdb.ZAdd(ctx, "one", zset1...)
	// 	rdb.ZAdd(ctx, "two", zset2...)
	// 	require.Greater(t, rdb.ZCard(ctx, "one").Val(), int64(100))
	// 	require.Greater(t, rdb.ZCard(ctx, "two").Val(), int64(100))
	// 	rdb.ZUnionStore(ctx, "dest", &redis.ZStore{Keys: []string{"one", "two"}})
	// 	oldScore := float64(0)
	// 	for _, z := range rdb.ZRangeWithScores(ctx, "dest", 0, -1).Val() {
	// 		require.GreaterOrEqual(t, z.Score, oldScore)
	// 		oldScore = z.Score
	// 	}
	// })

	// t.Run("ZSET commands don't accept the empty strings as valid score", func(t *testing.T) {
	// 	util.ErrorRegexp(t, rdb.Do(ctx, "zadd", "myzset", "", "abc").Err(), ".*not.*float.*")
	// })

	// stressTests(t, rdb, ctx, "skiplist")
}
