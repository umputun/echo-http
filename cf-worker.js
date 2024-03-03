// This is a simple Cloudflare Worker that echoes back the request details as a JSON response.
// Replicates the same behavior provided in go code

addEventListener("fetch", (event) => {
    event.respondWith(handleRequest(event.request));
});

async function handleRequest(request) {
    const url = new URL(request.url);
    // Filter headers to exclude Cloudflare-specific and other non-essential headers
    const filteredHeaders = {};
    for (const [key, value] of request.headers) {
        if (!key.toLowerCase().startsWith('cf-') && key.toLowerCase() !== 'x-real-ip') {
            filteredHeaders[key] = value;
        }
    }

    // Use CF-Connecting-IP as the remote_addr if available, otherwise leave as an empty string
    const remoteAddr = request.headers.get('cf-connecting-ip') || '';

    const echo = {
        message: "echo", // Adjust as needed or use environment variables for dynamic responses
        request: `${request.method} ${url.pathname}`,
        host: request.headers.get("host") || "", // Directly use the 'host' header from the request
        headers: filteredHeaders,
        remote_addr: remoteAddr, // Set to CF-Connecting-IP or leave empty if not available
    };

    return new Response(JSON.stringify(echo), {
        headers: { "Content-Type": "application/json" },
    });
}
