package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Path string `json:"path"`
}

type File struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	uuid := c.PostForm("uuid")

	for _, file := range files {
		err := c.SaveUploadedFile(file, "/tmp/share/"+uuid+".png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	ImgPost(uuid)

	c.JSON(http.StatusOK, gin.H{"message": "success!!"})
}

func ImgPost(uuid string) {
	url := "http://app:8008/mosaic/"
	request := new(Request)
	request.Path = "/tmp/share/" + uuid + ".png"
	req_json, _ := json.Marshal(request)
	log.Printf("[+] %s\n", string(req_json))
	res, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer([]byte(req_json)),
	)
	if err != nil {
		fmt.Println("[!] " + err.Error())
	} else {
		fmt.Println("[*] " + res.Status)
	}
	defer res.Body.Close()
}

func Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	err := os.Remove(fmt.Sprintf("/tmp/share/%s.png", uuid))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted!", uuid)})
}

func dirwalk(dir string) (files []File, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		path = strings.Replace(path, "images/", "http://localhost:8888/", 1)
		size := info.Size()
		f := File{
			Path: path,
			Size: size,
		}
		files = append(files, f)
		return nil
	})
	if err != nil {
		return
	}
	files = files[2:]
	return
}

func List(c *gin.Context) {
	files, err := dirwalk("/tmp/share")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}
