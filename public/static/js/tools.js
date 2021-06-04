function sendRequest(url, data, successCallback = function () {}, errorCallback = function () {}) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", url);
    xmlHttp.setRequestHeader("Content-Type", "application/json");
    xmlHttp.send(JSON.stringify(data));

    xmlHttp.onload = function(event) {
        if (event.status >= 200 && event.status < 300) {
            successCallback();
        } else {
            errorCallback();
        }
    };
}