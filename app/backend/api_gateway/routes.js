const ROUTES = [
    {
        url: '/free',
        auth: false,
        creditCheck: true,
        rateLimit: {
            windowMs: 0.5*60*1000,
            max: 2
        },
        proxy: {
            target: "https://www.google.com",
            changeOrigin: true,
            pathRewrite: {
                [`^/free`]: '',
            },
        }
    },
    {
        url: '/premium',
        auth: true,
        creditCheck: true,
        proxy: {
            target: "https://www.google.com",
            changeOrigin: true,
            pathRewrite: {
                [`^/premium`]: '',
            },
        }
    }
]

exports.ROUTES = ROUTES;