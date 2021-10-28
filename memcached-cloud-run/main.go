package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gomodule/redigo/redis"
)

const (
	REDIS_ENDPOINT = "127.0.0.1:6379"
	MEMCACHED_ENDPOINT = "127.0.0.1:11211"
)

var (
	rPool *redis.Pool
	mClient *memcache.Client
)

func main()  {
	rPool = &redis.Pool{
		MaxIdle: 10,
		Dial:    func() (redis.Conn, error) {
			return redis.Dial("tcp", REDIS_ENDPOINT)
		},
	}
	mClient = memcache.New(MEMCACHED_ENDPOINT)
	m := http.NewServeMux()
	m.HandleFunc("/redis", func (w http.ResponseWriter, r *http.Request) {
		conn := rPool.Get()
		defer conn.Close()
		counter, err := redis.Int(conn.Do("INCR", "visits"))
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf(`couldn't fetch "visits"... error: %v`, err),
				http.StatusInternalServerError,
			)
			return
		}
		fmt.Fprintf(w, "Visiter: %d", counter)
	})
	m.HandleFunc("/memcached", func(w http.ResponseWriter, r *http.Request) {
		mClient.Set(
			&memcache.Item{
				Key: "current-time",
				Value: []byte(fmt.Sprintf("%d", time.Now().Unix())),
			},
		)
		item, err := mClient.Get("current-time")
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf(`couldnt fetch "current-time"... error: %v`, err),
				http.StatusInternalServerError,
			)
			return
		}
		fmt.Fprintf(w, "Current time: %s", item.Value)
	})
	log.Fatal(http.ListenAndServe(":8080", m))
}
