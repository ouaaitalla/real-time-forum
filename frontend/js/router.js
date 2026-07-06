import { renderLogin } from "./pages/login.js";
import { renderRegister } from "./pages/register.js";
import { renderHome } from "./pages/home.js";
export function showLoginPage() {
    renderLogin();
}

export function showRegisterPage() {
    renderRegister();
}

export function showHomePage() {
    renderHome();
}