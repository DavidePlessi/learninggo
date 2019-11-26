# Declaration
`type newType []string`: this mean that newType is equal to []string.  

# Receiver functions
Example:
```
func (d deck) print () {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```
Usage
```
func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") //Returns a new slice!!!

	cards.print()
}
```
