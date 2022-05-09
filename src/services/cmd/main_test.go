package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	fmt.Println("testOne running")
	if testingFunc() == true {
		t.Fatal("Throwing an error from testOne")
	} else {
		fmt.Println("test succeeded no errors")
	}
}
