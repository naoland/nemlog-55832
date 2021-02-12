package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/last_price", func(c *gin.Context) {
		price, err := fetchLastPrice()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{
			"price": price,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func fetchLastPrice() (price float64, err error) {
	uri := "https://api.zaif.jp/api/1/last_price/xem_jpy"
	req, _ := http.NewRequest("GET", uri, nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	// {"last_price": 41.7796}

	response := new(ResponseZaifInfo)
	err = json.Unmarshal(byteArray, &response)

	if err != nil {
		log.Printf("マーシャルエラー: %#v\n", err) // 構造調べ
		return
	}
	price = response.LastPrice
	fmt.Printf("%.3f\n", price)
	// 41.780
	return
}

type ResponseZaifInfo struct {
	LastPrice float64 `json:"last_price"`
}
