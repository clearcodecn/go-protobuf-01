package main

import (
	"bytes"
	"grpc-starter/protobuild/zip"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1. 客户端上传文件，读到body里面，直接获取到.
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		// 写. 文件夹 爸爸 目录
		tmpDir := os.TempDir()
		dir := filepath.Join(tmpDir, "pb")
		os.RemoveAll(dir)
		os.MkdirAll(dir, 0777)
		defer os.RemoveAll(dir)

		zipFile := filepath.Join(dir, "pb.zip")
		err = ioutil.WriteFile(zipFile, data, 0777)
		if err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}
		// 2. 解压缩文件
		err = zip.UnZip(zipFile, dir)
		if err != nil {
			w.Write([]byte("unzip failed: " + err.Error()))
			w.WriteHeader(400)
			return
		}

		shfile := filepath.Join(dir, "build.sh")
		cmd := exec.Command("bash", "-c", shfile)

		var buf bytes.Buffer
		cmd.Stdout = &buf
		cmd.Stderr = &buf

		if err := cmd.Run(); err != nil {
			w.Write([]byte("exec failed: " + err.Error()))
			w.WriteHeader(400)
			return
		}

		data, err = zip.Zip(dir)
		if err != nil {
			w.Write([]byte("zip failed: " + err.Error()))
			w.WriteHeader(400)
			return
		}

		w.Write(data)
	})

	http.ListenAndServe(":9300", nil)
}
