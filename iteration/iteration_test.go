package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q, but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: bbb
}

func TestCompare(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Expected %d, but got %d", want, got)
		}
	}

	t.Run("test equal strings", func(t *testing.T) {
		compared := strings.Compare("hello, world", "hello, world")
		expected := 0

		assertCorrectMessage(t, compared, expected)
	})

	t.Run("test string a lexicographically larger than string b", func(t *testing.T) {
		compared := strings.Compare("hello, world!", "hello, world")
		expected := 1

		assertCorrectMessage(t, compared, expected)
	})

	t.Run("test string a lexicographically smaller than string b", func(t *testing.T) {
		compared := strings.Compare("hello!", "hello, world")
		expected := -1

		assertCorrectMessage(t, compared, expected)
	})
}
