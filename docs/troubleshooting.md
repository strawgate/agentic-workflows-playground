# Troubleshooting

Use these checks if the example integration does not enroll or send data as expected.

## Verify configuration

- Confirm `api_url` points to the expected endpoint.
- Confirm `api_key` is set when the external service requires authentication.
- Increase `timeout` if the endpoint is slow to respond.

## Verify connectivity

- Ensure the host running `elastic-agent` can reach `api.example.com`.
- Check local firewall or proxy settings if enrollment cannot connect.

## Verify the agent version

This example requires `elastic-agent` 8.0 or later.
