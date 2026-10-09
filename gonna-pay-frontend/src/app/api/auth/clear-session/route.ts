import { NextResponse } from 'next/server'

function getAppUrl() {
  if (process.env.NEXT_PUBLIC_APP_URL) return process.env.NEXT_PUBLIC_APP_URL
  if (process.env.VERCEL_URL) return `https://${process.env.VERCEL_URL}`
  return 'http://localhost:3000'
}

export async function GET() {
  const response = NextResponse.redirect(
    new URL('/signin', getAppUrl())
  )

  response.cookies.set('access_token', '', { maxAge: 0, path: '/' })
  response.cookies.set('refresh_token', '', { maxAge: 0, path: '/' })

  return response
}
