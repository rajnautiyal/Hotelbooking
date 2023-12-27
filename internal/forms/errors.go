package forms

type errors map[string][]string

// add the error in the map
func (e errors) Add(feild, message string) {
	e[feild] = append(e[feild], message)
}

// return the first error message
func (e errors) Get(feild string) string {
	message := e[feild]
	if len(message) == 0 {
		return ""
	} else {
		return message[0]
	}
}
