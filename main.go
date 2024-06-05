package main

import (
	"kitexcall-issue/server"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		server.RunEchoServer("127.0.0.1:10001")
		wg.Done()
	}()

	go func() {
		server.RunMultipleServer("127.0.0.1:10002")
		wg.Done()
	}()

	wg.Wait()
}
