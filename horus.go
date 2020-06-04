package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"time"
	//"sync"
	"strconv"
	"strings"

	"github.com/opesun/goquery"
	"gopkg.in/yaml.v2"
)

// Request http
type Request struct {
	DataString string
}

//yml設定
var config = &conf{}

type conf struct {
	MA1  float64 `yaml:"MA1"`
	MA2  float64 `yaml:"MA2"`
	MA3  float64 `yaml:"MA3"`
	MV1  float64 `yaml:"MV1"`
	MV2  float64 `yaml:"MV2"`
	MV3  float64 `yaml:"MV3"`
	MAX  int     `yaml:"MAX"`
	DAYS int     `yaml:"DAYS"`
	LOT  float64 `yaml:"LOT"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	return c
}

func main() {
	config.getConf()

	stockInfo := StockInfo()

	for i := 0; i < len(stockInfo); i++ {
		stock := strings.Split(stockInfo[i], ",")
		stockNumber := stock[0]
		stockName := stock[1]

		stockHistory := StockHistory(stockNumber)

		if len(stockHistory) > config.MAX {
			array1 := strings.Split(stockHistory, " ")
			MA := strings.Split(array1[4], ",")
			MV := strings.Split(array1[5], ",")
			MA1 := AverageCalculation(MA, config.MA1)
			MA2 := AverageCalculation(MA, config.MA2)
			MA3 := AverageCalculation(MA, config.MA3)

			MV1 := AverageCalculation(MV, config.MV1)
			MV2 := AverageCalculation(MV, config.MV2)
			MV3 := AverageCalculation(MV, config.MV3)

			if MA1 > MA2 && MA2 > MA3 && MV1 > MV2 && MV2 > MV3 && Condition3(MV) == true {
				fmt.Println(stockNumber + " " + stockName)
			}
		}
	}
}

// StockInfo 取得股票代碼
func StockInfo() []string {
	url := "https://stock.wespai.com/p/3752"
	p, _ := goquery.ParseUrl(url)
	tbody := p.Find("tbody").Text()
	tbody = strings.Replace(tbody, " ", "", -1)
	tbody = strings.Replace(tbody, "\n", ",", -1)
	tbody = strings.Replace(tbody, ",,,", " ", -1)
	tbody = strings.Replace(tbody, ",,", "", -1)
	array1 := strings.Split(tbody, " ")
	return array1
}

// StockHistory 取得個股歷史資料
func StockHistory(number string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://just2.entrust.com.tw/Z/ZC/ZCW/CZKC1.djbcd?a="+number+"&b=D&c="+strconv.Itoa(config.MAX), nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	res, err := client.Do(req)
	if err != nil {
		defer res.Body.Close()
		return ""
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

// AverageCalculation 條件1和2的計算
func AverageCalculation(data []string, average float64) float64 {
	value := 0.0
	if int(average) <= len(data) {
		for i := 0; i < int(average); i++ {
			v, _ := strconv.ParseFloat(data[len(data)-i-1], 64)
			value = value + v
		}
		value = value / average
	}
	return value
}

// Condition3 D天內，每日的成交量不低於X張
func Condition3(data []string) bool {
	value := 0
	if config.DAYS <= len(data) {
		for i := 0; i < config.DAYS; i++ {
			v, _ := strconv.ParseFloat(data[len(data)-i-1], 64)
			if v >= config.LOT {
				value = value + 1
			}
		}
		if value == 5 {
			return true
		}
		return false
	}
	return false
}
