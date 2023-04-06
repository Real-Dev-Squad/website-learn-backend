# The health route

Version: v1
All routes will include `/v1` as a prefix.

Route: `/health`

- Method: `GET`
- Response:

```json
{
    "message": "server is up and running"
}
```

- Error Response:

```text
404 Page Not Found
```

The `/health` endpoint can be used to check if the server is up and running.
The error response will be returned if the requested URL doesn't exist on the server.
