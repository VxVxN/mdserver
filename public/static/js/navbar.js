if (document.getElementById("logOut") !== null) {
    document.getElementById("logOut").onclick = function () {
        const successCallback = function () {
            window.location.href = "/";
            return false;
        }

        sendRequest("/log_out", {}, successCallback);
    };
}