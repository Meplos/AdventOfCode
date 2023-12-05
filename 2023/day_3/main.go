package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const FILENAME = "./input.txt"
const DOT = "."

type ParseAction = func(string)

type Cache struct {
	value     string
	abscissas []int
}

func (c *Cache) Clear() {
	c.value = ""
	c.abscissas = make([]int, 0)
}

func (c *Cache) Empty() bool {
	return c.value == ""
}

func (c *Cache) Append(v string, x int) {
	c.value += v
	c.abscissas = append(c.abscissas, x)
}

type Position struct {
	x int
	y int
}

type Gear struct {
	value    string
	position Position
}

type Number struct {
	value string
	x     []int
	y     int
}

func (n *Number) IsNeighborOf(p Position) bool  {
  xmin := n.x[0] -1
  xmax := n.x[len(n.x)-1] +1
  ymin := n.y -1 
  ymax := n.y +1 

  return p.x >= xmin && p.x <= xmax && p.y >= ymin && p.y <= ymax
}

func (s *Number) GetNeighbors() []Position {
	neighbors := make([]Position, 0)
	xmin := s.x[0]
	xmax := s.x[len(s.x)-1]

	for y := s.y - 1; y <= s.y+1; y++ {
		for x := xmin - 1; x <= xmax+1; x++ {
      if IsCenter(x,y,s) {
        continue
      }
      position := Position{x,y}
      neighbors = append(neighbors, position)
		}

	}

	return neighbors
}

var cache Cache = Cache{"", make([]int, 0)}
var CurrentPosition Position = Position{0, 0}
var engine = make(map[Position]Gear, 0)
var numberMap = make([]Number, 0)

func main() {
  log.Printf("Day 3 : Gear Ratio")

	file, oserr := os.Open(FILENAME)
	if oserr != nil {
		log.Fatalf("Canno't open file %v", FILENAME)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
    if len(line) <= 0 {
continue
    }
		characters := strings.Split(line, "")
		log.Printf("%v", characters)

		for _, c := range characters {
			var action ParseAction
			if IsDot(c) {
				action = ParseDot
			} else if IsNumber(c) {
				action = ParseNumber
			} else {
        
				action = ParseEngine
			}

			Parse(action, c)
			MoveLeft()
		}
		MoveDown()
	}


  var counter int 

  for _, n := range numberMap {

    if !HasGearAsNeigbors(n) {
      continue
    }

    intValue,NaN := strconv.Atoi(n.value)
    if NaN != nil {
      log.Fatalf("Not a Number : %v", n)
    }
    counter += intValue
    
  }
  var power int
  for p,g := range engine {
    if g.value != "*" {
      continue
    }
    var neigbor = make([]Number, 0)
    
    for _,n := range numberMap {
      if n.IsNeighborOf(p) {
        neigbor = append(neigbor, n)
      }
    }

    if len(neigbor) < 2 {
      continue
    }

    var gearpower int = 1
    for _, n := range neigbor {
      value, _ := strconv.Atoi(n.value)
      gearpower *= value
    }
    power += gearpower
 
  }

  log.Printf("Result of sum : %v", counter)
  log.Printf("Result of powers : %v", power)
}

func MoveLeft() {
	CurrentPosition = Position{
		CurrentPosition.x + 1,
		CurrentPosition.y,
	}
}
func MoveDown() {
	CurrentPosition = Position{
		0,
		CurrentPosition.y + 1,
	}
}

func IsNumber(c string) bool {
  numbers := []string{"1","2","3","4","5","6","7","8","9","0"}
  return slices.Contains(numbers, c)
}

func IsDot(c string) bool {
	return c == DOT
}

func ParseDot(c string) {
	SaveNumber()
	cache.Clear()
	return
}

func SaveNumber() {
	if cache.Empty() {
		return
	}
	number := Number{cache.value, cache.abscissas, CurrentPosition.y}
	numberMap = append(numberMap, number)
}

func ParseNumber(c string) {
	cache.Append(c, CurrentPosition.x)
	return
}

func ParseEngine(c string) {
	gear := Gear{c, CurrentPosition}
	engine[CurrentPosition] = gear

	SaveNumber()
	cache.Clear()
	return
}

func Parse(action ParseAction, c string) {
	action(c)
}

func IsCenter(x int, y int, s *Number) bool {
	return s.y == y && slices.Contains(s.x, x) 
}

func HasGearAsNeigbors(n Number) bool {
  neigbors := n.GetNeighbors()
  for _, neigbor := range neigbors {
    _, ok := engine[neigbor]
    if ok {
      return true
    }
  }
	return false

}
