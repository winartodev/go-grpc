package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	todoRepository "github.com/winartodev/go-grpc/repository/mysql"
	todoRepositoryMock "github.com/winartodev/go-grpc/repository/mysql/mocks"
	"github.com/winartodev/go-grpc/types"
)

type todoUsecaseMock struct {
	TodoRepository *todoRepositoryMock.TodoRepositoryInterface
}

func newTodoUsecaseMock() todoUsecaseMock {
	return todoUsecaseMock{
		TodoRepository: new(todoRepositoryMock.TodoRepositoryInterface),
	}
}

var (
	mockTime = time.Date(2020, 10, 25, 0, 0, 0, 0, time.UTC)

	dataMock = types.Task{
		ID:          1,
		Description: "Create Task",
		Completed:   false,
		CreatedAt:   &mockTime,
		UpdatedAt:   &mockTime,
	}

	dataMockList = []types.Task{
		{
			ID:          1,
			Description: "Create Task",
			Completed:   false,
			CreatedAt:   &mockTime,
			UpdatedAt:   &mockTime,
		},
	}
)

func TestNewTodoUsecase(t *testing.T) {
	type args struct {
		todoRepository *todoRepository.TodoRepository
	}
	tests := []struct {
		name string
		args args
		want TodoUsecaseInterface
	}{
		{
			name: "",
			args: args{
				todoRepository: &todoRepository.TodoRepository{},
			},
			want: &TodoUsecase{
				&todoRepository.TodoRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodoUsecase(tt.args.todoRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoUsecase_Create(t *testing.T) {
	todoUsecase := newTodoUsecaseMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})

	defer monkey.UnpatchAll()

	type fields struct {
		TodoRepository todoRepository.TodoRepositoryInterface
	}
	type args struct {
		ctx  context.Context
		data types.Task
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult *types.Task
		wantErr    bool
		mock       func()
	}{
		{
			name: "Success Create Task",
			fields: fields{
				TodoRepository: todoUsecase.TodoRepository,
			},
			args: args{
				ctx:  ctx,
				data: dataMock,
			},
			wantResult: &dataMock,
			wantErr:    false,
			mock: func() {
				todoUsecase.TodoRepository.On("Create", ctx, dataMock).Return(dataMock.ID, nil).Times(1)
				todoUsecase.TodoRepository.On("GetByID", ctx, dataMock.ID).Return(&dataMock, nil).Times(1)
			},
		},
		{
			name: "Failed Create Task",
			fields: fields{
				TodoRepository: todoUsecase.TodoRepository,
			},
			args: args{
				ctx:  ctx,
				data: dataMock,
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				todoUsecase.TodoRepository.On("Create", ctx, dataMock).Return(dataMock.ID, fmt.Errorf("asdf")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()

		t.Run(tt.name, func(t *testing.T) {

			tuc := &TodoUsecase{
				TodoRepository: tt.fields.TodoRepository,
			}
			gotResult, err := tuc.Create(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TodoUsecase.Create() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTodoUsecase_GetByID(t *testing.T) {
	todoUsecase := newTodoUsecaseMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})

	defer monkey.UnpatchAll()

	type fields struct {
		TodoRepository todoRepository.TodoRepositoryInterface
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult *types.Task
		wantErr    bool
		mock       func()
	}{
		{
			name: "Success Retrive Data By ID",
			fields: fields{
				TodoRepository: todoUsecase.TodoRepository,
			},
			args: args{
				ctx: ctx,
				id:  int64(1),
			},
			wantResult: &dataMock,
			wantErr:    false,
			mock: func() {
				todoUsecase.TodoRepository.On("GetByID", ctx, int64(1)).Return(&dataMock, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tuc := &TodoUsecase{
				TodoRepository: tt.fields.TodoRepository,
			}
			gotResult, err := tuc.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TodoUsecase.GetByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTodoUsecase_GetAll(t *testing.T) {
	todoUsecase := newTodoUsecaseMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})

	defer monkey.UnpatchAll()

	type fields struct {
		TodoRepository todoRepository.TodoRepositoryInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []types.Task
		wantErr    bool
		mock       func()
	}{
		{
			name: "Success Retrive All Data",
			fields: fields{
				TodoRepository: todoUsecase.TodoRepository,
			},
			args: args{
				ctx: ctx,
			},
			wantResult: dataMockList,
			wantErr:    false,
			mock: func() {
				todoUsecase.TodoRepository.On("GetAllTaskDB", ctx).Return(dataMockList, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()

		t.Run(tt.name, func(t *testing.T) {
			tuc := &TodoUsecase{
				TodoRepository: tt.fields.TodoRepository,
			}
			gotResult, err := tuc.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TodoUsecase.GetAll() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTodoUsecase_Update(t *testing.T) {
	todoUsecaseMock := newTodoUsecaseMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})

	defer monkey.UnpatchAll()

	type fields struct {
		TodoRepository todoRepository.TodoRepositoryInterface
	}
	type args struct {
		ctx  context.Context
		id   int64
		data types.Task
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult *types.Task
		wantErr    bool
		mock       func()
	}{
		{
			name: "Success Update Task",
			fields: fields{
				TodoRepository: todoUsecaseMock.TodoRepository,
			},
			args: args{
				ctx:  ctx,
				id:   int64(1),
				data: dataMock,
			},
			wantResult: &dataMock,
			wantErr:    false,
			mock: func() {
				todoUsecaseMock.TodoRepository.On("UpdateByIDDB", ctx, int64(1), dataMock).Return(nil)
				todoUsecaseMock.TodoRepository.On("GetByID", ctx, int64(1)).Return(&dataMock, nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tuc := &TodoUsecase{
				TodoRepository: tt.fields.TodoRepository,
			}
			gotResult, err := tuc.Update(tt.args.ctx, tt.args.id, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TodoUsecase.Update() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTodoUsecase_Delete(t *testing.T) {
	todoUsecaseMock := newTodoUsecaseMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	type fields struct {
		TodoRepository todoRepository.TodoRepositoryInterface
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "Success Delete Task",
			fields: fields{
				TodoRepository: todoUsecaseMock.TodoRepository,
			},
			args: args{
				ctx: ctx,
				id:  int64(1),
			},
			wantErr: false,
			mock: func() {
				todoUsecaseMock.TodoRepository.On("GetByID", ctx, int64(1)).Return(&dataMock, nil)
				todoUsecaseMock.TodoRepository.On("DeleteByIDDB", ctx, int64(1)).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tuc := &TodoUsecase{
				TodoRepository: tt.fields.TodoRepository,
			}
			if err := tuc.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
