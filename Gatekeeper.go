package main

import (
    "fmt"
    "time"
)

func main() {
    HuidigeTijd := time.Now().Hour()
    var groet string

    switch {
    case HuidigeTijd >= 7 && HuidigeTijd < 12:
        groet = "Goedemorgen"
    case HuidigeTijd >= 12 && HuidigeTijd < 18:
        groet = "Goedemiddag"
    case HuidigeTijd >= 18 && HuidigeTijd < 23:
        groet = "Goedenavond"
    default:
        fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten")
        return
    }

    fmt.Println(groet + "! Welkom bij Fonteyn Vakantieparken")

}
