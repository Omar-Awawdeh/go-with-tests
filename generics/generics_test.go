package generics

import "testing"

func TestAssertFunction(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 0)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})

	t.Run("asserting boolean", func(t *testing.T) {
		AssertFalse(t, false)
		AssertTrue(t, true)
	})
}
