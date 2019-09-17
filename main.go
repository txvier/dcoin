/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/txvier/dcoin/base"
	"github.com/txvier/dcoin/cmd"
	_ "github.com/txvier/dcoin/logger"
	"os"
	"os/signal"
	"strings"
)

func AcceptDo(chs chan os.Signal) {
	base.PrintDcoinPrefix()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "quit" {
			chs <- os.Interrupt
			break
		} else {
			go func(line string) {
				args := strings.Split(line, " ")
				cmd.Execute(args)
				base.PrintDcoinPrefix()
			}(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		logrus.Info("the service start...")
	}
	chs := make(chan os.Signal)
	signal.Notify(chs, os.Interrupt, os.Kill)
	go AcceptDo(chs)
	for s := range chs {
		switch s {
		case os.Interrupt:
			fmt.Println("Bye")
			return
		case os.Kill:
			fmt.Print("kill")
			return
		}
	}
}
