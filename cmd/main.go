package main

import (
	"custom_step/logic"
	"flag"
)

var (
	phoneNum, pwd string
)

func main() {
	flag.StringVar(&phoneNum, "user", "", "user name")
	flag.StringVar(&pwd, "password", "", "password")
	flag.Parse()

	setter := logic.NewStepSetter()
	if err := setter.Do(phoneNum, pwd); err != nil {
		panic(err)
	}
}
