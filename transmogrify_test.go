package transmogrify

import (
	"bytes"
	"reflect"
	"testing"
)

var encodeTests = []struct {
	cmds string
	in   []byte
	out  []byte
}{
	{"H", []byte("1234567890"), []byte("0987654321")},
	{"H", []byte("qwertyuiop"), []byte("poiuytrewq")},
	{"H", []byte("asdfghjkl;"), []byte(";lkjhgfdsa")},
	{"H", []byte("zxcvbnm,./"), []byte("/.,mnbvcxz")},
	{"H,H", []byte("1234567890"), []byte("1234567890")},
	{"H,H", []byte("qwertyuiop"), []byte("qwertyuiop")},
	{"H,H", []byte("asdfghjkl;"), []byte("asdfghjkl;")},
	{"H,H", []byte("zxcvbnm,./"), []byte("zxcvbnm,./")},

	{"V", []byte("1234567890"), []byte("zxcvbnm,./")},
	{"V", []byte("qwertyuiop"), []byte("asdfghjkl;")},
	{"V", []byte("asdfghjkl;"), []byte("qwertyuiop")},
	{"V", []byte("zxcvbnm,./"), []byte("1234567890")},
	{"V,V", []byte("1234567890"), []byte("1234567890")},
	{"V,V", []byte("qwertyuiop"), []byte("qwertyuiop")},
	{"V,V", []byte("asdfghjkl;"), []byte("asdfghjkl;")},
	{"V,V", []byte("zxcvbnm,./"), []byte("zxcvbnm,./")},

	{"1", []byte("1234567890"), []byte("/123456789")},
	{"2", []byte("1234567890"), []byte("./12345678")},
	{"41", []byte("1234567890"), []byte("/123456789")},
	{"42", []byte("1234567890"), []byte("./12345678")},
	{"-1", []byte("1234567890"), []byte("234567890q")},
	{"-2", []byte("1234567890"), []byte("34567890qw")},
	{"1,-1", []byte("1234567890"), []byte("1234567890")},
}

func TestEncode(t *testing.T) {
	for _, tt := range encodeTests {
		r := bytes.NewReader(tt.in)
		trans := New(r, tt.cmds)
		out := trans.encode(tt.in)
		if !reflect.DeepEqual(tt.out, out) {
			t.Errorf("encode(%s) => %s, want %s", tt.cmds, out, tt.out)
		}
	}
}

func TestShift(t *testing.T) {
}

func TestReverse(t *testing.T) {
}

func TestHorizontal(t *testing.T) {
}

func TestVertical(t *testing.T) {
}
