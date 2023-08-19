package handler

import (
	"context"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/winartodev/go-grpc/types"
	"github.com/winartodev/go-grpc/usecase"
	todoUsecaseMock "github.com/winartodev/go-grpc/usecase/mocks"
	"github.com/winartodev/go-grpc/util"
	"github.com/winartodev/protobuff-collections/todolist"
	"google.golang.org/grpc"
)

type todoHandlerMock struct {
	todolist.UnimplementedTodoServer
	TodoUsecase *todoUsecaseMock.TodoUsecaseInterface
}

func newTodoHandler() todoHandlerMock {
	return todoHandlerMock{
		TodoUsecase: new(todoUsecaseMock.TodoUsecaseInterface),
	}
}

func TestNewTodoHandler(t *testing.T) {
	type args struct {
		grpcServer  *grpc.Server
		todoUsecase usecase.TodoUsecaseInterface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				grpcServer:  grpc.NewServer(),
				todoUsecase: &usecase.TodoUsecase{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewTodoHandler(tt.args.grpcServer, tt.args.todoUsecase)
		})
	}
}

func TestTodoHandler_CreateTask(t *testing.T) {
	todoHandlerMock := newTodoHandler()
	ctx := context.Background()

	mockTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	rpcData := &todolist.Task{
		Id:          1,
		Description: "Description",
		Completed:   false,
		CreatedAt:   mockTime.Unix(),
		UpdatedAt:   time.Time{}.Unix(),
	}

	data := types.Task{
		ID:          1,
		Description: "Description",
		Completed:   false,
		CreatedAt:   &mockTime,
		UpdatedAt:   &time.Time{},
	}

	type fields struct {
		UnimplementedTodoServer todolist.UnimplementedTodoServer
		TodoUsecase             usecase.TodoUsecaseInterface
	}
	type args struct {
		ctx context.Context
		req *todolist.CreateTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *todolist.CreateTaskResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "Success Create Task GRPC",
			fields: fields{
				UnimplementedTodoServer: todolist.UnimplementedTodoServer{},
				TodoUsecase:             todoHandlerMock.TodoUsecase,
			},
			args: args{
				ctx: ctx,
				req: &todolist.CreateTaskRequest{
					Task: rpcData,
				},
			},
			want: &todolist.CreateTaskResponse{
				Task: rpcData,
			},
			wantErr: false,
			mock: func() {
				todoHandlerMock.TodoUsecase.On("Create", ctx, util.TransformTaskData(rpcData)).Return(&data, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			th := &TodoHandler{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				TodoUsecase:             tt.fields.TodoUsecase,
			}
			got, err := th.CreateTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoHandler.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoHandler.CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoHandler_DeleteTask(t *testing.T) {
	todoHandlerMock := newTodoHandler()
	ctx := context.Background()

	type fields struct {
		UnimplementedTodoServer todolist.UnimplementedTodoServer
		TodoUsecase             usecase.TodoUsecaseInterface
	}
	type args struct {
		ctx context.Context
		req *todolist.DeleteTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *todolist.DeleteTaskResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "Sucess Delete Task GRPC",
			fields: fields{
				UnimplementedTodoServer: todolist.UnimplementedTodoServer{},
				TodoUsecase:             todoHandlerMock.TodoUsecase,
			},
			args: args{
				ctx: ctx,
				req: &todolist.DeleteTaskRequest{
					Id: int64(1),
				},
			},
			want:    &todolist.DeleteTaskResponse{},
			wantErr: false,
			mock: func() {
				todoHandlerMock.TodoUsecase.On("Delete", ctx, int64(1)).Return(nil).Times(1).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			th := &TodoHandler{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				TodoUsecase:             tt.fields.TodoUsecase,
			}
			got, err := th.DeleteTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoHandler.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoHandler.DeleteTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoHandler_GetListTask(t *testing.T) {
	todoHandlerMock := newTodoHandler()
	ctx := context.Background()

	mockTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	data := []types.Task{
		{
			ID:          1,
			Description: "Description",
			Completed:   false,
			CreatedAt:   &mockTime,
			UpdatedAt:   &mockTime,
		},
	}

	type fields struct {
		UnimplementedTodoServer todolist.UnimplementedTodoServer
		TodoUsecase             usecase.TodoUsecaseInterface
	}
	type args struct {
		ctx context.Context
		req *todolist.GetListOfTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *todolist.ListOfTasksResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "Get List Task GRPC",
			fields: fields{
				UnimplementedTodoServer: todolist.UnimplementedTodoServer{},
				TodoUsecase:             todoHandlerMock.TodoUsecase,
			},
			args: args{
				ctx: ctx,
				req: &todolist.GetListOfTaskRequest{},
			},
			want: &todolist.ListOfTasksResponse{
				Task: []*todolist.Task{
					{
						Id:          1,
						Description: "Description",
						Completed:   false,
						CreatedAt:   mockTime.Unix(),
						UpdatedAt:   mockTime.Unix(),
					},
				},
			},
			mock: func() {
				todoHandlerMock.TodoUsecase.On("GetAll", ctx).Return(data, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			th := &TodoHandler{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				TodoUsecase:             tt.fields.TodoUsecase,
			}
			got, err := th.GetListTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoHandler.GetListTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoHandler.GetListTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoHandler_GetTaskByID(t *testing.T) {
	todoHandlerMock := newTodoHandler()
	ctx := context.Background()

	mockTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	data := types.Task{
		ID:          1,
		Description: "Description",
		Completed:   false,
		CreatedAt:   &mockTime,
		UpdatedAt:   &mockTime,
	}

	type fields struct {
		UnimplementedTodoServer todolist.UnimplementedTodoServer
		TodoUsecase             usecase.TodoUsecaseInterface
	}
	type args struct {
		ctx context.Context
		req *todolist.GetTaskByIDRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *todolist.GetTaskByIDResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "Sucess Get Task By ID GRPC",
			fields: fields{
				UnimplementedTodoServer: todolist.UnimplementedTodoServer{},
				TodoUsecase:             todoHandlerMock.TodoUsecase,
			},
			args: args{
				ctx: ctx,
				req: &todolist.GetTaskByIDRequest{
					Id: int64(1),
				},
			},
			want: &todolist.GetTaskByIDResponse{
				Task: util.TransformTaskDataRPC(&data),
			},
			wantErr: false,
			mock: func() {
				todoHandlerMock.TodoUsecase.On("GetByID", ctx, int64(1)).Return(&data, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			th := &TodoHandler{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				TodoUsecase:             tt.fields.TodoUsecase,
			}
			got, err := th.GetTaskByID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoHandler.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoHandler.GetTaskByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoHandler_UpdateTask(t *testing.T) {
	todoHandlerMock := newTodoHandler()
	ctx := context.Background()

	mockTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	data := types.Task{
		Completed:   true,
		Description: "Update Description",
		CreatedAt:   &mockTime,
		UpdatedAt:   &mockTime,
	}

	type fields struct {
		UnimplementedTodoServer todolist.UnimplementedTodoServer
		TodoUsecase             usecase.TodoUsecaseInterface
	}
	type args struct {
		ctx context.Context
		req *todolist.UpdateTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *todolist.UpdateTaskResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "Sucess Update Task GRPC",
			fields: fields{
				UnimplementedTodoServer: todolist.UnimplementedTodoServer{},
				TodoUsecase:             todoHandlerMock.TodoUsecase,
			},
			args: args{
				ctx: ctx,
				req: &todolist.UpdateTaskRequest{
					Id:          int64(1),
					Description: "Update Description",
					Completed:   true,
				},
			},
			want: &todolist.UpdateTaskResponse{
				Task: &todolist.Task{
					Description: "Update Description",
					Completed:   true,
					CreatedAt:   mockTime.Unix(),
					UpdatedAt:   mockTime.Unix(),
				},
			},
			wantErr: false,
			mock: func() {
				todoHandlerMock.TodoUsecase.On("Update", ctx, int64(1), types.Task{
					Completed:   true,
					Description: "Update Description",
				}).Return(&data, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			th := &TodoHandler{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				TodoUsecase:             tt.fields.TodoUsecase,
			}
			got, err := th.UpdateTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoHandler.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoHandler.UpdateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
