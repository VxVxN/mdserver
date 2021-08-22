const container = document.getElementById('signInModal');
const signInModal = new bootstrap.Modal(container);

if (document.getElementById("signIn") !== null) {
    document.getElementById("signIn").onclick = function () {
        document.getElementById("password").classList.remove("is-invalid");
        document.getElementById("username").classList.remove("is-invalid");
        signInModal.show();
    };
}

document.getElementById("signInBtn").onclick = function () {
    const usernameInput = document.getElementById("username");
    const passwordInput = document.getElementById("password");

    const errorDiv = document.getElementById("errorDiv");

    const username = usernameInput.value;
    const password = passwordInput.value;

    if (username === "") {
        errorDiv.innerText = "The username cannot be empty";
        usernameInput.classList.add("is-invalid");
        return;
    }

    if (password === "") {
        errorDiv.innerText = "The password cannot be empty";
        passwordInput.classList.add("is-invalid");
        return;
    }

    const successCallback = function () {
        signInModal.hide();
        window.location.href = "/";
        return false;
    }

    const errorCallback = function () {
        errorDiv.innerText = "Incorrect username or password";
        return false;
    }

    const data = {username: username, password: password};
    sendRequest("/sign_in", data, successCallback.bind(this, window), errorCallback);
};
