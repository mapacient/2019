package uuid

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r_s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(r_s)
}

//Uuid16 0x 16 0-F
func Uuid16() string {
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
