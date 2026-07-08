import { argument } from "../state.js"

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
    return await response.json();
}


export async function createPost(postData) {
    const response = await fetch("http://localhost:8080/creatPost", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(postData)
    });

    const data = await response.json();

    return data;
}

