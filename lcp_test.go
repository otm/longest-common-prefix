package lcp

import (
	"fmt"
	"testing"
)

var findtests = []struct {
	needle   string
	haystack []string
	out      []string
}{
	{"t", []string{"tes", "test", "testing", "footest", "testfoo"}, []string{"tes", "test", "testing", "testfoo"}},
	{"te", []string{"tes", "test", "testing", "footest", "testfoo"}, []string{"tes", "test", "testing", "testfoo"}},
	{"tes", []string{"tes", "test", "testing", "footest", "testfoo"}, []string{"tes", "test", "testing", "testfoo"}},
	{"test", []string{"tes", "test", "testing", "footest", "testfoo"}, []string{"test", "testing", "testfoo"}},
	{"test", []string{"abs", "asdb", "asdg", "afd", "asdfe"}, []string{}},
}

func TestFind(t *testing.T) {
	for _, test := range findtests {
		actual, _ := Find(test.needle, test.haystack)

		if len(test.out) != len(actual) {
			t.Errorf("Find(%v,%v) = %v; want %v", test.needle, test.haystack, actual, test.out)
			return
		}

	LOOP:
		for i, a := range actual {
			if a != test.out[i] {
				t.Errorf("Find(%v,%v) = %v; want %v", test.needle, test.haystack, actual, test.out)
				break LOOP
			}
		}
	}

}

func ExampleFind() {
	matches, err := Find("test", []string{"tes", "test", "testing", "footest", "testfoo"})
	if err != nil {
		//handle err
	}

	fmt.Println(matches)
	// Output: [test testing testfoo]
}
