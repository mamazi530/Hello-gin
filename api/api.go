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

//Callout service
func Callout(c *gin.Context) {

	recJson := make(map[string]interface{}) //注意该结构接受的内容
	err := c.BindJSON(&recJson)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	res, err := sendOut(recJson)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer res.Body.Close()
	//outMsg, err := json.Marshal(string(body))

	c.String(http.StatusOK, string(body))

}

//sendOut request will return string from db or service out
func sendOut(json map[string]interface{}) (res *http.Response, err error) {
	dc := config.DensoConfig{}
	err = dc.SetConf(json["env"].(string))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := &http.Client{}
	url := "https://" + dc.Host + json["endPoint"].(string)

	req := &http.Request{}

	method := json["method"].(string)

	if json["body"] != nil {
		req, err = http.NewRequest(method, url, strings.NewReader(json["body"].(string)))
		req.Header.Add("Content-Type", json["contentType"].(string))

	} else {
		req, err = http.NewRequest(method, url, nil)

	}

	req.Header.Add("Authorization", dc.Auth)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err = client.Do(req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil

}
