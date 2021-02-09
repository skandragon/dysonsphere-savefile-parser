package main

import (
	"fmt"
	"math"
)

// Buffer holds the data and current read position for parsing a
// binary blob of data.  It's not an efficient system, but it
// is direct.
//
// Errors are not returned, to simplify code...
type Buffer struct {
	data []byte
	pos  int
}

// MakeBuffer returns a new Buffer, position 0, with the provided content.
func MakeBuffer(content []byte) *Buffer {
	return &Buffer{
		data: content,
		pos:  0,
	}
}

func (b *Buffer) checklen(c int) {
	if b.pos+c > len(b.data) {
		panic(fmt.Sprintf("Attempt to read past end of buffer:  position %d, len %d", b.pos, c))
	}
}

// Length returns the total length of the buffer's content, read
// and unread.
func (b *Buffer) Length() int {
	return len(b.data)
}

// GetBytes returns the requested number of bytes.
// Will panic() if insufficient data remains.
func (b *Buffer) GetBytes(c int) []byte {
	b.checklen(c)
	ret := b.data[b.pos : b.pos+c]
	b.pos += c
	return ret
}

// GetInt64le returns an int64, encoded as little-endian.
// Will panic() if insufficient data remains.
func (b *Buffer) GetInt64le() int64 {
	b.checklen(8)
	var ret int64 = 0
	ret |= int64(b.data[b.pos])
	b.pos++
	ret |= int64(b.data[b.pos]) << 8
	b.pos++
	ret |= int64(b.data[b.pos]) << 16
	b.pos++
	ret |= int64(b.data[b.pos]) << 24
	b.pos++
	ret |= int64(b.data[b.pos]) << 32
	b.pos++
	ret |= int64(b.data[b.pos]) << 40
	b.pos++
	ret |= int64(b.data[b.pos]) << 48
	b.pos++
	ret |= int64(b.data[b.pos]) << 56
	b.pos++
	return ret
}

// GetUInt64le returns an uint64, encoded as little-endian.
// Will panic() if insufficient data remains.
func (b *Buffer) GetUInt64le() uint64 {
	b.checklen(8)
	var ret uint64 = 0
	ret |= uint64(b.data[b.pos])
	b.pos++
	ret |= uint64(b.data[b.pos]) << 8
	b.pos++
	ret |= uint64(b.data[b.pos]) << 16
	b.pos++
	ret |= uint64(b.data[b.pos]) << 24
	b.pos++
	ret |= uint64(b.data[b.pos]) << 32
	b.pos++
	ret |= uint64(b.data[b.pos]) << 40
	b.pos++
	ret |= uint64(b.data[b.pos]) << 48
	b.pos++
	ret |= uint64(b.data[b.pos]) << 56
	b.pos++
	return ret
}

// GetInt32le returns an int32, encoded as little-endian.
// Will panic() if insufficient data remains.
func (b *Buffer) GetInt32le() int32 {
	b.checklen(8)
	var ret int32 = 0
	ret |= int32(b.data[b.pos])
	b.pos++
	ret |= int32(b.data[b.pos]) << 8
	b.pos++
	ret |= int32(b.data[b.pos]) << 16
	b.pos++
	ret |= int32(b.data[b.pos]) << 24
	b.pos++
	return ret
}

// GetUInt32le returns an uint32, encoded as little-endian.
// Will panic() if insufficient data remains.
func (b *Buffer) GetUInt32le() uint32 {
	b.checklen(8)
	var ret uint32 = 0
	ret |= uint32(b.data[b.pos])
	b.pos++
	ret |= uint32(b.data[b.pos]) << 8
	b.pos++
	ret |= uint32(b.data[b.pos]) << 16
	b.pos++
	ret |= uint32(b.data[b.pos]) << 24
	b.pos++
	return ret
}

// GetByte returns a byte.
// Will panic() if insufficient data remains.
func (b *Buffer) GetByte() byte {
	b.checklen(1)
	ret := b.data[b.pos]
	b.pos++
	return ret
}

// GetBoolean reads one byte from the buffer, and if it is
// 0 will return false, and true for all other values.
// Will panic() if insufficient data remains.
func (b *Buffer) GetBoolean() bool {
	s := b.GetByte()
	return s != 0
}

// GetVarint will return an integer encoded as a variable number
// of bytes.  A maximum of 5 bytes will be consumed, as the
// final return value is an int32.
// Will panic() if insufficient data remains.
func (b *Buffer) GetVarint() int32 {
	var count int32 = 0
	var shift = 0
	var by byte = 0
	for ok := true; ok; ok = by&0x80 == 0x80 {
		if shift == 5*7 {
			panic(fmt.Errorf("Got more than 5 bytes for a varint"))
		}
		by = b.GetByte()
		count |= int32(by&0x7f) << shift
		shift += 7
	}
	return count
}

// GetString returns a varint-prefixed string.
// Will panic() if insufficient data remains.
func (b *Buffer) GetString() string {
	strlen := b.GetVarint()
	s := b.GetBytes(int(strlen))
	return string(s)
}

// GetFloat32 returns a float32, encoded as little endian format.
// Will panic() if insufficient data remains.
func (b *Buffer) GetFloat32() float32 {
	s := b.GetUInt32le()
	return math.Float32frombits(s)
}

// GetFloat64 returns a float64, encoded as little endian format.
// Will panic() if insufficient data remains.
func (b *Buffer) GetFloat64() float64 {
	s := b.GetUInt64le()
	return math.Float64frombits(s)
}
