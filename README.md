# ws-rpc
How to set up a gorilla/websocket server to process JSON-RPC calls

Most of the interesting stuff is in internal/common/rwc.go, which contains an implementaiton of an io.ReadWriteCloser, required by jsonrpc.

gorilla/websocket provides a ReadMessage/WriteMessage pair, or NextReader/NextWriter, which produce an io.Reader and an io.WriteCloser respectively, which can be used to read and  write entire messages. Since it requires the reader to read until io.EOF and the writer to call Close() after writing, the common.ReadWriteCloser adapter exists to cope with the requirements, of which the codec knows nothing about.

The example also implements minimal Ping/Pong handling. The server sends Pings and expects a Pong within a specific time period. If it doesn't receive the Pong, it assumes the client is gone. Once it receives the Pong, it extends the Read deadline, so that a hanging Read operation keeps waiting for more requests from the client.
