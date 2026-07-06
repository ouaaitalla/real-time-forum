export function registerTemplate() {
    return `
        <div class="register-container">
            <h1>Register</h1>

            <form id="register-form">

                <label for="nickname">Nickname</label>
                <input
                    type="text"
                    id="nickname"
                    name="nickname"
                    required
                >

                <label for="firstname">First Name</label>
                <input
                    type="text"
                    id="firstname"
                    name="firstname"
                    required
                >

                <label for="lastname">Last Name</label>
                <input
                    type="text"
                    id="lastname"
                    name="lastname"
                    required
                >

                <label for="age">Age</label>
                <input
                    type="number"
                    id="age"
                    name="age"
                    required
                >

                <label for="gender">Gender</label>
                <select id="gender" name="gender" required>
                    <option value="">Choose...</option>
                    <option value="male">Male</option>
                    <option value="female">Female</option>
                </select>

                <label for="email">Email</label>
                <input
                    type="email"
                    id="email"
                    name="email"
                    required
                >

                <label for="password">Password</label>
                <input
                    type="password"
                    id="password"
                    name="password"
                    required
                >

                <button type="submit">Register</button>

            </form>

            <p>
                Already have an account?
                <a href="#" id="go-login">Login</a>
            </p>
        </div>
    `;
}