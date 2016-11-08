package main

import (
  "io"
  "fmt"
  "os"

  "golang.org/x/net/websocket"
)

func main() {
  server := os.Args[1]

  fmt.Println("Connecting to", server)
  ws, err := websocket.Dial(server, "", "http://localhost")

  if err != nil {
    fmt.Println(err)
    return
  }

  go func() {
    for _, err := io.Copy(os.Stdout, ws); err != nil; {}
    fmt.Println(err)
    os.Exit(0)
  } ()

  for _, err := io.Copy(ws, os.Stdin); err != nil; {}
  fmt.Println(err)
  os.Exit(0)
}
