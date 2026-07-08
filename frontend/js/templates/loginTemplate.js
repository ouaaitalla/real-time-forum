export function loginTemplate() {
    return `
        <div class="login-container">
            <h1>Login</h1>

            <form id="login-form">

                <label for="identifier">Email or Nickname</label>
                <input
                    type="text"
                    id="identifier"
                    name="identifier"
                    placeholder="Enter your email or nickname"
                    required
                >

                <label for="password">Password</label>
                <input
                    type="password"
                    id="password"
                    name="password"
                    placeholder="Enter your password"
                    required
                >
                 <div id="login-error" class="error-message"></div>
                <button type="submit">Login</button>

            </form>

            <p>
                Don't have an account?
                <a href="#" id="go-register">Register</a>
            </p>
        </div>
    `;
}

