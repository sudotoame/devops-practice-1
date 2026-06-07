package logic

import "testing"

func TestNumber(T *testing.T) {
	logic := NewLogic()
	want := 1

	logic.Incriment()
	have := logic.NumberOutput()

	if have != int64(want) {
		T.Fatalf("want: %v, we have: %v", want, have)
	}
}
