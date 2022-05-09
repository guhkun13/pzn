package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	// "strings"
)

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	
	for {
		fmt.Print("Input number -> ")
		input, _ := reader.ReadString('\n')
		st := time.Now()
		
		input = strings.ReplaceAll(input, "\n","")
		endLoop, err := strconv.Atoi(input)
		
		if err != nil {
			fmt.Println("err masbro", err)
		}
		
		PrintFizzBuzz(endLoop)
		// endtime 
		fmt.Println("Execution time = ", time.Since(st))
	}
}


func PrintFizzBuzz(endLoop int) {
	fmt.Println("start Fizz Buzz masbrooo end at " , endLoop)
	
	fmt.Printf("Rules \n 1. Print Fizz if idx mod 3 = 0 \n 2. Print Buzz if idx mod 5 = 0 \n 3. Print Fizz Buzz if idx mod 3 and 5 = 0 \n --- Start --- \n")
	
	var results []string
	
	for idx := 1; idx <= endLoop; idx ++ {		
		fzText := GetFizzBuzzText(idx)
		fmt.Printf("idx = %d , txt = %s \n", idx, fzText)
		results = append(results, fzText)
	}
	fmt.Println("END RESULT")
	fmt.Println(results)
	
	fmt.Println("=========================================")
}

func GetFizzBuzzText(n int) string{
	
	fz := CheckFizzBuzz(n)
	
	var result string
	
	if fz.IsFizz {
		result = "Fizz"
		} else if fz.IsBuzz {
			result = "Buzz"
			} 
			
			if fz.IsFizz && fz.IsBuzz {
				result = "FizzBuzz"
			}
			
			if result == "" {
				result = strconv.Itoa(n)
			}
			
			return result
		}
		
		type FizzBuzz struct {
			IsFizz bool `def:"false"`
			IsBuzz bool `def:"false"`
			IsFizzBuzz bool `def:"false"`
		}
		
		func CheckFizzBuzz(n int) FizzBuzz {
			
			fz := new(FizzBuzz)
			
			if n%3 == 0 {
				fz.IsFizz = true
				}	
				
				if n%5 == 0 {
					fz.IsBuzz = true
				}
				
				if fz.IsFizz && fz.IsBuzz {
					fz.IsFizzBuzz = true
				}
				
				return *fz
			}