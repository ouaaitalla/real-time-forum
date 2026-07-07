
export async function fetchPosts(limit = 20, cursor = null) {
    let url = `http://localhost:8080/posts?limit=${limit}`;

    if (cursor !== null) {
        url += `&cursor=${cursor}`;
    }

    const response = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    });
    data = await response.json();
    let nextcursor = data.nextcursor
    return data
}
