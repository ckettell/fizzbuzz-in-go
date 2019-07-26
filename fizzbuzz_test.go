package main

import "testing"

// Testing that 3 should equal 0

func TestIsDivisibleByThree(t *testing.T){
  if IsDivisibleByThree(3) == true {
    t.Logf("Success!")
  }
}

func TestIsDivisibleByFive(t *testing.T){
  if IsDivisibleByFive(5) == true {
    t.Logf("Success!")
  }
}

func TestIsNotDivisibleByFive(t *testing.T){
  if IsDivisibleByFive(10) == false {
    t.Logf("Expected True")
  }
}

func TestIsNotDivisibleByThree(t *testing.T){
  if IsDivisibleByThree(5) == false {
    t.Logf("Expected True")
  }
}
