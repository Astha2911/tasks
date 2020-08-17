package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Fetch all album fields with the HGETALL command. Wrapping this
	// in the redis.Values() function transforms the response into type
	// []interface{}, which is the format we need to pass to
	// redis.ScanStruct() in the next step.
	values, err := redis.Values(conn.Do("HGETALL", "album:1"))
	if err != nil {
		log.Fatal(err)
	}

	var album Album
	err = redis.ScanStruct(values, &album)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", album)
}
