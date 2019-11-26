# Go routines
A routine in go is like a process, we can run code in a new routine using the key 'go':  
```
links := []string {
    "http://google.com",
    "http://facebook.com",
    "http://stackoverflow.com",
    "http://golang.org",
    "http://amazon.com",
}

for _, link := range links {
    go checkLink(link)
}
```

By default Go tires to uses only one CPU core and have a Scheduler that manage the routines.  
Scheduler runs one routine until it finishes or makes a blocking call.  
When we have multiple CPU Core each one can run one single go routine at a time.  
# Channels
Channel is a communication link between go routines.  
To create a channel: `c := make(chan string)`
# Sending data through channels
- `channel <- 5` Send the value into this channel.  
- `myNumber <- channel` Wait for a value to be sent into the channel. When we get one assign the value to 'myNumber'.  
- `fmt.Println(<- channel)` same as before but use directly as parameter
