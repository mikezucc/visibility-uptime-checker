Get simple http request information on endpoint
![screenshoot](https://i.imgur.com/tt9NLdy.png)

Every endpoint queried will be automatically queried periodically and results are available in a persistent cache
![screenshoot](https://i.imgur.com/odl8lgE.png)

# visibility-uptime-checker
Endpoint uptime checker for visibility platform

## TODO
- [x] Add in persistence for when an endpoint was first measured
- [x] Display basic stats for endpoints
- [ ] Interactive JSON viewer to have cleaner browsing of history
- [x] Add timer to continually query provided endpoints (include HttpRequest timeout)

## Setup

1. Install Go
2. Clone
3. `go build`
4. `./visibility-uptime-checker`
5. Visit `localhost:3008/uh_oh.html`
6. Enter in an address *WITH* a protocol

## Deps

1. github.com/googollee/go-socket.io
2. go.uber.org/zap
