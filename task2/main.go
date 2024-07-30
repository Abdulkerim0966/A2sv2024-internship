package main

import (
	"fmt"
	"regexp"
	"strings"
)

func freCounter(s string) map[string]int{
	re := regexp.MustCompile(`[^\w\s]`)
	s =re.ReplaceAllString(s,"")
	wordCount :=make(map[string]int)

	words := strings.Fields(s)
    
    // Count the occurrences of each word
    for _, word := range words {
        wordCount[word]++
	}
	return wordCount
}
func palindromeChecker(s string)bool{
	re := regexp.MustCompile(`[^\w]`)
	s =re.ReplaceAllString(s,"")
	i:=0
	j:=len(s)-1
	for i<j{
		if s[i] !=s[j] {
			return false
		}
		i+=1
		j-=1
	}
	return true
}   

func main(){
	fmt.Println(freCounter("this is the second task! for the project"))
	fmt.Println(palindromeChecker("fas@tt saf*&"))
	

}