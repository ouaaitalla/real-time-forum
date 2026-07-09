import { homeTemplate } from "../templates/homeTemplate.js";
import { fetchPosts } from "../api/posts.js";
import { argument } from "../state.js";
import { postCard } from "../templates/postTemplate.js";
import {createPostModal} from "../templates/postTemplate.js"
import { createPost } from "../api/posts.js";
import { fetchPost } from "../api/posts.js";
import { commentCard } from "../components/commentCard.js";


export async function renderHome() {
    console.log('hello')
    const app = document.getElementById("app");

    app.innerHTML = homeTemplate();
    const postsContainer = document.getElementById("posts-container");
    try {
    const data = await fetchPosts(argument);
    argument.nextcursor = 22
    data.forEach((post) => {
        postsContainer.innerHTML += postCard(post);
    });
    } catch (error) {
        console.error("hello:", error);
        postsContainer.innerHTML = "<p>Error loading posts</p>";
    }

    const modal = document.getElementById("create-post-modal");

    document
        .getElementById("create-post-btn")
        .addEventListener("click", () => {
            console.log("button clicked");
            modal.classList.remove("hidden");
        });

    document
        .getElementById("close-modal")
        .addEventListener("click", () => {
            modal.classList.add("hidden");
        });

        const createPostForm = document.getElementById("create-post-form");
        createPostForm.addEventListener("submit", async (e) => {
            e.preventDefault();
            console.log("clicked")
            const postData = {
                title: document.getElementById("post-title").value,
                category: document.getElementById("post-category").value,
                content: document.getElementById("post-content").value
        };

        try {
            const data = await createPost(postData);
            console.log(data);
        } catch (err) {
            console.error(err);
        }

    });
    document.addEventListener("click", async (e) => {
        const card = e.target.closest(".post-card");
        if (!card) return;
        const id = card.dataset.id;
        const data = await fetchPost(id);
        document.getElementById("modal-post").innerHTML =
            postCard(data.post);
        document.getElementById("comments-container").innerHTML =
            data.comments
                .map(comment => commentCard(comment))
                .join("");
        document
            .getElementById("post-modal")
            .classList.remove("hidden");
    });
    document
        .getElementById("close-post-modal")
        .addEventListener("click", () => {

            document
                .getElementById("post-modal")
                .classList.add("hidden");
        });

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
