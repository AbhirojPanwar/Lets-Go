package main

import "fmt"

func main() {

i := 1

for i<=10{
fmt.Println(i)
i=i+1
}	

for j := 7;j<=9;j++{
	fmt.Println(j)
}

for{
	fmt.Println("Inf 1")
	fmt.Println("Inf 2")
	break
}

}