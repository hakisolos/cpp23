package src

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func UninstallClang() {
	fmt.Println("Uninstalling Clang...")

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("sudo", "apt-get", "remove", "-y", "clang-19", "libc++-19-dev", "libc++abi-19-dev")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error removing packages: %v\n", err)
		}
		exec.Command("sudo", "apt-get", "autoremove", "-y").Run()

	case "darwin":
		cmd := exec.Command("brew", "uninstall", "llvm")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	case "windows":
		cmd := exec.Command("winget", "uninstall", "LLVM.LLVM")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	default:
		fmt.Printf("Automatic uninstall not supported for %s. Please remove Clang manually.\n", runtime.GOOS)
	}
}
