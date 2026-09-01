import { NextResponse } from 'next/server'

export async function GET() {
  const response = NextResponse.redirect(
    new URL('/signin', process.env.NEXT_PUBLIC_APP_URL || 'http://localhost:3000')
  )

  response.cookies.set('access_token', '', { maxAge: 0, path: '/' })
  response.cookies.set('refresh_token', '', { maxAge: 0, path: '/' })

  return response
}
