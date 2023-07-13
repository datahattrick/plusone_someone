var socket = new WebSocket("ws://localhost:3000/ws/123?v=1.0")

let connect = () => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected")
    };

    socket.onmessage = msg => {
        console.log(msg);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

};

let sendMsg = (msg: string ) => {
    console.log("sending msg: ", msg);
    socket.send(msg);
};

export {connect, sendMsg};