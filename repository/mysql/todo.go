package mysql

import (
	"context"
	"database/sql"

	"github.com/winartodev/go-grpc/types"
)

type TodoRepository struct {
	DB *sql.DB
}

type TodoRepositoryInterface interface {
	Create(ctx context.Context, data types.Task) (id int64, err error)
	GetByID(ctx context.Context, id int64) (result *types.Task, err error)
	GetAllTaskDB(ctx context.Context) (result []types.Task, err error)
	UpdateByIDDB(ctx context.Context, id int64, data types.Task) (err error)
	DeleteByIDDB(ctx context.Context, id int64) (err error)
}

func NewTodoRepository(db *sql.DB) TodoRepositoryInterface {
	return &TodoRepository{
		DB: db,
	}
}

func (tr *TodoRepository) Create(ctx context.Context, data types.Task) (id int64, err error) {
	stmt, err := tr.DB.Prepare(CreateTaskQuery)
	if err != nil {
		return id, err
	}

	res, err := stmt.Exec(data.Description, data.Completed, data.CreatedAt)
	if err != nil {
		return id, err
	}

	return res.LastInsertId()
}

func (tr *TodoRepository) GetByID(ctx context.Context, id int64) (result *types.Task, err error) {
	row := tr.DB.QueryRow(GetTaskByID, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var task types.Task
	err = row.Scan(&task.ID, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (tr *TodoRepository) GetAllTaskDB(ctx context.Context) (result []types.Task, err error) {
	rows, err := tr.DB.Query(GetAllTask)
	if err != nil {
		return nil, err
	}

	var tasks []types.Task
	for rows.Next() {
		var task types.Task
		err := rows.Scan(&task.ID, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, err
}

func (tr *TodoRepository) UpdateByIDDB(ctx context.Context, id int64, data types.Task) (err error) {
	stmt, err := tr.DB.Prepare(UpdateTaskQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(data.Description, data.Completed, data.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TodoRepository) DeleteByIDDB(ctx context.Context, id int64) (err error) {
	stmt, err := tr.DB.Prepare(DeleteTaskQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
