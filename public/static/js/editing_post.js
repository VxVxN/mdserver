const currentPageUrl = window.location.href;
const splitUrl = currentPageUrl.split('/');
const fileName = splitUrl[splitUrl.length-1];

document.getElementById("savePost").onclick = function () {
    const data = {name:fileName, text: postText.value};
    sendRequest("/save", data, function (window) {
        window.location.href = "/";
        return false;
    }(window));
}

document.getElementById("previewTab").onclick = function () {
    const splitFileName = fileName.split('#');
    const previewFileName = splitFileName[0];

    document.getElementById("editing").classList.remove('d-flex');

    const data = {name:previewFileName};
    const successCallback = function (response) {
        document.getElementById("preview").innerHTML = response;
        return false;
    }
    sendRequest("/preview", data, successCallback);
}

document.getElementById("editingTab").onclick = function () {
    document.getElementById("editing").classList.add('d-flex');

}