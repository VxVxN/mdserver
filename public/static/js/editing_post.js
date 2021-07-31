let cursorLocation;

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
        hljs.initHighlighting.called = false;
        hljs.initHighlighting();
        return false;
    }
    sendRequest("/preview", data, successCallback);
}

document.getElementById("editingTab").onclick = function () {
    document.getElementById("editing").classList.remove('d-none');
}

document.getElementById("fileUploadBtn").onclick = function () {
    const textarea = document.getElementById('postText');
    cursorLocation = textarea.selectionStart;
    document.getElementById('fileUpload').click();
}

document.getElementById('fileUpload').addEventListener('change', function () {
    const textarea = document.getElementById('postText');

    const image = this.files[0];
    const successCallback = function () {
        const imagePath = `![](/static/images/${image.name})`;

        const text = textarea.value.substring(0, cursorLocation) + imagePath + textarea.value.substring(cursorLocation);

        textarea.value = text;
        textarea.focus();
        return false;
    };

    sendRequestWithFile(window.location.href+'/image_upload', image, successCallback);
});

document.getElementById('postText').onkeydown = function (event) {
    if (event.key === 'Tab') {
        const textarea = document.getElementById('postText');
        const fourSpace = '    ';
        const currentCursorPosition = textarea.selectionStart
        textarea.value = textarea.value.substring(0, currentCursorPosition) + fourSpace + textarea.value.substring(currentCursorPosition);
        textarea.selectionStart = currentCursorPosition + fourSpace.length;
        textarea.selectionEnd = currentCursorPosition + fourSpace.length;
        return false;
    }
}