---
name: "Stale Issues Investigator"
description: "Triage stale issues and recommend closure or action"
on:
  schedule:
    - cron: '0 9 * * 1'  # Monday 9AM UTC
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
    title-prefix: "[stale-issues] "
    labels: [stale, automated]
    max: 1
    expires: 3d

timeout-minutes: 45
---

Investigate stale issues and recommend which should be closed or updated.

## Definition of Stale

An issue is stale if:
- No comments for 60+ days
- No PR linked
- Not labeled with: `enhancement`, `feature-request`, `roadmap`, `pinned`

## Data Gathering

1. List all open issues
2. Filter for potentially stale ones (older than 60 days)
3. For each candidate, check:
   - Last comment date
   - Linked PRs
   - Current labels
   - Whether it's tracked elsewhere

## Issue Format

```
# Stale Issues Triage Report

Analyzed X open issues. Found Y candidates for closure.

## Recommend Closing (no activity, not tracked)

| Issue | Last Activity | Reason |
|-------|---------------|--------|
| #123 | 2024-01-15 | No activity, enhancement not implemented |

## Recommend Updating (needs maintainer response)

| Issue | Last Activity | Reason |
|-------|---------------|--------|
| #456 | 2024-02-20 | Needs clarification from author |

## Already Tracked Elsewhere

| Issue | Duplicate Of |
|-------|--------------|
| #789 | #123 (closed) |

## Action Items

- [ ] Close X issues
- [ ] Add waiting-for-feedback label to Y issues
- [ ] Link duplicates to canonical issues
```

Call `create_issue` with the triage report, or `noop` if no stale issues found.