# CLAUDE.md — philips-wiz-bulb-cli

Cobra CLI talking HTTP to philips-wiz-bulb-core.

## Conventions

- Go 1.23
- `unset GOROOT; export GOPROXY=https://proxy.golang.org,direct` before any `go` command
  (the system `~/.profile` exports a stale GOROOT)
- Per-repo: `git config user.email 52166434+deepanshutr@users.noreply.github.com`
- Run `go vet`, `go test ./... -count=1`, `staticcheck ./...` before commit

## Layout

```
cmd/philips-wiz-bulb/main.go
internal/cli/root.go        # cobra
internal/core/client.go     # HTTP client + tests
internal/config/config.go   # env (PHILIPS_WIZ_BULB_CORE_URL)
```

## Install

```bash
go build -o ~/.local/bin/philips-wiz-bulb ./cmd/philips-wiz-bulb
```
