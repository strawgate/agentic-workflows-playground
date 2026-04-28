---
name: "Duplicate Issue Detector"
description: "Find potential duplicate issues and suggest consolidation"
on:
  schedule:
    - cron: '0 14 * * 1-5'  # Weekdays 2PM UTC
  workflow_dispatch:
    inputs:
      lookback_window:
        description: 'Lookback window for recent issues (e.g., "30 days ago")'
        required: false
        default: '30 days ago'
        type: string

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
    title-prefix: "[duplicate-issues] "
    labels: [duplicate, automated]
    max: 1
    expires: 7d

timeout-minutes: 45
---

Find potential duplicate issues and suggest consolidation.

## Data Gathering

1. List issues from the last `${{ inputs.lookback_window }}`
2. Group by:
   - Same error message or symptom
   - Same feature/area
   - Similar reproduction steps

## What to Look For

- Issues describing the same error message
- Issues with similar reproduction steps
- Issues in the same component/area
- Feature requests for the same capability

## Issue Format

```
# Duplicate Issues Report

Found X potential duplicate groups in Y recent issues.

## Group 1: [Description]

| Issue | Title | Similarity |
|-------|-------|------------|
| #123 | "Login broken" | High |
| #456 | "Cannot authenticate" | High |

**Recommendation:** Close #456 as duplicate of #123

## Group 2: [...]

## Action Items

- [ ] Close #456 as duplicate of #123
- [ ] Link related issues together
- [ ] Add `duplicate` label to closed issues
```

Call `create_issue` with the report, or `noop` if no duplicates found.