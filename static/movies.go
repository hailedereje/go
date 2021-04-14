package static

import (
	// "fmt"
	"go_project/models"
)


var Movies [] models.Movie

func MovieSeeder() {
	movie_1 := models.Movie{ID:"12",Isbn:"ha-123",Title: "Haile-the-greate",
							Director:&models.Director{FirstName: "haile",LastName: "dereje"}}
	movie_2 := models.Movie{ID:"13",Isbn:"ha-124",Title: "Haile-the-hilarious",
							Director:&models.Director{FirstName: "Tirsit",LastName: "dereje"}}
	
	Movies = append(Movies, movie_1,movie_2)

	// for _,k := range Movies{
	// 	fmt.Println(k)
	// }
}
/*
ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
*/