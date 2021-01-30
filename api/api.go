package api

import (
	"Hello-gin/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//
func Callout(c *gin.Context){
	fmt.Println("Starting the application...")
	json := make(map[string]interface{}) //注意该结构接受的内容
	err := c.BindJSON(&json)
	if err != nil{
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	dc := config.DensoConfig{}
	err =dc.SetConf(json["env"].(string))
	if err != nil{
		 c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	url := "https://"+dc.Host+json["endPoint"].(string)
	//method := json["method"].(string)

	client := &http.Client {}
	req := &http.Request{}
	payload := &strings.Reader{}
	if json["body"] != nil {
		payload = strings.NewReader(json["body"].(string))
	}else{
		payload = nil
	}

	req, err = http.NewRequest(json["method"].(string), url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", dc.Auth)
	if json["body"] != nil {
		req.Header.Add("Content-Type", json["content-type"].(string))
	}


	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))
	c.JSON(res.StatusCode, gin.H{
		"method": json["method"],
		"endPoint": json["endPoint"],
		"message": string(body),
	})
	// ...
	fmt.Println("Terminating the application...")
}
