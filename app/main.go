package main

import (
	"github.com/vearutop/monorepo-playground/deep/another"
	monorepo "github.com/vearutop/monorepo-playground/go"
)

func main() {
	println(another.Bar, monorepo.Foo)
}
