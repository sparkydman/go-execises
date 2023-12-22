# Go execises for developers
This repository compose of small Go projects to sharpen Go developers. Each folder is a new challege.
## Coverage
- Basic Go programing logics
- Goroutines
- Channel
- Pointer
- Unit tests
- And More
## Execise #1: Quiz Game
Develop a timer base quiz where the user answers questions util their time is up. The program should automatically exit the quiz and output the user result.

*Question and corresponding answer should be read from CSV file defaulted to `problems.csv` at onces before the quiz start. The default quiz duration should be `30 seconds` and user should able to pass the CSV file path and duration.*

### CSV sample
```
2*2,4
3+4,7
12*3,24
34+4,38
2*8,16
9+12,21
7+8,15
0+9,9
0*4,0
```
### Usage
```shell
go run main.go

or 

go run main.go -csv="problems.csv" -duration=30
```