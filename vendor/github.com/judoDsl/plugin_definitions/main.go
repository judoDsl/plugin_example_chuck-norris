// judo plugin definitions contains the definitions needed to implement a new plugin to be loaded within the judoDsl.
package plugin_definitions

import "github.com/judoDsl/interpolator"

/*
Factory Pattern, takeing the command to be executed and the vars needed by the command.

Params:

  - string: Name of the command to be executed.
  - interpolator.VarsContent: Vars with the value interpolated needed to execute the command.

Return:
  - Command: Return an implementation of the _Command Pattern_. This implementation doesn't admit any params and encapsulate the logic by the selected command to be executed. \
*/
type CommandFactory func(string, interpolator.VarsContent) Command

// Interface Command to be implemented for each command within the plugin.
type Command interface {
	// It encapsulates the logic of the command and return the result of the execution or error.
	Execute() (interpolator.VarsContent, error)
}
