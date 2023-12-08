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
      count ++
    }
  }

 return count
}

func (c *Card) GetWinningMatches() []int  {
  result := make([]int, 0)
  for _, n := range c.draw {
    if slices.Contains(c.winning, n) {
      result = append(result, n)
    }
  }
  return result
}

func (c *Card) GetPoints() int  { 
    points := c.GetNumberOfWinning()
    if points > 1 {
      points = int(math.Pow(2,float64(points -1))) 
    }
  return points
}

func (c *Card) Copy() Card  {
  return Card{
    c.winning, 
    c.draw,
  }
}

var nbOfCopy = make(map[int]int, 0)
func main()  {
  log.Printf("Day 4 : Scratchcard")

  file, filenotfound := os.Open(FILENAME)
  if filenotfound != nil {
    log.Panicf("File %v not found", FILENAME)
    return
  }

  scanner := bufio.NewScanner(file)

  total := 0

  var totalPart2 int
  var i int = 1;
  for scanner.Scan() {
    line := scanner.Text()
    if len(line)  <= 0 {
      continue
    }
    card := parseCard(line)
    nbOfCopy[i]++
    nbOfMatches := card.GetNumberOfWinning()

    for y := 1; y <= nbOfMatches; y++ {
      idx := i+y
      nbOfCopy[idx] += nbOfCopy[i]
    }

    points := card.GetPoints()
    totalPart2 += nbOfCopy[i]
    total += points
    i++
  }



  log.Printf("Total of part 1 is : %v", total)
  log.Printf("Total of part 2 is : %v", totalPart2)
  
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
