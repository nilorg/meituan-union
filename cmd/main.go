package main

import (
	"log"

	union "github.com/nilorg/meituan-union"
)

func main() {
	conf := union.NewConfig("app_key", "signature_key")
	client, err := union.NewClient(conf)
	if err != nil {
		log.Printf("创建客户端错误:%s\n", err)
		return
	}
	uoReq := &union.GenerateLinkRequest{
		ActID:     "33",
		SID:       "xxxxxx",
		LinkType:  "4",
		ShortLink: "0",
	}

	uoResp, err := client.GenerateLink(uoReq)
	if err != nil {
		log.Printf("生成推广链接错误:%s\n", err)
		// return
	}
	log.Printf("结果：%v\n", uoResp)
}
