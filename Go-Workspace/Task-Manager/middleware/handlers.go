package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mdNoman21/Go/Beginner-Projects/Task-Manager/models"
)

func createConnection() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "tasksdb",
	}
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatalf("Unable to decode request body. %v", err)
	}
	id := insertTask(task)
	res := models.Response{
		TaskID:  id,
		Message: "Task successfully added",
	}
	json.NewEncoder(w).Encode(res)
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := getAllTasks()
	if err != nil {
		log.Fatalf("Unable to get all the tasks. %v", err)
	}
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatalf("unable to decode the request body. %v", err)
	}
	updatedRows := updateTask(id, &task)
	msg := fmt.Sprintf("Stock updated successfully. Total rows/records affected %v", updatedRows)
	res := models.Response{
		TaskID:  id,
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	deletedRows := deleteTask(id)
	msg := fmt.Sprintf("Stock deleted successfully. Total rows affected %v", deletedRows)
	res := models.Response{
		TaskID:  id,
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func CompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := getCompletedTasks()
	if err != nil {
		log.Fatalf("Unable to get completed tasks. %v", err)
	}
	json.NewEncoder(w).Encode(tasks)
}
func SearchTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	task, err := searchTask(id)
	if err != nil {
		log.Fatalf("Unable to search the task %v", err)
	}
	json.NewEncoder(w).Encode(task)
}
func TaskStats(w http.ResponseWriter, r *http.Request) {
	tasks, err := getAllTasks()
	if err != nil {
		log.Fatalf("Unable to get all the tasks. %v", err)
	}
	allTasks, completedTasks := taskStats(tasks)
	stats := models.TaskStats{
		TotalTasks:     allTasks,
		CompletedTasks: completedTasks,
	}
	json.NewEncoder(w).Encode(stats)
}
func getAllTasks() ([]models.Task, error) {
	db := createConnection()
	defer db.Close()
	var tasks []models.Task
	rows, err := db.Query(`SELECT * FROM tasks`)
	if err != nil {
		log.Fatalf("Unable to fetch tasks %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			log.Fatalf("Unable to fetch row %v", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func taskStats(tasks []models.Task) (int, int) {
	db := createConnection()
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM tasks`)
	if err != nil {
		log.Fatalf("Unable to get tasks %v", err)
	}
	defer rows.Close()
	var allTasks, completedTasks int
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
			continue
		}
		allTasks++
		if task.Completed {
			completedTasks++
		}
	}
	return allTasks, completedTasks
}

func insertTask(task models.Task) string {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO tasks (title, description, completed) VALUES(?, ?, ?)`
	result, err := db.Exec(sqlStatement, task.Title, task.Description, task.Completed)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error retrieving last inserted ID. %v", err)
	}
	taskID := strconv.FormatInt(lastInsertID, 10)
	fmt.Printf("Inserted a single record with taskID %v\n", taskID)
	return taskID
}

func updateTask(id string, task *models.Task) int64 {
	db := createConnection()
	defer db.Close()
	// Convert the boolean value to an integer (0 or 1)
	completedInt := 0
	if task.Completed {
		completedInt = 1
	}
	sqlStatement := `UPDATE tasks SET title=?, description=?, completed=? WHERE taskID=?`
	res, err := db.Exec(sqlStatement, task.Title, task.Description, completedInt, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows affected %v", rowsAffected)
	return rowsAffected

}
func deleteTask(id string) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM tasks WHERE taskID=?`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows affected %v", rowsAffected)
	return rowsAffected
}
func searchTask(id string) (models.Task, error) {
	var task models.Task
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM tasks WHERE taskID=?`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&task.TaskID, &task.Title, &task.Description, &task.Completed)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows are returned!")
		return task, nil
	case nil:
		return task, nil
	default:
		log.Fatalf("Unable to scacn the row %v", err)
	}
	return task, err
}
func getCompletedTasks() ([]models.Task, error) {
	db := createConnection()
	defer db.Close()
	var tasks []models.Task
	rows, err := db.Query(`SELECT * FROM tasks WHERE completed=true`)
	if err != nil {
		log.Fatalf("Unable to fetch tasks %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.TaskID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			log.Fatalf("Unable to fetch row %v", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
