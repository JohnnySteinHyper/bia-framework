package main

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
	"log"
)

var app = cli.NewApp()

func info() {
  app.Name = "Go Programmers CLI"
  app.Usage = "A CLI for new Go Programmers"
  app.Author = "JohnnyStein" 
  app.Version = "1.0.0"
}

var pizza = []string{"Enjoy your pizza with some delicious"}

func commands() {
  app.Commands = []cli.Command{
    {
      Name:    "peppers",
      Aliases: []string{"p"},
      Usage:   "Add peppers to your pizza",
      Action: func(c *cli.Context) { 
        pe := "peppers"
        peppers := append(pizza, pe)
        m := strings.Join(peppers, " ")
        fmt.Println(m)
      },
    },
    {
      Name:    "pineapple",
      Aliases: []string{"pa"},
      Usage:   "Add pineapple to your pizza",
      Action: func(c *cli.Context) { 
        pa := "pineapple"
        pineapple := append(pizza, pa)
        m := strings.Join(pineapple, " ")
        fmt.Println(m)
      },
    },
    {
      Name:    "cheese",
      Aliases: []string{"c"},
      Usage:   "Add cheese to your pizza",
      Action: func(c *cli.Context) { 
        ch := "cheese"
        cheese := append(pizza, ch)
        m := strings.Join(cheese, " ")
        fmt.Println(m)
      },
    },
  }
}

func main() {
	info()
	commands()

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

