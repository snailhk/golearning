package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func grade(score int) string {
	var g string
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("wrong score: %d", score))
	}
	return g
}
func convertToBin(n int) string {
	var result string
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	//const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Printf("%s \n", contents)
	//}

	grade(59)
	grade(60)
	grade(70)
	grade(80)
	grade(90)
	//grade(101)
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(5000),
		convertToBin(199),
	)
	printFile("abc.txt")
}
