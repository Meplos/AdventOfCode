package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const FILENAME = "input.txt"

type Card struct {
  winning []int
  draw []int
}

func (c *Card) GetNumberOfWinning() int  {
  count := 0
  for _, w := range c.winning {
    if slices.Contains(c.draw, w) {
      log.Printf("%v contains %v", c.draw, w)
      count ++
    }
  }

 return count
}

func main()  {
  log.Printf("Day 4 : Scratchcard")

  file, filenotfound := os.Open(FILENAME)
  if filenotfound != nil {
    log.Panicf("File %v not found", FILENAME)
    return
  }

  scanner := bufio.NewScanner(file)

  total := 0
  for scanner.Scan() {
    line := scanner.Text()
    if len(line)  <= 0 {
      continue
    }

    card := parseCard(line)
    
    points := card.GetNumberOfWinning()
    if points > 1 {
      points = int(math.Pow(2,float64(points -1))) 
    }
    
    log.Printf("Card = %v", points)

    total += points

  }

  log.Printf("Total is : %v", total)
  
}


func parseCard(card string) Card {
  allNumbers := strings.Split(strings.Split(card,": ")[1]," | ")
  winningStr := strings.Split(allNumbers[0], " ")
  drawStr := strings.Split(allNumbers[1], " ")

  winning := StringsToInts(winningStr)
  draw := StringsToInts(drawStr)
  
  return Card{
    winning,
    draw,
  }

}

func StringsToInts(str []string) []int {
  res := make([]int, 0)

  for _, c := range str {
    if (c == "") { 
      continue
    }
    i,_ := strconv.Atoi(c)
    res = append(res, i )
  }

  return res
}
