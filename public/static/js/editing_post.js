document.getElementById("savePost").onclick = function () {
    const currentPageUrl = decodeURI(window.location.href);
    const splitUrl = currentPageUrl.split('/');
    const fileName = splitUrl[splitUrl.length-1];
    const dirName = splitUrl[splitUrl.length-2];

    const successCallback = function () {
        window.location.href = "/";
        return false;
    };

    const data = {
        dir_name:dirName,
        file_name: fileName,
        text: document.getElementById("postText").value,
    };
    sendRequest("/save_post", data, successCallback.bind(this, window));
}

document.getElementById("previewTab").onclick = function () {

    document.getElementById("editing").classList.add('d-none');

    const previewText = document.getElementById("postText").value;

    const data = {text:previewText};
    const successCallback = function (response) {
        document.getElementById("preview").innerHTML = response;
        return false;
    }
    sendRequest("/preview", data, successCallback);
}

document.getElementById("editingTab").onclick = function () {
    document.getElementById("editing").classList.remove('d-none');
}