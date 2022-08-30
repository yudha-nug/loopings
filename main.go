package main
import "fmt"

func main() {
	fmt.Println("--------------------")
	for i := 0; i < 3; i++ {
		fmt.Println("Angka ", i)
	}

	fmt.Println("--------------------")
	var i int = 0
	for i < 3 {
		fmt.Println("Angka ", i)
		i++
	}	

	fmt.Println("--------------------")	
	var j int = 0
	for {
		fmt.Println("Angka ", j)
		j++
		if j==3 {
			break
		}
	}
}
