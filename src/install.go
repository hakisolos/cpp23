package src

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InstallClang() error {
	fmt.Println("Clang 19 not found. Attempting to install...")
	switch runtime.GOOS {
	case "linux":
		// Specifically for Ubuntu/Debian as requested
		steps := [][]string{
			{"sudo", "apt-get", "update"},
			{"sudo", "apt-get", "install", "-y", "lsb-release", "wget", "software-properties-common", "gnupg"},
			{"bash", "-c", "wget https://apt.llvm.org/llvm.sh && chmod +x llvm.sh && sudo ./llvm.sh 19 && rm llvm.sh"},
			{"sudo", "apt-get", "install", "-y", "libc++-19-dev", "libc++abi-19-dev"},
		}
		for _, step := range steps {
			cmd := exec.Command(step[0], step[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	case "darwin":
		return exec.Command("brew", "install", "llvm").Run()
	case "windows":
		return exec.Command("winget", "install", "LLVM.LLVM").Run()
	}
	return nil
}
