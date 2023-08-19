package mysql

var (
	CreateTaskQuery = `INSERT INTO task (id, description, complete, created_at, updated_at) VALUES (NULL, ?, ?, ?, NULL);`

	GetTaskByID = `SELECT id, description, complete, created_at, updated_at FROM task WHERE id = ?;`

	GetAllTask = `SELECT id, description, complete, created_at, updated_at FROM task;`

	UpdateTaskQuery = `UPDATE task SET description = ?, complete = ?, updated_at = ? WHERE id = ?;`

	DeleteTaskQuery = `DELETE FROM task WHERE id = ?;`
)
