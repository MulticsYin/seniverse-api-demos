/*******************************************************************************************************
访问URL：
https://api.seniverse.com/v3/weather/now.json?key=使用自己的Key&location=zhuhai&language=zh-Hans&unit=c
返回数据：
{'results': [{'last_update': '2017-12-08T09:25:00+08:00',
              'location': {'country': 'CN',
                           'id': 'WEBY8Q5HHUCU',
                           'name': '珠海',
                           'path': '珠海,珠海,广东,中国',
                           'timezone': 'Asia/Shanghai',
                           'timezone_offset': '+08:00'},
              'now': {'code': '0', 'temperature': '17', 'text': '晴'}}]}

注:
  1. 使用个人开发者帐号返回的数据编写，服务数据不完全，不过看懂了这个程序，其他的都一样啦。
  2. 运行程序： go run WeatherNow.go
******************************************************************************************************/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

var API = "https://api.seniverse.com/v3/weather/now.json" // API URL，可替换为其他 URL
var KEY = "use Key for youself"                           // API key
var LOCATION = "zhuhai"                                   // 所查询的位置，可以使用城市拼音、v3 ID、经纬度等
var LANGUAGE = "zh-Hans"                                  // 查询结果的返回语言
var UNIT = "c"                                            // 单位

// URL Access url.
var URL = API + "?" + "key=" + KEY + "&" + "location=" + LOCATION + "&" + "language=" + LANGUAGE + "&" + "unit=" + UNIT

// ResponseJSON Store response data 'results'.
type ResponseJSON struct {
	Results []WeatherNow `json:"results,omitempty"`
}

// WeatherNow Basic slice data to 'results'.
type WeatherNow struct {
	LastUpdate string         `json:"last_update,omitempty"`
	Location   LocationStruct `json:"location,omitempty"`
	Now        NowStruct      `json:"now,omitempty"`
}

// LocationStruct Store data 'location'.
type LocationStruct struct {
	Country        string `json:"country,omitempty"`
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Path           string `json:"path,omitempty"`
	TimeZone       string `json:"timezone"`
	TimeZoneOffset string `json:"timezone_offset,omitempty"`
}

// NowStruct Store data 'now'.
type NowStruct struct {
	Code        string `json:"code,omitempty"`
	Temperature string `json:"temperature,omitempty"`
	Text        string `json:"text,omitempty"`
}

func main() {
	resp, err := http.Get(URL)
	if nil != err {
		fmt.Println("HTTP Get error.")
	}
	defer resp.Body.Close()

	ResponseData, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Println("ioutil.ReadAll error.")
	}
	fmt.Printf("%s\n", ResponseData)
	fmt.Println("Type of ResponseData: ", reflect.TypeOf(ResponseData))

	var Response ResponseJSON
	if err := json.Unmarshal([]byte(ResponseData), &Response); nil != err {
		fmt.Println("Unmarshal response weather now data error")
	}

	fmt.Println(Response.Results[0].LastUpdate)
	fmt.Println(Response.Results[0].Location)
	fmt.Println(Response.Results[0].Now)
	fmt.Println(Response.Results[0].Now.Code)
	fmt.Println(Response.Results[0].Now.Temperature)
	fmt.Println(Response.Results[0].Now.Text)
}
