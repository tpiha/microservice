package main

import (
	"testing"
)

func TestProcessDiff(t *testing.T) {
	a := 55.5
	b := 35.25

	wm := initWorkerManager()
	c := wm.processDiff(a, b)

	if c != 2025 {
		t.Errorf("got %d, want %d", c, 2025)
	}
}

func TestAddDatapoint(t *testing.T) {
	wm := initWorkerManager()

	for i := 0; i < 10; i++ {
		p := &Payload{}
		wm.addDatapoint(p)
	}

	if len(wm.Jobs) != 10 {
		t.Errorf("got %d, want %d", len(wm.Jobs), 10)
	}
}
