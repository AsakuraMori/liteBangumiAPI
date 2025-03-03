package bgmAPI

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func SearchEpisodesById(Id string, typeName string) ([]byte, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port
	str := strconv.Itoa(port)
	server := &http.Server{
		Addr: ":" + str,
	}
	out := searchEpisodesById(server, Id, typeName)
	time.Sleep(1 * time.Second)

	if len(out) == 0 {
		errMsg := errors.New("out is nil")
		return nil, errMsg
	}

	return out, nil

}
func searchEpisodesById(server *http.Server, Id, typeName string) []byte {
	go func() { // ListenAndServe是阻塞函数，要放在goroutine里跑
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	jsData, err := searchEpisodes(Id, typeName)
	if err != nil {
		panic(err)
		return nil
	}
	jsonData, jsonErr := json.MarshalIndent(jsData, "", "\t")
	if jsonErr != nil {
		panic(jsonErr)
		return nil
	}
	//fmt.Println(string(jsonData))
	go func() {
		time.Sleep(10 * time.Second)

		// 创建一个超时上下文
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 优雅地关闭服务器
		if err := server.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	return jsonData
}
func searchEpisodes(Id, typeName string) (map[string]interface{}, error) {
	keyword := Id
	if keyword == "" {
		errMsg := errors.New("Keyword is required")
		return nil, errMsg
	}
	if len(typeName) == 0 {
		typeName = ""
	}
	sType := 0
	switch typeName {
	case "本篇":
		sType = 0
	case "特别篇":
		sType = 1
	case "OP":
		sType = 2
	case "ED":
		sType = 3
	case "预告/宣传/广告":
		sType = 4
	case "MAD":
		sType = 5
	case "其他":
		sType = 6
	default:
		sType = 0
	}

	baseURL := "https://api.bgm.tv/v0/episodes"
	params := url.Values{}
	params.Add("subject_id", keyword)
	params.Add("type", strconv.Itoa(sType))

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("User-Agent", UserAgent)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		errMsg := errors.New(strconv.Itoa(resp.StatusCode))
		return nil, errMsg
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据到 map[string]interface{}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
