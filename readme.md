# A Tour of Go

## https://go.dev/tour/methods/4
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

## https://go.dev/tour/methods/6
## https://go.dev/tour/methods/7
## https://go.dev/tour/methods/8
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

### https://go.dev/tour/methods/12
In Go, an interface value is a two-part data structure: a type and a value. 
An interface is truly nil when it has not been assigned a concrete value or type. The type and value components of the interface are both empty.
```go

```