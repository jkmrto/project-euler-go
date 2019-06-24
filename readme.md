# Project Euler Go

Personal Resolutions of some of the problems of [project euler](https://projecteuler.net/problem=4). Just to learn Golang. 

# Execute go program

```
go run exercise.go
```

### To export/import custom packages on Go:
We should make the function available on outside putting the first letter of the function on capital letter:

```Go
package utils

func IsDivisibleByAny(dividend int, dividers []int) bool {
.....
}

```
Let's load it from other package inciated the path relative to the current project path:

``` Go
import (
	utils "./utils"
)
```
### Tests

Run `utils` tests

```Go
( cd utils && go test )
```
