package src

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/hakisolos/cpp23/models"
)

func GetConfig() models.Config {
	home, _ := os.UserHomeDir()
	pcmLocation := filepath.Join(home, ".cspark_std.pcm")

	switch runtime.GOOS {
	case "darwin":
		return models.Config{
			ClangBin: "/opt/homebrew/opt/llvm/bin/clang++",
			StdPath:  "/opt/homebrew/opt/llvm/share/libc++/v1/std.cppm",
			PcmPath:  pcmLocation,
		}
	case "windows":
		return models.Config{
			ClangBin: "clang++",
			StdPath:  `C:\Program Files\LLVM\share\libc++\v1\std.cppm`,
			PcmPath:  filepath.Join(os.Getenv("APPDATA"), "cspark_std.pcm"),
		}
	default: // Linux
		return models.Config{
			ClangBin: "clang++-19",
			StdPath:  "/usr/lib/llvm-19/share/libc++/v1/std.cppm",
			PcmPath:  pcmLocation,
		}
	}
}

func IsSourceFile(path string) bool {
	ext := filepath.Ext(path)
	for _, v := range []string{".cpp", ".cppm", ".cc", ".cxx"} {
		if ext == v {
			return true
		}
	}
	return false
}
