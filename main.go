package main

import "fmt"

// "go_project/routes"
// "go_project/static"
func main() {
	
	// static.MovieSeeder()
	// routes.Router()	
	name := "haile dereje"
	var ptr *string = &name

	*ptr = "solomon"
	fmt.Println(name)

}