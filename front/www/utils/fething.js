export async function post(url, request) {
    let response = await fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json;charset=utf-8",
        },
        body: JSON.stringify(request),
    });

    if (!response.ok) {
        return JSON.stringify({ error: "Failed to fetch data" });
    }

    return response.json();
}

export async function get(url, request) {
    let response = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json;charset=utf-8",
        },
        body: JSON.stringify(request),
    });

    if (!response.ok) {
        return JSON.stringify({ error: "Failed to fetch data" });
    }

    return response.json();
}
