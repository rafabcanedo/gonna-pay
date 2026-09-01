import { API_BASE_URL } from '../constants'
import { ApiError, ForbiddenError, InternalServerError, NetworkError, NotFoundError, ServiceUnavailableError, UnauthorizedError, ValidationError } from '../errors/api.error'

interface NextInternalError extends Error {
  digest: string
}

function isNextInternalError(error: unknown): error is NextInternalError {
  return (
    error instanceof Error &&
    'digest' in error &&
    typeof (error as NextInternalError).digest === 'string' &&
    (error as NextInternalError).digest.startsWith('NEXT_')
  )
}

function getBaseUrl() {
  if (typeof window !== 'undefined') return ''

  const envUrl = process.env.NEXT_PUBLIC_APP_URL || process.env.APP_URL
  if (envUrl) return envUrl.replace(/\/$/, '')

  return API_BASE_URL
}

async function getServerCookieHeader(): Promise<string> {
  try {
    const { cookies } = await import('next/headers')
    const cookieStore = await cookies()
    return cookieStore.getAll().map(c => `${c.name}=${c.value}`).join('; ')
  } catch {
    return ''
  }
}

async function tryRefresh(): Promise<boolean> {
  try {
    const res = await fetch('/api/auth/refresh', {
      method: 'POST',
      credentials: 'include',
    })
    return res.ok
  } catch {
    return false
  }
}

export async function apiCall<T>(
  endpoint: string,
  options: RequestInit = {},
  retry = true
): Promise<T> {
  try {
    const isServer = typeof window === 'undefined'
    const baseUrl = getBaseUrl()
    const url = `${baseUrl}/api${endpoint}`

    const serverCookies = isServer ? await getServerCookieHeader() : ''

    const res = await fetch(url, {
      headers: {
        'Content-Type': 'application/json',
        ...(serverCookies ? { Cookie: serverCookies } : {}),
        ...options.headers,
      },
      ...options,
      cache: 'no-store',
      credentials: 'include',
    })

    const text = await res.text()
    const data = text ? JSON.parse(text) : null

    if (!res.ok) {
      switch (res.status) {
        case 400:
          if (data?.fields) {
            throw new ValidationError(data.fields)
          }
          throw new ApiError(400, data?.error || 'Bad request', data)

        case 401:
          if (retry) {
            if (typeof window !== 'undefined') {
              const refreshed = await tryRefresh()
              if (refreshed) return apiCall<T>(endpoint, options, false)
              window.location.href = '/signin'
            } else {
              const { redirect } = await import('next/navigation')
              redirect('/api/auth/clear-session')
            }
          }
          throw new UnauthorizedError()

        case 403:
          throw new ForbiddenError()

        case 404:
          throw new NotFoundError(data?.resource || 'Resource')

        case 500:
          throw new InternalServerError()

        case 503:
          throw new ServiceUnavailableError()

        default:
          throw new ApiError(res.status, data?.error || 'Unknown error', data)
      }
    }

    return data as T
  } catch (error) {
    if (error instanceof ApiError) throw error
    if (isNextInternalError(error)) throw error
    if (error instanceof TypeError) {
      throw new NetworkError()
    }
    throw new ApiError(500, 'Unknown error', error)
  }
}