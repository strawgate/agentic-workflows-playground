# Agentic Workflows Playground

This repository is a testing ground for agentic workflows powered by [GitHub Agentic Workflows (gh-aw)](https://github.com/github/gh-aw).

## Workflows

| Workflow | Description | Schedule |
|---------|-------------|----------|
| [daily-status.md](.github/workflows/daily-status.md) | Daily repository status reports | Weekdays 2AM UTC |
| [bug-hunter.md](.github/workflows/bug-hunter.md) | Find reproducible bugs | Weekdays 11AM UTC |
| [docs-patrol.md](.github/workflows/docs-patrol.md) | Detect documentation drift | Weekdays 10AM UTC |
| [stale-issues.md](.github/workflows/stale-issues.md) | Triage stale issues | Monday 9AM UTC |
| [duplicate-issue-detector.md](.github/workflows/duplicate-issue-detector.md) | Find duplicate issues | Weekdays 2PM UTC |
| [code-quality-audit.md](.github/workflows/code-quality-audit.md) | Static analysis reports | Sunday 3AM UTC |
| [breaking-change-detector.md](.github/workflows/breaking-change-detector.md) | Detect breaking changes | Weekdays 10AM UTC |
| [dependency-review.md](.github/workflows/dependency-review.md) | Review dependencies | Monday 8AM UTC |

## Setup

These workflows use GitHub App authentication. See the [main agentic-workflows repo](https://github.com/strawgate/agentic-workflows) for setup instructions.

## Development

```bash
# Install dependencies
gh extension install github/gh-aw

# Compile workflows
gh aw compile

# Run locally
gh aw run <workflow-name>

# View logs
gh aw logs <workflow-name>
```

## License

MIT