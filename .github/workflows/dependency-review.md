---
name: "Dependency Review"
description: "Review dependencies for outdated packages and vulnerabilities"
on:
  schedule:
    - cron: '0 8 * * 1'  # Monday 8AM UTC
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
    allowed: [list_issues, create_issue]

safe-outputs:
  github-app:
    client-id: ${{ vars.APP_ID }}
    private-key: ${{ secrets.APP_PRIVATE_KEY }}
  create-issue:
    title-prefix: "[dependencies] "
    labels: [dependencies, automated]
    max: 1
    expires: 7d

timeout-minutes: 30
---

Review dependencies for outdated packages and known vulnerabilities.

## Analysis Steps

1. Identify dependency files: package.json, requirements.txt, go.mod, Cargo.toml, Gemfile, etc.
2. Check for outdated dependencies (compare to latest versions)
3. Check for known security vulnerabilities
4. Identify unused dependencies

## What to Look For

- Dependencies with security advisories
- Major version updates available
- Dependencies not imported/used
- Very old dependencies (no maintenance)
- Dependencies with breaking changes available

## What to Skip

- Dev-only dependencies (if build succeeds)
- Lock files (only review source manifests)
- Private packages (can't check versions)

## Issue Format

```
# Dependency Review

## Security Updates Needed

| Package | Current | Severity | Advisory |
|---------|---------|----------|----------|
| lodash | 4.17.20 | High | CVE-2021-xxxx |

## Major Updates Available

| Package | Current | Latest | Breaking? |
|---------|---------|--------|-----------|
| express | 4.17.1 | 5.0.0 | Yes |

## Unused Dependencies

| Package | Reason to Remove |
|---------|------------------|
| moment | Not imported in codebase |

## Action Items

- [ ] Update [package] to [version]
- [ ] Remove unused [package]
- [ ] Review breaking changes in [major update]
```

Call `create_issue` or `noop` if all dependencies are current and secure.