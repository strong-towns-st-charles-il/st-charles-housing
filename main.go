package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type House struct {
	PropertyURL    string  // 0
	PropertyPrice  float64 // 1
	SoldDate       string  // 2
	LivingArea     float64 // 3
	LivingAreaUnit string  // 4
	LotSize        float64 // 5
	LotSizeUnit    string  // 6
	PricePerSqft   float64 // 7
	BrokerName     string  // 8
	BrokerAgent    string  // 9
	StreetAddress  string  // 10
	City           string  // 11
	State          string  // 12
	Zip            string  // 13
	Country        string  // 14
	Bedrooms       int     // 15
	Bathrooms      float64 // 16
}

func main() {
	fileName := flag.String("f", "sold.csv", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var houses []House

	for i := 1; i < len(input); i++ {
		houses = append(houses, parse(input[i]))
	}
	var affordablePrice float64 = 231528
	var affordableHouses []House
	var total float64
	var atotal float64
	var sqft float64
	var asqft float64
	var costArea float64
	var acostArea float64
	for _, h := range houses {
		total += h.PropertyPrice
		sqft += h.LivingArea
		costArea += h.PricePerSqft
		if h.PropertyPrice < affordablePrice {
			affordableHouses = append(affordableHouses, h)
			atotal += h.PropertyPrice
			asqft += h.LivingArea
			acostArea += h.PricePerSqft
		}
	}

	n := len(houses)
	m := len(affordableHouses)

	fmt.Printf("total number of houses sold in 2022: %v\n", n)
	fmt.Printf("total number of affordable: %v\n", m)
	fmt.Printf("affordable house faction: %.2f\n", float64(m)/float64(n))
	fmt.Printf("\naverage house price: $%.2f\n", total/float64(n))
	fmt.Printf("average affordable house price: $%.2f\n", atotal/float64(m))

	bubbleSort(houses)
	bubbleSort(affordableHouses)

	fmt.Printf("\nmedian house price: $%v\n", houses[n/2].PropertyPrice)
	fmt.Printf("median affordable house price: $%v\n", affordableHouses[m/2].PropertyPrice)

	avgArea := sqft / float64(n)
	avgAreaA := asqft / float64(m)
	fmt.Printf("\naverage house sqft: %.2f\n", avgArea)
	fmt.Printf("average affordable house sqft: %.2f\n", avgAreaA)
	fmt.Printf("%.2f\n", avgArea/avgAreaA)

	avgCost := costArea / float64(n)
	avgCostA := costArea / float64(m)

	fmt.Printf("\naverage house cost/sqft: %.2f\n", avgCost)
	fmt.Printf("average affordable house cost/sqft: %.2f\n", avgCostA)
	fmt.Printf("%.2f\n", avgCost/avgCostA)
}

func read(path string) ([]string, error) {
	var input []string
	dat, err := os.Open("data/" + path)
	if err != nil {
		return nil, err
	}
	defer dat.Close()

	s := bufio.NewScanner(dat)

	for s.Scan() {
		input = append(input, s.Text())
	}

	return input, nil
}

func swap(xp, yp *House) {
	tmp := *xp
	*xp = *yp
	*yp = tmp
}

func bubbleSort(hs []House) {
	for i := 0; i < len(hs); i++ { //
		for j := 0; j < len(hs)-1; j++ {
			if hs[j].PropertyPrice > hs[j+1].PropertyPrice {
				swap(&hs[j], &hs[j+1])
			}
		}
	}
}

func parse(s string) House {
	h := strings.Split(s, ",")

	price, err := strconv.ParseFloat(h[1], 64)
	if err != nil {
		price = 0
	}

	area, err := strconv.ParseFloat(h[3], 64)
	if err != nil {
		area = 0
	}

	lot, err := strconv.ParseFloat(h[5], 64)
	if err != nil {
		lot = 0
	}

	ppsqft, err := strconv.ParseFloat(h[7], 64)
	if err != nil {
		ppsqft = 0
	}

	bed, err := strconv.Atoi(h[15])
	if err != nil {
		bed = 0
	}

	bath, err := strconv.ParseFloat(h[16], 64)
	if err != nil {
		bath = 0
	}

	return House{h[0], price, h[2], area, h[4], lot, h[6], ppsqft, h[8], h[9], h[10], h[11], h[12], h[10], h[11], bed, bath}
}
