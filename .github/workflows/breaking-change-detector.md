---
name: "Breaking Change Detector"
description: "Detect breaking changes in recent commits"
on:
  schedule:
    - cron: '0 10 * * 1-5'  # Weekdays 10AM UTC
  workflow_dispatch:
    inputs:
      base_branch:
        description: 'Base branch to compare against'
        required: false
        default: 'main'
        type: string

permissions:
  contents: read
  issues: read
  pull-requests: read

engine:
  id: claude
  model: anthropic/claude-3-5-sonnet-20241022
  env:
    ANTHROPIC_BASE_URL: https://api.minimax.io/anthropic

tools:
  github:
    mode: remote
    allowed: [list_issues, create_issue, list_pull_requests, get_pull_request]

safe-outputs:
  github-app:
    client-id: ${{ vars.APP_ID }}
    private-key: ${{ secrets.APP_PRIVATE_KEY }}
  create-issue:
    title-prefix: "[breaking-change] "
    labels: [breaking-change, automated]
    max: 1
    expires: 7d

timeout-minutes: 45
---

Detect breaking changes in recent commits.

## What is a Breaking Change

- Removed or renamed public APIs
- Changed function signatures
- Removed CLI flags
- Changed config file formats
- Changed behavior that affects backward compatibility
- Version bumps that drop support

## Data Gathering

1. Run `git log --since="7 days ago" --oneline --stat` for recent commits
2. For each commit, examine the diff for breaking patterns
3. Check changelog/version history if present

## What to Look For

- Removed `export`, `pub`, `public` declarations
- Changed type signatures
- Removed CLI flags/commands
- Config file schema changes
- Error message changes
- Removed environment variables

## What to Skip

- Internal refactors with no API change
- Documentation-only changes
- Test file changes
- Build script changes that don't affect users

## Issue Format

```
# Breaking Change Report

Found X potential breaking changes in recent commits.

## Breaking Changes

### 1. [Description]
**Commit:** [SHA]
**Type:** [API/CLI/Config/Behavior]
**Impact:** [Who/what is affected]
**Recommendation:** [How to migrate]

## Migration Steps

- [ ] Update [affected component]
- [ ] Add deprecation notice
- [ ] Update migration guide
```

Call `create_issue` or `noop` if no breaking changes detected.