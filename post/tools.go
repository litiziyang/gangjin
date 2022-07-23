//go:build tools
// +build tools

//go:generate go run github.com/99designs/gqlgen init
package main

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/vektah/dataloaden"
)
