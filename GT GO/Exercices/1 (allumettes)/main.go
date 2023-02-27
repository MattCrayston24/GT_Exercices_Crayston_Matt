package main

import (
    "fmt"
)

func main() {
    var nballu int

    for {
        fmt.Print("Choisis un nombre d'allumettes (au minimum 4):")
        fmt.Scanln(&nballu)
        if nballu >= 4 {
            break
        }
        fmt.Println("Le nombre d'allumettes est inférieur à 4, il faut qu'il soit supérieur, à toi !")
    }

    var player int = 1
    var lastnb int = nballu

    for lastnb > 0 {
        fmt.Println("Il reste", lastnb, "allumettes.")

        var choice int
        for {
            fmt.Print("player ", player, ", Choisis un nombre entre 1 et 3, cela sera ton nombre d'allumettes ! ")
            fmt.Scanln(&choice)
            if choice >= 1 && choice <= 3 && choice <= lastnb {
                break
            }
            fmt.Println("Fais attention à bien choisir un nombre d'allumettes entre 1 et 3 !")
        }

        lastnb -= choice

        if player == 1 {
            player = 2
        } else {
            player = 1
        }
    }

    fmt.Println("Le player", player, "a perdu !")
}