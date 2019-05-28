package uuid

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var r *rand.Rand

func get_rand() {
	var v atomic.Value = atomic.Value{}
	if v.Load() == nil {
		//		println("v load nil")
		r_s := rand.NewSource(time.Now().UnixNano())
		r = rand.New(r_s)

		v.Store(r)
	}
}

func Uuid16() string {
	if r == nil {
		get_rand()
	}
	sl := make([]byte, 8)
	_, err := r.Read(sl)
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", sl)
	return s
}

//Uuid 0x l 0-F
func Uuid(l int) string {
	if r == nil {
		get_rand()
	}
	if l < 1 {
		panic(errors.New("l should > 0"))
	}
	count_byte := (l + 1) >> 1
	sl := make([]byte, count_byte)
	_, err := r.Read(sl)
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", sl)
	fmt.Println("len is ", l, "truth is ", len(s), "#######", s)
	return s[:l]
}
