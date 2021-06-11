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

document.addEventListener('click',function(e){
    if (e.target.classList.contains('deleteDirectory')) {
        document.getElementById("deleteDirectoryName").innerText = 'Вы уверены что хотите удалить директорию ' + e.target.dataset.name + '?';

        const container = document.getElementById('deleteDirectoryModal');
        const modal = new bootstrap.Modal(container);
        modal.show();

        document.getElementById("deleteDirectory").dataset.name = e.target.dataset.name;
    }
})

document.getElementById("deleteDirectory").onclick = function () {
    console.log(this.dataset.name)

    const successCallback = function () {
        window.location.reload();
        return false;
    }

    const data = {name:this.dataset.name};
    sendRequest("/delete_directory", data, successCallback.bind(this, window));
}