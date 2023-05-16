package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Download(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET")
		w.Write([]byte(string("method is GET")))
	case "POST":
		fmt.Println("POST")
		r.ParseForm()
		imgFile, _, err := r.FormFile("file") //获取文件内容
		if err != nil {
			log.Fatal(err)
		}
		defer imgFile.Close()
		imgName := ""
		files := r.MultipartForm.File //获取表单中的信息

		for k, v := range files {
			for _, vv := range v {
				fmt.Println(k + ":" + vv.Filename) //获取文件名
				imgName = vv.Filename
			}
		}
		saveFile, _ := os.Create("./" + imgName)
		defer saveFile.Close()
		io.Copy(saveFile, imgFile) //保存
		w.Write([]byte("successfully saved"))
	default:
		fmt.Print("error method")
	}
}

func main() {
	server := &http.Server{
		ReadTimeout:  200 * time.Second,
		WriteTimeout: 200 * time.Second,
		Addr:         "127.0.0.1:8888",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/post", Download)
	server.Handler = mux
	server.ListenAndServe()
}

/*
使用nginx代理
client_max_body_size 8192m;
location /post {proxy_pass http://127.0.0.1:8888;}
*/
