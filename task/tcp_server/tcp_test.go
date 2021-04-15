package tcp_server

import "testing"

func BenchmarkSendTcpRequest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SendTcpRequest()
	}
}
