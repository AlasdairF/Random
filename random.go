package randbo

import (
	"math/rand"
	"time"
	"sync"
)

const (
	bufferSize = 32
	byte2print = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-`
)

var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, bufferSize)
    },
}

var auto rand.Source = rand.NewSource(time.Now().UnixNano())

func Read(p []byte) {
	var val uint64
	var on uintptr
	div := len(p) / 8
	for i:=0; i<div; i++ {
		val = uint64(auto.Int63())
		p[on] = byte(val)
		p[on + 1] = byte(val >> 8)
		p[on + 2] = byte(val >> 16)
		p[on + 3] = byte(val >> 24)
		p[on + 4] = byte(val >> 32)
		p[on + 5] = byte(val >> 40)
		p[on + 6] = byte(val >> 48)
		p[on + 7] = byte(val >> 56)
		on += 8
	}
	switch (uintptr(len(p)) - on) {
		case 0:
			return
		case 1:
			p[on] = byte(auto.Int63())
		case 2:
			val = uint64(auto.Int63())
			p[on] = byte(val)
			p[on + 1] = byte(val >> 8)
		case 3:
			val = uint64(auto.Int63())
			p[on] = byte(val)
			p[on + 1] = byte(val >> 8)
			p[on + 2] = byte(val >> 16)
		case 4:
			val = uint64(auto.Int63())
			p[on] = byte(val)
			p[on + 1] = byte(val >> 8)
			p[on + 2] = byte(val >> 16)
			p[on + 3] = byte(val >> 24)
		case 5:
			val = uint64(auto.Int63())
			p[on] = byte(val)
			p[on + 1] = byte(val >> 8)
			p[on + 2] = byte(val >> 16)
			p[on + 3] = byte(val >> 24)
			p[on + 4] = byte(val >> 32)
		case 6:
			val = uint64(auto.Int63())
			p[on] = byte(val)
			p[on + 1] = byte(val >> 8)
			p[on + 2] = byte(val >> 16)
			p[on + 3] = byte(val >> 24)
			p[on + 4] = byte(val >> 32)
			p[on + 5] = byte(val >> 40)
		case 7:
			val = uint64(auto.Int63())
			p[on] = byte(val)
			p[on + 1] = byte(val >> 8)
			p[on + 2] = byte(val >> 16)
			p[on + 3] = byte(val >> 24)
			p[on + 4] = byte(val >> 32)
			p[on + 5] = byte(val >> 40)
			p[on + 6] = byte(val >> 48)
	}
	return
}

func Printable(p []byte) {
	var val uint64
	var on uintptr
	div := len(p) / 8
	for i:=0; i<div; i++ {
		val = uint64(auto.Int63())
		p[on] = byte2print[byte(val)]
		p[on + 1] = byte2print[byte(val >> 8)]
		p[on + 2] = byte2print[byte(val >> 16)]
		p[on + 3] = byte2print[byte(val >> 24)]
		p[on + 4] = byte2print[byte(val >> 32)]
		p[on + 5] = byte2print[byte(val >> 40)]
		p[on + 6] = byte2print[byte(val >> 48)]
		p[on + 7] = byte2print[byte(val >> 56)]
		on += 8
	}
	switch (uintptr(len(p)) - on) {
		case 0:
			return
		case 1:
			p[on] = byte2print[byte(auto.Int63())]
		case 2:
			val = uint64(auto.Int63())
			p[on] = byte2print[byte(val)]
			p[on + 1] = byte2print[byte(val >> 8)]
		case 3:
			val = uint64(auto.Int63())
			p[on] = byte2print[byte(val)]
			p[on + 1] = byte2print[byte(val >> 8)]
			p[on + 2] = byte2print[byte(val >> 16)]
		case 4:
			val = uint64(auto.Int63())
			p[on] = byte2print[byte(val)]
			p[on + 1] = byte2print[byte(val >> 8)]
			p[on + 2] = byte2print[byte(val >> 16)]
			p[on + 3] = byte2print[byte(val >> 24)]
		case 5:
			val = uint64(auto.Int63())
			p[on] = byte2print[byte(val)]
			p[on + 1] = byte2print[byte(val >> 8)]
			p[on + 2] = byte2print[byte(val >> 16)]
			p[on + 3] = byte2print[byte(val >> 24)]
			p[on + 4] = byte2print[byte(val >> 32)]
		case 6:
			val = uint64(auto.Int63())
			p[on] = byte2print[byte(val)]
			p[on + 1] = byte2print[byte(val >> 8)]
			p[on + 2] = byte2print[byte(val >> 16)]
			p[on + 3] = byte2print[byte(val >> 24)]
			p[on + 4] = byte2print[byte(val >> 32)]
			p[on + 5] = byte2print[byte(val >> 40)]
		case 7:
			val = uint64(auto.Int63())
			p[on] = byte2print[byte(val)]
			p[on + 1] = byte2print[byte(val >> 8)]
			p[on + 2] = byte2print[byte(val >> 16)]
			p[on + 3] = byte2print[byte(val >> 24)]
			p[on + 4] = byte2print[byte(val >> 32)]
			p[on + 5] = byte2print[byte(val >> 40)]
			p[on + 6] = byte2print[byte(val >> 48)]
	}
	return
}

func String(length int) string {
	if length > bufferSize {
		p := make([]byte, length)
		Printable(p)
		return string(p)
	} else {
		if length <= 0 {
			return ``
		}
		p := pool.Get().([]byte)
		defer pool.Put(p)
		Printable(p[0:length])
		return string(p[0:length])
	}
}
