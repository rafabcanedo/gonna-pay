import { NextRequest, NextResponse } from 'next/server'
import { API_BASE_URL } from '../../constants'

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ id: string }> }
) {
  try {
    const { id } = await params
    const accessToken = request.cookies.get('access_token')?.value

    const res = await fetch(`${API_BASE_URL}/group/${id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        ...(accessToken ? { Cookie: `access_token=${accessToken}` } : {}),
      },
      cache: 'no-store',
    })

    const data = await res.json()

    if (!res.ok) {
      return NextResponse.json(
        { error: data.error || 'Failed to fetch group' },
        { status: res.status }
      )
    }

    return NextResponse.json(data)
  } catch (error) {
    console.error('[GET /api/groups/:id] Error:', error)
    return NextResponse.json({ error: 'Internal server error' }, { status: 500 })
  }
}

export async function PATCH(
  request: NextRequest,
  { params }: { params: Promise<{ id: string }> }
) {
  try {
    const { id } = await params
    const body = await request.json()
    const accessToken = request.cookies.get('access_token')?.value

    const res = await fetch(`${API_BASE_URL}/group/${id}`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
        ...(accessToken ? { Cookie: `access_token=${accessToken}` } : {}),
      },
      body: JSON.stringify(body),
    })

    const data = await res.json()

    if (!res.ok) {
      return NextResponse.json(
        { error: data.error || 'Failed to update group' },
        { status: res.status }
      )
    }

    return NextResponse.json(data)
  } catch (error) {
    console.error('[PATCH /api/groups/:id] Error:', error)
    return NextResponse.json({ error: 'Internal server error' }, { status: 500 })
  }
}

export async function DELETE(
  request: NextRequest,
  { params }: { params: Promise<{ id: string }> }
) {
  try {
    const { id } = await params
    const accessToken = request.cookies.get('access_token')?.value

    const res = await fetch(`${API_BASE_URL}/group/${id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        ...(accessToken ? { Cookie: `access_token=${accessToken}` } : {}),
      },
    })

    if (!res.ok) {
      const data = await res.json()
      return NextResponse.json(
        { error: data.error || 'Failed to delete group' },
        { status: res.status }
      )
    }

    return NextResponse.json({ message: 'group deleted successfully' })
  } catch (error) {
    console.error('[DELETE /api/groups/:id] Error:', error)
    return NextResponse.json({ error: 'Internal server error' }, { status: 500 })
  }
}
