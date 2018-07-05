package main

import "testing"

func TestSum(t *testing.T) {
	t.Fatalf("Exected response version to be %d but it is %d", DefaultVersion, DefaultVersion)
}
