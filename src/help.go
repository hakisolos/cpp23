package src

import "fmt"

const (
	AppVersion = "0.0.1"
	OutputName = "app"
)

func ShowHelp() {
	fmt.Printf("cspark %s - C++23 Modules Build Tool\n\n", AppVersion)
	fmt.Println("Usage:")
	fmt.Println("  cspark [files...]       Compile specific .cpp or .cppm files")
	fmt.Println("  cspark -all             Compile all C++ files in current directory")
	fmt.Println("  cspark --version        Show version information")
	fmt.Println("  cspark --help           Show this help message")
	fmt.Println("\nConfiguration:")
	fmt.Println("  Uses Clang 19+ with libc++ for C++23 'import std' support.")
}
