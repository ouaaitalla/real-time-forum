export function hideError(id) {

    document.getElementById(id).style.display = "none";
}
export function showError(message, id) {

    const div = document.getElementById(id);

    div.textContent = message;

    div.style.display = "block";
}