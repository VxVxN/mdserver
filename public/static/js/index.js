const signInContainer = document.getElementById('signInModal');
const signInModalIndex = new bootstrap.Modal(signInContainer);

if (document.getElementById("signIn") !== null) {
    document.getElementById("signIn").onclick = function () {
        document.getElementById("password").classList.remove("is-invalid");
        document.getElementById("username").classList.remove("is-invalid");
        signInModalIndex.show();
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
        signInModalIndex.hide();
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

document.getElementById("createDirectory").onclick = function () {
    const directoryName = document.getElementById("directoryName").value;
    if (directoryName === '') {
        return;
    }

    const successCallback = function () {
        window.location.reload();
        return false;
    }

    const data = {name:directoryName};
    sendRequest("/create_directory", data, successCallback.bind(this, window));
};

document.getElementById("createDirectoryBtn").onclick = function () {
    document.getElementById("directoryName").value = '';
};

document.getElementById("deleteModalBtn").onclick = function () {
    const successCallback = function () {
        window.location.reload();
        return false;
    }

    if (this.dataset.type === 'directory') {
        const data = {name:this.dataset.name};
        sendRequest("/delete_directory", data, successCallback.bind(this, window));
    } else {
        const data = {dir_name: this.dataset.dirname, file_name: this.dataset.name};
        sendRequest("/delete_post", data, successCallback.bind(this, window));
    }
}

const renameMedals = document.getElementsByName("renameModal");
for(let i=0; i<renameMedals.length; i++){
    renameMedals[i].addEventListener("click", function(e){
        rename(e);
    }, false);
}

const deleteMedals = document.getElementsByName("deleteModal");
for(let i=0; i<deleteMedals.length; i++){
    deleteMedals[i].addEventListener("click", function(e){
        del(e);
    }, false);
}

const createPost = document.getElementsByName("createPost");
for(let i=0; i<createPost.length; i++){
    createPost[i].addEventListener("click", function(e){
        document.getElementById("createPostName").value = '';
        document.getElementById("createPostModalBtn").dataset.dirname = e.target.dataset.dirname;
    }, false);
}

document.getElementById("renameModalBtn").onclick = function () {
    const newName = document.getElementById("renameModalName").value;

    const successCallback = function () {
        window.location.reload();
        return false;
    }

    if (this.dataset.type === 'directory') {
        const data = {old_name:this.dataset.name, new_name:newName};
        sendRequest("/rename_directory", data, successCallback.bind(this, window));
    } else {
        const data = {dir_name: this.dataset.dirname, old_file_name: this.dataset.name, new_file_name: newName};
        sendRequest("/rename_post", data, successCallback.bind(this, window));
    }
}

document.getElementById("createPostModalBtn").onclick = function () {
    const dirName = this.dataset.dirname;
    const fileName = document.getElementById("createPostName").value;
    if (fileName === '') {
        return false;
    }

    const successCallback = function () {
        window.location.reload();
        return false;
    }

    const data = {dir_name: dirName, file_name:fileName};
    sendRequest("/create_post", data, successCallback.bind(this, window));
}

function rename(e) {
    document.getElementById("renameModalName").value = '';
    if (e.target.dataset.type === 'directory') {
        document.getElementById("renameModalLabel").innerText = 'Are you sure you want to rename the directory ' + e.target.dataset.name + '?';
        document.getElementById("renameModalTitle").innerText = "Rename the directory: " +  e.target.dataset.name;
    } else {
        document.getElementById("renameModalLabel").innerText = 'Are you sure you want to rename the file ' + e.target.dataset.name + '?';
        document.getElementById("renameModalTitle").innerText = "Rename the file: " +  e.target.dataset.name;

        document.getElementById("renameModalBtn").dataset.dirname = e.target.dataset.dirname;
    }

    const container = document.getElementById('renameModal');
    const modal = new bootstrap.Modal(container);
    modal.show();

    document.getElementById("renameModalBtn").dataset.name = e.target.dataset.name;
    document.getElementById("renameModalBtn").dataset.type = e.target.dataset.type;
}

function del(e) {
    if (e.target.dataset.type === 'directory') {
        document.getElementById("deleteModalLabel").innerText = 'Are you sure you want to delete the directory ' + e.target.dataset.name + '?';
        document.getElementById("deleteModalTitle").innerText = "Delete the directory:" +  e.target.dataset.name;
    } else {
        document.getElementById("deleteModalLabel").innerText = 'Are you sure you want to delete the file ' + e.target.dataset.name + '?';
        document.getElementById("deleteModalTitle").innerText = "Delete the file: " +  e.target.dataset.name;

        document.getElementById("deleteModalBtn").dataset.dirname = e.target.dataset.dirname;
    }

    const container = document.getElementById('deleteModal');
    const modal = new bootstrap.Modal(container);
    modal.show();

    document.getElementById("deleteModalBtn").dataset.name = e.target.dataset.name;
    document.getElementById("deleteModalBtn").dataset.type = e.target.dataset.type;
}