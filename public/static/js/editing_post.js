savePost.onclick = function () {
    const currentPageUrl = window.location.href
    const splitUrl = currentPageUrl.split('/');
    const fileName = splitUrl[splitUrl.length-1]

    const data = {name:fileName, text: postText.value}
    sendRequest("/save", data, function (window) {
        window.location.href = "/";
        return false;
    }(window))
}