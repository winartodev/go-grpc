package handler

import (
	"context"

	"github.com/winartodev/go-grpc/types"
	"github.com/winartodev/go-grpc/usecase"
	"github.com/winartodev/go-grpc/util"
	"github.com/winartodev/protobuff-collections/todolist"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type TodoHandler struct {
	todolist.UnimplementedTodoServer
	TodoUsecase usecase.TodoUsecaseInterface
}

func NewTodoHandler(grpcServer *grpc.Server, todoUsecase usecase.TodoUsecaseInterface) {
	todoHandler := &TodoHandler{
		TodoUsecase: todoUsecase,
	}

	todolist.RegisterTodoServer(grpcServer, todoHandler)

	reflection.Register(grpcServer)
}

func (th *TodoHandler) CreateTask(ctx context.Context, req *todolist.CreateTaskRequest) (*todolist.CreateTaskResponse, error) {

	data := util.TransformTaskData(req.Task)

	task, err := th.TodoUsecase.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	res := util.TransformTaskDataRPC(task)

	return &todolist.CreateTaskResponse{
		Task: res,
	}, nil
}

func (th *TodoHandler) DeleteTask(ctx context.Context, req *todolist.DeleteTaskRequest) (*todolist.DeleteTaskResponse, error) {
	err := th.TodoUsecase.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &todolist.DeleteTaskResponse{}, nil
}

func (th *TodoHandler) GetListTask(ctx context.Context, req *todolist.GetListOfTaskRequest) (*todolist.ListOfTasksResponse, error) {
	tasks, err := th.TodoUsecase.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var todolistTasks []*todolist.Task
	for _, task := range tasks {
		taskTmp := util.TransformTaskDataRPC(&task)
		todolistTasks = append(todolistTasks, taskTmp)
	}

	return &todolist.ListOfTasksResponse{
		Task: todolistTasks,
	}, nil
}

func (th *TodoHandler) GetTaskByID(ctx context.Context, req *todolist.GetTaskByIDRequest) (*todolist.GetTaskByIDResponse, error) {
	task, err := th.TodoUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res := util.TransformTaskDataRPC(task)

	return &todolist.GetTaskByIDResponse{
		Task: res,
	}, nil
}

func (th *TodoHandler) UpdateTask(ctx context.Context, req *todolist.UpdateTaskRequest) (*todolist.UpdateTaskResponse, error) {
	task, err := th.TodoUsecase.Update(ctx, req.Id, types.Task{
		Completed:   req.Completed,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	res := util.TransformTaskDataRPC(task)

	return &todolist.UpdateTaskResponse{
		Task: res,
	}, nil
}
