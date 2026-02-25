/**
 * ohmytoken — Track your LLM token consumption in one line.
 * 
 * Usage:
 *   import { track, init } from 'ohmytoken'
 *   init('omt_your_key')
 *   await track('gpt-4o', 500, 200)
 */

const DEFAULT_ENDPOINT = 'https://api.ohmytoken.dev/api/v1/ingest'

let _apiKey = ''
let _endpoint = DEFAULT_ENDPOINT

export function init(apiKey, endpoint = DEFAULT_ENDPOINT) {
  _apiKey = apiKey
  _endpoint = endpoint
}

export async function track(model, promptTokens, completionTokens, options = {}) {
  if (!_apiKey) throw new Error('Call init("omt_your_key") first')

  const body = {
    model,
    prompt_tokens: promptTokens,
    completion_tokens: completionTokens,
    ...options,
  }

  try {
    const res = await fetch(_endpoint, {
      method: 'POST',
      headers: { 'X-API-Key': _apiKey, 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    })
    return await res.json()
  } catch {
    return { status: 'error' }
  }
}

export class OhMyToken {
  constructor(apiKey, endpoint = DEFAULT_ENDPOINT) {
    this.apiKey = apiKey
    this.endpoint = endpoint
  }

  async track(model, promptTokens, completionTokens, options = {}) {
    const body = {
      model,
      prompt_tokens: promptTokens,
      completion_tokens: completionTokens,
      ...options,
    }

    try {
      const res = await fetch(this.endpoint, {
        method: 'POST',
        headers: { 'X-API-Key': this.apiKey, 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
      })
      return await res.json()
    } catch {
      return { status: 'error' }
    }
  }
}
