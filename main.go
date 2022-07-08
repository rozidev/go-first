package main

import (
	"fmt"
	"log"

	"github.com/imroc/req/v3"
)

func main() {
	for i := 0; i < 40; i++ {
		client := req.C().EnableForceHTTP1() // Use C() to create a client.
		// client.SetProxyURL("http://myproxy:8080")
		// package get proxy
		resp, err := client.R().
			SetHeader("accept", "application/json").
			SetHeaders(map[string]string{
				"User-Agent": "Mozilla/5.0 (Linux; Android 10; MI 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.61 Mobile Safari/537.36200",
				// get
			}).
			Get("https://shopee.co.id/api/v4/item/get?itemid=5195830258&shopid=180687053")
		if err != nil {
			log.Fatal(err)
		}
		if resp.GetStatusCode() != 200 {
			log.Fatal("Status code is not 200")
		}
		fmt.Print(resp.GetStatus())
		// add space
		fmt.Println("")
		// fmt.Println(resp.GetStatusCode())
	}
}
