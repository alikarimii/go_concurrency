package mutex

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func StartNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get() // get connection from pool
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn) // put it back to pool when finish my job
			conn.Close()
		}
	}()
	return &wg
}
func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: mockConnectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func mockConnectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}
