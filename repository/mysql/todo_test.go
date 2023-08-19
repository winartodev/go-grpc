package mysql

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/winartodev/go-grpc/types"
)

var (
	mockTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	dataMock = types.Task{
		ID:          1,
		Description: "Test",
		Completed:   false,
		CreatedAt:   &mockTime,
		UpdatedAt:   &mockTime,
	}
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("%s", err)
	}

	return db, mock
}

func TestNewTodoRepository(t *testing.T) {
	db, _ := NewMock()

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want TodoRepositoryInterface
	}{
		{
			name: "Success Call Todo Repository",
			args: args{
				db: db,
			},
			want: &TodoRepository{
				DB: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodoRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_Create(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx  context.Context
		data types.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
		mock    func()
	}{
		{
			name: "Success Create Task",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				data: dataMock,
			},
			wantId:  int64(1),
			wantErr: false,
			mock: func() {
				dbmock.ExpectPrepare(regexp.QuoteMeta(CreateTaskQuery)).
					ExpectExec().
					WithArgs(dataMock.Description, dataMock.Completed, dataMock.CreatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()

		t.Run(tt.name, func(t *testing.T) {
			tr := &TodoRepository{
				DB: tt.fields.DB,
			}
			gotId, err := tr.Create(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("TodoRepository.Create() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestTodoRepository_GetByID(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	type fields struct {
		DB *sql.DB
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
			name: "Success Get Task By ID",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: ctx,
				id:  int64(1),
			},
			wantResult: &dataMock,
			wantErr:    false,
			mock: func() {
				dbmock.ExpectQuery(regexp.QuoteMeta(GetTaskByID)).
					WithArgs(dataMock.ID).
					WillReturnRows(
						dbmock.NewRows([]string{"id", "description", "complete", "created_at", "updated_at"}).
							AddRow(dataMock.ID, dataMock.Description, dataMock.Completed, dataMock.CreatedAt, dataMock.UpdatedAt),
					)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TodoRepository{
				DB: tt.fields.DB,
			}
			gotResult, err := tr.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TodoRepository.GetByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTodoRepository_GetAllTaskDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	type fields struct {
		DB *sql.DB
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
			name: "Success Retrive All Task",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: ctx,
			},
			wantResult: []types.Task{
				dataMock,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectQuery(regexp.QuoteMeta(GetAllTask)).
					WillReturnRows(
						dbmock.NewRows([]string{"id", "description", "complete", "created_at", "updated_at"}).
							AddRow(dataMock.ID, dataMock.Description, dataMock.Completed, dataMock.CreatedAt, dataMock.UpdatedAt),
					)
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TodoRepository{
				DB: tt.fields.DB,
			}
			gotResult, err := tr.GetAllTaskDB(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.GetAllTaskDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TodoRepository.GetAllTaskDB() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTodoRepository_UpdateByIDDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx  context.Context
		id   int64
		data types.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "Success Retrive All Task",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				id:   1,
				data: dataMock,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectPrepare(regexp.QuoteMeta(UpdateTaskQuery)).
					ExpectExec().
					WithArgs(dataMock.Description, dataMock.Completed, dataMock.UpdatedAt, int64(1)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TodoRepository{
				DB: tt.fields.DB,
			}
			if err := tr.UpdateByIDDB(tt.args.ctx, tt.args.id, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.UpdateByIDDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoRepository_DeleteByIDDB(t *testing.T) {
	db, dbmock := NewMock()
	ctx := context.Background()

	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})
	defer monkey.UnpatchAll()

	type fields struct {
		DB *sql.DB
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
			name: "Success Retrive All Task",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: false,
			mock: func() {
				dbmock.ExpectPrepare(regexp.QuoteMeta(DeleteTaskQuery)).
					ExpectExec().
					WithArgs(int64(1)).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		defer tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			tr := &TodoRepository{
				DB: tt.fields.DB,
			}
			if err := tr.DeleteByIDDB(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.DeleteByIDDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
