
document.getElementById("editPost").onclick = function () {
    const currentPageUrl = decodeURI(window.location.href);
    const splitUrl = currentPageUrl.split('/');
    const dirName = splitUrl[splitUrl.length-2];
    const fileName = splitUrl[splitUrl.length-1];

    window.location.href = `/edit/${dirName}/${fileName}`;
}