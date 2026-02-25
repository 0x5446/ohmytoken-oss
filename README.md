# ohmytoken

**Your AI spending, visualized as pixel art. In real-time.**

[![ohmytoken](https://ohmytoken.dev/preview.png)](https://ohmytoken.dev)

Track your LLM token consumption across Claude, GPT, Gemini, DeepSeek, and any model. See it as a colorful pixel bead board or GitHub-style contribution heatmap.

**[ohmytoken.dev](https://ohmytoken.dev)** | **[Community](https://github.com/0x5446/ohmytoken/discussions)**

---

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

// After each LLM call:
await track('claude-sonnet-4-20250514', 1000, 400)
```

### Go

```go
import "github.com/0x5446/ohmytoken-oss/sdk/go"

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

### Claude Code (zero code)

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

- Pixel bead board with 7 shapes and 7 fill patterns
- GitHub-style contribution heatmap
- Input/output token split (light = cheap input, dark = expensive output)
- Time views: today, 7 days, month, year, all-time
- Multi-key tracking (separate stats per API key)
- Achievements, leaderboards, gallery, wrapped reports
- Share cards with QR codes
- Embeddable SVG badges

## API

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/v1/ingest` | POST | Simple JSON ingestion |
| `/api/v1/traces` | POST | OTLP trace ingestion |

Auth: `X-API-Key: omt_your_key` header.

Request body:
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

We collect: **model name**, **token counts**, **timestamp**. That's it.

We never collect: prompts, responses, code, files, API keys, or any content.

## Repo Structure

```
ohmytoken-oss/
├── sdk/
│   ├── python/       # pip install ohmytoken
│   ├── javascript/   # npm install ohmytoken
│   └── go/           # go get github.com/0x5446/ohmytoken-oss/sdk/go
├── openclaw-plugin/  # ClawHub skill for OpenClaw
└── docs/             # API documentation
```

## License

MIT

---

Made with pixels and love. [ohmytoken.dev](https://ohmytoken.dev)
