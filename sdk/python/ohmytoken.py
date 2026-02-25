"""ohmytoken — Track your LLM token consumption in one line."""

import urllib.request
import json
from typing import Optional

DEFAULT_ENDPOINT = "https://api.ohmytoken.dev/api/v1/ingest"


class OhMyToken:
    """Lightweight client for ohmytoken.dev token tracking."""

    def __init__(self, api_key: str, endpoint: str = DEFAULT_ENDPOINT):
        self.api_key = api_key
        self.endpoint = endpoint

    def track(
        self,
        model: str,
        prompt_tokens: int,
        completion_tokens: int,
        cost: Optional[float] = None,
        label: Optional[str] = None,
    ) -> dict:
        """Report token usage. Call this after each LLM API call."""
        payload = {
            "model": model,
            "prompt_tokens": prompt_tokens,
            "completion_tokens": completion_tokens,
        }
        if cost is not None:
            payload["cost"] = cost
        if label:
            payload["label"] = label

        data = json.dumps(payload).encode()
        req = urllib.request.Request(
            self.endpoint,
            data=data,
            headers={
                "X-API-Key": self.api_key,
                "Content-Type": "application/json",
            },
            method="POST",
        )
        try:
            resp = urllib.request.urlopen(req, timeout=5)
            return json.loads(resp.read())
        except Exception:
            return {"status": "error"}


# Convenience function
_client: Optional[OhMyToken] = None


def init(api_key: str, endpoint: str = DEFAULT_ENDPOINT):
    """Initialize the global ohmytoken client."""
    global _client
    _client = OhMyToken(api_key, endpoint)


def track(model: str, prompt_tokens: int, completion_tokens: int, **kwargs) -> dict:
    """Track token usage using the global client."""
    if _client is None:
        raise RuntimeError("Call ohmytoken.init('omt_your_key') first")
    return _client.track(model, prompt_tokens, completion_tokens, **kwargs)
