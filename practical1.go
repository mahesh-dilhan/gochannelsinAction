// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_BICYCLE_PRICE float32 = 10

func main() {
	var bikechannel = make(chan string)
	var bikewebsite []string = []string{"lazada", "shopee", "q0010"}

	for site := range bikewebsite {
		go checkfordeal(bikewebsite[site], bikechannel)
	}
	alert(bikechannel)
}

func checkfordeal(site string, c chan string) {

	for {
		time.Sleep(time.Second * 1)
		var deal = rand.Float32() * 50
		if deal < MAX_BICYCLE_PRICE {
			c <- site
			break
		}
	}
}

func alert(c chan string) {
	fmt.Println("we have a deal here", <-c)
}
