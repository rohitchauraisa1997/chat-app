var socket = new WebSocket("ws://localhost:3002/ws")
let connect = (cb) => {
    // connects us to the websocket endpoint
    console.log("Attempting connection");

    socket.onopen = () => {
        console.log("Successfully connected");
    };

    socket.onmessage = msg => {
        console.log(msg);
        // triggers a callback whenever it rxs a new message from 
        // the websocket connection.
        cb(msg);
    };

    socket.onclose = event => {
        console.log("Socket connection closed: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error = ", error);
    };

};

let sendMsg = msg => {
    // sends message from frontend to our webserver
    console.log("sending message = ", msg);
    socket.send(msg);
};

export { connect, sendMsg };