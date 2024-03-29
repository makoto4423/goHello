package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int // total visits per domain
		domains []string       // unique domain names
		total   int            // total visits to all domains
		lines   int            // number of parsed lines (for the error messages)
	)

	sum = make(map[string]int)

	// Scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		lines++

		// Parse the fields
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, lines)
			return
		}

		domain := fields[0]

		// Sum the total visits per domain
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}

		// Collect the unique domains
		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		// Keep track of total and per domain visits
		total += visits
		sum[domain] += visits
	}

	// Print the visits per domain
	sort.Strings(domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

	// Let's handle the error
	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}

func Mk() {
	// in := bufio.NewScanner(os.Stdin)
	// for in.Scan() {
	// 	fmt.Println(in.Text())
	// }
	// fmt.Printf("num")
	query := os.Args[1]

	rx := regexp.MustCompile("[^a-z]+")

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	if words[query] {
		fmt.Print(query + "get the word")
		return
	}
	fmt.Print("find nothing")

}
