package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(string []string,avgtime int) int{
    if avgtime == 0{
        avgtime = 2
    }
    return avgtime * len(string)
}
// TODO: define the 'Quantities()' function
func Quantities(string []string) (int,float64){
    noodles := 0
    sauce := 0.0
    for i := 0; i < len(string); i++{
        if string[i] == "noodles"{
            noodles += 50
        }
        if string[i] == "sauce"{
            sauce += 0.2
        }
    }
    return noodles,sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend []string,mine []string){
    f := len(friend)
    m := len(mine)
    mine[m-1] = friend[f-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slice []float64,mumber int) []float64{
    result := []float64{}
    result = append(result,slice...)
    for i := 0; i < len(slice); i++{
        result[i] = (slice[i] * float64(mumber)) / 2.0
    }
    return result
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
