package main 
import "fmt"

func main() { 

    qtd := 0
    fmt.Scan (&qtd)
    animais := make ([]int, qtd)

    solteiros := make (map[int] int)
    for i := range animais {
        fmt.Scan(&animais[i])
    }

    pares := 0
    for _, animal := range animais {
        if solteiros [-animal] > 0 {
            solteiros [-animal] --
            pares ++
        } else {
            solteiros [animal] ++
        }
    }

    fmt.Println (pares)
}
