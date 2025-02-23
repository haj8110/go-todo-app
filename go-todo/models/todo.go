package models

import (
	"go-todo/database"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Create a new Todo
func (t *Todo) CreateTodo() error {
	sqlStatement := `INSERT INTO todos (title, completed) VALUES (?, ?)`
	stmt, err := database.DB.Prepare(sqlStatement) // Prepare statement for the insert
	if err != nil {
		return err
	}
	defer stmt.Close() // Ensure the statement is closed after the execution

	// Execute the statement
	result, err := stmt.Exec(t.Title, t.Completed)
	if err != nil {
		return err
	}

	// Get the last inserted ID and update the Todo struct
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = int(id)

	// log.Printf("Todo created with ID: %d", t.ID)
	return nil
}

// Get all Todos
func GetTodos() ([]Todo, error) {
	sqlStatement := `SELECT id, title, completed FROM todos`
	rows, err := database.DB.Query(sqlStatement) // Execute query
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed when function exits

	var todos []Todo
	// Iterate over the rows and scan the data into Todo structs
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
// UpdateStatus updates the completed status of a todo in the database
func (t *Todo) UpdateStatus(id string, completed bool) error {
	// Prepare SQL query to update the todo status
	sqlStatement := `UPDATE todos SET completed = ? WHERE id = ?`
	_, err := database.DB.Exec(sqlStatement, completed, id)
	if err != nil {
		// log.Printf("Error updating todo status: %v", err)
		return err
	}
	return nil
}
// UpdateTitle updates the title of a todo in the database
func (t *Todo) UpdateTitle(id string, title string) error {
	// Prepare SQL query to update the todo title
	sqlStatement := `UPDATE todos SET title = ? WHERE id = ?`
	_, err := database.DB.Exec(sqlStatement, title, id)
	if err != nil {
		return err
	}
	return nil
}
// DeleteTodo deletes a todo from the database
func DeleteTodo(id string) error {
	// Prepare SQL query to delete the todo
	sqlStatement := `DELETE FROM todos WHERE id = ?`
	_, err := database.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
// MarkAllTodosCompleted updates the completed field of all todos in the database
func MarkAllTodosCompleted() error {
	// Prepare SQL query to mark all todos as completed
	sqlStatement := `UPDATE todos SET completed = true`
	_, err := database.DB.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
