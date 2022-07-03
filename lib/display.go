package lib

import (
  "fmt"
//  "strings"
)

/*
func isEmpty(arr []any) bool {
  if (len(arr) > 0) {
    return true
  } else {
    return false
  }
}

func punctuate(input string) (output string) {

}*/

func Display(resp OxfordResponse) {
  fmt.Printf("%s\n", resp.ID)
}
