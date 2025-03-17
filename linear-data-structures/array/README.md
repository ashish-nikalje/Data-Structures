# Array Data Structure

## Overview
This program implements a simple **array-based data structure** in Go, providing basic operations such as insertion, deletion, searching, and displaying elements.

## Features
- **Insert Element**: Add a new element to the array.
- **Delete Element**: Remove a specific element from the array.
- **Search Element**: Find an element within the array.
- **Display List**: Print all elements in the array.

## File Structure
```
array/
├── command/
│   ├── array.go  # Contains ArrayList implementation
│   ├── cmd.go    # Handles user input and program control
└── main.go       # Entry point for the application
```

## Usage
Run the program using:
```sh
go run main.go
```

### Example User Interaction
```
Welcome to the Array List Program! This program allows you to perform the following operations:
1. Insert an element
2. Display the list
3. Delete an element
4. Search for an element
5. Exit
Please enter the number of the operation you would like to perform:
```

### Example Commands:
- **Insert an element**: Enter `1` and input the value.
- **Display the list**: Enter `2` to see all elements.
- **Delete an element**: Enter `3` and input the value to delete.
- **Search for an element**: Enter `4` and input the value to search.
- **Exit**: Enter `5` to quit the program.

## Code Overview
### Struct Definition
```go
type ArrayList struct {
    data []int
}
```

### Insertion Function
```go
func (a *ArrayList) InsertElement() {
    var data int
    fmt.Print("enter the data to insert: ")
    fmt.Scan(&data)
    a.data = append(a.data, data)
    fmt.Println("element inserted successfully.")
}
```

### Deletion Function
```go
func (a *ArrayList) DeleteElement() {
    var data int
    fmt.Print("enter the data to delete: ")
    fmt.Scan(&data)
    for i, v := range a.data {
        if v == data {
            a.data = append(a.data[:i], a.data[i+1:]...)
            fmt.Println("element deleted successfully.")
            return
        }
    }
    fmt.Println("element not found.")
}
```