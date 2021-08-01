package functions

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

/* 何の日かを表す構造体 */
type Day struct {
	/* タイトル */
	Title string

	/* 説明 */
	Description string
}

func GetDocument(url string) (document *goquery.Document, err error) {
	// Getリクエスト
	res, err := http.Get(url)
	if err != nil {
		log.Printf("HTTPリクエストエラー url: %v\n", url)
		return nil, err
	}
	defer res.Body.Close()

	// 読み取り
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("読み取りエラー\n")
		return nil, err
	}

	// 文字コード判定
	detector := chardet.NewTextDetector()
	detectedResult, err := detector.DetectBest(buf)
	if err != nil {
		log.Printf("文字コードの読み取りに失敗\n")
		return nil, err
	}

	// 文字コードの変換
	bReader := bytes.NewReader(buf)
	reader, err := charset.NewReaderLabel(detectedResult.Charset, bReader)

	// HTMLをパース
	document, err = goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Printf("HTMLのパースに失敗")
		return nil, err
	}

	return document, nil
}

/* 「今日は何の日」を取得する */
func GetToday() (day *Day, err error) {
	const url string = "https://kids.yahoo.co.jp/today/"

	document, err := GetDocument(url)
	if err != nil {
		return nil, err
	}

	dateinfoSelection := document.Find("div#contents > div#main > div#time > div#mainInner > div#dateinfo")
	dateDtlSelection := dateinfoSelection.Find("dl#dateDtl")

	title := dateDtlSelection.Find("dt > span").First().Text()
	description := dateDtlSelection.Find("dd").Text()

	return &Day{Title: title, Description: description}, nil
}
