# Example Integration

This integration connects to external services.

## Configuration

| Setting | Description | Default |
|---------|-------------|---------|
| api_url | API endpoint URL | https://api.example.com |
| api_key | Authentication key | - |
| timeout | Request timeout | 30s |

## Usage

```bash
elastic-agent enroll --url=https://api.example.com
```

## Features

- Data collection
- Real-time monitoring
- Alerting

## Requirements

- elastic-agent 8.0+
- Network access to api.example.com

## Troubleshooting

See [Troubleshooting Guide](troubleshooting.md)
