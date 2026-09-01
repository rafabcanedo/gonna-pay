import { NextRequest, NextResponse } from 'next/server'
import { API_BASE_URL } from '../../constants'

export async function POST(request: NextRequest) {
  try {
    const refreshToken = request.cookies.get('refresh_token')?.value

    const res = await fetch(`${API_BASE_URL}/auth/refresh`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...(refreshToken ? { Cookie: `refresh_token=${refreshToken}` } : {}),
      },
    })

    const data = await res.json()

    if (!res.ok) {
      return NextResponse.json(data, { status: res.status })
    }

    const response = NextResponse.json(data, { status: 200 })

    const cookies = res.headers.getSetCookie()
    cookies.forEach((cookie) => {
      const modified = cookie.startsWith('refresh_token=')
        ? cookie.replace('Path=/auth', 'Path=/api/auth')
        : cookie
      response.headers.append('Set-Cookie', modified)
    })

    return response
  } catch (error) {
    console.error('[POST /api/auth/refresh] Error:', error)
    return NextResponse.json({ error: 'Internal server error' }, { status: 500 })
  }
}
