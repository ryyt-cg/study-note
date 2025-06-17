package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func main() {

	// Initialize a variable of type user by positional parameter
	// Must assign values to all fields.
	// Not prefer way because not knowing what the values are assigned to the fields
	rossi := user{"rossi", "rossi@email.com", 56, true}

	// Declare a variable of type user and initialize some or all the fields
	// prefer way because of code readability and knowing what the values are assigned to the fields
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	// You may omit assign some values to the fields.
	// Omitted fields will be assigned the zero value of the field's type
	barry := user{
		name:  "barry",
		email: "barr@email.com",
	}

	fmt.Printf("%v, %T\n", rossi, rossi)
	fmt.Printf("%v, %T\n", lisa, lisa)
	fmt.Printf("%v, %T\n", barry, barry)
}
