package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Path string `json:"path"`
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

// func ImgPost(uuid string) error {
//     log.Println("img-post")
// 	url := "http://localhost:8008/mosaic/"
// 	json := `{"path":"/tmp/share/` + uuid + `.png"}`
// 	request, err := http.NewRequest(
// 		"POST",
// 		url,
// 		bytes.NewBuffer([]byte(json)),
// 	)
// 	if err != nil {
// 		return err
// 	}
//
// 	request.Header.Set("Content-Type", "application/json")
//
// 	client := &http.Client{}
// 	resp, err := client.Do(request)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
//
// 	return err
// }

// func ImgPost(uuid string) {
// 	link := "http://localhost:8008/mosaic/"
// 	ps := url.Values{}
// 	ps.Add("path", "/tmp/share/"+uuid+".png")
// 	fmt.Println(ps.Encode())
//
// 	res, err := http.PostForm(link, ps)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	defer res.Body.Close()
// 	body, _ := io.ReadAll(res.Body)
//
// 	log.Println(string(body))
//
// }

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
