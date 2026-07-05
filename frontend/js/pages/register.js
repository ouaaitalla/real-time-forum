import { registerTemplate } from "../templates/registerTemplate.js";
import { showLoginPage } from "../router.js";

export function renderRegister() {
    const app = document.getElementById("app");

    app.innerHTML = registerTemplate();

    const loginBtn = document.getElementById("go-login");

    loginBtn.addEventListener("click", (e) => {
        e.preventDefault();
        showLoginPage();
    });
}