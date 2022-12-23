# interpolator

The _interpolator_ project is part as _judoDSL_, the main goal of _interpolator_ is to help you to interpolate your vars inside the string and evaluate functions related with this vars. It is dased on the 'go templates' and using the large list of functions provided by the [Sprig Functions Project](https://masterminds.github.io/sprig/).


## How-to install:

To install the dependency:

```bash
go get -u github.com/judoDSL/interpolator
```

## How-to use:

You can see an example about how to use it. 

```go
package main

import "github.com/judoDSL/interpolator"

func main () {
	values := make(map[string] interface{})
	values["name"] = "            Jose                 "
	values["main_topic"] = "restore the snyderverse"
	values["favorite_superhero"] = "batman who laughs"
	interpolator.Do("I'm {{ .name | trim }} and I want to {{ .main_topic | upper  }} because I would like to see a film related with {{ .favorite_superhero | title }}", values).Println()
}
```




The result of this esecution is:

```bash
[jose78@~/ws/interpolator] $  go run main.go 
I'm Jose and I want to RESTORE THE SNYDERVERSE because I would like to see a film related with Batman Who Laughs
```
