export function commentCard(comment) {
    return `
        <article class="comment-card">

            <div class="comment-header">
                <strong>${comment.nickname}</strong>
                <span>${comment.created_at}</span>
            </div>

            <p class="comment-content">
                ${comment.content}
            </p>

        </article>
    `;
}
