/**
 * Created by Vim
 * User: lock
 * Date: 2018/10/25
 * Time: 14:11
 */
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ParseXmlData struct {
	XMLName     xml.Name     `xml:"yodaodict"`
	RawInput    string       `xml:"return-phrase"`
	CustomTrans CustomNode   `xml:"custom-translation"`
	WebTrans    WebTransList `xml:"yodao-web-dict"`
}

type CustomNode struct {
	Type        string        `xml:"type"`
	Translation []Translation `xml:"translation"`
}

type WebTransList struct {
	TransNode []WebTransNode `xml:"web-translation"`
}

type WebTransNode struct {
	Key   string      `xml:"key"`
	Trans []TransNode `xml:"trans"`
}

type TransNode struct {
	Value string `xml:"value,CDATA"`
}

type Translation struct {
	Content string `xml:"content,CDATA"`
}

func HttpGet(url string, ch chan []byte) {
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println("please enter ctrl+c ,request error:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("please enter ctrl+c ,io read error:", err)
	}
	ch <- body
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("please input the translation word !")
		fmt.Println("example: dict hello")
		os.Exit(0)
	}
	word := strings.Join(args, " ")
	requestUrl := fmt.Sprintf("https://dict.youdao.com//fsearch?client=deskdict&keyfrom=chrome.extension&q=%s&pos=-1&doctype=xml&xmlVersion=3.2&dogVersion=1.0&vendor=unknown&appVer=3.1.17.4208&le=eng", url.QueryEscape(word))
	ch := make(chan []byte)
	go HttpGet(requestUrl, ch)

	xmlObject := ParseXmlData{}
	xml.Unmarshal(<-ch, &xmlObject)
	//fmt.Println("您的输入:", strings.TrimSpace(xmlObject.RawInput))
	fmt.Println("\033[1;31m*******英汉翻译:*******\033[0m")
	for _, v := range xmlObject.CustomTrans.Translation {
		fmt.Println(strings.TrimSpace(v.Content))
	}
	fmt.Println("\033[1;31m*******网络释义:*******\033[0m")
	for _, v := range xmlObject.WebTrans.TransNode {
		key := strings.TrimSpace(v.Key)
		for _, vv := range v.Trans {
			value := strings.TrimSpace(vv.Value)
			fmt.Println(key, ":", value)
			break
		}
	}
}
