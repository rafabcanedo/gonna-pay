import { NextRequest, NextResponse } from 'next/server'
import { API_BASE_URL } from '../../../../constants'

export async function DELETE(
  request: NextRequest,
  { params }: { params: { id: string; contactId: string } }
) {
  try {
    const { id, contactId } = params
    const accessToken = request.cookies.get('access_token')?.value

    const res = await fetch(`${API_BASE_URL}/group/${id}/member/${contactId}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        ...(accessToken ? { Cookie: `access_token=${accessToken}` } : {}),
      },
    })

    if (!res.ok) {
      const data = await res.json()
      return NextResponse.json(
        { error: data.error || 'Failed to remove member' },
        { status: res.status }
      )
    }

    return NextResponse.json({ message: 'member removed successfully' })
  } catch (error) {
    console.error('[DELETE /api/groups/:id/member/:contactId] Error:', error)
    return NextResponse.json({ error: 'Internal server error' }, { status: 500 })
  }
}

export async function POST(
  request: NextRequest,
  { params }: { params: { id: string; contactId: string } }
) {
  try {
    const { id, contactId } = params
    const accessToken = request.cookies.get('access_token')?.value

    const res = await fetch(`${API_BASE_URL}/group/${id}/member/${contactId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        ...(accessToken ? { Cookie: `access_token=${accessToken}` } : {}),
      },
    })

    if (!res.ok) {
      const data = await res.json()
      return NextResponse.json(
        { error: data.error || 'Failed to add member' },
        { status: res.status }
      )
    }

    return NextResponse.json({ message: 'member added successfully' })
  } catch (error) {
    console.error('[POST /api/groups/:id/member/:contactId] Error:', error)
    return NextResponse.json({ error: 'Internal server error' }, { status: 500 })
  }
}
