# The health route

Version: v1
All routes will include `/v1` as a prefix.

## Available Routes

- [`/health`](#route-health)
- [`/health/dashboard`](#route-healthdashboard)

### Route: `/health`

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

### Route: `health/dashboard`

The `/health/dashboard` endpoint is an authenticated route that returns a welcome message and the user ID of the logged-in user.
If the user is not authenticated, a 401 error response will be returned with a message indicating that the user is unauthenticated.
In case of any internal server error, a 500 error response will be returned with a message indicating that an internal server error occurred.

- Method: `GET`
- Response:

```json
{
    "message": "Welcome to the authenticated route",
    "userId": "<user_id of logged in user>"
}
```

- Error Responses:

  - Unauthenticated:
    - status: 401

    ```json
    {
        "message": "unauthenticated user"
    }
    ```

  - Internal Server Error
    - status: 500

    ```json
    {
        "message": "An internal server error occurred"
    }
    ```
