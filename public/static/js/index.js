const container = document.getElementById('signInModal');
const signInModal = new bootstrap.Modal(container);

document.getElementById("signIn").onclick = function () {
    document.getElementById("password").classList.remove("is-invalid");
    signInModal.show();
};

document.getElementById("logOut").onclick = function () {
    const successCallback = function () {
        window.location.href = "/";
        return false;
    }

    sendRequest("/log_out", {}, successCallback);
};

document.getElementById("signInBtn").onclick = function () {
    const passwordInput = document.getElementById("password");
    const errorDiv = document.getElementById("errorDiv");

    const password = passwordInput.value;
    document.getElementById("createPostName").value = '';

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
        errorDiv.innerText = "Incorrect password";
        passwordInput.classList.add("is-invalid");
        return false;
    }

    const data = {password:password};
    sendRequest("/sign_in", data, successCallback.bind(this, window), errorCallback);
};

document.getElementById("createDirectory").onclick = function () {
    const directoryName = document.getElementById("directoryName").value;
    if (directoryName === '') {
        return;
    }
    document.getElementById("directoryName").value = "";

    const successCallback = function () {
        window.location.reload();
        return false;
    }

    const data = {name:directoryName};
    sendRequest("/create_directory", data, successCallback.bind(this, window));
};

document.getElementById("cancelSaveDirectory").onclick = function () {
    document.getElementById("directoryName").value = "";
};

document.getElementById("cancelPostModalBtn").onclick = function () {
    document.getElementById("createPostName").value = "";
};

document.getElementById("cancelRenameModalBtn").onclick = function () {
    document.getElementById("renameModalName").value = "";
};

document.addEventListener('click',function(e){
    if (e.target.classList.contains('deleteModal')) {
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
    if (e.target.classList.contains('createPost')) {
        document.getElementById("createPostModalBtn").dataset.dirname = e.target.dataset.dirname;
    }
    if (e.target.classList.contains('renameModal')) {
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
})

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

document.getElementById("renameModalBtn").onclick = function () {
    const newName = document.getElementById("renameModalName").value;

    const successCallback = function () {
        document.getElementById("renameModalName").value = "";
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
    document.getElementById("createPostName").value = '';

    const successCallback = function () {
        window.location.reload();
        return false;
    }

    const data = {dir_name: dirName, file_name:fileName};
    sendRequest("/create_post", data, successCallback.bind(this, window));
}