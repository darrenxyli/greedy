package redis

import (
	"strings"

	"github.com/garyburd/redigo/redis"
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
func (client *Client) Get(dbName string) {
	c := client.Pool.Get()

	defer c.Close()

	c.Do("LPOP", dbName)
}
