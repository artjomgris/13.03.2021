package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

type Action struct {
	Act string `json:"action"`
	Id  int    `json:"id"`
}

type Update struct {
	Act  Action
	Text string `json:"body"`
}

var DB = []string{"Hello Wordl!", "Text", "Lorem Ipsum", "Posts", "Rolling Stone", "Windows"}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	buf := make([]byte, 2000)
	n, err := conn.Read(buf)
	if err != nil {
		conn.Close()
		return
	}
	var action Action
	data := string(buf[:n])
	for {
		i := strings.Index(data, "},{")
		if i > 4 {
			err := json.Unmarshal([]byte(data[:i+1]), &action)
			if err != nil {
				panic(err)
			}
			fmt.Println(action.Act)
			if len(action.Act) < 4 {
				conn.Write([]byte("update"))
			} else {
				conn.Write([]byte(action.Act))
			}

			data = data[i+2:]
			i = 0
		} else {
			break
		}
	}
	err = json.Unmarshal([]byte(data), &action)
	if err != nil {
		panic(err)
	}
	conn.Write([]byte(action.Act))

	conn.Close()
}
