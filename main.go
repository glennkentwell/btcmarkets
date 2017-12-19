package main

import (
	"fmt"

	"github.com/glennkentwell/btcmarketsgo"
	ccg "github.com/glennkentwell/cryptoclientgo"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

var client *btcmarketsgo.BTCMarketsClient

func init() {

	var err error
	client, err = btcmarketsgo.NewDefaultClient(btcmarketsgo.GetKeys("api.secret"))
	log.SetLevel(log.ErrorLevel)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//got, err := client.GetOrderBook("BTC", "AUD")
	//got, err := btcmarketsgo.BTCMarketsClient{}.GetOrderBook("ETH", "BTC")
	//log.Info("Open orders output:")
	// got, err := client.OrdersDetails(765500586)
	//got, err := client.GetOpenOrders("BCH", "AUD")
	// print(got, err)

	//Ticker example
	/*quit := make(chan bool)
	client.Ticker(func(tr btcmarketsgo.TickResponse, err error) {
		fmt.Printf("%+v\n", tr)
	}, time.Second, quit)
	log.Info("quiting after 50 seconds")
	time.Sleep(time.Second * 5 * 10)
	quit <- true
	log.Info("quit")*/
	got, err := client.GetBalances()
	if err != nil {
		fmt.Println("Ere be erraah %v", err)
		return
	}
	var total float64
	for _,v := range(got) {
		t := ccg.ConvertToFloat(v.TotalBalance)
		if (t > 0) {
			tr, err := client.Tick("AUD", v.Currency)
			if err != nil {
				fmt.Println("%v", err)
			}
			price := tr.LastPrice
			fmt.Printf("%s: %f @ $%f = $%f\n", v.Currency, t, price, t * price)
			total += t * price
		}
	}
	fmt.Printf("Total: $%f", total)
}

func print(got interface{}, err error) {
	if err != nil {
		fmt.Println(err)
	}
	config := spew.NewDefaultConfig()
	config.Indent = "\t"
	config.Dump(got)
}
