
export async function fetchPosts() {
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
    return data
}
