package main

import (
	"fmt"
	"liteBangumiAPI/bgmAPI"
)

func main() {

	bgmAPI.Token = "YOUR_TOKEN"    // 替换为你的 token
	bgmAPI.UserAgent = "UserAgent" // 设置参考：https://github.com/bangumi/api/blob/master/docs-raw/user%20agent.md

	//所有接口返回的都是一个[]byte和一个err。[]byte是返回体，可以进行json解析，err表示错误，如果为nil，那么说明没有错误

	//jData, err := bgmAPI.SearchName("CLANNAD", "书籍") //字符串搜索，前者为作品名，后者为类型。如果没有指定类型，那么将跨越所有类型进行搜索
	//jData, err := bgmAPI.SearchId("1191") //ID搜索，参数为ID名
	jData, err := bgmAPI.SearchCalendar() // 每日放送，没有参数
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jData))

}
