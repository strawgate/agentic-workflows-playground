# Coding Conventions

This repository follows these conventions for agentic workflows.

## Issue Quality Standards

1. **Evidence-based**: Every finding must have concrete evidence (file paths, line numbers, commit SHAs)
2. **Actionable**: Each issue must have specific, actionable next steps
3. **Not noise**: If in doubt, prefer `noop` over filing a low-value issue

## Issue Labels

- `automated` - Created by an agentic workflow
- `bug` - Bug report from bug-hunter
- `documentation` - Docs issue from docs-patrol
- `breaking-change` - Breaking change detected
- `dependencies` - Dependency update needed
- `code-quality` - Code quality issue
- `duplicate` - Duplicate issue
- `stale` - Stale issue

## Issue Expiration

Issues created by workflows expire after 7 days to prevent stale bot-like behavior.

## Noop Policy

**Noop is the expected outcome most days.** This indicates a healthy codebase.

When calling `noop`, include a brief message explaining what was analyzed:
- "No bugs found in last 28 days of commits"
- "All dependencies are current"
- "No documentation drift detected"