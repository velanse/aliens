## Aliens invasion

### Usage of a program  
Example:

`go run main.go -map testdata/test01.txt -N=2`

parameters:  
- N (int): Number of aliens  
- map (string): A path to file that contains a map of cities
- debug (bool): Set to true in order to output debug messages

Testing:  

`go test ./...`

Run Benchmarks:

`go test ./... -bench .`

### Assumptions
- Output file contains full information about the cities (all the cities are listed with inbound and outbound connections)
- More than one alien can land into one city and destroy it immediately
- An alien can land into the dead city - this way he dies immediately

### Things to improve
- Use one of the existing logging libraries to support log levels
- Encapsulate the random numbers generator in a separate interface. So it could be replaced with a mock that generates deterministic values which will help to create more sophisticated tests.
