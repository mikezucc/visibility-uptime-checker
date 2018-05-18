![screenshoot](https://i.imgur.com/tt9NLdy.png)

# visibility-uptime-checker
Endpoint uptime checker for visibility platform

## TODO
- [x] Add in persistence for when an endpoint was first measured
- [x] Display basic stats for endpoints
- [ ] Interactive JSON viewer to have cleaner browsing of history

## Setup

1. Install Go
2. Clone
3. `go build`
4. `./visibility-uptime-checker`
5. Visit `localhost:3003/uh_oh.html`
6. Enter in an address *WITH* a protocol

## Deps

1. github.com/googollee/go-socket.io
2. go.uber.org/zap
