package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type IntInString struct {
  value int 
  text string
}


var maps = []IntInString {
  {
    value: 1,
    text: "one",
  },
  {
    value: 2,
    text : "two",
  },
  {
    value: 3,
    text: "three",
  },
  {
    value: 4,
    text: "four",
  },
  {
    value: 5,
    text: "five",
  },
  {
    value: 6,
    text: "six",
  },
  {
    value: 7, 
    text: "seven",
  },
  {
    value: 8, 
    text: "eight",
  },
  {
    value: 9,
    text: "nine",
  },
}


func main()  {
  file, err := os.Open("input.txt")
  
  if err != nil {
    log.Fatal("Cannot read example file")
    file.Close()
    return
  }

  sum := 0
  scanner := bufio.NewScanner(file)



  for scanner.Scan() {
    line := scanner.Text()
    
    var fullIntString string
    var finds = make(map[int]IntInString)
    var keys = make([]int, 0);

    numberRegExp := regexp.MustCompile("[1-9]")
    intergersIndexes := numberRegExp.FindAllStringIndex(line,-1);

    for _,idx := range intergersIndexes {
      value := fmt.Sprintf("%c",line[idx[0]])
      intVal,_ := strconv.Atoi(value)

      finds[idx[0]] = maps[intVal-1]
      keys = append(keys, idx[0])
    }

    for _, number := range maps {
      // exp := fmt.Sprintf("(?m)%s", number)
      reg := regexp.MustCompile(number.text)
      indexes := reg.FindAllStringIndex(line,-1)
      if indexes == nil {
        continue
      }
      for _,matches := range indexes {
        fmt.Println(matches[0])
        fmt.Printf("Match found in %s : for %s at index %v\n", line, number.text, matches[0])
        finds[matches[0]] = number
        keys = append(keys, matches[0])
      }
    }

    sort.Ints(keys)
    
    for _,key := range keys {
      fullIntString = fmt.Sprintf("%s%v",fullIntString, finds[key].value)
      
    } 
    fmt.Printf("%s --> %s\n",line ,fullIntString)
    if len(fullIntString) <= 0 {
      continue
    } 

    result := fmt.Sprintf("%c%c",fullIntString[0], fullIntString[len(fullIntString)-1])

    fmt.Println(result)

    intVal, _ := strconv.Atoi(result)
    sum += intVal
  }

  fmt.Printf("Sum of all integer lines is : %v", sum)
}
