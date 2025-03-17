package command

import "fmt"

// InsertElement inserts a new element into the array list.
func (a *ArrayList) InsertElement() {
	var data int
	fmt.Print("enter the data to insert: ")
	fmt.Scan(&data)

	a.data = append(a.data, data)
	fmt.Println("element inserted successfully.")
}

// DeleteElement removes an element with the given value from the array list.
func (a *ArrayList) DeleteElement() {
	if len(a.data) == 0 {
		fmt.Println("list is empty.")
		return
	}

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

// DisplayList prints all elements of the array list.
func (a *ArrayList) DisplayList() {
	if len(a.data) == 0 {
		fmt.Println("list is empty.")
		return
	}

	fmt.Print("array list: ")
	for _, v := range a.data {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// SearchElement searches for an element in the array list.
func (a *ArrayList) SearchElement() {
	var data int
	fmt.Print("enter the data to search: ")
	fmt.Scan(&data)

	for _, v := range a.data {
		if v == data {
			fmt.Println("element found.")
			return
		}
	}
	fmt.Println("element not found.")
}
