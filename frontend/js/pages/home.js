import { homeTemplate } from "../templates/homeTemplate.js";
import { fetchPosts } from "../api/posts.js";
import { argument } from "../state.js";
import { postCard } from "../components/postCard.js";


export async function renderHome() {
    const app = document.getElementById("app");

    app.innerHTML = homeTemplate();
    const postsContainer = document.getElementById("posts-container");
    window.addEventListener("scroll", handleScroll);

    try {
    const data = await fetchPosts(argument);
    argument.nextcursor = data.nextcursor;
    data.posts.forEach((post) => {
        postsContainer.innerHTML += postCard(post);
    });
    } catch (error) {
        console.error("Error fetching posts:", error);
        postsContainer.innerHTML = "<p>Error loading posts. Please try again later.</p>";
    }

}

async function loadPosts() {

    if (argument.isLoading || !argument.hasMore) return;

    argument.isLoading = true;

    const postsContainer = document.querySelector(".posts-container");

    try {

        const data = await fetchPosts(argument);

        postsContainer.insertAdjacentHTML(
            "beforeend",
            data.posts.map(post => postCard(post)).join("")
        );

        nextCursor = argument.nextCursor;
        hasMore = argument.hasMore;

    } catch (err) {
        console.error(err);
    }

    argument.isLoading = false;
}

function handleScroll() {

    const scrollPosition = window.innerHeight + window.scrollY;

    const bottom = document.body.offsetHeight - 300;

    if (scrollPosition >= bottom) {
        loadPosts();
    }

}
