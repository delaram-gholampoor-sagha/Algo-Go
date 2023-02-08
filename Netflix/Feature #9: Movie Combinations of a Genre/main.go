package main
import ( 
  "fmt"
  "strings" )


func letterCombinations(genres []string) []string {
    //Return empty array if input is empty
    if len(genres) == 0{
        var temp []string
        return temp
    }
    // Mapping the genres to their corresponding movies
    movies := map[string][]string{
        "Family" : {"Frozen","Kung fu Panda", "Ice Age"}, 
        "Action" : {"Iron Man","Wonder Woman","Avengers"},
        "Fantasy" : {"Jumangi", "Lion King", "Tarzan"},
        "Comedy" : {"Coco", "The Croods", "Vivi","Pets"},
        "Horror" : {"Oculus", "Sinister","Insidious","Annebelle"} }
    // Declaring the combinations variable
    var combinations []string
    // Declaring the backtrack function
    var backtrack func(int,[]string)
   
   // Use backtrack function to generate all possible combinations
    backtrack = func(index int, path []string){
        // If the length of path and genres is same, 
        // we have a complete combinations
        if len(path) == len(genres) {
            combinations = append(combinations, strings.Join(path, ""))
            return
        }
        // Using the index and genres[index], get the list of movies
        possible_movies := movies[string(genres[index])]
        // Loop through the movies and generate combinations
        for i := 0; i < len(possible_movies); i++{
            // Add the movie to our current path
            path = append(path, string(possible_movies[i])+";")
            // Move on to the next genre
            backtrack(index + 1, path)
            // Before moving onto the next movie, 
            // backtrack by removing the movie from the path
            if len(path) > 0{
                // RemoveIndex is used to remove the element from the path
                path = RemoveIndex(path, len(path)-1)
            }    
        }    
    }   
    // Initiate backtracking with an empty path and starting index of 0
    var path []string
    backtrack(0, path)
    return combinations
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {

    //Example 1
    genres := []string{"Action"}
    combinations := letterCombinations(genres)
    output := "\""+strings.Join(combinations, `","`) + `"`
    fmt.Println("Output 1:")
    fmt.Println(output)

    //Example 2
    genres = []string{"Family", "Action"}
    combinations = letterCombinations(genres)
    output = "\""+strings.Join(combinations, `","`) + `"`
    fmt.Println("Output 2:")
    fmt.Println(output)

    //Example 3
    genres = []string{"Horror", "Comedy"}
    combinations = letterCombinations(genres)
    output = "\""+strings.Join(combinations, `","`) + `"`
    fmt.Println("Output 3:")
    fmt.Println(output)

    //Example 4
    genres = []string{"Horror", "Fantasy", "Comedy", "Family"}
    combinations = letterCombinations(genres)
    output = "\""+strings.Join(combinations, `","`) + `"`
    fmt.Println("Output 4:")
    fmt.Println(output)
}