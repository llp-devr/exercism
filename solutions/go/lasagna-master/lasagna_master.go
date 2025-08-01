package lasagna

/*
import (
    "fmt"
)
*/
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }

    return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
    var noodles int = 0
    var sauce float64 = 0.0
    
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 50
        } else if layers[i] == "sauce"  {
            sauce += 0.2
        }
    }

    return noodles, sauce
}

func AddSecretIngredient(friendsList []string , myList []string) {
    friendsListIdx := len(friendsList) -1
    myListIdx := len(myList) -1
    myList[myListIdx] = friendsList[friendsListIdx]

} 

func ScaleRecipe(quantities []float64, portions int) []float64 {
    buffer := []float64{}

    for i := 0; i < len(quantities); i++ {
        newQuantitie := quantities[i] * float64(portions) / 2.0
        buffer = append(buffer, newQuantitie)
    }


    return buffer
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
