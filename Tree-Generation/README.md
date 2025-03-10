For not generates a tree based on game rules

Requirements go 1.23.4+ (probably will build with older go versions)

Build guide:

1. clone the repository 
2. run `go build`


Can be ran without arguments, will generate 5 random integers based on game rules and run through them till all endpoints are found and then will calculate winners, and result point count for each endpoint.

Run examples without arguments:

`./game.exe`

Run examples with arguments:
`./game.exe 44580` 
or
`./game.exe 44580 44600 ....`
