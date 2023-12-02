package main

import (
    "regexp"
    "fmt"
)

func reg() {
    var re = regexp.MustCompile(`one`)
    var str = `kldjfsldjfodfonelskdjfmeoone`
    
    if len(re.FindStringIndex(str)) > 0 {
        fmt.Println(re.FindString(str),"found at index",re.FindStringIndex(str)[0])
    }
}


