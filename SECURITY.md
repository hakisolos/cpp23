# Security Policy

## Reporting Security Issues

If you discover a security vulnerability, **please email** the project maintainers directly instead of using the issue tracker.

Include:
- Description of the vulnerability
- Affected versions/configurations
- Potential impact
- Suggested fix (if you have one)

We will acknowledge receipt within 48 hours and provide a status update within 5 days.

## Scope

Security vulnerabilities that may affect this project:

- **Setup scripts** - Command injection, arbitrary code execution
- **Build process** - Unsafe shell operations, dependency issues
- **Documentation** - Misleading installation instructions

## Security Best Practices

As a user of this project:

1. Always run official installation scripts
2. Verify package integrity when installing Clang/libc++
3. Keep your Ubuntu/Debian system updated
4. Report suspicious behavior or unexpected build failures

## Supported Versions

| Version | Supported |
|---------|-----------|
| Latest | ✅ Yes |
| Prior releases | ⚠️ Case-by-case |

## Disclaimer

This project is provided "as-is". Users are responsible for:
- Security of their build environment
- Verification of downloaded packages
- Testing in their own systems before deployment
