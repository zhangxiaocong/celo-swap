
package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	utils "uniswap/utils"
)
type Market struct {
	Data struct {
		Pair struct {
			Token1Price string `json:"token1Price"`
		} `json:"pair"`
	} `json:"data"`
}

func main() {
	utils.Startredis()
	s := gocron.NewScheduler()
	s.Every(15).Seconds().Do(getPrice)
	//s.Every(10).Seconds().Do(showprice)
	result:=utils.GetValue("celo")
	fmt.Println("celo",result)
	sc := s.Start() // keep the channel
	<-sc
}
func showprice()  {
	utils.GetValue("celo")
}


func getPrice() string {
	client := &http.Client{}
	//查询价格
	//{"query":"{\n  pair(id:\"0xb460f9ae1fea4f77107146c1960bb1c978118816\") {\n    id\n    token0Price\n    token1Price\n  }\n}","variables":null}
	//var data = strings.NewReader(`{"query":"{\n bundle(id:1) {  celoPrice }\n}\n","variables":null}`)
	var data = strings.NewReader(`{"query":"{\n  pair(id:\"0xb460f9ae1fea4f77107146c1960bb1c978118816\") {\n     token1Price\n  }\n}"}`)
	//查询所有的交易对
	//var data = strings.NewReader(`{ "query": "{ubeswapFactories(first: 5) { id pairCount totalVolumeUSD totalVolumeCELO} tokens(first: 5) { id symbol name tradeVolume}}" }`)
	req, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/ubeswap/ubeswap", data)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
	market := Market{}
	err = json.Unmarshal([]byte(body), &market)
	if err != nil {
		fmt.Println(err)
	}
	price:=market.Data.Pair.Token1Price
	//save in redis
	utils.SaveValue(price)
	return price
}
