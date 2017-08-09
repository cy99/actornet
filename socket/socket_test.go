package socket

import "testing"

func TestServer(t *testing.T) {

	t.Log("server")

	Listen("127.0.0.1:8081")
}

func TestClient(t *testing.T) {

	t.Log("Client")

	Connect("127.0.0.1:8081")
}
