## Map

* Maps are used to store data values in key:value pairs.
* The default value of a map is nil.

```go
package main
import "fmt"

func main() {
	// Create a map with a key of type string and a value of type int.
	dict := make(map[string]int)

	// Create a map with a key and value of type string.
	// Initialize the map with 2 key/value pairs.
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}
	
	fmt.Println("dict: ", dict)
	fmt.Println("dict2: ", dict2)
	fmt.Println("subjectMarks: ", subjectMarks)
}
```

Change value of a map

```go
package main
import "fmt"

func main() {

  // create a map
  capital := map[string]string{ "Nepal": "Kathmandu", "US": "New York"}
  fmt.Println("Initial Map: ", capital)

  // change value of US to Washington DC
  capital["US"] = "Washington DC"

  fmt.Println("Updated Map: ", capital)
}
```

Add Element of Go Map Element

```go
package main
import "fmt"

func main() {

  // create a map
  students := map[int]string{1: "John", 2: "Lily"}
  fmt.Println("Initial Map: ", students)

  // add element with key 3
  students[3] = "Robin"

  // add element with key 5
  students[5] = "Julie"

  fmt.Println("Updated Map: ", students)
}

```

Delete an element in map

```go
package main
import "fmt"

func main() {

  // create a map
  personAge := map[string]int{"Hermione": 21, "Harry": 20, "John": 25}
  fmt.Println("Initial Map: ", personAge)

  // remove element of map with key John
  delete(personAge, "John")

  fmt.Println("Updated Map: ", personAge)
}
```

```go

```
