export function renderNavbar() {
    return `
        <header class="navbar">

            <div class="navbar-logo">
                <a href="#/">Forum</a>
            </div>
            <nav class="navbar-links">
                <a href="#/">Home</a>
                <a href="#/messages">Messages</a>
                <button id="logout-btn">Logout</button>
            </nav>

        </header>
    `;
}

export function homeTemplate() {
    return `
        ${renderNavbar()}

        <main class="home-container">

            <section class="posts-container">
                <!-- Posts will be rendered here -->
            </section>
            <div id="posts-container"></div>

            <div id="loading" style="display: none;">
                Loading...
            </div>

        </main>
    `;
}