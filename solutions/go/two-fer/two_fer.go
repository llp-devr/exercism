package twofer
import "fmt"

// ShareWith generates a sharing message based on the given name.
// If the input name is an empty string, it defaults to "you".
func ShareWith(name string) string {
    // If the name is an empty string, set it to the default value "you"
    if name == "" {
        name = "you"
    }

    // Use fmt.Sprintf to create a formatted string with the given name
    // and return the result
	return fmt.Sprintf("One for %s, one for me.", name)
}
