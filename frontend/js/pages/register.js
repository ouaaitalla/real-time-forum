import { registerTemplate } from "../templates/registerTemplate.js";
import { showLoginPage } from "../router.js";
import { getregisterData } from "../api/auth.js";
import { registerapi } from "../api/auth.js";
import { hideError, showError } from "../utils/render.js";

export function renderRegister() {
    const app = document.getElementById("app");
    app.innerHTML = registerTemplate();
    const form = document.getElementById("register-form");
    form.addEventListener("submit", async (event) => {
        event.preventDefault();
    
        getregisterData();
        const data = await registerapi();
        console.log(data.type)
        if (data.type === "error"){
            showError(data.text, "register-error")
        }else{
            showLoginPage();
            // hideError("register-error");
        }
    });
    const loginBtn = document.getElementById("go-login");

    loginBtn.addEventListener("click", (e) => {
        e.preventDefault();
        showLoginPage();
    });
}