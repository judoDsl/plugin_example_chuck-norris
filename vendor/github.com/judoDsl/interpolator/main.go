/*
The main goal of interpolator is to help you to interpolate your vars inside the string and evaluate functions related with this vars. It is dased on the 'go templates' and using the large list of functions provided by the Sprig Functions Project.
it is an example ahout how to use it:

	package main

	import "github.com/judoDSL/interpolator"

	func main () {
		values := make(map[string] interface{})
		values["name"] = "            Jose                 "
		values["main_topic"] = "restore the snyderverse"
		values["favorite_superhero"] = "batman who laughs"
		interpolator.Do("I'm {{ .name | trim }} and I want to {{ .main_topic | upper  }} because I would like to see a film related with {{ .favorite_superhero | title }}", values).Println()
	}

It would be the result of the execution:

	[jose78@~/ws/test_interpolator] $  go run main.go
	I'm Jose and I want to RESTORE THE SNYDERVERSE because I would like to see a film related with Batman Who Laughs
*/
package interpolator

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// Custom type of map[string]interface{}
type VarsContent map[string]interface{}

type content struct {
	value string
}

// Interface that provides the functionality to easily work with the interpolated string
type Content interface {
	// Retrieve the string with the vvalues interpolated.
	Get() string
	// Write to console
	Print()
	// Write to console
	Println()
	// Create an error
	Error() error
}

func (body content) Get() string {
	return body.value
}

func (body content) Print() {
	fmt.Print(body.value)
}

func (body content) Println() {
	fmt.Println(body.value)
}

func (body content) Error() error {
	return fmt.Errorf(body.value)
}

// Given a string with the templates, it is interpolated with the value of the vars.
func Do[T ~map[string]interface{}](str string, vars T) Content {
	funcMap := sprig.FuncMap()
	tmpl, err := template.New("tmpl").Funcs(funcMap).Parse(str)
	if err != nil {
		panic(err)
	}
	var tmplBytes bytes.Buffer
	err = tmpl.Execute(&tmplBytes, vars)
	if err != nil {
		panic(err)
	}
	return content{value: tmplBytes.String()}
}
