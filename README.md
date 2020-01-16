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

### Notes  
- Since Go standard library provides us preudo-random number generator, we can rely on this to make the output of 
the Alien Invasion always the same. In order to achieve constant result we need to keep in mind to operate on the
structures to keep an same order of data. Having the deterministic output helps us in testing.
In order to have a random result of an Aliens invasion (closer to a real world scenario) - we provide the default 
source to deterministic state using rand.Seed function that we run in our main.go file. 

### Things to improve
- Use one of the existing logging libraries to support log levels