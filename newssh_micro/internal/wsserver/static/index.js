const form = document.getElementById("connectForm");

form.addEventListener("submit", function (event) {
    event.preventDefault();

    const userID = document.getElementById("userID").value;
    const serverID = document.getElementById("serverID").value;

    if (!userID || !serverID) {
        return;
    }

    window.location.href =
        `/terminal?user_id=${encodeURIComponent(userID)}&server_id=${encodeURIComponent(serverID)}`;
});