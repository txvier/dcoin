package main

import (
	"bufio"
	"fmt"
	"github.com/txvier/dcoin/cmd"
	"os"
	"os/signal"
	"strings"
)

func AcceptDo(chs chan os.Signal) {
	r := bufio.NewReader(os.Stdin)
	//r.read
	fmt.Print("dcoin> ")
	for {
		if line, _, err := r.ReadLine(); err != nil {
			panic(err)
		} else if string(line) == "quit" {
			chs <- os.Interrupt
			break
			//close(chs)
		} else {
			go func(line string) {
				args := strings.Split(line, " ")
				cmd.Execute(args)
				fmt.Print("dcoin> ")
			}(string(line))
		}
	}
}

func main() {
	chs := make(chan os.Signal)
	signal.Notify(chs, os.Interrupt, os.Kill)
	go AcceptDo(chs)
	for s := range chs {
		switch s {
		case os.Interrupt:
			fmt.Println("Bye")
			os.Exit(0)
		case os.Kill:
			fmt.Print("kill")
			os.Exit(0)
		}
	}
}
