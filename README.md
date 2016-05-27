# datawell-gae

This attempts to replace EventLog with a GAE-hosted Go app. We'll see.

    $ goapp version
    go version go1.6.1 (appengine-1.9.37)

## Data Model

The core data model is comprised of events and tags. Tag maps should be
generally presented to the user as a comma-delimited list of names, with
optional colon-suffixed numbers. The number `1` is the default, so should not be
explicitly shown to the user. E.g., `home, ibuprofen tablet:2, coffee` and
`home:1, ibuprofen tablet:2, coffee:1` are equivalent, but the former is
preferred.

### Event (type `event`)

Events do not have a natural key, only a surrogate one.

* `timestamp` - some sort of indication of when the event happened. minute
  granularity, unclear if it's wall clock or an instant.
* `tags` - a `map[string]float` with tag names for keys and arbitrary numbers
  for values.
* `notes` - an arbitrary string

### Tag (type `tag`)

Tags use their name as a natural key, and do not have a surrogate one.

* `name` - the name of the tag, which cannot be numeric
* `tags` - a `map[string]float` with tag names for keys and arbitrary numbers
  for values.
* `notes` - an arbitrary string

## JS API

The JS API exists under the `/jsapi` path. All requests require authentication,
and will return 401 if the user isn't authenticated. All responses, error or
success, are in JSON format. Currently, a `.json` extension is *not* allowed
on request paths, though I hope to make it optional in the future.

The API follows the [http://jsonapi.org/](http://jsonapi.org/format) spec. I think.

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

### `GET /current_user`

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
        },
        "meta": {
            "logout_url": "..."
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
                    "timestamp": "<TBD>",
                    "tags": {
                        "coffee": 1,
                        "home": 1,
                        "ibuprofen tablet": 2
                    },
                    "notes": ""
                }
            },
            ...
        ]
    }

### `POST /events`

Creates a new event for the logged in user.

Sample request body:

    {
        "data": {
            "type": "event",
            "attributes": {
                "timestamp": "<TBD>",
                "tags": {
                    "soccer": 1,
                    "beaverton": 1
                },
                "notes": ""
            }
        }
    }

Sample response:

    {
        "data": {
            "id": "987643210",
            "type": "event",
            "attributes": {
                "timestamp": "<TBD>",
                "tags": {
                    "soccer": 1,
                    "beaverton": 1
                },
                "notes": ""
            }
        }
    }

### `GET /events/:id`

Returns data about a single event of the logged in user.

Sample response:

    {
        "data": {
            "type": "event",
            "id": "7890123456",
            "attributes": {
                "timestamp": "<TBD>",
                "tags": {
                    "manhattan": 1,
                    "the old gold": 1
                },
                "notes": ""
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
                "timestamp": "<TBD>",
                "tags": {
                    "soccer": 0.5,
                    "beaverton": 1
                },
                "notes": ""
            }
        }
    }

Sample response:

*the same as GETting the resource, post update*

### `DELETE /events/:id`

Deletes an existing event of the logged in user.

No request or response body is used.

### `GET /tags`

Returns the logged-in user's available tags. Note that since tags do not have a
surrogate key, the tag name is both the resource's ID and one of its attributes.

The following query parameters are supported:

* `fields[tag]` - a comma-delimited list of fields names to include. If omitted,
  all fields will be included. E.g., `fields[tag]=name`.

Sample response:

    {
        "data": [
            {
                "type": "tag",
                "id": "home",
                "attributes": {
                    "name": "coffee",
                    "tags": {
                        "caffeine": 95,
                        "calories: 22
                    },
                    "notes": "calories based on a couple sugar packets"
                }
            },
            ...
        ]
    }

### `GET /tagsets`

Returns a list of suggested tagsets based on the user's historical data.

The following query parameters are supported:

* `page[size]` - an integer for how many tagsets to return at most. If not
  provided, "a few" will be returned. E.g., `page[size]=7`.
* `filter[tags]` - a comma-delimited list of tags that must be included. If not
  provided, no filtering will be performed. E.g., `filter[tags]=coffee`.

Sample response

    {
        "data": [
            {
                "type": "tagset",
                "id": "...",
                "attributes": {
                    "tags": ["coffee", "desk"],
                }
            },
            {
                "type": "tagset",
                "id": "...",
                "attributes": {
                    "tags": ["home", "coffee", "ibuprofen tablet"],
                }
            },
            ...
        ]
    }

