package main

import (
	"fmt"
	"strconv"
	"strings"
)

var Answer int = 0

func main() {
	input:= "18623-26004,226779-293422,65855-88510,868-1423,248115026-248337139,903911-926580,97-121,67636417-67796062,24-47,6968-10197,193-242,3769-5052,5140337-5233474,2894097247-2894150301,979582-1016336,502-646,9132195-9191022,266-378,58-91,736828-868857,622792-694076,6767592127-6767717303,2920-3656,8811329-8931031,107384-147042,941220-969217,3-17,360063-562672,7979763615-7979843972,1890-2660,23170346-23308802"
	convertProductCodes(input)
	fmt.Println(Answer)
}

func convertProductCodes (input string){
	pairings := strings.Split(input, ",")
	for _, p := range pairings{
		pair:= strings.Split(p, "-")
		start, _:= strconv.Atoi(pair[0])
		end, _:= strconv.Atoi(pair[1])
		for i := start; i <= end; i++ {
			compare(i)
		}
	}

}

func compare(productCode int) {
	limit := len(strconv.Itoa(productCode)) / 2
	runes := []rune(strconv.Itoa(productCode))
	for i := limit; i >= 1; i-- {
		if len(strconv.Itoa(productCode))%i == 0 {
			newRunes := []rune{}
			for j:=0;j < i; j++ {
				newRunes = append(newRunes, runes[j])
			}
			s := string(newRunes)
			if concatenate(s, len(strconv.Itoa(productCode))/i) == strconv.Itoa(productCode) {
				Answer += productCode
				break
			}
		}
	}
}

func concatenate(partial string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += partial
	}
	return result
}
