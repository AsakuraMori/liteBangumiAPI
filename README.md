# liteBangumiAPI

包装了部分BangumiAPI到go里面。采用http通信形式获取数据。

目前仅支持条目搜索查看，角色搜索查看，人物搜索查看。

详细API见下：

## 全局变量设置

liteBangumiAPI有两个全局变量，必须在调用前设置。填写全局变量如下：

``` go
bgmAPI.Token = "YOUR_ACCESS_TOKEN"
bgmAPI.UserAgent = "YOUR User-Agent"
```

说明：

1.bangumi使用OAuth 2.0，token格式已经在内部写好了，这里token直接填写你从bangumi获取的token就行。**【不需要写入Bearer】**。

2.UserAgent的形式，请参考https://github.com/bangumi/api/blob/master/docs-raw/user%20agent.md

## go封装的API接口：

所有的接口返回都是一个[]byte和一个err。[]byte是返回体，可以进行json解析。err表示错误，如果为nil，那么说明没有错误。

### bgmAPI.SearchSubjectByName([标题名], [类型名], [最大数量])

该API可以搜索所有的条目。如果你想规定搜索类型，请在类型名里写【一个】类型名。如果你想全局搜搜，请在类型名用""表示，类型名只能是以下字符串：

```
书籍、动漫、音乐、游戏、三次元
```

如果你想规定搜索的数量，可以直接修改最大数量，不超过25条。

调用示例：

```go
jData, err := bgmAPI.SearchName("缘之空", "", 25)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(string(jData))
```

###  bgmAPI.SearchSubjectById([ID])

通过指定的ID搜索相关条目的详细内容。

调用示例：

``` go
jData1, err1 := bgmAPI.SearchSubjectId("1191")
if err1 != nil {
	fmt.Println(err1)
	return
}
```

### bgmAPI.SearchPersonsByName([人名])

通过指定人名（现实人名，公司等），搜索有关的人物条目。

调用示例：

``` go
jData3, err3 := bgmAPI.SearchPerson("Ceui")
if err3 != nil {
	fmt.Println(err3)
	return
}
fmt.Println(string(jData3))
```

### bgmAPI.SearchPersonsById([ID])

通过指定人名/公司名的ID来获取对应的详细内容。

调用示例：

```go
jData4, err4 := bgmAPI.SearchPersonId("6141")
if err4 != nil {
	fmt.Println(err4)
	return
}
fmt.Println(string(jData4))
```

### bgmAPI.SearchCharactersByName([角色名])

通过指定的角色名，搜索有关角色条目

调用示例：

```go
jData5, err5 := bgmAPI.SearchCharacters("天王寺瑚太朗")
if err5 != nil {
	fmt.Println(err5)
	return
}
fmt.Println(string(jData5))
```

### bgmAPI.SearchCharactersById([ID])

通过指定角色ID，显示角色详细内容

调用示例：

```go
jData6, err6 := bgmAPI.SearchCharactersId("12062")
if err6 != nil {
	fmt.Println(err6)
	return
}
fmt.Println(string(jData6))
```

### bgm.SearchEpisodesById([ID], [类型名])

指定作品ID和类型来搜索章节。类型名只能是以下字符串：

```
本篇、特别篇、OP、ED、预告/宣传/广告、MAD、其他
```

如果需要全局搜索，可指定类型名为""

调用示例：

``` go
jData7, err7 := bgmAPI.SearchEpisodesById("464376", "本篇")
if err7 != nil {
	fmt.Println(err7)
	return
}
fmt.Println(string(jData7))
```

### bgmAPI.SearchpisodesByEpisodesId[章节ID])

指定章节ID来获取当前章节的详细信息。

调用示例：

```go
jData8, err8 := bgmAPI.SearchEpisodesByEpisodesId("1354646")
if err8 != nil {
	fmt.Println(err8)
	return
}
fmt.Println(string(jData8))
```

### bgmAPI.SearchUserByName([userName])

指定username来获取用户基本信息。

调用示例：

```go
jData9, err9 := bgmAPI.SearchUserByName("soratane")
if err9 != nil {
	fmt.Println(err9)
	return
}
fmt.Println(string(jData9))
```

### bgmAPI.SearchCalendar()

每日放送，没有参数

调用示例：

```go
jData2, err2 := bgmAPI.SearchCalendar()
if err2 != nil {
	fmt.Println(err2)
	return
}
fmt.Println(string(jData2))
```

