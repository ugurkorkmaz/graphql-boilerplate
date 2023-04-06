//go:build mage
// +build mage

// GraphQL Boilerplate for Fullstack
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// compiles the binary
func Build() error {
	mg.Deps(Install, Frontend)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "app", ".")
	return cmd.Run()
}

// dependencies for the gomod
func Install() error {
	// Colorize the output
	fmt.Println("Installing dependencies...")
	cmd := exec.Command("go", "mod", "tidy")
	return cmd.Run()
}

// generates code from graphql schema
func Gqlgen() error {
	fmt.Println("Generating gqlgen...")
	// gqlgen generate, See https://gqlgen.com/getting-started/
	cmd := exec.Command("go", "run", "github.com/99designs/gqlgen@latest")
	return cmd.Run()
}

// builds the frontend templates
func Frontend() error {
	fmt.Println("Generating templates...")
	// npm build --prefix templates
	cmd := exec.Command("npm", "run", "build", "--prefix", "template")
	return cmd.Run()
}

// runs the binary
func Run() error {
	fmt.Println("Running...")
	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// project tests and lints
func Test() error {
	fmt.Println("Testing...")
	cmd := exec.Command("go", "test", "./...")
	return cmd.Run()
}

// Clean removes the binary
func Clean() error {
	fmt.Println("Cleaning...")
	return os.Remove("app")
}
