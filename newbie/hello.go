package main 
import "fmt"
import "math"
func main(){
	//This is the way of how concatenate two strings in Go
	const pet string = "gopher"
	const age = 500
	fmt.Println("go "+"lang"+" "+pet)	
	const value =3e20/age
	fmt.Println(int64(value))
	fmt.Println(math.Sin(age))
	//This is the way of how add and substract
	fmt.Println("5+5 =",5+5)

	//This is the way of how divide 7 from 2
	fmt.Println("7.5/3.2 =",5.5/3.2)
	var a string ="init"
	fmt.Println(a)
	var nombre,apellido string  ="fernando","segura"
	fmt.Println("My name is "+nombre+" "+apellido)
	//This is the way of how declare a boolean variable, this default value is false
	var d bool
	fmt.Println(d)
	var f = 5
	fmt.Println(f)
	h:="Short"
	fmt.Println(h)

	var h string="testeo"
	fmt.Println(h)
}