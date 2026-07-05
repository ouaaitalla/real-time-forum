import { loginTemplate } from "../templates/loginTemplate.js";
import { showRegisterPage } from "../router.js";

export function renderLogin() {
    const app = document.getElementById("app");

    app.innerHTML = loginTemplate();

    const registerBtn = document.getElementById("go-register");

    registerBtn.addEventListener("click", (e) => {
        e.preventDefault();
        showRegisterPage();
    });
}