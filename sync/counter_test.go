package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter three times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := NewCounter()
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

func assertCount(t testing.TB, counter *Counter, expected int) {
	t.Helper()
	if counter.Value() != expected {
		t.Errorf("got %v expected %v", counter.Value(), expected)
	}
}
