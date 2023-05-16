package main

// Q: what is the most common word (ignoring case) in sherlock.txt?
// Word frequency

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	filePath := `/Users/ryanbulloch/go/src/gofoundation/Day2/freq/sherlock.txt`
	// we're passing errors alongside values (var, err) as an error value ideally is "nil"
	file, err := os.Open(filePath)
	// reminder this is always the starting point
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	w, err := mostCommon(file, 3)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println(w)

	wordFrequency(file)
	//mapDemo()
}

// Change mostCommon to return the most common n words (e.g. func mostCommon(r io.Reader, n int) ([]string, error))
func mostCommon(r io.Reader, n int) (string, error) {
	// remember - freqs[string]int map
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}

	for k, v := range freqs {
		if v > 1000 {
			fmt.Println(k, v)
		}

	}

	return MaxWords(freqs)

}

// func mostCommon(r io.Reader) (string, error) {
// 	freqs, err := wordFrequency(r)
// 	if err != nil {
// 		return "", err
// 	}

// 	return MaxWords(freqs)
// }

// Note - variables will always run before the func main, as well as func init() {}
// "Who's on first?" -> [Who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func mapDemo() {
	//maps are kv pairs, like dictionary's in other languages
	var stocks map[string]float64 // string key is 'symbol' -> float64 value is 'price'
	symbol := "TTWO"
	price := stocks[symbol]
	// maps will always return something even if they're empty
	fmt.Printf("%s -> $%.2f\n", symbol, price)

	stocks = map[string]float64{
		symbol: 136.73,
		"AAPL": 172.35,
	}

	// but if you check with the 'ok' syntax it can tell you if something is empty/missing
	if price, ok := stocks[symbol]; ok {
		fmt.Printf("%s -> $%.2f\n", symbol, price)
	} else {
		fmt.Printf("%s not found \n", symbol)
	}

	// range over the stock key
	for k := range stocks {
		fmt.Println(k)
	}

	// range over the stock key and stock value
	for k, v := range stocks {
		fmt.Println(k, "->", v)
	}
	// range over just the value, need the underscore to access the primary keys
	for _, v := range stocks {
		fmt.Println(v)
	}

	delete(stocks, "AAPL")
	delete(stocks, "AAPL") // no panic if you delete something that doesn't exist
	fmt.Println(stocks)

}

// we'll consume the wordFrequenncy functions return of a map of strings and word count and use it here

// receive 'freqs' map from wordFrequency() -- string and number
func MaxWords(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}
	// set 2 variables, int and string
	maxNumber, maxWords := 0, ""
	// set 2 variables to range over the freqs data
	for word, count := range freqs {
		if count > maxNumber {
			maxNumber, maxWords = count, word
		}
	}

	return maxWords, nil
}

// we want the frequency / count of words with this function

// getting io reader as r, putting into a map with index of string, value of int, and an error if needed
func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // key is word -> value is count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // current line
		// for the value
		for _, w := range words {
			// ++ is incriment - without lower would look like adding to the map: freqs[w] as we're looping on each word and adding to the int/value of the map
			freqs[strings.ToLower(w)]++

		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	// returns the full map after the for loop has been ran with Scan()
	return freqs, nil

}
