package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hakisolos/cpp23/src"
)

func main() {
	cfg := src.GetConfig()

	args := os.Args[1:]
	if len(args) == 0 {
		src.ShowHelp()
		return
	}

	switch args[0] {
	case "--help", "-h":
		src.ShowHelp()
		return
	case "--version", "-v":
		fmt.Printf("cspark version %s\n", src.AppVersion)
		return
	case "--uninstall":
		fmt.Print("This will remove cspark and the precompiled modules. Are you sure? (yes/no): ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "yes" {
			fmt.Println("Uninstall cancelled.")
			return
		}

		if _, err := os.Stat(cfg.PcmPath); err == nil {
			os.Remove(cfg.PcmPath)
			fmt.Println("Removed precompiled modules.")
		}

		exePath, err := os.Executable()
		if err == nil {
			fmt.Printf("Removing binary at %s...\n", exePath)
			cmd := exec.Command("sudo", "rm", exePath)
			if err := cmd.Run(); err != nil {
				fmt.Printf("Note: Could not remove binary automatically. Please run: sudo rm %s\n", exePath)
			}
		}

		fmt.Print("Do you also want to uninstall Clang 19? (yes/no): ")
		var clangConfirm string
		fmt.Scanln(&clangConfirm)
		if clangConfirm == "yes" {
			src.UninstallClang()
		}

		fmt.Println("Uninstallation complete.")
		return
	}

	if _, err := exec.LookPath(cfg.ClangBin); err != nil {
		if err := src.InstallClang(); err != nil {
			fmt.Printf("Failed to install Clang: %v\n", err)
			os.Exit(1)
		}
	}

	if _, err := os.Stat(cfg.PcmPath); os.IsNotExist(err) {
		src.LoadingAction("Precompiling C++23 Standard Library", func() error {
			return exec.Command(cfg.ClangBin, "-std=c++23", "-stdlib=libc++", "--precompile", "-o", cfg.PcmPath, cfg.StdPath).Run()
		})
	}

	var targets []string
	if args[0] == "-all" {
		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && src.IsSourceFile(path) {
				targets = append(targets, path)
			}
			return nil
		})
	} else {
		for _, arg := range args {
			if src.IsSourceFile(arg) {
				targets = append(targets, arg)
			}
		}
	}

	if len(targets) == 0 {
		fmt.Println("No valid C++ source files found.")
		return
	}

	compileArgs := []string{"-std=c++23", "-stdlib=libc++", "-fmodule-file=std=" + cfg.PcmPath}
	compileArgs = append(append(compileArgs, targets...), "-o", src.OutputName)

	if len(targets) == 1 {
		fmt.Printf("Building %s -> %s\n", targets[0], src.OutputName)
	} else {
		fmt.Printf("Building [%d files] -> %s\n", len(targets), src.OutputName)
	}

	src.LoadingAction("Compiling", func() error {
		cmd := exec.Command(cfg.ClangBin, compileArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	})

	fmt.Printf("Build successful: ./%s\n", src.OutputName)
}
