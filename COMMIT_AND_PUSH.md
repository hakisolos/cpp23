# Quick Commands

## Before Pushing

✅ **Check status**:
```bash
cd /home/haki/cpp23
git status
```

✅ **Add all files**:
```bash
git add .
```

✅ **Commit**:
```bash
git commit -m "Initial open source setup: Add documentation, templates, and contribution guidelines"
```

✅ **Push to GitHub**:
```bash
git push origin main
# (replace 'main' with your branch name if different)
```

## After Pushing

📝 **Update these in GitHub web interface**:
1. Go to Settings → Topics → Add: `cpp23` `modules` `clang` `linux`
2. Go to Settings → General → Update repository description
3. Go to Settings → Features → Enable "Discussions"

## Important URLs to Update

Before pushing, find and replace `yourusername` with your actual GitHub username in:
- README.md (lines with github.com links)
- CONTRIBUTING.md (if any references)

Example:
```
FROM: https://github.com/yourusername/cpp23
TO:   https://github.com/YOUR_USERNAME/cpp23
```

## Files Created

| File | Purpose |
|------|---------|
| README.md | Main documentation (updated with badges & status) |
| CONTRIBUTING.md | Contribution guidelines |
| DEVELOPMENT.md | Developer setup guide |
| CODE_OF_CONDUCT.md | Community standards |
| SECURITY.md | Security policy |
| CHANGELOG.md | Version history |
| LICENSE | MIT license |
| .gitignore | Git ignore patterns |
| .github/ISSUE_TEMPLATE/bug_report.md | Bug report template |
| .github/ISSUE_TEMPLATE/feature_request.md | Feature request template |
| .github/PULL_REQUEST_TEMPLATE.md | Pull request template |
| .github/config.yml | GitHub configuration |
| PUSH_CHECKLIST.md | This setup checklist |

---

**You're ready to push!** 🚀
