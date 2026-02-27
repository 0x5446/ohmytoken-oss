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
  <a href="https://clawhub.ai/0x5446/ohmytoken">ClawHub</a> &bull;
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

Track your LLM token consumption across Claude, GPT, Gemini, DeepSeek, and any model. Every token becomes a colorful pixel on a retro bead board. Input tokens are light, output tokens are dark — so you can see at a glance where your tokens go.

## Quick Start

### Step 1: Get your API key (10 seconds)

1. Go to **[ohmytoken.dev](https://ohmytoken.dev)**
2. Click **GOOGLE** or **GITHUB** to sign in
3. Your API key appears on the welcome screen — copy it (looks like `omt_a1b2c3...`)

### Step 2: Track your tokens

Pick your language/tool below and add 2-3 lines of code. That's it.

---

### Claude Code — 0 code changes, just env vars

```bash
# Add to ~/.zshrc (or ~/.bashrc):
export CLAUDE_CODE_ENABLE_TELEMETRY=1
export OTEL_METRICS_EXPORTER=otlp
export OTEL_EXPORTER_OTLP_ENDPOINT=https://api.ohmytoken.dev/api
export OTEL_EXPORTER_OTLP_PROTOCOL=http/json
export OTEL_EXPORTER_OTLP_HEADERS="x-api-key=omt_YOUR_KEY"
# Then: source ~/.zshrc && claude
```

Only aggregated token counts are sent — no prompts, no code, no content.

### OpenClaw — 1 command

```bash
openclaw skill install @0x5446/ohmytoken
```

Then add your key to `openclaw.json`:
```json
{ "skills": { "ohmytoken": { "config": { "api_key": "omt_YOUR_KEY" } } } }
```

### curl — any language, any agent

```bash
curl -X POST https://api.ohmytoken.dev/api/v1/ingest \
  -H "X-API-Key: omt_YOUR_KEY" \
  -H "Content-Type: application/json" \
  -d '{"model":"gpt-4o","prompt_tokens":500,"completion_tokens":200}'
```

Call this after each LLM API call. Works from any language.

### Python SDK

```bash
# Copy the single-file SDK into your project:
curl -o ohmytoken.py https://raw.githubusercontent.com/0x5446/ohmytoken-oss/main/sdk/python/ohmytoken.py
```

```python
import ohmytoken

ohmytoken.init("omt_YOUR_KEY")
ohmytoken.track("gpt-4o", prompt_tokens=500, completion_tokens=200)
```

### JavaScript SDK

```bash
# Copy the SDK into your project:
curl -o ohmytoken.js https://raw.githubusercontent.com/0x5446/ohmytoken-oss/main/sdk/javascript/ohmytoken.js
```

```javascript
import { init, track } from './ohmytoken.js'
init('omt_YOUR_KEY')
await track('gpt-4o', 500, 200)
```

### Go SDK

```go
// Copy sdk/go/ohmytoken.go into your project
client := ohmytoken.New("omt_YOUR_KEY")
client.Track("gpt-4o", 500, 200)
```

### Step 3: Watch your board

Open **[ohmytoken.dev](https://ohmytoken.dev)** — beads appear in real-time as you use AI.

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

We collect **only**: model name, prompt token count, completion token count. That's 3 fields.

We **never** collect: prompts, responses, code, files, API keys, or any content whatsoever. Your IP is visible during the HTTP connection but is not logged or stored.

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
