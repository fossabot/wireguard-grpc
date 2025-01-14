[![Build Status](https://github.com/ezh/wireguard-grpc/actions/workflows/go.yml/badge.svg)](https://github.com/ezh/wireguard-grpc/actions?query=workflow%3Ago)
[![Go Report Card](https://goreportcard.com/badge/github.com/ezh/wireguard-grpc)](https://goreportcard.com/report/github.com/ezh/wireguard-grpc)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fezh%2Fwireguard-grpc.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fezh%2Fwireguard-grpc?ref=badge_shield)

# Generate

`go generate -v`

# Test

## Unit tests

`ginkgo ./...`

## Integration tests

`ginkgo -tags=integration -v test/...`

# Build

`go build cmd/main.go`


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fezh%2Fwireguard-grpc.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fezh%2Fwireguard-grpc?ref=badge_large)