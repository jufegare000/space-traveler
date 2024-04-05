package test

import "testing"

func TestAdd(testing *testing.T) {

	var got int = 10
	want := uint(10)

	if uint(got) != want {
		testing.Errorf("got %v, wanted %v", got, want)
	}
}
