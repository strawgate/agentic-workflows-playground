---
name: "Docs Patrol"
description: "Detect code changes that require documentation updates"
on:
  schedule:
    - cron: '0 10 * * 1-5'  # Weekdays 10AM UTC
  workflow_dispatch:
    inputs:
      lookback_window:
        description: 'Git lookback window (e.g., "7 days ago")'
        required: false
        default: '7 days ago'
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
    allowed: [list_issues, create_issue, search_code]

safe-outputs:
  github-app:
    client-id: ${{ vars.APP_ID }}
    private-key: ${{ secrets.APP_PRIVATE_KEY }}
  create-issue:
    title-prefix: "[docs-patrol] "
    labels: [documentation, automated]
    max: 1
    close-older-key: "[docs-patrol]"
    close-older-issues: true
    expires: 7d

timeout-minutes: 60
---

Detect documentation drift — code changes that require corresponding documentation updates.

**Noop is the expected outcome most days.** Only file when documentation is concretely wrong.

## Data Gathering

1. Run `git log --since="${{ inputs.lookback_window }}" --oneline --stat` to get recent commits
2. If no commits in the window, call `noop`
3. Scan for documentation files: README.md, CONTRIBUTING.md, docs/, etc.

## What to Look For

1. **Public API changes** — new/removed functions, CLI flags, config options
2. **Behavioral changes** — altered defaults, changed error messages
3. **New features** — anything users need to know about
4. **Dependency/tooling changes** — version bumps, new dependencies
5. **Structural changes** — moved/renamed files referenced in docs

## How to Analyze

For each impactful change:
- Read the full diff
- Read current documentation
- Check if docs were updated in the same/subsequent commit
- Check for open issues/PRs tracking the update

## What to Skip

- Purely internal refactors
- Changes where docs were already updated
- Test-only changes
- Minor changes where existing docs are still correct

## Issue Format

```
Recent code changes require documentation updates:

## Changes Requiring Docs Updates

### 1. [Description]
**Commit:** [SHA]
**What changed:** [Description]
**Docs impact:** [Which file needs what changes]

### 2. [...]

## Suggested Actions
- [ ] [Specific actionable checkbox for each]
```

Call `create_issue` or `noop` if nothing concrete found.