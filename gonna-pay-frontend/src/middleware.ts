import { MiddlewareConfig, NextRequest, NextResponse } from "next/server";

// Work with public routes, definition routes which the user can access without authentication
const publicRoutes = ["/signin", "/signup", "/forgot-your-password"] as const;

// Route where user is redirected when user is not authenticated
const REDIRECT_WHEN_NOT_AUTHENTICATED_ROUTE = "/signin";
// Route where user is redirected when user is authenticated
const REDIRECT_WHEN_AUTHENTICATED_ROUTE = "/dashboard";

// Here we search if the access_token cookie exists, but we do not validate
// the JWT signature. The real validation happens in the Fastify backend.
export function middleware(request: NextRequest) {
  const path = request.nextUrl.pathname;
  const authToken = request.cookies.get("access_token");

  // Check if the route is public
  const isPublicRoute = publicRoutes.some((route) => path.startsWith(route));

  // User NOT authenticated
  if (!authToken) {
    // If trying to access private route, redirect to signin
    if (!isPublicRoute) {
      const redirectUrl = request.nextUrl.clone();
      redirectUrl.pathname = REDIRECT_WHEN_NOT_AUTHENTICATED_ROUTE;
      return NextResponse.redirect(redirectUrl);
    }
    // If in public route, allow
    return NextResponse.next();
  }

  // User AUTHENTICATED
  if (authToken) {
    // If trying to access public route, redirect to dashboard
    if (isPublicRoute) {
      const redirectUrl = request.nextUrl.clone();
      redirectUrl.pathname = REDIRECT_WHEN_AUTHENTICATED_ROUTE;
      return NextResponse.redirect(redirectUrl);
    }
    // If in private route, allow
    return NextResponse.next();
  }

  return NextResponse.next();
}

// Next middleware config default
export const config: MiddlewareConfig = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico, sitemap.xml, robots.txt (metadata files)
     */
    "/((?!api|_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt).*)",
  ],
};
