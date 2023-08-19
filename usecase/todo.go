package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	todoRepository "github.com/winartodev/go-grpc/repository/mysql"
	"github.com/winartodev/go-grpc/types"
)

type TodoUsecase struct {
	TodoRepository todoRepository.TodoRepositoryInterface
}

type TodoUsecaseInterface interface {
	Create(ctx context.Context, data types.Task) (result *types.Task, err error)
	GetByID(ctx context.Context, id int64) (result *types.Task, err error)
	GetAll(ctx context.Context) (result []types.Task, err error)
	Update(ctx context.Context, id int64, data types.Task) (result *types.Task, err error)
	Delete(ctx context.Context, id int64) (err error)
}

func NewTodoUsecase(todoRepository todoRepository.TodoRepositoryInterface) TodoUsecaseInterface {
	return &TodoUsecase{
		TodoRepository: todoRepository,
	}
}

func (tuc *TodoUsecase) Create(ctx context.Context, data types.Task) (result *types.Task, err error) {
	now := time.Now()
	data.CreatedAt = &now

	id, err := tuc.TodoRepository.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	result, err = tuc.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (tuc *TodoUsecase) GetByID(ctx context.Context, id int64) (result *types.Task, err error) {
	return tuc.TodoRepository.GetByID(ctx, id)
}

func (tuc *TodoUsecase) GetAll(ctx context.Context) (result []types.Task, err error) {
	return tuc.TodoRepository.GetAllTaskDB(ctx)
}

func (tuc *TodoUsecase) Update(ctx context.Context, id int64, data types.Task) (result *types.Task, err error) {
	task, err := tuc.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, fmt.Errorf("task with id %d was not found", id)
	}

	if data.Description != "" {
		task.Description = data.Description
	}

	task.Completed = data.Completed

	now := time.Now()
	task.UpdatedAt = &now

	err = tuc.TodoRepository.UpdateByIDDB(ctx, id, *task)
	if err != nil {
		return nil, err
	}

	result, err = tuc.GetByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return result, err
}

func (tuc *TodoUsecase) Delete(ctx context.Context, id int64) (err error) {
	task, err := tuc.GetByID(ctx, id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if task == nil {
		return fmt.Errorf("task with id %v not found", id)
	}

	return tuc.TodoRepository.DeleteByIDDB(ctx, id)
}
