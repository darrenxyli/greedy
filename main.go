package main

import (
	"fmt"
	"runtime"
	"sync"

	dedis "github.com/darrenxyli/greedy/database/redis"
	"github.com/garyburd/redigo/redis"
)

// 生成连接池
var pool = dedis.NewPool("192.80.146.5", "6379")

func publish(channel, value interface{}) {
	c := pool.Get()
	defer c.Close()
	c.Do("PUBLISH", channel, value)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := pool.Get()

	defer c.Close()
	var wg sync.WaitGroup
	wg.Add(2)

	psc := redis.PubSubConn{Conn: c}

	// This goroutine receives and prints pushed notifications from the server.
	// The goroutine exits when the connection is unsubscribed from all
	// channels or there is an error.
	go func() {
		defer wg.Done()
		for {
			switch n := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("Sub1: Message: %s %s\n", n.Channel, n.Data)
			case redis.PMessage:
				fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
			case redis.Subscription:
				fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
				if n.Count == 0 {
					return
				}
			case error:
				fmt.Printf("error: %v\n", n)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			switch n := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("Sub2: Message: %s %s\n", n.Channel, n.Data)
			case redis.PMessage:
				fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
			case redis.Subscription:
				fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
				if n.Count == 0 {
					return
				}
			case error:
				fmt.Printf("error: %v\n", n)
				return
			}
		}
	}()

	// This goroutine manages subscriptions for the connection.
	go func() {
		defer wg.Done()

		psc.Subscribe("example")

		// The following function calls publish a message using another
		// connection to the Redis server.

		publish("example", "hello")
		publish("example", "world")

		for index := 0; index < 100; index++ {
			publish("example", index)
		}

		// Unsubscribe from all connections. This will cause the receiving
		// goroutine to exit.
		psc.Unsubscribe()
	}()

	go func() {
		defer wg.Done()

		psc.Subscribe("example")

		// The following function calls publish a message using another
		// connection to the Redis server.

		publish("example", "hello")
		publish("example", "world")

		for index := 0; index < 200; index++ {
			publish("example", index)
		}

		// Unsubscribe from all connections. This will cause the receiving
		// goroutine to exit.
		psc.Unsubscribe()
	}()

	wg.Wait()
}
