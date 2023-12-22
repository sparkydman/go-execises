# Go execises for developers
This repository compose of small Go projects to sharpen Go developers. Each folder is a new challege.
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
cd quize

go run main.go

or 

go run main.go -csv="problems.csv" -duration=30
```
## Execise #2: URL-Shortner
Create a url-shortener that should map path to actual URL. The app should be able to lookup the url on a map and a yaml file and redirect to the url if there's match else stick to the default package. 
## Sample
```go
    mapPathUrl := map[string]string{
        "/gg": "https://google.com",
        "/tb":  "https://www.youtube.com",
    }
```
```yaml
- path: /yaml-doc
  url: https://pkg.go.dev/gopkg.in/yaml.v3
- path: /go-play
  url: https://pkg.go.dev
```
### Usage
```
cd url-shortener

go run main.go
```