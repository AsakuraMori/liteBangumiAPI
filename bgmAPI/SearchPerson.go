package bgmAPI

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
)

func SearchPersonsByName(keyWord string) ([]byte, error) {
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
	out := searhPersonsByName(server, keyWord)
	time.Sleep(1 * time.Second)

	if len(out) == 0 {
		errMsg := errors.New("out is nil")
		return nil, errMsg
	}

	return out, nil

}
func searhPersonsByName(server *http.Server, keyWord string) []byte {
	go func() { // ListenAndServe是阻塞函数，要放在goroutine里跑
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	jsData, err := searchPerson(keyWord)
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

func searchPerson(keyWord string) (map[string]interface{}, error) {
	keyword := keyWord
	// 构造请求体
	requestBody := map[string]interface{}{
		"keyword": keyword,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", "https://api.bgm.tv/v0/search/persons", bytes.NewBuffer(jsonBody))
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
