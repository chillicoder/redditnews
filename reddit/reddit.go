package main

import (
  "fmt"
  "github.com/chillicoder/redditnews"
  "log"
)

func main() {
  items, err := redditnews.Get("elixir")
  if err != nil {
    log.Fatal(err)
  }

  for _, item := range items {
    fmt.Println(item)
  }
}

