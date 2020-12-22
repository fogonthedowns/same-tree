package main

import (
	"testing"
)

func TestFun(t *testing.T) {

	two := TreeNode{Val: 2}
	three := TreeNode{Val: 3}
	one := TreeNode{Val: 1, Left: &two, Right: &three}
	a := TreeNode{Val: 2}
	b := TreeNode{Val: 3}
	c := TreeNode{Val: 1, Left: &a, Right: &b}

	out := isSameTree(&one, &c)
	if out != true {
		t.Errorf("got %t, want %t", out, true)
	}
}
