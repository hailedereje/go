package main

import (

"go_project/routes"
"go_project/static"

)
func main() {
	
	static.MovieSeeder()
	routes.Router()	

}