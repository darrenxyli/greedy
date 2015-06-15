package redis

import (
	"github.com/garyburd/redigo/redis"
	"strings"
)

// Client is client of redis
type Client struct {
	Pool *redis.Pool
}

// NewClient create a new client
func NewClient(ip string, port string) *Client {
	return &Client{Pool: NewPool(ip, port)}
}

// NewPool create a pool with size 80
func NewPool(ip string, port string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", strings.Join([]string{ip, port}, ":"))
			if err != nil {
				panic(err.Error())
			}
			return con, err
		},
	}
}

// Put is rpush
func (client *Client) Put(dbName string, value interface{}) {
	c := client.Pool.Get()

	defer c.Close()

	c.Do("RPUSH", dbName, value)

}

// Get is lpop
func (client *Client) Get(dbName string) (string, error) {
	c := client.Pool.Get()

	defer c.Close()

	return redis.String(c.Do("LPOP", dbName))
}

// QueueSize get the size of queue
func (client *Client) QueueSize(dbName string) (int, error) {

	c := client.Pool.Get()

	defer c.Close()

	return redis.Int(c.Do("LLEN", dbName))
}
