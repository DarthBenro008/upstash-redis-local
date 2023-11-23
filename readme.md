

![release](https://img.shields.io/github/v/release/DarthBenro008/upstash-redis-local)
[![GitHub License](https://img.shields.io/github/license/DarthBenro008/upstash-redis-local)](https://github.com/DarthBenro008/upstash-redis-local/blob/master/LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/DarthBenro008/upstash-redis-local/issues/new/choose)

# upstash-redis-local

> A local webserver for testing and development using [`@upstash/redis`](https://github.com/upstash/upstash-redis)

## 🤔 upstash-redis-local?

The `upstash-redis-local` command is used to initiate a local web server, serving a REST API compatible with [Upstash REST Api](https://docs.upstash.com/redis/features/restapi), in front of a real Redis database instance. This setup is intended for testing purposes.

This project is inspired by [upstashdis](https://github.com/mna/upstashdis) but uses `fasthttp` instead for better performance and provides some better logging along with a docker image, and build release!

## 💻 Usage

```bash
upstash-redis-local v1.0
A local server that mimics upstash-redis for local testing purposes!

       * Connect to any local redis of your choice for testing
       * Comlpetely mimics the upstash REST API https://docs.upstash.com/redis/features/restapi

Developed by Hemanth Krishna (https://github.com/DarthBenro008)

USAGE:
        upstash-redis-local
        upstash-redis-local --token upstash --addr :8000 --redis :6379

ARGUMENTS:
        --token TOKEN   The API token to accept as authorised (default: upstash)
        --addr  ADDR    Address for the server to listen on (default: :8000)
        --redis ADDR    Address to your redids server (default: :6379)
        --help          Prints this message

```

### Arguments

1. **`--token <STRING>`**: The API token to accept as authorised (default: upstash)
2. **`--addr <ADDR>`**:Address for the server to listen on (default: :8000)
3. **`--redis <ADDR>`**: Address to your redis server (default: :6379)
4. **`--help`**: Prints a help message


## ⬇ Installation

### Using Homebrew

```bash
brew tap DarthBenro008/upstash-redis-local
brew install upstash-redis-local
```

### Using Docker

```bash
docker run -rm -p 8000:8000 darthbenro008/upstash-redis-local:latest
```


### Building from source

1. You can simple do a `docker build . -t upstash-redis-local:development` after cloning this repo
2. You can run a `make build` and find the binary inside the `bin/` folder


### Manual Installation

You can also download the binary and install it manually.

- Go to [releases page](https://github.com/DarthBenro008/upstash-redis-local/releases) and grab the latest release of upstash-redis-local.
- Download the latest release of upstash-redis-local specific to your OS.
- If you are on Linux/MacOS, make sure you move the binary to somewhere in your `$PATH` (e.g. `/usr/local/bin`).


## 🤝 Contributions

- Feel Free to Open a PR/Issue for any feature or bug(s).
- Make sure you follow the [community guidelines](https://docs.github.com/en/github/site-policy/github-community-guidelines).
- Feel free to open an issue to ask a question/discuss anything about upstash-redis-local.
- Have a feature request? Open an Issue!


## ⚖ License

Copyright 2023 Hemanth Krishna

Licensed under MIT License : https://opensource.org/licenses/MIT

<p align="center">Made with ❤ , and a single cup of kofi</p>