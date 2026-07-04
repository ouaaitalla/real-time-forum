async function renderFeed() {
  const root = viewRoot();
  root.innerHTML = `
    <h1 class="view-title">Feed</h1>
    <section class="feed-layout">
      <div>
        <form id="post-form" class="composer form-grid">
          <div class="two-col form-grid">
            <label>Title
              <input name="title" minlength="3" required>
            </label>
            <label>Category
              <input name="category" required>
            </label>
          </div>
          <label>Content
            <textarea name="content" rows="4" minlength="5" required></textarea>
          </label>
          <p class="error-text" id="post-error"></p>
          <button type="submit">Publish</button>
        </form>
        <div id="post-list" class="post-list"><div class="loading">Loading posts...</div></div>
      </div>
      <aside id="comments-panel" class="comments">
        <div class="empty">Select a post</div>
      </aside>
    </section>
  `;
  document.querySelector("#post-form").addEventListener("submit", createPost);
  await loadPosts();
}

async function loadPosts() {
  const list = document.querySelector("#post-list");
  try {
    const data = await App.api("/api/posts");
    App.state.posts = data.posts || [];
    list.innerHTML = App.state.posts.length ? App.state.posts.map(renderPost).join("") : `<div class="empty">No posts yet</div>`;
    list.querySelectorAll("[data-post-id]").forEach((node) => {
      node.addEventListener("click", () => openPost(node.dataset.postId));
    });
  } catch (err) {
    list.innerHTML = `<div class="error-text">${escapeHTML(err.message)}</div>`;
  }
}

function renderPost(post) {
  return `
    <article class="post" data-post-id="${escapeHTML(post.id)}">
      <h2 class="post-title">${escapeHTML(post.title)}</h2>
      <p>${escapeHTML(post.content)}</p>
      <div class="meta">
        <span class="tag">${escapeHTML(post.category)}</span>
        by ${escapeHTML(post.author_nickname)} · ${formatDate(post.created_at)} · ${post.comment_count || 0} comments
      </div>
    </article>
  `;
}

async function createPost(event) {
  event.preventDefault();
  const form = event.currentTarget;
  const button = form.querySelector("button");
  const error = document.querySelector("#post-error");
  button.disabled = true;
  error.textContent = "";
  try {
    await App.api("/api/posts", { method: "POST", body: JSON.stringify(Object.fromEntries(new FormData(form).entries())) });
    form.reset();
    await loadPosts();
    showToast("Post published");
  } catch (err) {
    error.textContent = err.message;
  } finally {
    button.disabled = false;
  }
}
