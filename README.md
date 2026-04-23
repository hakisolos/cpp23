# C++23 Modules Setup on Linux (Clang 20 + libc++)

A complete guide to setting up `import std;` with a permanent `compile` command that works anywhere on your Linux system.

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

# Check if std module exists
if [ ! -f "$STD_PCM" ]; then
    echo "✗ Std module not found at: $STD_PCM"
    echo "Please run the setup script first."
    exit 1
fi

# Check for input file
if [ -z "$1" ]; then
    echo "Usage: compile <file.cpp> [output_name]"
    echo "Examples:"
    echo "  compile main.cpp          # Creates './main'"
    echo "  compile main.cpp myapp    # Creates './myapp'"
    exit 1
fi

INPUT_FILE="$1"

# Check if input exists
if [ ! -f "$INPUT_FILE" ]; then
    echo "✗ Error: File '$INPUT_FILE' not found"
    exit 1
fi

# Determine output name
if [ -n "$2" ] && [ "$2" != "-o" ]; then
    OUTPUT_NAME="$2"
else
    OUTPUT_NAME="${INPUT_FILE%.cpp}"
fi

# Compile
echo "Compiling $INPUT_FILE -> $OUTPUT_NAME"
clang++-20 -std=c++23 -stdlib=libc++ \
    -fmodule-file=std="$STD_PCM" \
    "$INPUT_FILE" -o "$OUTPUT_NAME"

if [ $? -eq 0 ]; then
    echo "✓ Success! Run: ./$OUTPUT_NAME"
else
    echo "✗ Compilation failed"
fi
EOF

sudo chmod +x /usr/local/bin/compile
```

## How to Use (After Setup)

### Basic Usage

```bash
# Compile any C++ file
compile myprogram.cpp

# Run it
./myprogram

# Or compile with specific output name
compile myprogram.cpp myapp
./myapp
```

### Example

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

### Issue: 'compile: command not found'

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

**Created by:** Haki  
**Date:** 2026  
**Distro:** Ubuntu (Clang 20.1.8)
