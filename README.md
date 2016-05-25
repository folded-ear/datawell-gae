# datawell-gae

This attempts to replace EventLog with a GAE-hosted Go app. We'll see.

    $ goapp version
    go version go1.6.1 (appengine-1.9.37)

## JS API

The JS API exists under the `/jsapi` path. All requests require authentication,
and will return 401 if the user isn't authenticated. All responses, error or
success, are in JSON format. Currently, a `.json` extension is *not* allowed
on request paths, though I hope to make it optional in the future.

All 401 responses will indicate a login URL based on the request's referrer:

    {
        "errors": [
            {
                "status": "401",
                "title": "Unauthorized",
                "meta": {
                    "login_url": "..."
                }
            }
        ]
    }

All other errors follow the same pattern, and multiple errors may be bundled
together if appropriate.

The following endpoints are available:

### `GET /user`

Returns the logged-in user's information.

Sample response:

    {
        "data": {
            "type": "user"
            "id": "1234567890",
            "attributes": {
                "email": "test@example.com",
                "name": "Test User"
            }
        }
    }

### `GET /events`

Returns logged-in user's events in reverse chronological order. Eventually this
will support paging and filtering (and maybe sorting?), but not right now!

Sample response:

    {
        "data": [
            {
                "type": "event",
                "id": "4567890123",
                "attributes": {
                    "timestamp": "",
                    "tags": "coffee, home, ibuprofen tablet:2"
                }
            },
            {
                "type": "event",
                "id": "7890123456",
                "attributes": {
                    "timestamp": "",
                    "tags": "manhattan, the old gold"
                }
            }
        ]
    }

### `GET /events/:id`

Returns data about a single event of the logged in user.

Sample response:

    {
        "data": {
            "type": "event",
            "id": "7890123456",
            "attributes": {
                "timestamp": "",
                "tags": "manhattan, the old gold"
            }
        }
    }

### `POST /events`

Creates a new event for the logged in user.

Sample request body:

    {
        "data": {
            "type": "event",
            "attributes": {
                "timestamp": "",
                "tags": "soccer, beaverton"
            }
        }
    }

Sample response:

    {
        "data": {
            "id": "987643210",
            "type": "event",
            "attributes": {
                "timestamp": "",
                "tags": "soccer, beaverton"
            }
        }
    }

### `PATCH /events/:id`

Updates an existing event of the logged in user.

Sample request body:

    {
        "data": {
            "id": "987643210",
            "type": "event",
            "attributes": {
                "timestamp": "",
                "tags": "soccer:0.5, beaverton"
            }
        }
    }

Sample response:

*the same as GETting the resource, post update*

### `DELETE /events/:id`

Deletes an existing event of the logged in user.

No request or response body is used.
