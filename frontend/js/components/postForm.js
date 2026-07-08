export function postModal() {
    return `
        <div id="post-modal" class="modal hidden">

            <div class="modal-content post-modal-content">

                <button id="close-post-modal" class="close-btn">
                    &times;
                </button>

                <div id="modal-post"></div>

                <hr>

                <h3>Comments</h3>

                <form id="comment-form">

                    <textarea
                        id="comment-content"
                        placeholder="Write a comment..."
                        required
                    ></textarea>

                    <button type="submit">
                        Comment
                    </button>

                </form>

                <div id="comments-container"></div>

            </div>

        </div>
    `;
}
