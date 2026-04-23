# Development Guide

## Development Setup

### Prerequisites
- Ubuntu 22.04+ (or Debian-based distro)
- Clang 20 with libc++ modules support
- Basic knowledge of C++23 and modules

### Initial Setup

```bash
# Clone and enter the repository
git clone https://github.com/yourusername/cpp23.git
cd cpp23

# Run the installation guide from README.md
# This will set up Clang 20 and the std module cache
```

## Project Structure

```
cpp23/
├── README.md              # Main documentation
├── DEVELOPMENT.md         # This file
├── CONTRIBUTING.md        # Contribution guidelines
├── test/
│   ├── main.cpp          # Test source file
│   └── main              # Compiled test binary
├── build_simple.sh        # Build script example
└── .github/
    ├── ISSUE_TEMPLATE/    # Issue templates
    └── PULL_REQUEST_TEMPLATE.md
```

## Using the `compile` Command

After installation, use the enhanced `compile` command for both single-file and multi-file projects.

### Single File Compilation

```bash
# For simple single-file programs
compile myprogram.cpp          # Creates './myprogram'
compile myprogram.cpp myapp    # Creates './myapp'
```

### Multi-File Project Compilation

For projects with multiple C++ files and/or C++ modules (`.cppm` or `.ixx` files):

```bash
# Automatically detects all modules and files, builds them all
compile

# Output binary is created as 'my_app'
./my_app
```

**What the enhanced compile script does for multi-file projects:**
1. Detects all module interface files (`.cppm`, `.ixx`)
2. Compiles each module to a precompiled module file (`.pcm`)
3. Compiles all `.cpp` files to object files
4. Links all objects together into the final executable
5. Provides color-coded output for easy debugging

### Running Tests

```bash
# For single test file (simple test)
cd test
compile main.cpp main
./main

# For multi-file tests (if test directory has multiple files)
cd test
compile
./my_app
```

## Testing Checklist

When making changes, verify:

- [ ] Code compiles with `clang++-20` using the `compile` command
- [ ] Test program runs without errors
- [ ] Works for both single-file and multi-file projects (if applicable)
- [ ] Color output displays correctly
- [ ] Module detection works (if adding `.cppm` files)
- [ ] Documentation is updated
- [ ] Changes work on Ubuntu 22.04+ (primary target)
- [ ] README examples still work

## Creating New Tests

Add test files to the `test/` directory following this pattern:

```cpp
import std;

// Your test code here
int main() {
    std::println("Test passed!");
    return 0;
}
```

## Building Documentation

The documentation is in Markdown format. To preview locally:

```bash
# Using any Markdown viewer, or
cat README.md
```

## Common Issues During Development

### "Std module not found"
- Verify `~/.cache/clang20-std/std.pcm` exists
- Run the setup from README.md again

### Compilation fails with "unknown flag"
- Ensure Clang 20 is installed: `clang++-20 --version`
- Check libc++ headers: `ls /usr/lib/llvm-20/include/c++/v1/`

### Test binary doesn't run
- Verify libc++ runtime is installed
- Check with: `ldd test/main`

### "Compile: command not found" or module detection issues
- Ensure `/usr/local/bin/compile` is executable: `ls -la /usr/local/bin/compile`
- Reinstall the compile command from the README.md setup steps
- For multi-file projects, ensure all modules use `.cppm` or `.ixx` extensions

## Submitting Changes

1. Create a feature branch: `git checkout -b feature/my-feature`
2. Make your changes and test thoroughly
3. Update documentation as needed
4. Create a Pull Request with:
   - Clear description of changes
   - List of tested platforms
   - Any breaking changes noted
   - References to related issues

## Performance Considerations

- Module precompilation (`std.pcm`) should be done once and cached
- Subsequent compilations use the cached module for speed
- Monitor compilation times in benchmarks

## Compatibility Matrix

This project targets:

| Component | Version | Status |
|-----------|---------|--------|
| Clang | 20+ | Primary |
| libc++ | 20+ | Primary |
| Ubuntu | 22.04+ | Primary |
| GCC | N/A | Not supported (no modules support) |

## Questions?

- Check existing issues and discussions
- Open a new issue with the `question` label
- Reach out to maintainers

## Additional Resources

- [C++23 Standard Modules](https://en.cppreference.com/w/cpp/language/modules)
- [Clang Documentation](https://clang.llvm.org/)
- [libc++ Documentation](https://libcxx.llvm.org/)
