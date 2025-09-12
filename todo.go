package main

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json: "completed"`
}

var todos []Todo

var currentID int

func init() {
	todos = append(todos, Todo{
		ID:        1,
		Title:     "Learn Go",
		Completed: false,
	})

	todos = append(todos, Todo{
		ID:        2,
		Title:     "Learn Go",
		Completed: false,
	})

	todos = append(todos, Todo{
		ID:        3,
		Title:     "Learn Go",
		Completed: false,
	})

	todos = append(todos, Todo{
		ID:        4,
		Title:     "Learn Go",
		Completed: false,
	})
	currentID = len(todos)
}
