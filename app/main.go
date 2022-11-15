package main

import (
	"github.com/vearutop/monorepo-playground/deep/another"
	monorepo "github.com/vearutop/monorepo-playground/go"
)

func main() {
	println(another.Bar, monorepo.Foo)
}

// First tag is created as v2.0.0.

// go get -u
// go: downloading github.com/vearutop/monorepo-playground v2.0.0+incompatible
// go: downloading github.com/vearutop/monorepo-playground/deep/another v0.0.0-20221115195248-9e8b3614d0ac
// go: downloading github.com/vearutop/monorepo-playground/go v0.0.0-20221115195248-9e8b3614d0ac
// go: upgraded github.com/vearutop/monorepo-playground/deep/another v0.0.0-20221115195156-82371b5008e4 => v0.0.0-20221115195248-9e8b3614d0ac
// go: upgraded github.com/vearutop/monorepo-playground/go v0.0.0-20221115195156-82371b5008e4 => v0.0.0-20221115195248-9e8b3614d0ac

// go get -u github.com/vearutop/monorepo-playground/go@v2.0.0
// go: module github.com/vearutop/monorepo-playground@v2.0.0 found (v2.0.0+incompatible), but does not contain package github.com/vearutop/monorepo-playground/go

// go get -u github.com/vearutop/monorepo-playground/deep/another@v2.0.0
// go: module github.com/vearutop/monorepo-playground@v2.0.0 found (v2.0.0+incompatible), but does not contain package github.com/vearutop/monorepo-playground/deep/another

// .........

// Second tag is created as v1.0.0

// go get -u github.com/vearutop/monorepo-playground/deep/another@v1.0.0
// go: downloading github.com/vearutop/monorepo-playground v1.0.0
// go: module github.com/vearutop/monorepo-playground@v1.0.0 found, but does not contain package github.com/vearutop/monorepo-playground/deep/another

// go get -u github.com/vearutop/monorepo-playground/go@v1.0.0
// go: module github.com/vearutop/monorepo-playground@v1.0.0 found, but does not contain package github.com/vearutop/monorepo-playground/go

// go get -u
// go: downloading github.com/vearutop/monorepo-playground/deep/another v0.0.0-20221115195614-77518b3db090
// go: downloading github.com/vearutop/monorepo-playground/go v0.0.0-20221115195614-77518b3db090
// go: upgraded github.com/vearutop/monorepo-playground/deep/another v0.0.0-20221115195248-9e8b3614d0ac => v0.0.0-20221115195614-77518b3db090
// go: upgraded github.com/vearutop/monorepo-playground/go v0.0.0-20221115195248-9e8b3614d0ac => v0.0.0-20221115195614-77518b3db090
