import { renderLogin } from "./pages/login.js";
import { renderRegister } from "./pages/register.js";

export function showLoginPage() {
    renderLogin();
}

export function showRegisterPage() {
    renderRegister();
}
