# Test Data for Agentic Workflows

This repo is a test target for agentic workflows.

## Creating Test Issues

To test workflows like stale-issues and duplicate-issue-detector, manually create issues in this repo via the GitHub UI:

### Stale Issues (60+ days old)
Create issues and don't update them for 60+ days to test stale issue detection.

### Duplicate Issues
Create issues with similar titles/descriptions to test duplicate detection.

### Bug Reports
Create bug reports with clear reproduction steps.

### Documentation Issues
Create issues about outdated docs.

## Workflows Running Against This Repo

The main workflows are in `strawgate/agentic-workflows` and target this repo via:

```yaml
# In daily-status.md
repos: 'strawgate/agentic-workflows-playground'
```

## Test Checklist

- [ ] Daily Status Report - daily at 2AM UTC
- [ ] Bug Hunter - weekdays 11AM UTC
- [ ] Docs Patrol - weekdays 10AM UTC
- [ ] Stale Issues - Monday 9AM UTC
- [ ] Duplicate Issues - weekdays 2PM UTC
