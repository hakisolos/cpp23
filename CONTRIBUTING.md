# Contributing to C++23 Modules Setup

Thank you for your interest in contributing! This project is actively seeking help to improve C++23 modules support on Linux. Here's how you can contribute:

## Ways to Contribute

### 1. **Testing & Bug Reports**
- Test the setup on different Linux distributions
- Report issues with specific Clang/libc++ versions
- Document any edge cases you encounter

### 2. **Improving Documentation**
- Clarify confusing instructions
- Add platform-specific guides
- Document workarounds you've discovered

### 3. **Code & Scripts**
- Improve build scripts
- Add error handling
- Create helper tools
- Refactor existing code

### 4. **Examples & Tests**
- Create example C++23 programs
- Write test cases (see `test/` directory)
- Test different module configurations

### 5. **Platform Support**
- Document setup for different distros (Fedora, Arch, etc.)
- Test on older/newer Ubuntu versions
- Provide ARM/alternative architecture support

## Getting Started

1. **Fork the repository**
2. **Create a branch** for your feature/fix:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Set up your environment** using the [DEVELOPMENT.md](DEVELOPMENT.md) guide
4. **Make your changes** and test thoroughly using the `compile` command:
   ```bash
   # For single-file tests
   compile test_file.cpp test_binary
   ./test_binary
   
   # For multi-file projects
   compile
   ./my_app
   ```
5. **Commit with clear messages**:
   ```bash
   git commit -m "Add: brief description of changes"
   ```
6. **Push to your fork** and **create a Pull Request**

## Pull Request Guidelines

- Use the PR template (provided automatically)
- Link related issues
- Describe what was tested
- Note any breaking changes
- Include examples for new features

## Code Style

- Follow existing code conventions
- Add comments for complex sections
- Test on multiple systems if possible
- Update documentation with your changes

## Issue Guidelines

- Search existing issues first
- Use templates for bug reports and features
- Include system info (distro, Clang version, etc.)
- Provide reproducible examples when possible

## Help Wanted 🙏

This project **especially needs help with:**
- Testing on non-Ubuntu Debian derivatives
- Automating the std.pcm build process
- Cross-platform compatibility
- Performance optimization
- Documentation improvements
- Example programs using C++23 features

## Questions?

Open an issue or start a discussion! We're here to help.

## Recognition

Contributors will be recognized in the README. Thank you for making this project better!
