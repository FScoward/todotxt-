package main

import (
	"strings"
	"testing"
)

type Type string

const (
	PRJ  Type = Type("PRJ")
	TODO      = Type("TODO")
	DONE      = Type("DONE")
	DOC       = Type("DOC")
)

func TestParser(t *testing.T) {
	var projectLine = "PRJ this is a title 2023-11-01"
	var head = strings.Split(projectLine, " ")
	if Type(head[0]) != PRJ {
		t.Error("", head[0], "")
	}
}
