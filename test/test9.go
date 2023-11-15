package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func req(pipelineId string) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建 GET 请求
	req, err := http.NewRequest("GET", "http://192.168.1.224:8000", nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(body))
}
