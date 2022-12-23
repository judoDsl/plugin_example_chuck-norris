# plugin_definitions

_plugin_definitions_ contains the definitions needed to implement a new plugin to be loaded within the _jusdo_dsl_.

## Minimum definitions

### _CommandFactory_ function:

This function implement the _Factory Pattern_ over the string argument, this string is the name of the command to be loaded and returned as implementation of the _Command Pattern_.
 
#### Params

* *string*: Name of the command to be executed.

* **interpolator.VarsContent*: Vars with the value interpolated needed to execute the command.

#### Return
*  *Command*: Return an implementation of the _Command Pattern_. This implementation doesn't admit any params and encapsulate the logic by the selected command to be executed. 

```go
func(string, interpolator.VarsContent) Command
```