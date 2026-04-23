# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- ✨ Initial open-source release
- 📝 Comprehensive documentation
- 🧪 Basic test suite
- 📋 Contributing guidelines
- 🛡️ Security policy
- 📜 Code of Conduct
- 🔄 Issue and PR templates
- 📚 Development guide
- 🚀 **Enhanced `compile` command** - Now supports both single-file and multi-file projects
  - Automatic module detection and compilation (`.cppm`, `.ixx` files)
  - Multi-file project support with automatic dependency management
  - Color-coded output for better debugging
  - Seamless transition between simple and complex projects

### In Progress
- 🔄 Cross-distro testing (Ubuntu, Debian, Fedora, Arch)
- 🔄 Automated build tests
- 🔄 More C++23 feature examples
- 🔄 Performance benchmarks
- 🔄 Extended troubleshooting guide

### Known Issues
- Module caching requires manual setup (`~/.cache/clang20-std/std.pcm`)
- Only tested on Ubuntu 22.04+ with Clang 20
- libc++ modules still evolving; some features may be unstable

### Planned
- [ ] Automated installation script
- [ ] CMake integration example
- [ ] Prebuilt module binaries
- [ ] GitHub Actions CI/CD
- [ ] ARM/alternative architecture support
- [ ] GCC modules support (when available)
- [ ] Template library examples
- [ ] Performance optimization guide

## [0.1.0] - 2026-04-23

### Initial Release
- Basic setup guide for C++23 modules on Linux
- Clang 20 + libc++ configuration
- Build script and compile command
- Simple test program
- Documentation and examples

---

## Help Wanted 🙏

We're actively looking for contributors to help with:
- Testing on different distributions
- Improving documentation clarity
- Creating more examples
- Automating build processes
- Cross-platform support

If you'd like to help, see [CONTRIBUTING.md](CONTRIBUTING.md)!
