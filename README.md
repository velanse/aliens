##Aliens invasion

###Usage of a program  
Example:

`go run main.go -map testdata.txt -N=2`

parameters:  
- N (int): Number of aliens  
- map (string): A path to file that contains a map of cities

Testing:  

`go test ./...`

###Assumptions
- Output file contains full information about the cities (all the cities are listed with inbound and outbound connections)
- More than one alien can land into one city and destroy it immediately
- An alien can land into the dead city - this way he dies immediately

###Things to improve
- Implement bounded parallelism for the async version for an optimal goroutines number;
- Use some of the existing logging libraries that already support log levels
- For the async version implement the travel time for the real life simulation

