# Social cache cleaner

Before `socache`, we had to send to every single social network to clear the cache. Now we need to send a request to the `socache`, specifying a list of social networks in which we need to clear the cache!

[![Build Status](https://drone.io/github.com/iamsalnikov/socache/status.png)](https://drone.io/github.com/iamsalnikov/socache/latest)

## Installation

You can install `socache` by three ways:

1. go get
2. build from source

### Installation by `go get`

```bash
go get -u github.com/iamsalnikov/socache
```

### Build from source

```bash
git clone git@github.com:iamsalnikov/socache.git
cd socache
go build
```

## Usage

1. Start server. If you build socache from source, or install by `go get`:

```
socache [-host] [-port]
```

Default port is `9099`, host - `0.0.0.0`

2. Send query: `http://host:port/?url=<url>&net=<networks>`

- `url` - is url which cache you want to clear
- `networks` - list of networks ids (if you want specify a few networks, use comma as delimiter).

At the current moment we support 2 social networks: vk, facebook.

Example: `http://host:port/?url=https://gmail.com&net=vk,facebook`
