function sendRequest(url, data, successCallback, errorCallback) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", url);
    xmlHttp.setRequestHeader("Content-Type", "application/json");
    xmlHttp.send(JSON.stringify(data));

    xmlHttp.onload = function(event) {
        if (xmlHttp.status >= 200 && xmlHttp.status < 300) {
            successCallback(xmlHttp.response);
        } else {
            errorCallback();
        }
    };
}

function sendRequestWithFile(url, file, successCallback, errorCallback) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", url);
    const formData = new FormData();
    formData.append("image", file);
    xmlHttp.send(formData);

    xmlHttp.onload = function(event) {
        if (xmlHttp.status >= 200 && xmlHttp.status < 300) {
            successCallback(xmlHttp.response);
        } else {
            errorCallback();
        }
    };
}