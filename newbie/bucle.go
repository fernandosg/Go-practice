package main
import "fmt"
//In go language, not exist do while, or while, only for
func main(){
	alumnos:=1
	for alumnos<=3{
		fmt.Println(alumnos)
		alumnos=alumnos+1
	}
	test:=1
	// With this type of for is how i can  create something that is like bucle While
	for{
		if test > 3 {
			break
		}else{
			fmt.Println("To next ", test)
		}
		test=test+1
	}

	for calificaciones:=7;calificaciones<=9;calificaciones++ {
		fmt.Println(calificaciones)
	}
}