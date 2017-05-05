package main
import "fmt"

func main() {
fmt.Println("\nValues")
fmt.Println("go"+"lang")
fmt.Println("1+1",1+1)
fmt.Println("7.0/3.0",7.0/3.0)
fmt.Println(true && false)
fmt.Println(true || false)
fmt.Println(!true && !false)

fmt.Println("\nVars")

var a string="initialize vars"
fmt.Println(a)

var b,c int=1,2
fmt.Println(b,c)

var d bool 
fmt.Println(d)

var e int
fmt.Println(e)

f := "short form for initializing variables"
fmt.Println(f)

}
