package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	_ "github.com/badgerodon/simulator/kernel"
)

func main() {
	li, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			res, err := http.Get("http://127.0.0.1:9000")
			if err != nil {
				panic(err)
			}
			bs, _ := ioutil.ReadAll(res.Body)
			fmt.Println("RESPONSE:", string(bs), res.Header)
			res.Body.Close()
			time.Sleep(10 * time.Second)
		}
	}()

	http.Serve(li, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	}))
}
