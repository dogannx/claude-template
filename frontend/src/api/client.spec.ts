import { afterEach, describe, expect, it, vi } from 'vitest'
import { api } from './client'

function mockFetch(status: number, body: unknown) {
  const fetchMock = vi.fn().mockResolvedValue(
    new Response(JSON.stringify(body), { status }),
  )
  vi.stubGlobal('fetch', fetchMock)
  return fetchMock
}

afterEach(() => {
  vi.unstubAllGlobals()
})

describe('api', () => {
  it('returns parsed JSON on success', async () => {
    const fetchMock = mockFetch(200, { status: 'ok' })
    await expect(api<{ status: string }>('/api/v1/health')).resolves.toEqual({ status: 'ok' })
    expect(fetchMock.mock.calls[0][0]).toMatch(/\/api\/v1\/health$/)
  })

  it('throws the backend error envelope message', async () => {
    mockFetch(404, { error: { code: 404, message: 'not found' } })
    await expect(api('/api/v1/nope')).rejects.toThrow('not found')
  })

  it('falls back to the HTTP status when the body is not JSON', async () => {
    vi.stubGlobal('fetch', vi.fn().mockResolvedValue(new Response('oops', { status: 500 })))
    await expect(api('/x')).rejects.toThrow('HTTP 500')
  })
})
