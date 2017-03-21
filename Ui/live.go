package main

import (
	"fmt"
	//"io/ioutil"
	//"net/http"
	//	"os"
	"os/exec"
	//"path"
	"strings"
	//"time"
	"image"
	"image/jpeg"
	"os"
	"path"
)

func WidthHeightFile(picture string) string {
	total := len(picture)
	if total < 5 {
		return ""
	}
	var img image.Image
	file, err := os.Open(picture)
	if err != nil {
		return ""
	}

	img, err = jpeg.Decode(file)
	file.Close()

	bounds := img.Bounds()
	min, max := bounds.Min, bounds.Max
	height, width := max.Y-min.Y, max.X-min.X

	index := strings.LastIndex(picture, ".")
	ext := path.Ext(picture)
	var temp string = fmt.Sprintf("%s-%dx%d%s", picture[:index], width, height, ext)
	out, err := os.Create(temp)
	if err != nil {
		return path.Base(picture)
	}
	// write new image to file
	jpeg.Encode(out, img, nil)
	out.Close()
	//os.Rename(temp, name)
	return path.Base(temp)
}

func Checkvideo(content, style, scale, weight string) string {

	out := UPLOAD_PATH + "ok.jpg"
	var args = []string{content, style, scale, weight}
	cmd := exec.Command("/style/neural-style-master/niux.sh", args[0:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	fmt.Println(err)
	return out
}
func Checkv(content, style, scale, weight string) string {

	out := UPLOAD_PATH + "ok.jpg"
	var args = []string{content, style, scale, weight}
	cmd := exec.Command("/style/neural-style-master/final.sh", args[0:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	fmt.Println(err)
	return out
}
