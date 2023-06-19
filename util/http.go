package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/longbridgeapp/qiyu-dl/model"
)

func DownloadFile(url string, filepath string) error {

	// 创建 HTTP 客户端对象
	client := http.Client{}

	// 创建 HTTP 请求对象
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 发送 HTTP 请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建本地文件
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应内容写入本地文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func PostQiyu(urlStr string, body []byte, key, secret string) (*model.QiyuResult, error) {
	nowStr := fmt.Sprint(time.Now().Unix())
	checksum := EncodeToChecksum(secret, body, nowStr)

	values := url.Values{}
	values.Set("time", nowStr)
	values.Set("checksum", checksum)
	values.Set("appKey", key)
	resp, err := http.Post(urlStr+"?"+values.Encode(), "application/json;charset=utf-6", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("七鱼接口请求失败:%v", err)
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("七鱼返回内容解析失败:%v", err)
	}

	messageRes := &model.QiyuResult{}
	err = json.Unmarshal(respData, messageRes)
	if err != nil {
		return nil, fmt.Errorf("七鱼返回内容解析失败:%v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("七鱼接口返回 statusCode 为%d, message 为 %v", resp.StatusCode, messageRes)
	}
	return messageRes, nil
}
