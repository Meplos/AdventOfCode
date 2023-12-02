package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Color struct {
	label string
	max   int
}

func NewColor(label string, max int) Color {
	return Color{label, max}
}

type Red = Color

var red Red = NewColor("red", 12)

type Green = Color

var green Green = NewColor("green", 13)

type Blue = Color

var blue Blue = NewColor("blue", 14)

var colorEnum = map[string]Color{
	red.label:   red,
	green.label: green,
	blue.label:  blue,
}

type Draw struct {
	number int
	color  Color
}

type Game struct {
	id         int
	maxByColor map[Color]int
  valid bool
}

func (g *Game) Power() int  {
  return g.maxByColor[red] * g.maxByColor[green] * g.maxByColor[blue] 
}


func GetColor(color string) (Color, error) {
  result := Color{"err", 0}
  err := errors.New("Invalid color")
	switch color {
	case red.label:
    result = red
    err = nil
	case green.label:
   result = green
  err = nil
	case blue.label:
    result = blue
    err = nil
	}
return result, err
}

func GetGameId(input string) int {

	find := strings.Split(input, " ")[1]

	id, _ := strconv.Atoi(find)
	return id
}

func ParseGame(input string) Game {
	gameArr := strings.Split(input, ": ")
	id := GetGameId(gameArr[0])
	drawsStrings := strings.Split(gameArr[1], "; ")
	maxs := make(map[Color]int, 0)
  valid := true
	for _, draws := range drawsStrings {

		sorts := strings.Split(draws, ", ")
    for _, draw := range sorts {

      split := strings.Split(draw," ")
      color,_ := GetColor(split[1])
      number,_ := strconv.Atoi(split[0])
      if number > color.max {
        valid = false
      }

      if number > maxs[color] {
        maxs[color] = number
      }


    }
	}


	return Game{
		id,
    maxs,
    valid,
	}
}

func main() {

  input,_ := os.Open("input.txt")
  scanner := bufio.NewScanner(input)
  
  var sumOfIds int
  var powers int

  for scanner.Scan() {
    line := scanner.Text()
    if(len(line) <= 0 ) { 
      break;
    }
    game := ParseGame(line)

    if game.valid {
      sumOfIds += game.id
    }

    powers += game.Power()

  }

  fmt.Printf("Sum of valid game is %v\n", sumOfIds)
  fmt.Printf("Sum of power :  %v\n", powers)

}
