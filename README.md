# Readme is AI generated >_> too lazy
# C++23 Modules Setup on Linux (Clang 20 + libc++)

> ⚠️ **Project Status**: ACTIVE DEVELOPMENT & TESTING  
> This project is in active development. Testing and contributions are welcome! See [Contributing](#contributing) if you'd like to help.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Status: In Progress](https://img.shields.io/badge/Status-In%20Progress-orange)](https://github.com/yourusername/cpp23#contributing)
[![Help Wanted](https://img.shields.io/badge/Help-Wanted-brightgreen)](CONTRIBUTING.md)

A complete guide to setting up `import std;` with a permanent `compile` command that works anywhere on your Linux system.

## ⚠️ Current Status

This project is **actively seeking help** with:
- ✅ **Testing** - Test on different distros and report issues
- ✅ **Documentation** - Improve clarity and add examples  
- ✅ **Automation** - Improve build scripts and error handling
- ✅ **Examples** - Create more test programs showcasing C++23 features
- ✅ **Troubleshooting** - Document workarounds and edge cases

**Please open issues, contribute fixes, and share your experience!**

### Recent Updates

- 🚀 **Enhanced `compile` command** - Now intelligently handles both single-file and multi-file C++23 projects
- 🎨 **Color-coded build output** - Better error visibility with green/red/yellow highlighting
- 📦 **Module support** - Automatic detection and compilation of C++23 module interfaces (`.cppm`, `.ixx`)

## System Requirements

- **Ubuntu 22.04+** (or any Debian-based distro)
- **Internet connection** for downloading packages
- **~500MB disk space** for Clang and libc++

## One-Time Installation

### Step 1: Install Clang 20 and libc++

Run these commands in terminal:

```bash
# Add LLVM repository
wget -qO- https://apt.llvm.org/llvm-snapshot.gpg.key | sudo tee /etc/apt/trusted.gpg.d/apt.llvm.org.asc
sudo add-apt-repository "deb http://apt.llvm.org/$(lsb_release -cs)/ llvm-toolchain-$(lsb_release -cs)-20 main"
sudo apt update

# Install Clang 20 and libc++ with modules support
sudo apt install -y clang-20 libc++-20-dev libc++abi-20-dev lld-20 clang-tools-20
```

### Step 2: Create build script and generate std module

```bash
cd ~
mkdir cpp && cd cpp

cat > build_simple.sh << 'EOF'
#!/bin/bash
MODULE_DIR="/usr/lib/llvm-20/share/libc++/v1"
BUILD_DIR="build"

if [ ! -f "$BUILD_DIR/std.pcm" ]; then
    echo "Building std module (one time only)..."
    mkdir -p $BUILD_DIR
    clang++-20 -std=c++23 -stdlib=libc++ \
        --precompile -Xclang -emit-module-interface \
        -o $BUILD_DIR/std.pcm \
        $MODULE_DIR/std.cppm 2>/dev/null
fi

# Test compilation
cat > main.cpp << 'MAINEOF'
import std;

int main() {
    std::println("C++23 modules are working!");
    return 0;
}
MAINEOF

clang++-20 -std=c++23 -stdlib=libc++ \
    -fmodule-file=std=$BUILD_DIR/std.pcm \
    main.cpp \
    -o $BUILD_DIR/my_app \
    && $BUILD_DIR/my_app
EOF

chmod +x build_simple.sh
./build_simple.sh
```

**Expected output:**

```
Building std module (one time only)...
C++23 modules are working!
```

### Step 3: Cache the std module globally

```bash
# Create cache directory and copy the precompiled module
mkdir -p ~/.cache/clang20-std
cp ~/cpp/build/std.pcm ~/.cache/clang20-std/std.pcm

# Verify it exists
ls -lh ~/.cache/clang20-std/std.pcm
```

### Step 4: Install the permanent 'compile' command

```bash
sudo tee /usr/local/bin/compile > /dev/null << 'EOF'
#!/bin/bash

STD_PCM="$HOME/.cache/clang20-std/std.pcm"
PROJECT_DIR="$(pwd)"

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Check if std module exists
if [ ! -f "$STD_PCM" ]; then
    echo -e "${RED}✗ Std module not found at: $STD_PCM${NC}"
    echo "Please run: cp ~/cpp/build/std.pcm ~/.cache/clang20-std/"
    exit 1
fi

# Check if we're in a multi-file project (has .cppm or multiple .cpp files)
MODULE_FILES=$(find . -maxdepth 1 -type f \( -name "*.cppm" -o -name "*.ixx" \) 2>/dev/null | wc -l)
CPP_FILES=$(find . -maxdepth 1 -name "*.cpp" 2>/dev/null | wc -l)

# CASE 1: Single file compilation (original behavior)
if [ $MODULE_FILES -eq 0 ] && [ $CPP_FILES -eq 1 ] && [ -n "$1" ]; then
    INPUT_FILE="$1"
    
    if [ ! -f "$INPUT_FILE" ]; then
        echo -e "${RED}✗ Error: File '$INPUT_FILE' not found${NC}"
        exit 1
    fi
    
    if [ -n "$2" ] && [ "$2" != "-o" ]; then
        OUTPUT_NAME="$2"
    else
        OUTPUT_NAME="${INPUT_FILE%.cpp}"
    fi
    
    echo -e "${GREEN}Compiling $INPUT_FILE -> $OUTPUT_NAME${NC}"
    clang++-20 -std=c++23 -stdlib=libc++ \
        -fmodule-file=std="$STD_PCM" \
        "$INPUT_FILE" -o "$OUTPUT_NAME"
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Success! Run: ./$OUTPUT_NAME${NC}"
    else
        echo -e "${RED}✗ Compilation failed${NC}"
        exit 1
    fi
    exit 0
fi

# CASE 2: Multi-file project (new behavior)
echo -e "${GREEN}Detected multi-file project...${NC}"

# Find all module interface files
MODULE_FILES=$(find . -maxdepth 1 -type f \( -name "*.cppm" -o -name "*.ixx" \) 2>/dev/null)
REGULAR_CPP=$(find . -maxdepth 1 -name "*.cpp" 2>/dev/null)

# Step 1: Build each module interface
for module in $MODULE_FILES; do
    module_name=$(basename "$module" | sed 's/\.[^.]*$//')
    echo -e "${YELLOW}Building module: $module_name${NC}"
    
    clang++-20 -std=c++23 -stdlib=libc++ \
        -fmodule-file=std="$STD_PCM" \
        -c "$module" -o "${module_name}.pcm" \
        -Xclang -emit-module-interface 2>/dev/null
    
    if [ $? -ne 0 ]; then
        echo -e "${RED}Failed to build module: $module_name${NC}"
        exit 1
    fi
done

# Step 2: Build all cpp files into objects
OBJECTS=""
for cpp_file in $REGULAR_CPP; do
    obj_name="${cpp_file%.cpp}.o"
    echo -e "${YELLOW}Compiling: $cpp_file${NC}"
    
    # Collect module flags
    MODULE_FLAGS=""
    for pcm in *.pcm; do
        if [ -f "$pcm" ]; then
            mod_name="${pcm%.pcm}"
            MODULE_FLAGS="$MODULE_FLAGS -fmodule-file=$mod_name=$pcm"
        fi
    done
    
    clang++-20 -std=c++23 -stdlib=libc++ \
        -fmodule-file=std="$STD_PCM" \
        $MODULE_FLAGS \
        -c "$cpp_file" -o "$obj_name" 2>/dev/null
    
    if [ $? -ne 0 ]; then
        echo -e "${RED}Failed to compile: $cpp_file${NC}"
        exit 1
    fi
    OBJECTS="$OBJECTS $obj_name"
done

# Step 3: Link everything
echo -e "${YELLOW}Linking...${NC}"
clang++-20 -stdlib=libc++ $OBJECTS -o my_app

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Build successful! Run: ./my_app${NC}"
else
    echo -e "${RED}✗ Linking failed${NC}"
    exit 1
fi
EOF

sudo chmod +x /usr/local/bin/compile
```

## How to Use (After Setup)

### Single File Projects

For simple single-file C++23 programs:

```bash
# Compile any C++ file
compile myprogram.cpp

# Run it
./myprogram

# Or compile with specific output name
compile myprogram.cpp myapp
./myapp
```

### Multi-File Projects

For larger projects with multiple files and/or C++ modules (`.cppm` or `.ixx` files):

```bash
# Just run compile - it automatically detects and builds everything
compile

# Output binary is created as 'my_app'
./my_app
```

The enhanced `compile` script handles:
- ✅ Module interface compilation (`.cppm`, `.ixx`)
- ✅ Multi-file object compilation with module dependencies
- ✅ Automatic linking of all object files
- ✅ Color-coded output for easy debugging

### Example: Single File

```bash
cd ~
mkdir myproject && cd myproject

cat > hello.cpp << 'EOF'
import std;

int main() {
    std::println("Hello from C++23 modules!");
    
    std::vector<int> numbers = {1, 2, 3, 4, 5};
    for (int n : numbers) {
        std::println("Number: {}", n);
    }
    
    return 0;
}
EOF

compile hello.cpp
./hello
```

### Example: Multi-File Project with Modules

```bash
cd myproject

# Create a module interface
cat > math_utils.cppm << 'EOF'
export module math_utils;

export int add(int a, int b) {
    return a + b;
}
EOF

# Create implementation
cat > main.cpp << 'EOF'
import std;
import math_utils;

int main() {
    std::println("2 + 3 = {}", add(2, 3));
    return 0;
}
EOF

# Compile everything automatically
compile

# Run the built application
./my_app
```

## Optional: Create Aliases for Faster Workflow

Add these to your `~/.bashrc` or `~/.zshrc`:

```bash
# Quick compile and run
echo 'alias run="compile fek.cpp && ./fek"' >> ~/.bashrc

# Compile with current file
echo 'alias rc="compile main.cpp && ./main"' >> ~/.bashrc

# Reload bashrc
source ~/.bashrc
```

## Verify Installation

Run these commands to ensure everything is working:

```bash
# Check compile command exists
which compile
# Output: /usr/local/bin/compile

# Check std module exists
ls -lh ~/.cache/clang20-std/std.pcm
# Output: -rw-r--r-- 1 haki haki ~2.5M std.pcm

# Check compiler exists
which clang++-20
# Output: /usr/bin/clang++-20

# Test compilation
echo 'import std; int main() { std::println("OK"); }' > test.cpp
compile test.cpp && ./test.cpp
# Output: OK
```

## Troubleshooting

### Problem: "Std module not found"
```bash
# Check if module exists
ls -lh ~/.cache/clang20-std/std.pcm

# If missing, rebuild it
cd ~/cpp && ./build_simple.sh
```

### Problem: "clang++-20: command not found"
```bash
# Verify installation
clang++-20 --version

# If not found, reinstall:
sudo apt install -y clang-20 libc++-20-dev
```

### Problem: Compilation fails
- Check your C++23 syntax
- Verify you're using `import std;` not `#include`
- See [DEVELOPMENT.md](DEVELOPMENT.md) for more details

## Contributing

**We need your help!** Whether it's testing, documentation, or code:

- 🧪 [Report bugs](CONTRIBUTING.md)
- 📝 [Improve documentation](CONTRIBUTING.md)
- 🔧 [Submit fixes and features](CONTRIBUTING.md)
- 🤔 [Request features](CONTRIBUTING.md)

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Development

For local development and testing, see [DEVELOPMENT.md](DEVELOPMENT.md).

## Code of Conduct

Please read our [Code of Conduct](CODE_OF_CONDUCT.md) before contributing.

## Security

For security issues, see [SECURITY.md](SECURITY.md).

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Support

- Open an [issue](https://github.com/yourusername/cpp23/issues) for bugs
- Start a [discussion](https://github.com/yourusername/cpp23/discussions) for questions
- See [DEVELOPMENT.md](DEVELOPMENT.md) for common issues

## Acknowledgments

- Built for C++23 enthusiasts learning modules on Linux
- Inspired by the LLVM/Clang project
- Thanks to all contributors!

---

## Issue: 'compile: command not found'

**Solution:** Reinstall the compile script

```bash
sudo chmod +x /usr/local/bin/compile
```

### Issue: 'Std module not found'

**Solution:** Rebuild the std module

```bash
cd ~/cpp
./build_simple.sh
cp ~/cpp/build/std.pcm ~/.cache/clang20-std/std.pcm
```

### Issue: 'clang++-20: command not found'

**Solution:** Reinstall Clang

```bash
sudo apt install --reinstall clang-20
```

### Issue: 'undefined identifier: println'

**Solution:** Use `std::println` or add `using std::println;`

```cpp
import std;
int main() {
    std::println("hello");  // Correct
    // println("hello");    // Wrong - needs std::
    return 0;
}
```

## Files Installed on Your System

| Component | Location | Persists Reboot? |
|-----------|----------|------------------|
| compile command | `/usr/local/bin/compile` | Yes |
| std.pcm module | `~/.cache/clang20-std/std.pcm` | Yes |
| Build script | `~/cpp/build_simple.sh` | Yes |
| Clang compiler | `/usr/bin/clang++-20` | Yes |
| libc++ headers | `/usr/lib/llvm-20/share/libc++/v1/` | Yes |

## Uninstall (If Needed)

```bash
# Remove compile command
sudo rm /usr/local/bin/compile

# Remove cached std module
rm -rf ~/.cache/clang20-std

# Remove build directory
rm -rf ~/cpp/build

# Remove Clang (optional)
sudo apt remove clang-20 libc++-20-dev
```

## Notes

- The std module builds **once** and is reused for all your projects
- No CMake or Ninja needed after initial setup
- Works from **any directory** on your system
- Survives reboots, logouts, and system updates
- Supports all C++23 standard library features

---

**Last Updated**: April 2026  
**Current Phase**: Active Testing & Development  
**Created by:** Haki
