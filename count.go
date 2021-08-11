package main

import "fmt"

func main(){
  n := 5
  arrival := []int{1,2,3,3,4}
  duration := []int{2,2,1,3,1}
  if n >= 1 && n <= 50 && len(arrival) == len(duration) {
    fmt.Println(maxEvents(arrival, duration,n))
  }  
}

func maxEvents(arrival []int, duration []int, max int) int {
  currentTime := arrival[0]
  canPresent := 0
  for i:=0; i < max; i++ {
    if arrival[i] >= 1 && arrival[i] <= 1000 && duration[i] >= 1 && duration[i] <= 1000 {
      if currentTime <= arrival[i] {
        currentTime += duration[i]
        canPresent+=1
      } else {
        continue
      }
    } else {
      continue
    }
  }
  return canPresent
}