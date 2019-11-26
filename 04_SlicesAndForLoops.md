# Slices and arrays
Go has two base data structures for handling lists of records.  
- **Array**: can be thought of as a very basic or very primitive data structure for holding a list of records.
- **Slice**: is very similar to array but it's a little bit more fancy. It's got more features attached to it that you
would usually expect to be attached to an array.  

Both have to be defied as a datatype.  

# Declaration
Slice: `slice := []string{"value1", "value2"}`  

# For
Example: 
```
	cards := []string{"Ace of Diamonds", newCard()}
	for i, card := range cards {
		fmt.Println(i, card)
	}
```

# Mechanics
Access to an element -> `slice[index]`
Access to a sub slice -> `slice[startIndex:endIndexNotIncluding]`
