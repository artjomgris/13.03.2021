package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Action struct {
	Act string `json:"action"`
	Id  int    `json:"id"`
}

type Update struct {
	Act  Action
	Text string `json:"body"`
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var data []byte
	var i = 0
	for {
		var inp Action
		in := bufio.NewReader(os.Stdin)
		fmt.Println("Enter Action and Id")
		fmt.Fscanf(in, "%v %v\n", &inp.Act, &inp.Id)
		if inp.Act == "update" || inp.Act == "Update" {
			var upd Update
			upd.Act = inp
			fmt.Println("Enter Text (Use _ instead of spaces)")
			fmt.Scan(&upd.Text)
			form, err := json.Marshal(upd)
			if err != nil {
				fmt.Println(err)
				return
			}
			//fmt.Println(string(form))
			data = append(data, form...)
		} else {
			form, err := json.Marshal(inp)
			if err != nil {
				fmt.Println(err)
				return
			}
			//fmt.Println(string(form))
			data = append(data, form...)
		}
		i++
		fmt.Println("Enter stop to stop or Press \"enter\" to continue")
		var cont string
		fmt.Fscanf(in, "%v\n", &cont)
		if cont == "stop" {
			break
		}
		data = append(data, byte(','))
	}

	conn.Write(data)

	buf := make([]byte, 2000)
	for j := 0; j < i; j++ {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
	}

}
