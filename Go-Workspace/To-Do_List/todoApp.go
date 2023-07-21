// In this code we have the ToDoItem struct representing a single ToDO item,and the ToDoList struct representing the collection of ToDo items.
// The ToDoList struct has methods for each CRUD operations:

// CreateItem:Adds a new item to the list
// ReadAllItems:Prints all items in the list
// UpdateItem :UPdates an Item in the List based on the provided id
// DeleteItem : Deletes an item from the list based on provided id

// Inside the main function we demenostrate the usage of each operation:

// --Creating a new Item and adding it to the list
// --Reading all items in the list
// --Updating an item in the list
// --Deleting an item from the list
// --Reading all items after modifications

package main

import "fmt"

type ToDoItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
type ToDoList struct {
	Items []ToDoItem
}

func (list *ToDoList) CreateItem(item ToDoItem) {
	list.Items = append(list.Items, item)
	fmt.Println("Item created successfully.")
}

func (list *ToDoList) ReadAllItems() {
	if len(list.Items) == 0 {
		fmt.Println("No Items in the List.")
		return
	}

	for _, item := range list.Items {
		fmt.Println("ID:", item.ID)
		fmt.Println("Title:", item.Title)
		fmt.Println("Description:", item.Description)
		fmt.Println("Status:", item.Status)
		fmt.Println()
	}
}
func (list *ToDoList) UpdateItem(id string, updatedItem ToDoItem) {
	for i, item := range list.Items {
		if item.ID == id {
			list.Items[i] = updatedItem
			fmt.Println("Item Updated Successfully.")
			return
		}
	}
	fmt.Println("Item not found.")
}
func (list *ToDoList) DeleteItem(id string) {
	for i, item := range list.Items {
		if item.ID == id {
			list.Items = append(list.Items[:i], list.Items[i+1:]...)
			fmt.Println("Item Deleted Successfully.")
			return
		}
	}
	fmt.Println("Item not found.")

}

func main() {
	// for the ToDoList type we have already defined mehods CRUD,so any ToDoList type can access all methods
	todoList := ToDoList{
		Items: []ToDoItem{
			{ID: "1", Title: "Task 1", Description: "Description 1", Status: "Incomplete"},
			{ID: "2", Title: "Task 2", Description: "Description 2", Status: "Incomplete"},
			{ID: "3", Title: "Task 3", Description: "Description 3", Status: "Incomplete"},
		},
	}

	// Create an item
	newItem := ToDoItem{ID: "4", Title: "Task 4", Description: "Description 4", Status: "Incomplete"}
	todoList.CreateItem(newItem)

	// Read all items
	todoList.ReadAllItems()

	// Update an item
	itemID := "2" // Replace with the ID of the item you want to update
	updatedItem := ToDoItem{ID: itemID, Title: "Updated Task", Description: "Updated Description", Status: "Complete"}
	todoList.UpdateItem(itemID, updatedItem)

	// Delete an item
	itemIDToDelete := "3" // Replace with the ID of the item you want to delete
	todoList.DeleteItem(itemIDToDelete)

	// Read all items after modifications
	todoList.ReadAllItems()
}
