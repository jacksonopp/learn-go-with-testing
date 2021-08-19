package iteration

import "testing"

func assertCorrectString(t testing.TB, exp, got string) {
	t.Helper()
	if got != exp {
		t.Errorf("expected %q but got %q", exp, got)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("it should run 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectString(t, expected, repeated)
	})

	t.Run("it should run 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		assertCorrectString(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
