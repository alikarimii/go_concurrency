package mutex_test

import (
	"io/ioutil"
	"net"
	"testing"

	mutex "github.com/alikarimii/go_concurrency/3.mutex"
)

func init() {
	daemonStarted := mutex.StartNetworkDaemon()
	daemonStarted.Wait()
}
func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
