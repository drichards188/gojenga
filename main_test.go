package main

import (
	"fmt"
	"gojenga/gjServer"
	"testing"
)

func TestOne(t *testing.T) {
	fmt.Println("testOne running")
	if gjServer.testingFunc() == true {
		t.Fatal("Throwing an error from testOne")
	} else {
		fmt.Println("test succeeded no errors")
	}
}
