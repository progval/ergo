// Copyright (c) 2016- Daniel Oaks <daniel@danieloaks.net>
// released under the MIT license

package irc

// argsToStrings takes the arguments and splits them into a series of strings,
// each argument separated by delim and each string bounded by maxLength.
func argsToStrings(maxLength int, arguments []string, delim string) []string {
	var messages []string

	var buffer string
	for {
		if len(arguments) < 1 {
			break
		}

		if len(buffer) > 0 && maxLength < len(buffer)+len(delim)+len(arguments[0]) {
			messages = append(messages, buffer)
			buffer = ""
			continue
		}

		if len(buffer) > 1 {
			buffer += delim
		}
		buffer += arguments[0]
		arguments = arguments[1:]
	}

	if len(buffer) > 0 {
		messages = append(messages, buffer)
	}

	return messages
}
