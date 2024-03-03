// This is a simple Cloudflare Worker that echoes back the request details as a JSON response.
// Replicates the same behavior provided in go code

addEventListener("fetch", (event) => {
    event.respondWith(handleRequest(event.request));
});

async function handleRequest(request) {
    const url = new URL(request.url);
    const headers = {};
    for (const [key, value] of request.headers) {
        if (!key.startsWith("cf-") && key !== "x-real-ip") {
            const normalizedKey = normalizeHeaderCase(key);
            headers[normalizedKey] = value;
        }
    }

    const echo = {
        message: "echo",
        request: `${request.method} ${url.pathname}${url.search}`,
        host: url.host,
        headers: headers,
        remote_addr: request.headers.get("x-real-ip") || "",
    };

    return new Response(JSON.stringify(echo), {
        headers: { "Content-Type": "application/json" },
    });
}

function normalizeHeaderCase(header) {
    return header.split('-')
        .map(part => part.charAt(0).toUpperCase() + part.slice(1).toLowerCase())
        .join('-');
}
