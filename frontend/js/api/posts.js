import { argument } from "../state";

export async function fetchPosts() {
    let url = `http://localhost:8080/posts?limit=${argument.limit}`;

    if (argument.cursor !== null) {
        url += `&cursor=${argument.cursor}`;
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
