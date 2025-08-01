package lasagna

const (
	defaultPrepTime  = 2
	noodlesQuant     = 50
	sauceQuant       = 0.2
)

func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = defaultPrepTime
    }

    return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
    var noodles int = 0
    var sauce float64 = 0.0
    
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += noodlesQuant
        } else if layers[i] == "sauce"  {
            sauce += sauceQuant
    	}
    }

    return noodles, sauce
}

func AddSecretIngredient(friendsList []string , myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
} 

func ScaleRecipe(quantities []float64, portions int) []float64 {
    buffer := make([]float64, len(quantities))

    for i := 0; i < len(quantities); i++ {
        buffer[i] = quantities[i] * float64(portions) / 2.0
    }

    return buffer
}