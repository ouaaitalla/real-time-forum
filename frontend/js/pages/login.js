import { loginTemplate } from "../templates/loginTemplate.js";
import { showRegisterPage } from "../router.js";
import { getloginData } from "../api/auth.js";
import { loginapi } from "../api/auth.js";
import { showError } from "../utils/render.js";
import { renderHome } from "./home.js";

export function renderLogin() {
    const app = document.getElementById("app");

    app.innerHTML = loginTemplate();
    const form = document.getElementById("login-form");
    form.addEventListener("submit", async ()=>{
        event.preventDefault();

        getloginData();
        const data = await loginapi();
        if (data.type === "error"){
            showError(data.text, "login-error")
        }else {
            renderHome()
        }
        
    })
    const registerBtn = document.getElementById("go-register");

    registerBtn.addEventListener("click", (e) => {
        e.preventDefault();
        showRegisterPage();
    });
}
