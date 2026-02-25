<p align="center">
  <a href="https://ohmytoken.dev">
    <img src="https://ohmytoken.dev/preview.png" alt="ohmytoken" width="600" />
  </a>
</p>

<h1 align="center">ohmytoken</h1>

<p align="center">
  <strong>Your AI spending, visualized as pixel art. In real-time.</strong>
</p>

<p align="center">
  <a href="https://ohmytoken.dev">Website</a> &bull;
  <a href="https://github.com/0x5446/ohmytoken-oss/discussions">Community</a> &bull;
  <a href="https://clawhub.ai/0x5446/ohmytoken-tracker">ClawHub</a> &bull;
  <a href="#license">MIT License</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT License" />
  <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs Welcome" />
  <img src="https://img.shields.io/badge/python-3.8+-blue.svg" alt="Python 3.8+" />
  <img src="https://img.shields.io/badge/node-18+-green.svg" alt="Node 18+" />
  <img src="https://img.shields.io/badge/go-1.21+-00ADD8.svg" alt="Go 1.21+" />
</p>

---

Track your LLM token consumption across Claude, GPT, Gemini, DeepSeek, and any model. Every token becomes a colorful pixel on a retro bead board. Input tokens are light, output tokens are dark — because output costs 3-5x more and you should see that.

## Quick Start

### Python

```python
import ohmytoken

ohmytoken.init("omt_your_key")

# After each LLM call:
ohmytoken.track("gpt-4o", prompt_tokens=500, completion_tokens=200)
```

### JavaScript / TypeScript

```javascript
import { init, track } from 'ohmytoken'

init('omt_your_key')

await track('claude-sonnet-4-20250514', 1000, 400)
```

### Go

```go
client := ohmytoken.New("omt_your_key")
client.Track("gemini-2.0-flash", 800, 300)
```

### curl

```bash
curl -X POST https://api.ohmytoken.dev/api/v1/ingest \
  -H "X-API-Key: omt_your_key" \
  -H "Content-Type: application/json" \
  -d '{"model":"gpt-4o","prompt_tokens":500,"completion_tokens":200}'
```

### Claude Code (zero code changes)

```bash
export CLAUDE_CODE_ENABLE_TELEMETRY=1
export OTEL_EXPORTER_OTLP_ENDPOINT=https://api.ohmytoken.dev/api/v1/traces
export OTEL_EXPORTER_OTLP_HEADERS="x-api-key=omt_your_key"
```

### OpenClaw

```bash
openclaw skill install @0x5446/ohmytoken-tracker
```

Get your free API key at **[ohmytoken.dev](https://ohmytoken.dev)**.

## What You Get

| Feature | Description |
|---------|-------------|
| **Pixel Bead Board** | Tokens as colored blocks, NES retro style |
| **Contribution Heatmap** | GitHub-style green grid of daily usage |
| **Input/Output Split** | Light = cheap input, dark = expensive output |
| **7 Board Shapes** | Square, cat, heart, star, circle, diamond, mushroom |
| **7 Fill Patterns** | Spiral, rain, snake, center-out, random, diagonal, sequential |
| **5 Time Views** | Today, 7 days, month, year, all-time |
| **Multi-Key Tracking** | Separate stats per API key |
| **10 Achievements** | Night Owl, Millionaire, Board Clear... |
| **Leaderboards** | Total, efficiency, diversity rankings |
| **Share Cards** | Pixel art image + QR code + profile URL |
| **Embeddable Badges** | SVG badge for your README |

## API Reference

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/v1/ingest` | POST | Simple JSON token ingestion |
| `/api/v1/traces` | POST | OTLP trace ingestion |

**Auth:** `X-API-Key: omt_your_key` header.

**Request body:**

```json
{
  "model": "gpt-4o",
  "prompt_tokens": 500,
  "completion_tokens": 200,
  "cost": 0.003,
  "label": "my-bot"
}
```

Required: `model`, `prompt_tokens`, `completion_tokens`. Optional: `cost`, `label`.

## Privacy

We collect **only**: model name, token counts (input + output), timestamp.

We **never** collect: prompts, responses, code, files, API keys, IP addresses, or any content whatsoever.

Like a step counter — it counts steps, not conversations.

## Repo Structure

```
ohmytoken-oss/
├── sdk/
│   ├── python/       # Zero-dependency single-file client
│   ├── javascript/   # ESM module with init/track API
│   └── go/           # Typed client with timeout
├── openclaw-plugin/  # ClawHub skill for OpenClaw
├── docs/             # API documentation
└── LICENSE           # MIT
```

## Contributing

PRs are welcome! Whether it's a new SDK language, a bug fix, or documentation improvement.

1. Fork the repo
2. Create your branch (`git checkout -b feat/my-feature`)
3. Commit (`git commit -m 'feat: add something cool'`)
4. Push (`git push origin feat/my-feature`)
5. Open a Pull Request

## License

[MIT](LICENSE) — use it however you want.

---

<p align="center">
  Made with pixels and love &bull; <a href="https://ohmytoken.dev">ohmytoken.dev</a>
</p>
