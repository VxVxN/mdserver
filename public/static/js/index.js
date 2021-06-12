// const container = document.getElementById("createDirectoryModal");
// const modal = new bootstrap.Modal(container);
// modal.hide();

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

document.addEventListener('click',function(e){
    if (e.target.classList.contains('deleteModal')) {
        if (e.target.dataset.type === 'directory') {
            document.getElementById("deleteModalLabel").innerText = 'Вы уверены что хотите удалить директорию ' + e.target.dataset.name + '?';
            document.getElementById("deleteModalTitle").innerText = "Удалить директорию";
        } else {
            document.getElementById("deleteModalLabel").innerText = 'Вы уверены что хотите удалить файл ' + e.target.dataset.name + '?';
            document.getElementById("deleteModalTitle").innerText = "Удалить файл";

            document.getElementById("deleteModalBtn").dataset.dirname = e.target.dataset.dirname;
        }

        const container = document.getElementById('deleteDirectoryModal');
        const modal = new bootstrap.Modal(container);
        modal.show();

        document.getElementById("deleteModalBtn").dataset.name = e.target.dataset.name;
        document.getElementById("deleteModalBtn").dataset.type = e.target.dataset.type;
    }
    if (e.target.classList.contains('createPost')) {
        document.getElementById("createPostModalBtn").dataset.dirname = e.target.dataset.dirname;
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