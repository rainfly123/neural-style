package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	//	"path"
	//"sort"
	//"strconv"
	////	"strings"
	//	"time"
)

var logger *log.Logger

type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Url     string `json:"data"`
}

func writev3Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		io.WriteString(w, fmt.Sprintf("<html><head><title>我的第一个页面</title></head><body><form action=\"writev2\" method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file0'/><br/><<input type=\"file\" name='file1'/><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>"))
	} else {
		file, head, err := req.FormFile("file0")
		if err != nil {
			jsonres := JsonResponse{1, "参数错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		defer file.Close()
		temp := getFileName(head.Filename)
		uuidFile := UPLOAD_PATH + temp
		fW, err := os.Create(uuidFile)
		defer fW.Close()
		if err != nil {
			jsonres := JsonResponse{2, "系统错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		_, err = io.Copy(fW, file)
		if err != nil {
			jsonres := JsonResponse{2, "系统错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}

		file1, head1, err1 := req.FormFile("file1")
		if err1 != nil {
			jsonres := JsonResponse{1, "参数错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		defer file1.Close()
		temp = getFileName(head1.Filename)
		uuidFile1 := UPLOAD_PATH + temp
		fW1, err1 := os.Create(uuidFile1)
		defer fW1.Close()
		if err1 != nil {
			jsonres := JsonResponse{2, "系统错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		_, err1 = io.Copy(fW1, file1)
		if err1 != nil {
			jsonres := JsonResponse{2, "系统错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}

		scale := req.FormValue("scale")
		weight := req.FormValue("weight")
		fmt.Println(scale, weight)

		Checkv(uuidFile, uuidFile1, scale, weight)

		b:=(ACCESS_URL + "ok.jpg")
		io.WriteString(w, b)
	}
}



func writev2Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		io.WriteString(w, fmt.Sprintf("<html><head><title>我的第一个页面</title></head><body><form action=\"writev2\" method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file0'/><br/><<input type=\"file\" name='file1'/><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>"))
	} else {
		file, head, err := req.FormFile("file0")
		if err != nil {
			jsonres := JsonResponse{1, "参数错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		defer file.Close()
		temp := getFileName(head.Filename)
		uuidFile := UPLOAD_PATH + temp
		fW, err := os.Create(uuidFile)
		defer fW.Close()
		if err != nil {
			jsonres := JsonResponse{2, "系统错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		_, err = io.Copy(fW, file)
		if err != nil {
			jsonres := JsonResponse{2, "系统错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}

		file1, head1, err1 := req.FormFile("file1")
		if err1 != nil {
			jsonres := JsonResponse{1, "参数错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		defer file1.Close()
		temp = getFileName(head1.Filename)
		uuidFile1 := UPLOAD_PATH + temp
		fW1, err1 := os.Create(uuidFile1)
		defer fW1.Close()
		if err1 != nil {
			jsonres := JsonResponse{2, "系统错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		_, err1 = io.Copy(fW1, file1)
		if err1 != nil {
			jsonres := JsonResponse{2, "系统错误", ""}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}

		scale := req.FormValue("scale")
		weight := req.FormValue("weight")
		fmt.Println(scale, weight)

		Checkvideo(uuidFile, uuidFile1, scale, weight)

		b:=(ACCESS_URL + "ok.jpg")
		io.WriteString(w, b)
	}
}

var staticHandler http.Handler = http.FileServer(http.Dir(UPLOAD_PATH))

func StaticServer(w http.ResponseWriter, req *http.Request) {
	staticHandler.ServeHTTP(w, req)
}
func main() {

	http.HandleFunc("/", StaticServer)
	http.HandleFunc("/writev2", writev2Handle)
	http.HandleFunc("/writev3", writev3Handle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
	}
}
