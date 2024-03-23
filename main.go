package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
	"strings"
)

type Portfolio struct {
	cryptos []Crypto
}

type Crypto struct {
	name         string
	total        float64
	percentage   float64
	amount       float64
	currentPrice float64
}

func main() {
	fmt.Println()
	fmt.Println("             Portfolio Project               ")
	fmt.Println()
	fmt.Println("----------------------------------------------------------------------------")

	// Create Portfolio
	portfolio := Portfolio{}

	// Create accumulator
	var accumulator = 0.0

	//	Create List for totals
	var answer = true

	// redis connection
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:2424",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//	-- while(!answer) --
	for answer {
		crypto := Crypto{}

		//	Read and store name
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter name of your crpyto: ")
		name, _ := reader.ReadString('\n')
		crypto.name = strings.TrimSpace(name)

		//	Read and store amount
		fmt.Print("Enter amount of your crpyto: ")
		amount, err := reader.ReadString('\n')
		if err != nil {
			panic("Error reading amount of crypto")
		}
		//fmt.Println("Value of amount: ", amount)

		amountF, err := strconv.ParseFloat(strings.TrimSpace(amount), 64)
		if err != nil {
			panic("Error while parsing \"amount\" from float to string")
		}
		crypto.amount = amountF

		//	Read and store current price
		fmt.Print("Enter the current price of your crpyto: ")
		currentPrice, _ := reader.ReadString('\n')
		currentPriceF, err := strconv.ParseFloat(strings.TrimSpace(currentPrice), 64)
		if err != nil {
			panic("Error while parsing \"currentPrice\" from float to string")
		}
		crypto.currentPrice = currentPriceF

		//	Calculate total: (amount * current price)
		crypto.total = amountF * currentPriceF

		//	Add total to accumulator
		accumulator += crypto.total

		// store in portfolio
		portfolio.cryptos = append(portfolio.cryptos, crypto)

		fmt.Println()
		fmt.Println("Crypto Portfolio Data")
		fmt.Println("----------------------")
		fmt.Println("Grand Total: $", accumulator)
		fmt.Println()
		fmt.Println("Name: ", crypto.name)
		fmt.Println("Total: $", crypto.total)
		fmt.Println("Percentage: ", (crypto.total/accumulator)*100)
		fmt.Println("Amount: ", crypto.amount)
		fmt.Println("Current Price: $", crypto.currentPrice)
		fmt.Println()

		// update answer
		//	Read and store current price
		fmt.Print("Do you want to enter another Crypto? Type \"yes\" or \"no\": ")
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)

		//	Store Data:
		//	1)
		//	Add to redis sorted set
		//	Key: nameOfCrypto
		//	Val: total
		var nameAndTotal redis.Z
		nameAndTotal.Member = crypto.name
		nameAndTotal.Score = crypto.total

		err = rdb.ZAdd(ctx, "totals", nameAndTotal).Err()
		if err != nil {
			panic(err)
		}

		//	2)
		//	store portfolio in redis hash
		//	Key: nameOfCrypto
		//	Val: JSON for portfolio struct
		//

		//	-- Do you want to enter another? = answer --
		if strings.Compare(userAnswer, "no") == 0 {
			answer = false
		}

	}

	//	Get the reverse order of sorted set as list
	//
	//	Use returned list of cryptos to get the portfolio
	//struct from redis hash
	//
	//
	//	Calculate %: (portfolio.total % accumulator) x 100
	//	Round to 2 decimal places
	//
	//	update the % for the portfolio
	//
	//
	//	Portfolio struct is complete
	//	add to list
	//	iterate through list
	//
	//	Print to console
	//
	//
	//
	//	Upgrade:
	//	open file
	//	print row for each portfolio
	//	close file
	//	save to C drive

}
