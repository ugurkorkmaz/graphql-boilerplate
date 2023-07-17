package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

var appPort string

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set variable values from environment
	appPort = os.Getenv("APP_PORT")
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		Help()
		return
	}

	command := os.Args[1]
	switch command {
	case "build":
		err := Build()
		handleError("Failed to build binary", err)
	case "install":
		err := Install()
		handleError("Failed to install dependencies", err)
	case "update":
		err := UpdateDependencies()
		handleError("Failed to update dependencies", err)
	case "gqlgen":
		err := Gqlgen()
		handleError("Failed to generate gqlgen", err)
	case "run":
		err := Run()
		handleError("Failed to run binary", err)
	case "test":
		err := Test()
		handleError("Tests failed", err)
	case "clean":
		err := Clean()
		handleError("Failed to clean", err)
	case "all":
		err := Gqlgen()
		handleError("Failed to generate gqlgen", err)

		err = UpdateDependencies()
		handleError("Failed to update dependencies", err)

		err = Build()
		handleError("Failed to build binary", err)

		err = Run()
		handleError("Failed to run binary", err)

	default:
		fmt.Println("Invalid command")
		Help()
	}
}

func handleError(message string, err error) {
	if err != nil {
		log.Fatalf("Error: %s: %v", message, err)
	}
}

func Help() {
	fmt.Println("Available commands:")
	fmt.Println("build     - Compiles the binary")
	fmt.Println("install   - Installs dependencies for the go.mod file")
	fmt.Println("update    - Updates project dependencies")
	fmt.Println("gqlgen    - Generates code from GraphQL schema")
	fmt.Println("frontend  - Builds the frontend templates")
	fmt.Println("run       - Runs the binary")
	fmt.Println("test      - Runs project tests and lints")
	fmt.Println("clean     - Removes the binary")
	fmt.Println("all       - Runs all the commands in order")
	fmt.Println("help      - Shows available commands")
}

// compiles the binary
func Build() error {
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "app", ".")
	return cmd.Run()
}

// dependencies for the gomod
func Install() error {
	fmt.Println("Installing dependencies...")
	cmd := exec.Command("go", "mod", "install")
	return cmd.Run()
}

func UpdateDependencies() error {
	cmd := exec.Command("go", "get", "-u", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Updating dependencies...")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to update dependencies: %v", err)
	}

	fmt.Println("Dependencies updated successfully.")

	return nil
}

// generates code from graphql schema
func Gqlgen() error {
	fmt.Println("Generating gqlgen...")
	cmd := exec.Command("go", "run", "github.com/99designs/gqlgen@latest")
	return cmd.Run()
}

// runs the binary
func Run() error {
	fmt.Println("Running...")
	fmt.Println("Server is running at http://127.0.0.1:" + appPort)
	fmt.Println("Press Ctrl+C to stop the server.")

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
