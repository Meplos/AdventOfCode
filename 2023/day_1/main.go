package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main()  {
  file, err := os.Open("input.txt")
  
  if err != nil {
    log.Fatal("Cannot read input file")
    file.Close()
    return
  }

  sum := 0
  scanner := bufio.NewScanner(file)
  isCharRegExp := regexp.MustCompile("[a-zA-z]*")
  for scanner.Scan() {
    line := scanner.Text()
    fullIntString := isCharRegExp.ReplaceAllString(line, "")
  
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
