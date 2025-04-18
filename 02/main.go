package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

const (
	RUB = "RUB"
	EUR = "EUR"
	USD = "USD"
)

type Bank map[[2]string][2]float64

type QueueElement struct {
	Currency  string
	Amount    float64
	UsedBanks map[int]struct{}
}

func Ozon02() {
	var t int // количество наборов входных жанных
	fmt.Fscan(in, &t)

	for range t { // итерация по наборам
		var banks = make([]Bank, 3)
		for i := range 3 {
			bank := Bank{
				{RUB, USD}: {0, 0},
				{RUB, EUR}: {0, 0},
				{USD, RUB}: {0, 0},
				{USD, EUR}: {0, 0},
				{EUR, RUB}: {0, 0},
				{EUR, USD}: {0, 0},
			}

			var x, y float64
			fmt.Fscan(in, &x, &y)
			bank[[2]string{RUB, USD}] = [2]float64{x, y}
			fmt.Fscan(in, &x, &y)
			bank[[2]string{RUB, EUR}] = [2]float64{x, y}
			fmt.Fscan(in, &x, &y)
			bank[[2]string{USD, RUB}] = [2]float64{x, y}
			fmt.Fscan(in, &x, &y)
			bank[[2]string{USD, EUR}] = [2]float64{x, y}
			fmt.Fscan(in, &x, &y)
			bank[[2]string{EUR, RUB}] = [2]float64{x, y}
			fmt.Fscan(in, &x, &y)
			bank[[2]string{EUR, USD}] = [2]float64{x, y}

			banks[i] = bank
		}

		assa(banks)
	}
	defer out.Flush()
}

func main() {
	in = bufio.NewReader(strings.NewReader(aaa))
	Ozon02()
}

func assa(banks []Bank) {

	maxUSD := 0.0
	queue := list.New()
	queue.PushBack(QueueElement{Currency: "RUB", Amount: 1.0, UsedBanks: map[int]struct{}{}})

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		current := element.Value.(QueueElement)
		currentCur := current.Currency
		currentAmt := current.Amount
		used := current.UsedBanks

		if currentCur == "USD" {
			if currentAmt > maxUSD {
				maxUSD = currentAmt
			}
		}

		if len(used) >= 3 {
			continue
		}

		for bankIdx := range 3 {
			if _, exists := used[bankIdx]; exists {
				continue
			}
			bank := banks[bankIdx]
			for _, toCur := range []string{"USD", "EUR", "RUB"} {
				if toCur == currentCur {
					continue
				}
				key := [2]string{currentCur, toCur}
				if rates, ok := bank[key]; ok {
					x, y := rates[0], rates[1]
					newAmt := (currentAmt * y) / x
					newUsed := make(map[int]struct{})
					for k := range used {
						newUsed[k] = struct{}{}
					}
					for k := range used {
						newUsed[k] = struct{}{}
					}
					newUsed[bankIdx] = struct{}{}
					queue.PushBack(QueueElement{Currency: toCur, Amount: newAmt, UsedBanks: newUsed})
				}
			}
		}
	}
	fmt.Fprintf(out, "%g\n", maxUSD)
}
