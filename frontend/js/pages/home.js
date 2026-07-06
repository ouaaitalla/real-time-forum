import { homeTemplate } from "../templates/homeTemplate.js";

export function renderHome() {
    const app = document.getElementById("app");

    app.innerHTML = homeTemplate();
}
