package main

import (
	"fmt"
	"log"

	"github.com/corpix/uarand"
	"github.com/imroc/req/v3"
)

func main() {
	for i := 0; i < 10; i++ {
		client := req.C().EnableForceHTTP1() // Use C() to create a client.
		resp, err := client.R().
			SetHeader("accept", "application/json").
			SetHeaders(map[string]string{
				"User-Agent": uarand.GetRandom(),
				// get
			}).
			Get("https://shopee.co.id/api/v4/item/get?itemid=5195830258&shopid=180687053")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)
	}
}
