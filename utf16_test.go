package goutf16

import (
	"fmt"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type UTF16Suite struct{}

var _ = Suite(&UTF16Suite{})

var letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var letterASCII = []uint16{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122}

func (s *UTF16Suite) TestCount(c *C) {
	c.Assert(Count([]uint16{}, 112), DeepEquals, 0)
	c.Assert(Count([]uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423}, 112), DeepEquals, 0)
	c.Assert(Count([]uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423}, 8205), DeepEquals, 3)
	c.Assert(Count([]uint16{8205, 55357, 56425, 8205, 55357, 56423, 8205}, 8205), DeepEquals, 3)
	c.Assert(Count([]uint16{112}, 112), DeepEquals, 1)
}

func (s *UTF16Suite) TestIndex(c *C) {
	c.Assert(Index([]uint16{}, 112), DeepEquals, -1)
	c.Assert(Index([]uint16{110}, 112), DeepEquals, -1)
	c.Assert(Index([]uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423}, 112), DeepEquals, -1)
	c.Assert(Index([]uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423}, 8205), DeepEquals, 2)
	c.Assert(Index([]uint16{8205, 55357, 56425, 8205, 55357, 56423, 8205}, 8205), DeepEquals, 0)
	c.Assert(Index([]uint16{112}, 112), DeepEquals, 0)
}

func (s *UTF16Suite) TestJoin(c *C) {
	c.Assert(Join([][]uint16{{55357, 56424}, {55357, 56425}, {55357, 56423}, {55357, 56423}}, []uint16{8205}), DeepEquals, []uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423})
	c.Assert(Join([][]uint16{{56424}, {56425}, {56423}, {56423}}, []uint16{8205}), DeepEquals, []uint16{56424, 8205, 56425, 8205, 56423, 8205, 56423})
	c.Assert(Join([][]uint16{{}, {56424, 8205}, {56425, 8205}, {56423, 8205}, {56423}, {}}, []uint16{55357}), DeepEquals, []uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423, 55357})
	c.Assert(Join([][]uint16{{55357, 56424}, {}, {}, {55357, 56423}}, []uint16{8205}), DeepEquals, []uint16{55357, 56424, 8205, 8205, 8205, 55357, 56423})
	c.Assert(Join([][]uint16{{56424, 8205, 56425, 8205, 56423, 8205, 56423}}, []uint16{55357, 56423}), DeepEquals, []uint16{56424, 8205, 56425, 8205, 56423, 8205, 56423})
	c.Assert(Join([][]uint16{}, []uint16{55357, 56423}), DeepEquals, []uint16{})
	c.Assert(Join([][]uint16{{}, {56424, 8205}, {56425, 8205}, {56423, 8205}, {56423}, {}}, []uint16{55357, 56423, 55357}), DeepEquals, []uint16{55357, 56423, 55357, 56424, 8205, 55357, 56423, 55357, 56425, 8205, 55357, 56423, 55357, 56423, 8205, 55357, 56423, 55357, 56423, 55357, 56423, 55357})
}

func (s *UTF16Suite) TestEncodeStringToUTF16(c *C) {
	content := EncodeStringToUTF16(letter)
	c.Assert(content, DeepEquals, letterASCII)
	c.Assert(len(content), Equals, 52)
	content = EncodeStringToUTF16("你好")
	c.Assert(content, DeepEquals, []uint16{20320, 22909})
	c.Assert(len(content), Equals, 2)
	content = EncodeStringToUTF16("👨‍👩‍👧‍👧")
	c.Assert(content, DeepEquals, []uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423})
	c.Assert(len(content), Equals, 11)
	content = EncodeStringToUTF16("👨‍👩‍👧‍👧哈哈👋开心")
	c.Assert(content, DeepEquals, []uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423, 21704, 21704, 55357, 56395, 24320, 24515})
	c.Assert(len(content), Equals, 17)

	content = EncodeStringToUTF16(string([]rune{1, 2, 3, 4}))
	c.Assert(content, DeepEquals, []uint16{1, 2, 3, 4})

	content = EncodeStringToUTF16(string([]rune{0xffff, 0x10000, 0x10001, 0x12345, 0x10ffff}))
	c.Assert(content, DeepEquals, []uint16{0xffff, 0xd800, 0xdc00, 0xd800, 0xdc01, 0xd808, 0xdf45, 0xdbff, 0xdfff})

	content = EncodeStringToUTF16(string([]rune{'a', 'b', 0xd7ff, 0xd800, 0xdfff, 0xe000, 0x110000, -1}))
	c.Assert(content, DeepEquals, []uint16{'a', 'b', 0xd7ff, 0xfffd, 0xfffd, 0xe000, 0xfffd, 0xfffd})

	content = EncodeStringToUTF16(string([]rune{0xd800, 'a'}))
	c.Assert(content, DeepEquals, []uint16{0xfffd, 'a'})
}

func (s *UTF16Suite) TestDecodeUTF16ToString(c *C) {
	content := DecodeUTF16ToString(letterASCII)
	c.Assert(content, DeepEquals, letter)
	c.Assert(len(content), DeepEquals, 52)
	content = DecodeUTF16ToString([]uint16{20320, 22909})
	c.Assert(content, DeepEquals, "你好")
	content = DecodeUTF16ToString([]uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423})
	c.Assert(content, DeepEquals, "👨‍👩‍👧‍👧")
	c.Assert(len(content), Equals, 25)
	content = DecodeUTF16ToString([]uint16{55357, 56424, 8205, 55357, 56425, 8205, 55357, 56423, 8205, 55357, 56423, 21704, 21704, 55357, 56395, 24320, 24515})
	c.Assert(content, DeepEquals, "👨‍👩‍👧‍👧哈哈👋开心")
	c.Assert(len(content), Equals, 41)

	content = DecodeUTF16ToString([]uint16{1, 2, 3, 4})
	c.Assert(content, DeepEquals, string([]rune{1, 2, 3, 4}))

	content = DecodeUTF16ToString([]uint16{0xffff, 0xd800, 0xdc00, 0xd800, 0xdc01, 0xd808, 0xdf45, 0xdbff, 0xdfff})
	c.Assert(content, DeepEquals, string([]rune{0xffff, 0x10000, 0x10001, 0x12345, 0x10ffff}))

	content = DecodeUTF16ToString([]uint16{0xd800, 'a'})
	c.Assert(content, DeepEquals, string([]rune{0xfffd, 'a'}))
}

func BenchmarkEncodeStringToUTF16(b *testing.B) {
	opFile, err := os.Open("./benchmark/text")
	if err != nil {
		fmt.Println(err)
		return
	}
	byteValue, _ := ioutil.ReadAll(opFile)
	defer func() {
		_ = opFile.Close()
	}()

	content := string(byteValue)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		EncodeStringToUTF16(content)
	}
}

func BenchmarkDecodeUTF16ToString(b *testing.B) {
	opFile, err := os.Open("./benchmark/text")
	if err != nil {
		fmt.Println(err)
		return
	}
	byteValue, _ := ioutil.ReadAll(opFile)
	defer func() {
		_ = opFile.Close()
	}()

	content := EncodeStringToUTF16(string(byteValue))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		DecodeUTF16ToString(content)
	}
}
