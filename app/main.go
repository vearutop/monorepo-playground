package main

import (
	"github.com/vearutop/monorepo-playground/deep/another/v2"
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

// Second tag is created as v1.0.0.

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

// .........

// Third tag is created as v2.0.1.
// Another module is now deep/another/v2.

// go get -u github.com/vearutop/monorepo-playground/go@v2.0.1
// go: downloading github.com/vearutop/monorepo-playground v2.0.1+incompatible
// go: module github.com/vearutop/monorepo-playground@v2.0.1 found (v2.0.1+incompatible), but does not contain package github.com/vearutop/monorepo-playground/go

// go get -u github.com/vearutop/monorepo-playground/deep/another/v2
// go: downloading github.com/vearutop/monorepo-playground/deep/another/v2 v2.0.0-20221115200137-ebfd72605c0e
// go: added github.com/vearutop/monorepo-playground/deep/another/v2 v2.0.0-20221115200137-ebfd72605c0e

// go get -u github.com/vearutop/monorepo-playground/deep/another/v2@v2.0.0
// go: module github.com/vearutop/monorepo-playground@v2.0.0 found (v2.0.0+incompatible), but does not contain package github.com/vearutop/monorepo-playground/deep/another/v2

// go get -u github.com/vearutop/monorepo-playground/deep/another/v2@v2.0.1
// go: module github.com/vearutop/monorepo-playground@v2.0.1 found (v2.0.1+incompatible), but does not contain package github.com/vearutop/monorepo-playground/deep/another/v2

// go get -u github.com/vearutop/monorepo-playground/go/v2
// go: downloading github.com/vearutop/monorepo-playground/go v0.0.0-20221115200137-ebfd72605c0e
// go: module github.com/vearutop/monorepo-playground/go@upgrade found (v0.0.0-20221115200137-ebfd72605c0e), but does not contain package github.com/vearutop/monorepo-playground/go/v2

// go get -u github.com/vearutop/monorepo-playground/go@ebfd72605c0e
// go: upgraded github.com/vearutop/monorepo-playground/go v0.0.0-20221115195614-77518b3db090 => v0.0.0-20221115200137-ebfd72605c0e
