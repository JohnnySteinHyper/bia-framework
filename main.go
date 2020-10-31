package main

import "github.com/urfave/cli"

var app = cli.NewApp()

func main() {
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

func info() {
  app.Name = "Simple pizza CLI"
  app.Usage = "An example CLI for ordering pizza's"
  app.Author = "Jeroenouw" 
  app.Version = "1.0.0"
}