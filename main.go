package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	userProto "grpc-starter/proto"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var user = &userProto.User{
		Name:  "abc",
		Email: "aa@qq.com",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/protobuf")
		data, _ := proto.Marshal(user)
		w.Write(data)
	})

	go func() {
		time.Sleep(1 * time.Second)
		resp, err := http.Get("http://localhost:18080")
		if err != nil {
			return
		}
		defer resp.Body.Close()

		data, _ := ioutil.ReadAll(resp.Body)
		var respUser userProto.User
		proto.Unmarshal(data, &respUser)

		fmt.Println("respUser:", respUser)
	}()

	http.ListenAndServe(":18080", nil)
}
