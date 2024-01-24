package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating Single Instance Now!!!")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single Instance already created!")
	}

	return singleInstance
}
