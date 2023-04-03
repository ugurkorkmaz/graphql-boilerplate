//go:build mage
// +build mage

/*
This is a magefile.  It can be used to run any build related tasks
in a reproducible way.  See https://magefile.org for more info.
*/
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps, Gqlgen, Frontend)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "app", ".")
	return cmd.Run()
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "mod", "tidy")
	return cmd.Run()
}

// Generate gqlgen
func Gqlgen() error {
	fmt.Println("Generating gqlgen...")
	// gqlgen generate, See https://gqlgen.com/getting-started/
	cmd := exec.Command("go", "run", "github.com/99designs/gqlgen@latest")
	return cmd.Run()
}

// Generate templates for the frontend
func Frontend() error {
	fmt.Println("Generating templates...")
	fmt.Println("Frontend framework installation or build steps here...")
	return nil
}

// Watch for changes and rebuild the project automatically
func Watch() error {
	fmt.Println("Starting air...")
	// air, See https://github.com/cosmtrek/air#usage
	cmd := exec.Command("go", "run", "github.com/cosmtrek/air@latest")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	// Remove the binary
	os.RemoveAll("app")
}
