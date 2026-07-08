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
                        <option value="">Choose category</option>
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

                    <button type="submit">
                        Publish
                    </button>

                </form>

            </div>

        </div>
    `;
}