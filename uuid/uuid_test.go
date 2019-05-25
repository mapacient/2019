package uuid

import (
	"testing"
)

var test_count = 7

func Test_uuid16(t *testing.T) {
	t.Log("test")
	for i := 0; i < test_count; i++ {
		s := Uuid16()
		t.Log(s)
	}
}

func Test_uuid(t *testing.T) {
	t.Log("test")
	for i := 0; i < test_count; i++ {
		s := Uuid(3)
		t.Log(s)
	}

}
