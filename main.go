package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/longbridgeapp/qiyu-dl/util"
)

var exportUrl = "https://%s/openapi/export/session"
var checkUrl = "https://%s/openapi/export/session/check"

func main() {

	// 定义命令行参数
	dateStr := flag.String("date", "", "聊天记录的日期, 不能为空，格式是yyyy-mm-dd，如2024-05-23")
	output := flag.String("output", "", "指定文件存储位置")
	flag.Parse()

	// 判断是否输入了文件路径参数
	if *dateStr == "" {
		fmt.Println("请传 date 参数, 格式为 yyyy-mm-dd")
		os.Exit(1)
	}

	date, err := util.ConvertToDate(*dateStr)
	// 判断是否输入了文件路径参数
	if err != nil {
		fmt.Println("输入 date 格式不正确, 格式为 yyyy-mm-dd")
		os.Exit(1)
	}

	appKey := os.Getenv("QIYU_APP_KEY")
	if appKey == "" {
		fmt.Println("环境变量 QIYU_APP_KEY 是空的")
		os.Exit(1)
	}

	appSecret := os.Getenv("QIYU_APP_SECRET")
	if appSecret == "" {
		fmt.Println("环境变量 QIYU_APP_SECRET 是空的")
		os.Exit(1)
	}

	host := os.Getenv("QIYU_HOST")
	if host == "" {
		fmt.Println("环境变量 QIYU_HOST 是空的")
		os.Exit(1)
	}

	var dir string
	var filename = *dateStr + ".zip"
	if *output == "" {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			fmt.Println("获取当前命令执行目录错误, err:", err)
			os.Exit(1)
		}
	} else {
		dir, err = filepath.Abs(filepath.Dir(*output))
		if err != nil {
			fmt.Println("output 参数不合法, err:", err)
			os.Exit(1)
		}
		basename := filepath.Base(*output)
		if basename != "" && basename != "." {
			filename = basename
		}
	}
	*output = filepath.Join(dir, filename)

	body := map[string]string{
		"start": fmt.Sprint(date.UnixMilli()),
		"end":   fmt.Sprint(date.Add(time.Hour * 24).UnixMilli()),
	}

	bodyData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("处理请求body失败 err:" + err.Error())
		os.Exit(1)
	}
	messageRes, err := util.PostQiyu(fmt.Sprintf(exportUrl, host), bodyData, appKey, appSecret)
	if err != nil {
		fmt.Println("访问获取会话接口失败:", err)
		os.Exit(1)
	}
	if messageRes.Code != 200 {
		fmt.Println("七鱼接口返回业务错误:", messageRes)
		os.Exit(1)
	}

	checkBody := map[string]string{
		"key": messageRes.Message,
	}
	checkData, _ := json.Marshal(checkBody)

	for {
		time.Sleep(2 * time.Second)
		checkRes, err := util.PostQiyu(fmt.Sprintf(checkUrl, host), checkData, appKey, appSecret)
		if err != nil {
			fmt.Println("访问导出校验接口失败:", err)
			os.Exit(1)
		}

		// 如果错误码为14403，说明文件还在导出中，继续等待
		if checkRes.Code == 14403 {
			fmt.Println("文件导出中, 请稍等")
			continue
		} else if checkRes.Code == 200 {
			err = util.DownloadFile(checkRes.Message, *output)
			if err != nil {
				fmt.Println("文件下载失败, err:", err)
				os.Exit(1)
			}
			fmt.Printf("Download completed on %s\n", *output)
			break
		} else {
			fmt.Println("七鱼导出校验接口返回业务错误:", checkRes)
			os.Exit(1)
		}
	}

}
