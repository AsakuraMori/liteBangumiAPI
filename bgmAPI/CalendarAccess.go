package bgmAPI

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func SearchCalendar() ([]byte, error) {
	server := &http.Server{
		Addr: PortNumber,
	}
	out := searchCalendar(server)
	time.Sleep(1 * time.Second)

	if len(out) == 0 {
		errMsg := errors.New("out is nil")
		return nil, errMsg
	}

	return out, nil

}

func searchCalendar(server *http.Server) []byte {
	go func() { // ListenAndServe是阻塞函数，要放在goroutine里跑
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	jsData, err := calendarAccess()
	if err != nil {
		panic(err)
		return nil
	}
	jsonData, jsonErr := json.MarshalIndent(jsData, "", "\t")
	if jsonErr != nil {
		panic(jsonErr)
		return nil
	}
	go func() {
		time.Sleep(10 * time.Second) // 等待 1 秒，确保响应已发送

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
func calendarAccess() ([]map[string]interface{}, error) {

	// 构建 Bangumi API 请求 URL
	apiURL := "https://api.bgm.tv/calendar"

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
		errMsg := errors.New(string(resp.StatusCode))
		return nil, errMsg
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据到 map[string]interface{}
	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
