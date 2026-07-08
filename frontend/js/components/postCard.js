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