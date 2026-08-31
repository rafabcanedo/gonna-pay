import { NextRequest, NextResponse } from "next/server";
import { API_BASE_URL } from "../../constants";

export async function POST(request: NextRequest) {
  try {
    const body = await request.json();

    const res = await fetch(`${API_BASE_URL}/auth/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });

    const data = await res.json();

    if (!res.ok) {
      return NextResponse.json(data, { status: res.status });
    }

    const response = NextResponse.json(data, { status: 200 });

    const cookies = res.headers.getSetCookie();
    cookies.forEach((cookie) => {
      const modified = cookie.startsWith("refresh_token=")
        ? cookie.replace("Path=/auth", "Path=/api/auth")
        : cookie;
      response.headers.append("Set-Cookie", modified);
    });

    return response;
  } catch (error) {
    return NextResponse.json(
      { error: "Internal Server Error na API Route" },
      { status: 500 }
    );
  }
}
