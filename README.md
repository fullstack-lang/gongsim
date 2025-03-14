Repo archived, function is no as a lib in gong

A gong stack for developping simulations

### compile application
> cd ng; npm i; ng build;

### launch application

at the root of the repository (requires go >= 1.16)
> go run main.go

### expected result

launch browser on http://localhost:8080

# Architecture of a simulation

Gongsim allows implementation of Agent Based simuation.

For instance, let's simulation N lady bugs walking on a ruler of one meter.

```go
type Ladybug struct {


}
```




