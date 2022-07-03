package main

import (
  "os"

  "github.com/ayannw/define/lib"
)

func main() {

  definitions := lib.FetchDefinition(os.Args[1])

  lib.Display(definitions)

}
