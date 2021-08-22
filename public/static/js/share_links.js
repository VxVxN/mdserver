
const deleteMedals = document.getElementsByName("deleteModal");
for(let i=0; i<deleteMedals.length; i++){
    deleteMedals[i].addEventListener("click", function(e){
        const container = document.getElementById('deleteModal');
        const modal = new bootstrap.Modal(container);
        modal.show();

        document.getElementById("deleteModalBtn").dataset.link = e.target.dataset.link;
    }, false);
}

document.getElementById("deleteModalBtn").onclick = function () {
    const successCallback = function () {
        window.location.reload();
        return false;
    }

    const data = {link:this.dataset.link};
    sendRequest("/delete_share_link", data, successCallback.bind(this, window));
}