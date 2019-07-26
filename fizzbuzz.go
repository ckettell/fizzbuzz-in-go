package main

import (
  "fmt"
)

func IsDivisibleByThree(x int) (result bool) {
  if x%3 == 0 {
    result = true
  }
  return result
}

func IsDivisibleByFive(x int) (result bool) {
  if x%5 == 0 {
    result = true
  }
  return result
}

func NotDivisibleByFiveOrThree(x int) (result bool) {
  if x%3 != 0 && x%5 != 0 {
    result = true
  }
  return result
}

func main() {
  for i := 1; i < 100; i++{
    if IsDivisibleByThree(i){
      fmt.Printf("Fizz")
    }
    if IsDivisibleByFive(i){
      fmt.Printf("Buzz")
    }
    if NotDivisibleByFiveOrThree(i){
      fmt.Printf("%d", i)
    }
    fmt.Printf("\n")
  }

}
