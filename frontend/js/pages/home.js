import { homeTemplate } from "../templates/homeTemplate.js";
import { fetchPosts } from "../api/posts.js";
import { argument } from "../state.js";

export function renderHome() {
    const app = document.getElementById("app");

    app.innerHTML = homeTemplate();
    // const data = await fetchPosts(argument);
    // argument.nextcursor = data.nextcursor
}

