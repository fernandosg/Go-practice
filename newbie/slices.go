//Slices

//Create loop where i insert 10 values in slice
//Iterate over the slice and show every value
// Declare one slice of 5 string and init this slice with values String
// Print all values
//Take one slice from first and second index

package main
import(
	"fmt"
)
func main(){
	var numbers [] int 
	for i:=0;i<10;i++ {
		numbers=append(numbers,i*10)
	}

	for _,numero:= range numbers{
		fmt.Println(numero)
	}

	frut_array:=[5] string{"manzana","naranja","pera","sandia","aguacate"}
	for i,fruit:=range frut_array{
		fmt.Printf("Index: %d Fruit: %s\n",i,fruit)
	}

	//Tomar un slice de indice 1 y 2
	slice :=frut_array[1:3]
	for i,fruit:=range slice{
		fmt.Printf("Index: %d Frut: %s\n",i,fruit)
	}
}