document.getElementById("checkPassword").onclick = function () {
    const inputPassword = document.getElementById("password").value;
    if (inputPassword === '') {
        return;
    }
    const currentPageUrl = decodeURI(window.location.href);
    const splitUrl = currentPageUrl.split('/');
    const fileName = splitUrl[splitUrl.length-1];
    const dirName = splitUrl[splitUrl.length-2];

    const data = {password: inputPassword};
    sendRequest("/check_password", data, function (response) {
        if (JSON.parse(response).valid) {
            const successCallback = function () {
                window.location.href = "/";
                return false;
            };
            const data = {dir_name:dirName, file_name: fileName, text: document.getElementById("postText").value};
            sendRequest("/save_post", data, successCallback.bind(this, window));
        } else {
            alert("Invalid password");
        }
        document.getElementById("password").value = "";
        return false;
    });
}

document.getElementById("cancelPassword").onclick = function () {
    document.getElementById("password").value = "";
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