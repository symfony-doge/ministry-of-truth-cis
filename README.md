# Ministry of Truth CIS

Sanity index provider for vacancy aggregation engine [symfony-doge/veslo](https://github.com/symfony-doge/veslo).

An extremely fast microservice that takes a text and returns sanity index (SI) for it w/ search tags.
Name is a humorous reference to George Orwell's [1984](https://en.wikipedia.org/wiki/Nineteen_Eighty-Four).

### Installation

#### Docker

The preferred way to install is through [docker-compose](https://docs.docker.com/compose).
You need to have a [Docker](https://docs.docker.com/install) daemon at least [17.05.0-ce](https://docs.docker.com/engine/release-notes/#17050-ce) (with build-time `ARG` in `FROM`) to successfully cook container with application.

Run an automated deploy script for local development with Docker.

```
$ bin/deploy_dev.sh
```

#### Manual

You may want to simply clone and build the application in your local Go environment, here is a shortcut:

```
$ git clone git@github.com:symfony-doge/ministry-of-truth-cis.git motcis && cd "$_"
$ cp config/debug.yml.dist config/debug.yml
$ go get -d
$ go mod vendor
$ go build -mod vendor -o build/app
```

In this case you also need to download [Yandex MyStem](https://tech.yandex.ru/mystem) manually in order to use `/index` action.
Executable file must be placed according to path
`analysis.lemmatizator.mystem.executable` from the application's config `config/debug.yml` (`bin/mystem` by default).

Ensure your app instance will work as expected:

```
$ go test ./...
```

### API

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

### Changelog
All notable changes to this project will be documented in [CHANGELOG.md](CHANGELOG.md).