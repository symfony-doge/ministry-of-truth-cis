# Ministry of Truth CIS

Sanity index provider for the vacancy aggregation engine [symfony-doge/veslo](https://github.com/symfony-doge/veslo).

An extremely fast microservice that takes a text and returns sanity index (SI) for it w/ search tags.
Name is a humorous reference to George Orwell's [1984](https://en.wikipedia.org/wiki/Nineteen_Eighty-Four).

## Installation

### Docker

The preferred way to install is through [docker-compose](https://docs.docker.com/compose).
You need to have a [Docker](https://docs.docker.com/install) daemon at least [17.05.0-ce](https://docs.docker.com/engine/release-notes/#17050-ce) (with build-time `ARG` in `FROM`) to successfully cook container with application.

Run an automated deploy script for local development with Docker.

```
$ bin/deploy_dev.sh
```

### Manual

You may want to simply clone and build the application in your local Go environment, here is a shortcut:

```
$ git clone git@github.com:symfony-doge/ministry-of-truth-cis.git motcis && cd "$_"
$ cp config/debug.yml.dist config/debug.yml
$ go get -d
$ go mod vendor
$ go build -mod vendor -o build/app
```

Ensure your app instance will work as expected:

```
$ go test ./... -bench . -benchmem
```

## API

Run the application (for manual installation):

```
$ build/app
```

There are some flags available for customization:

| Flag | Default value | Description |
| :--- | :--- | :--- |
| -mode | debug | [Mode](https://github.com/gin-gonic/gin/blob/v1.4.0/mode.go#L15) for Gin Engine |
| -port | 9595 | Port to listen |

You can send requests to the endpoint [http://localhost:9595](http://localhost:9595) (default port).
Microservice allows you to use `POST` or `GET` methods, both json and query parameters are supported.
`config/debug.yml` contains a set of parameters used by application components,
such as log filenames and data directory. By default, the application writes to `var/logs/app/debug.log`
and `var/logs/app/debug-error.log`.

### Indexing a text

**URL**

`/index`

**Methods**

`POST`

**Request example**

```
$ curl --request POST \
       --header "Content-Type: application/json" \
       --data '{"locale":"ru", "context":{"description":"Молодая, динамично развивающаяся компания \"Рога и копыта\" ищет специалиста по X для выполнения Y."}}' \
       "http://localhost:9595/index"
```

**Response example**

```
{
    "status": "OK",
    "errors": [],
    "index": {
        "value": 95.5,
        "tags": {
            "soft": [
                {
                    "name": "young_dynamic",
                    "title": "Динамично развивающаяся компания",
                    "description": "Молодая, динамично развивающаяся компания возьмет в аренду степлер.",
                    "color": "#8db7ad",
                    "image_url": "https://cdn.veslo.it/current/images/tags/young_dynamic.jpg",
                    "group": "soft"
                }
            ]
        }
    }
}
```

### Receiving tag groups

**URL**

`/tag/groups`

**Methods**

`GET`, `POST`

**Request example**

```
$ curl "http://localhost:9595/tag/groups?locale=ru"
```

**Response example**

```
{
    "status": "OK",
    "errors": [],
    "tag_groups": [
        {
            "name": "soft",
            "color": "#3fa0db",
            "description": "Вода, спам и очепятки"
        },
        {
            "name": "hard",
            "color": "#d84343",
            "description": "Технические фейлы"
        },
        {
            "name": "lulz",
            "color": "#d834cd",
            "description": "Юмор и мемы"
        }
    ]
}
```

### Errors

The microservice will throw an error if you are not a gentleman :smirk:.
Stay patient and explore a table below with error codes and their meaning.

| Code | Type | Description |
| :--- | :--- | :--- |
| 1 | `main.handler_not_found` | No handler implemented for requested URI path. |
| 2 | `main.method_not_allowed` | Handler for requested URI path exists, but HTTP method is not allowed. |
| 3 | `request.binder.bad_request` | Whenever a request is not filled with valid parameters. For example, locale is not specified or not supported. |
| 4 | `main.internal_error` | Something terrible happened. Stay safe. |

Example of response with negative status:

```
{
	"status": "FAIL",
	"errors": [
		{
			"code": 3,
			"type": "request.binder.bad_request",
			"description": "Invalid request param 'Locale'."
		}
	]
}
```

### Changelog

All notable changes to this project will be documented in [CHANGELOG.md](CHANGELOG.md).
