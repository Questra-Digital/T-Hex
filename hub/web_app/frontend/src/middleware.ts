import { NextRequest, NextResponse } from 'next/server';
import { httpUtils, HttpRequestOptions } from '@/services/httpUtils';

export async function middleware(request: NextRequest) {
    const sessionId = request.cookies.get('sessionId')?.value;

    if (!sessionId) {
        return NextResponse.redirect(new URL('/login', request.url));
    }

    const options: HttpRequestOptions = {
        headers: {
            'Authorization': `Bearer ${sessionId}`
        }
    };

    const response = await httpUtils.post(process.env.NEXT_PUBLIC_API_GATEWAY_URL + '/verify-session', options);

    if (!response.success) {
        return NextResponse.redirect(new URL('/login', request.url));
    }

    return NextResponse.next();
}

export const config = {
    matcher: [
        // Apply only to specific protected routes
        '/getting_started/:path*',
        '/pipelines/:path*',
        '/test_sessions/:path*',
    ],
};

