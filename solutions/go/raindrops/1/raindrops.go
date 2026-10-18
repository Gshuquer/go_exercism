package raindrops
import "fmt"
func Convert(number int) string {
    factors := []struct {
        div   int
        sound string
    }{
        {3, "Pling"},
        {5, "Plang"},
        {7, "Plong"},
    }

    var result string
    for _, f := range factors {
        if number%f.div == 0 {
            result += f.sound
        }
    }

    if len(result) == 0 {
        return fmt.Sprintf("%d",number)
    }
    return result
}

