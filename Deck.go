package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type card struct {
    suit string
    value string
}

type deck []card

func newDeck() deck {
    values := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
                        "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
    suits := [4]string{"Spades", "Diamonds", "Clubs", "Hearts"}
    outputDeck := deck{}

    for _, suit := range suits {
        for _, value := range values {
            outputDeck = append(outputDeck, card{
                                                value: value,
                                                suit: suit,
                                                })
        }
    }
    return outputDeck
}

func (c card) print() {
    fmt.Println(c.value + " of " + c.suit)
}
func (d deck) print() {
    for _, card := range d {
        card.print()
    }
}

func (d *deck) deal(numberOfCards int) deck {
    hand := (*d)[:numberOfCards]
    *d = append((*d)[numberOfCards:])
    return hand

}

func (d *deck) addToDeck(hands []*deck) {
    for _, hand := range hands {
        for _, card := range *hand {
            *hand = append((*hand)[1:])
            *d = append(*d, card)
        }
    }

}

func (d deck) shuffle() {
    src := rand.NewSource(time.Now().UnixNano())
    rng := rand.New(src)
    deckLength := len(d)
    
    for i := range d {
        newIndex := rng.Intn(deckLength - 1)
        d[newIndex], d[i] = d[i], d[newIndex]
    }
} 

func (d deck) writeToFile(filename string) {
    err := os.WriteFile(filename, []byte(d.toString()), 0666)
    if err != nil {
        log.Fatal(err)
    }
}

func readDeckFromFile(filename string) deck {
    outputDeck := deck{}
    fileData, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    fileString := string(fileData)
    stringsSlice := strings.Split(fileString, ", ")
    for _, cardString := range stringsSlice {
        s := strings.Split(cardString, " of ")
        outputDeck = append(outputDeck, card{
                                        value: s[0],
                                        suit: s[1],
                                        })
    }
    return outputDeck
}

func (c card) toString() string {
    return c.value + " of " + c.suit
}

func (d deck) toString() string {
    output := ""
    for i, card := range d {
        output = output + card.toString() 
        if i < len(d) - 1 {
            output = output + ", "
        }
    }
    return output
}
