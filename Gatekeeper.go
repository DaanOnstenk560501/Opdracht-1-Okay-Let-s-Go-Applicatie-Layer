package main

import (
    "fmt"
    "strings"
    "time"
)

func main() {
    var kenteken string
    fmt.Print("Voer uw kenteken in: ")
    fmt.Scanln(&kenteken)

    kenteken = strings.ToUpper(kenteken)

    toegestaneKentekens := map[string]bool{
        "ABC123": true,
        "XYZ789": true,
        "LMN456": true,
    }

    if !toegestaneKentekens[kenteken] {
        fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
        fmt.Println("Druk op Enter om af te sluiten...")
        fmt.Scanln()
        return
    }

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
        fmt.Println("Druk op Enter om af te sluiten...")
        fmt.Scanln()
        return
    }

    fmt.Println(groet + "! Welkom bij Fonteyn Vakantieparken")
    fmt.Println("Druk op Enter om af te sluiten...")
    fmt.Scanln()
}