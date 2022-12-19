# judo_plugin_example-chuck-norris

It is an example of implementation of a simple plugin, in this case it implements only one command. 

Please, note we have used the *judoDSL/judo_utils* lib to generate the structure needed to be loaded from JudoDsl.

```go
import utils "github.com/judoDsl/plugin_definitions"



// The Factory Pattern 'commandFactory'
var commandFactory utils.CommandFactory = func(commandString string, vars interpolator.VarsContent) utils.Command {
	//...
	return command
}

// The Command Pattern 
func (joker joke) Execute() (interpolator.VarsContent, error) {
    //...
}

```