package transmogrify

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

var (
	// originalKeyboard is a 4x10 layout represented as a string
	// "1234567890",
	// "qwertyuiop",
	// "asdfghjkl;",
	// "zxcvbnm,./",
	originalKeyboard = []byte("1234567890qwertyuiopasdfghjkl;zxcvbnm,./")
)

// tmog performs transformations on a 4x10 qwerty keyboard
// and outputs an encoded stream of text.
//
// The 4 rows start with 1, q, a, z and extend to 0, p, ;, /, respectively.
//
// The transforms can be chained and the output is an encoded
// stream of text.
//
// Example transform commands are(comma seperated input):
//
// horizontal flip (H)
//	This transformation will flip all rows of the keyboard horizontally
//	(e.g., the 1 will swap with the 0, the 2 with the 9, etc.)
//
// vertical flip (V)
//	This transformation will flip all columns of the keyboard vertically
//	(e.g., the 1 will swap with the z, the q with the a, the 2 with the x, etc.)
//
// shift by n (+-int)
//	This transformation takes an integer N and performs a
//	linear shift of the keyboard.  Each key shifts N places to its right if N > 0
//	(and likewise to the left if N < 0).  If a key would move past its current row
//	then it will shift into the row below, and so on.  For example, for N = 5,
//	the last five keys (nm,./ would move into the first 5 places of the top row,
//	the 12345 would move 5 places to the right, 67890 would move to the start
//	of the 2nd row, and so on). Likewise, left-shifting keys past their current
//	rows would shift them back into the row above. Therefore, a single right and
//	a left shift would produce the same keyboard.
type tmog struct {
	Stream   io.Reader     // the text input to stream and encode.
	Commands []string      // an in order list of operations to perform. See spec.
	Mappings map[byte]byte // original key to new key mapping.
	current  []byte        // current keyboard layout
}

// New takes an io stream of text and a set of commands as a comma
// seperated string. If cmds is empty no initial transformation
// is made.
func New(in io.Reader, cmds string) *tmog {
	t := &tmog{
		Stream:   in,
		Commands: strings.Split(strings.ToUpper(cmds), ","),
		Mappings: map[byte]byte{},
		current:  []byte("1234567890qwertyuiopasdfghjkl;zxcvbnm,./"),
	}
	t.setMap()
	for _, cmd := range t.Commands {
		err := t.Transform(cmd)
		if err != nil {
			log.Printf("err: %s", err)
		}
	}
	return t
}

// Transform applies a command and modifies the internal
// keyboard layout.
func (t *tmog) Transform(cmd string) error {
	var err error
	switch cmd {
	case "H":
		err = t.horizontal()
	case "V":
		err = t.vertical()
	default:
		// shift
		n, err := strconv.Atoi(cmd)
		if err == nil {
			t.shift(n)
		}
	}
	t.setMap()
	return err
}

// setMap takes the current keyboard layout and produces
// mappings for easier lookup when encoding.
func (t *tmog) setMap() {
	for i, c := range originalKeyboard {
		t.Mappings[c] = t.current[i]
	}
}

// horizontal will flip all rows of the keyboard horizontally
// (e.g., the 1 will swap with the 0, the 2 with the 9, etc.)
func (t *tmog) horizontal() error {
	i := 0
	j := 9
	for i < j {
		t.current[i], t.current[j] = t.current[j], t.current[i]
		t.current[i+10], t.current[j+10] = t.current[j+10], t.current[i+10]
		t.current[i+20], t.current[j+20] = t.current[j+20], t.current[i+20]
		t.current[i+30], t.current[j+30] = t.current[j+30], t.current[i+30]
		i++
		j--
	}
	return nil
}

// vertical will flip all rows of the keyboard horizontally
// (e.g., the 1 will swap with the 0, the 2 with the 9, etc.)
func (t *tmog) vertical() error {
	for i := 0; i < 10; i++ {
		t.current[i], t.current[i+30] = t.current[i+30], t.current[i]
		t.current[i+10], t.current[i+20] = t.current[i+20], t.current[i+10]
	}
	return nil
}

// reverse reverses characters of input from start to end indexes.
func (t *tmog) reverse(start, end int) {
	for start < end {
		t.current[start], t.current[end] = t.current[end], t.current[start]
		start++
		end--
	}
}

// shift applies a shift to the keyboard of k characters.
// found here:
// http://stackoverflow.com/questions/22078728/rotate-array-left-or-right-by-a-set-number-of-positions-in-on-complexity
// but originaly from Programming Pearls.
func (t *tmog) shift(k int) error {
	l := len(t.current)
	switch {
	case k == 0:
		return nil
	case k < 0:
		k = k + l
	}
	if k > l {
		k = k % l
	}

	t.reverse(0, l-1)
	t.reverse(0, k-1)
	t.reverse(k, l-1)

	return nil
}

// encode applies the current keyboard layout to the given input
// bytes. TODO modify original or return copy?
func (t *tmog) encode(b []byte) []byte {
	out := bytes.ToLower(b)
	for i, c := range out {
		if m, ok := t.Mappings[c]; ok {
			b[i] = m
		}
	}
	return b
}

// Print streams out from the internal reader encoded text based on
// the previously initialized transforms.
func (t *tmog) Print() error {
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := t.Stream.Read(buf)
		if err != nil {
			if err == io.EOF {
				if n > 0 {
					fmt.Printf("%s", t.encode(buf))
				}
				return nil
			}
			return err
		}
		if n == 0 {
			continue
		}
		fmt.Printf("%s", t.encode(buf))
	}
	return nil
}
