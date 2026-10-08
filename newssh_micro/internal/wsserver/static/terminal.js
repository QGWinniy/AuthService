const terminalElement = document.getElementById("terminal");

const terminal = new Terminal({
    cursorBlink: true,
    fontSize: 15,
    scrollback: 5000,
});

terminal.open(terminalElement);

terminal.focus();


// ============================================================
// Получаем user_id / server_id
// ============================================================

const params = new URLSearchParams(window.location.search);

const userID = params.get("user_id");
const serverID = params.get("server_id");

if (!userID || !serverID) {
    terminal.write("Missing user_id or server_id\r\n");
    throw new Error("Missing user_id or server_id");
}


// ============================================================
// WebSocket
// ============================================================

const protocol =
    window.location.protocol === "https:"
        ? "wss:"
        : "ws:";

const ws = new WebSocket(
    `${protocol}//${window.location.host}/ws?user_id=${userID}&server_id=${serverID}`
);


// ============================================================
// SSH -> Browser
// ============================================================

ws.onmessage = function (event) {

    const message = JSON.parse(event.data);

    /*
        Go:

        []byte -> JSON -> Base64 string

        Поэтому message.result это Base64.
    */

    const bytes = base64ToBytes(message.result);

    /*
        xterm умеет принимать Uint8Array.
    */

    terminal.write(bytes);
};


// ============================================================
// Browser -> SSH
// ============================================================

terminal.onData(function (data) {

    if (ws.readyState !== WebSocket.OPEN) {
        return;
    }

    /*
        data содержит нажатие клавиши.

        Например:

        "a"
        "ls"
        "\r"
        "\u001b[A" // arrow up

        Превращаем строку в UTF-8 bytes.
    */

    const bytes = new TextEncoder().encode(data);

    /*
        Go ждёт:

        {
            "command": []byte
        }

        JSON для []byte должен содержать Base64.
    */

    ws.send(JSON.stringify({
        command: bytesToBase64(bytes)
    }));
});


// ============================================================
// Connection events
// ============================================================

ws.onopen = function () {
    console.log("WebSocket connected");
};

ws.onclose = function () {
    terminal.write(
        "\r\n\r\n[ SSH connection closed ]\r\n"
    );
};

ws.onerror = function (error) {
    console.error("WebSocket error:", error);

    terminal.write(
        "\r\n\r\n[ WebSocket error ]\r\n"
    );
};


// ============================================================
// Base64
// ============================================================

function bytesToBase64(bytes) {

    let binary = "";

    const chunkSize = 0x8000;

    for (
        let i = 0;
        i < bytes.length;
        i += chunkSize
    ) {

        const chunk = bytes.subarray(
            i,
            Math.min(i + chunkSize, bytes.length)
        );

        binary += String.fromCharCode(...chunk);
    }

    return btoa(binary);
}


function base64ToBytes(base64) {

    const binary = atob(base64);

    const bytes = new Uint8Array(
        binary.length
    );

    for (
        let i = 0;
        i < binary.length;
        i++
    ) {
        bytes[i] = binary.charCodeAt(i);
    }

    return bytes;
}