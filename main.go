package main

import (
	"encoding/json"
	"fmt"

	"github.com/judoDsl/interpolator"
	utils "github.com/judoDsl/plugin_definitions"

	"io/ioutil"
	"net/http"
)

/*
This plugin implements the commands:
	* joke: get a joke  of Chuck Norris usiong a specific category
*/

const (
	CMD_JOKE = "joke"
)

type joke struct {
	vars interpolator.VarsContent
}

func main() {

	vars := interpolator.VarsContent{}
	vars["category"] = "career"
	content, _ := commandFactory(CMD_JOKE, vars).Execute()
	fmt.Print(content)
}

// Command JOKE: it returns a joke related with the category param. List of params:
//   - category: The category must be some one of this values ["animal","career","celebrity","dev","explicit","fashion","food","history","money","movie","music","political","religion","science","sport","travel"], also admit an empty value
func (joker joke) Execute() (interpolator.VarsContent, error) {

	if category, ok := joker.vars["category"]; !ok {
		return nil, fmt.Errorf("error, plugin '%s' the category params are not setted", CMD_JOKE)
	} else {
		url := fmt.Sprintf("https://api.chucknorris.io/jokes/random?category=%s", category)
		if response, err := http.Get(url); err != nil {
			return nil, fmt.Errorf("error, calling rest service: %s", err)
		} else {
			if responseData, err := ioutil.ReadAll(response.Body); err != nil {
				return nil, fmt.Errorf("error, reading boy: %s", err)
			} else {
				jsonMap := interpolator.VarsContent{}

				if err = json.Unmarshal(responseData, &jsonMap); err != nil {
					return nil, fmt.Errorf("error, building the Json fromn data: %s", err)
				}
				return jsonMap, nil
			}
		}
	}
}

var commandFactory utils.CommandFactory = func(commandString string, vars interpolator.VarsContent) utils.Command {
	var command utils.Command = nil
	switch commandString {
	case CMD_JOKE:
		command = joke{vars: vars}
	}

	return command
}
