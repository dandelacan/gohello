package main

import "fmt"

func main(){
	helloArray := [2]string{}
	helloArray[0]="hello"
	helloArray[1]="BJSS"
	for _, s := range helloArray{
		fmt.Print(s, " ")
	}
}