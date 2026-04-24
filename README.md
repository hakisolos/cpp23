# cspark

cspark is a high-performance build tool written in Go, designed to simplify the use of C++23 Modules with Clang. It automates the setup of the standard library module (import std;) and provides a streamlined interface for compiling modern C++ projects.

## Quick Installation

Visit the [releases page](https://github.com/hakisolos/cpp23/releases) to choose the binary for your operating system, then use the appropriate command below:

### Linux

```bash
curl -sL https://github.com/hakisolos/cpp23/releases/download/<release-version>/<binary-release> -o cspark && chmod +x cspark && sudo mv cspark /usr/local/bin/
```

Example:
```bash
curl -sL https://github.com/hakisolos/cpp23/releases/download/v0.0.1/cspark-linux-binary -o cspark && chmod +x cspark && sudo mv cspark /usr/local/bin/
```

### macOS

```bash
curl -sL https://github.com/hakisolos/cpp23/releases/download/<release-version>/<binary-release> -o cspark && chmod +x cspark && sudo mv cspark /usr/local/bin/
```

Example:
```bash
curl -sL https://github.com/hakisolos/cpp23/releases/download/v0.0.1/cspark-macos-binary -o cspark && chmod +x cspark && sudo mv cspark /usr/local/bin/
```

### Windows

```powershell
Invoke-WebRequest -Uri "https://github.com/hakisolos/cpp23/releases/download/<release-version>/<binary-release>" -OutFile cspark.exe
Move-Item -Path cspark.exe -Destination "C:\Windows\System32\" -Force
```

Example:
```powershell
Invoke-WebRequest -Uri "https://github.com/hakisolos/cpp23/releases/download/v0.0.1/cspark-windows-binary.exe" -OutFile cspark.exe
Move-Item -Path cspark.exe -Destination "C:\Windows\System32\" -Force
```

## Features

- **Automatic Environment Setup**: Detects if Clang 19+ is installed; if not, it offers to install it for you.
- **One-Time Precompilation**: Automatically precompiles the C++23 Standard Library module (std.pcm) on its first run.
- **Functional File Discovery**: Effortlessly finds and links .cpp, .cppm, and .cc files.
- **Cross-Platform**: Designed to work across Linux, macOS, and Windows.

## Usage

### 1. Compile Specific Files

To compile a specific set of files (including modules), simply list them:

```bash
cspark main.cpp math.cppm math.cpp
```

### 2. Compile Everything

To automatically find and compile all C++ source and module files in the current directory:

```bash
cspark -all
```

### 3. Check Version

```bash
cspark --version
```

### 4. Uninstall

```bash
cspark --uninstall
```

### 5. Help Menu

```bash
cspark --help
```

## How it Works

1. **Environment Check**: On the first run, cspark ensures clang++-19 and libc++ are present.
2. **Standard Module**: It locates your system's std.cppm and compiles it into a .pcm file stored in your home directory (~/.cspark_std.pcm).
3. **Compilation**: It executes the compilation string using the required C++23 flags: `-std=c++23 -stdlib=libc++ -fmodule-file=std=...`

## Requirements

- **OS**: Linux (Ubuntu/Debian recommended), macOS, or Windows.
- **Compiler**: Clang 19 or later (cspark can automate this installation on most systems).

## Pro Tip for Developers

If you are adding cspark to a CI/CD pipeline, you can use the `-all` flag to ensure all new modules added to your directory are automatically tracked and compiled without updating your build scripts.

## Acknowledgments

cspark is built on top of exceptional open-source projects:

- **[LLVM Project](https://llvm.org/)** - The infrastructure and compiler framework
- **[Clang](https://clang.llvm.org/)** - The C/C++ compiler front-end  
- **[libc++](https://libcxx.llvm.org/)** - The C++ standard library implementation

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to contribute.

## License

cspark is released under the MIT License. See [LICENSE](LICENSE) for details.

---

**Created by:** [Haki](https://haki.top)  
**Email:** hakixer@gmail.com
