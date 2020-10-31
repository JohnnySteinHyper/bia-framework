package main

import "github.com/urfave/cli"

var app = cli.NewApp()

func main() {
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

var pizza = []string{"Enjoy your pizza with some delicious"}

func info() {
  app.Name = "Go Programmers CLI"
  app.Usage = "A CLI for new Go Programmers"
  app.Author = "JohnnyStein" 
  app.Version = "1.0.0"
}