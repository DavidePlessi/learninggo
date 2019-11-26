# Interfaces
```
type bot interface {
    getGreeting() string
}
```
Interfaces are implicit so all types that fulfill the interface requirements automatically are associated.  
Interfaces can have interfaces specified:
```
type bot interface {
    interface1
    interface2
}
```
