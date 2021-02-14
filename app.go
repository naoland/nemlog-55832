package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseZaifInfo Zaifからのレスポンス
type ResponseZaifInfo struct {
	LastPrice float64 `json:"last_price"`
}

func main() {
	r := gin.Default()
	r.GET("/last_price", func(c *gin.Context) {
		price, err := fetchLastPrice()
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, "XEM最終価格: %.3f", price)
	})
	r.Run("0.0.0.0:3000")
}

// Zaif取引所の最終価格を取得して、価格を返します
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

	response := new(ResponseZaifInfo)
	err = json.Unmarshal(byteArray, &response)

	if err != nil {
		log.Printf("マーシャルエラー: %#v\n", err) // 構造調べ
		return
	}
	price = response.LastPrice
	return
}
