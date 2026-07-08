export function createPostModal() {
    return `
        <div id="create-post-modal" class="modal hidden">

            <div class="modal-content">

                <button id="close-modal" class="close-btn">
                    &times;
                </button>

                <h2>Create Post</h2>

                <form id="create-post-form">

                    <label>Title</label>
                    <input
                        type="text"
                        id="post-title"
                        maxlength="100"
                        required
                    >

                    <label>Category</label>
                    <select id="post-category" required>
                        <option value="General">General</option>
                        <option value="Programming">Programming</option>
                        <option value="Gaming">Gaming</option>
                        <option value="AI">AI</option>
                    </select>

                    <label>Content</label>
                    <textarea
                        id="post-content"
                        rows="8"
                        maxlength="1000"
                        required
                    ></textarea>

                    <button type="submit" id="submit">
                        Publish
                    </button>

                </form>

            </div>

        </div>
    `;
}


export function postCard(post) {
    return `
        <article class="post-card" data-id="${post.id}">

            <div class="post-header">

                <div class="post-user">
                    <img
                        src="${post.avatar || "/assets/default-avatar.png"}"
                        alt="Avatar"
                        class="post-avatar"
                    >

                    <div>
                        <h3 class="post-author">${post.nickname}</h3>
                        <span class="post-date">${post.created_at}</span>
                    </div>
                </div>

            </div>

            <div class="post-body">

                <h2 class="post-title">${post.title}</h2>

                <p class="post-content">
                    ${post.content}
                </p>

            </div>

            <div class="post-footer">

                <button class="post-action like-btn">
                    👍 ${post.likes}
                </button>

                <button class="post-action dislike-btn">
                    👎 ${post.dislikes}
                </button>

                <button class="post-action comments-btn">
                    💬 ${post.comments}
                </button>

            </div>

        </article>
    `;
}

