package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"strings"
//)
//
//func main() {
//	coin:="CELO"
//	coinprice:=getPrice(coin)
//	fmt.Println("price is ",coinprice)
//}
//
//
//func getPrice(coin string)string{
//	client := &http.Client{}
//	//var data = strings.NewReader(`"query":"{pair(id:"0xb460f9ae1fea4f77107146c1960bb1c978118816") { token1Price} }"}`)
//	var data = strings.NewReader(`"query":"{tokens(first: 1) { id symbol name tradeVolume}}"`)
//	req, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2",data)
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//	bodyText, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%s\n", bodyText)
//	return string(bodyText)
//}