package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	digits := 0
	format := "%d\n"
	total, _ := new(big.Float).SetPrec(128).SetString("0")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := new(big.Float).SetPrec(128).SetString(line)
		if strings.Contains(line, ".") {
			s := strings.Split(line, ".")
			if 2 != len(s) {
				log.Fatal("Invalid input value: " + line)
			}
			digits = len(s[1])
			format = "%." + fmt.Sprintf("%d", digits) + "f\n"
		}
		total = new(big.Float).SetPrec(128).Add(total, value)
	}
	fmt.Printf(format, total)
}
