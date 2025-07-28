package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"task4/controller"
)

const (
	token string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTM3MTY2MzksImlkIjoxLCJ1c2VybmFtZSI6ImdpbiJ9.QolDiHrOEZB2sqpZCPHaofsyoRTbNkuXGIla0bFRXxQ"
)

func Get(url string, data url.Values) *controller.Response {
	if data != nil {
		url = fmt.Sprintf("%s?%s", url, data.Encode())
	}
	// 发起请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return nil
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return nil
	}
	defer respClose(resp)
	// 处理响应
	return respBody(resp)
}

func Post(url string, data interface{}) *controller.Response {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("JSON编码失败: %v\n", err)
		return nil
	}
	// 发起请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return nil
	}
	defer respClose(resp)
	// 处理响应
	return respBody(resp)
}

func respBody(resp *http.Response) *controller.Response {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应数据失败: %v\n", err)
		return nil
	}
	response := controller.Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("关闭请求失败: %v\n", err)
		return nil
	}
	return &response
}

func respClose(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		fmt.Printf("关闭请求失败: %v\n", err)
	}
}
