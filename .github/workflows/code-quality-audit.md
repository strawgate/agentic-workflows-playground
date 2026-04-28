---
name: "Code Quality Audit"
description: "Run static analysis and report code quality findings"
on:
  schedule:
    - cron: '0 3 * * 0'  # Sunday 3AM UTC
  workflow_dispatch:

permissions:
  contents: read
  issues: read

engine:
  id: claude
  model: anthropic/claude-3-5-sonnet-20241022
  env:
    ANTHROPIC_BASE_URL: https://api.minimax.io/anthropic

tools:
  github:
    mode: remote
    allowed: [list_issues, create_issue, list_pull_requests]

safe-outputs:
  github-app:
    client-id: ${{ vars.APP_ID }}
    private-key: ${{ secrets.APP_PRIVATE_KEY }}
  create-issue:
    title-prefix: "[code-quality] "
    labels: [code-quality, automated]
    max: 1
    expires: 7d

timeout-minutes: 60
---

Run static analysis and report code quality findings.

## Analysis Steps

1. Run `git log --since="7 days ago" --oneline --stat` to see recent changes
2. Analyze the codebase for:
   - Complex functions (high cyclomatic complexity)
   - Long functions (100+ lines)
   - Large files (500+ lines)
   - TODO/FIXME comments
   - Code smells (duplication, dead code)
   - Potential bugs (null dereferences, resource leaks)

## What to Look For

- Functions that are too complex to test
- Files that are too large to maintain
- TODOs that represent known debt
- Patterns that violate language best practices
- Security concerns (hardcoded secrets, SQL injection patterns)

## What to Skip

- Generated code
- Test files (unless critical issues)
- Vendor/third-party code
- Build artifacts

## Issue Format

```
# Code Quality Report

Analyzed codebase for quality issues.

## High Priority Findings

### 1. [File: line] - [Issue]
**Severity:** High
**Description:** [...]
**Recommendation:** [...]

## Medium Priority

## Technical Debt

| Item | Location | Type |
|------|----------|------|
| TODO | src/main.go:45 | Unimplemented feature |

## Action Items

- [ ] Refactor [complex function]
- [ ] Break up [large file]
- [ ] Address TODO in [file]
```

Call `create_issue` with the report, or `noop` if no significant issues found.