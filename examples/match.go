package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jjshoe/go-grep"
)

func main() {
	l := []string{"a", "aa", "b", "ccc"}

	n := grep.String("^a$", l)

	spew.Dump(l)
	spew.Dump(n)
}
