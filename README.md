# philips-wiz-bulb-cli

Cobra CLI for [philips-wiz-bulb-core](https://github.com/deepanshutr/philips-wiz-bulb-core).

```
philips-wiz-bulb list
philips-wiz-bulb state    [target]
philips-wiz-bulb on       [target]
philips-wiz-bulb off      [target]
philips-wiz-bulb bri      <10-100>     [target]
philips-wiz-bulb temp     <2200-6500>  [target]
philips-wiz-bulb color    <r> <g> <b>  [target]
philips-wiz-bulb scene    <name|id>    [target]
philips-wiz-bulb discover
philips-wiz-bulb name     <mac|ip>     <new-name>
```

`target` accepts a MAC (with or without colons), IPv4, friendly name, or
`all`. Reads `PHILIPS_WIZ_BULB_CORE_URL` (default `http://127.0.0.1:8766`).

See [`docs/superpowers/specs/2026-05-14-philips-wiz-bulb-stack-design.md`](docs/superpowers/specs/2026-05-14-philips-wiz-bulb-stack-design.md)
for the full design.

## Sibling repos

- [philips-wiz-bulb-core](https://github.com/deepanshutr/philips-wiz-bulb-core) — Python daemon
- [philips-wiz-bulb-mcp](https://github.com/deepanshutr/philips-wiz-bulb-mcp) — MCP stdio server for Claude Code
