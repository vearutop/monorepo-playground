package main

import (
	"github.com/vearutop/monorepo-playground/deep/another/v2"
	monorepo "github.com/vearutop/monorepo-playground/go/v3"
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

// .........

// Go module is now go/v3
// No v3.x.x tags are available.

// go mod tidy
// go: finding module for package github.com/vearutop/monorepo-playground/go/v3
// go: downloading github.com/vearutop/monorepo-playground/go/v3 v3.0.0-20221115201755-a0e8e409bcfb
// go: found github.com/vearutop/monorepo-playground/go/v3 in github.com/vearutop/monorepo-playground/go/v3 v3.0.0-20221115201755-a0e8e409bcfb

// .........

// Tag is created as v3.0.2.

// go get -u github.com/vearutop/monorepo-playground/go/v3@v3.0.2
// go: downloading github.com/vearutop/monorepo-playground v3.0.2+incompatible
// go: module github.com/vearutop/monorepo-playground@v3.0.2 found (v3.0.2+incompatible), but does not contain package github.com/vearutop/monorepo-playground/go/v3

// .........

// Tag is created as go/v3.0.5.

// Upgrading to newly tagged version fails at first as it does not seem to pull new tags.

// go get -u github.com/vearutop/monorepo-playground/go/v3@v3.0.5
// go: github.com/vearutop/monorepo-playground/go/v3@v3.0.5: invalid version: resolves to version v3.0.6-0.20221115201755-a0e8e409bcfb (v3.0.5 is not a tag)

// Upgrading to @latest works.

// go get -u github.com/vearutop/monorepo-playground/go/v3@latest
// go: downloading github.com/vearutop/monorepo-playground/go/v3 v3.0.5
// go: upgraded github.com/vearutop/monorepo-playground/go/v3 v3.0.0-20221115201755-a0e8e409bcfb => v3.0.5

// Upgrading to tag directly works after that too.

// go get -u github.com/vearutop/monorepo-playground/go/v3@v3.0.5
// <no output>
