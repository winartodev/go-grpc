package util

import (
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/winartodev/go-grpc/types"
	"github.com/winartodev/protobuff-collections/todolist"
)

func TestTransformTaskData(t *testing.T) {
	mockTime := time.Date(2020, 10, 25, 0, 0, 0, 0, time.UTC)
	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})

	defer monkey.UnpatchAll()

	rpcData := &todolist.Task{
		Id:          1,
		Description: "Description",
		Completed:   false,
		CreatedAt:   mockTime.Unix(),
		UpdatedAt:   mockTime.Unix(),
	}

	unixTime := time.Unix(mockTime.Unix(), 0)

	taskData := types.Task{
		ID:          1,
		Description: "Description",
		Completed:   false,
		CreatedAt:   &unixTime,
		UpdatedAt:   &unixTime,
	}

	type args struct {
		rpcdata *todolist.Task
	}
	tests := []struct {
		name       string
		args       args
		wantResult types.Task
	}{
		{
			name: "Success Transform RPC to Data",
			args: args{
				rpcdata: rpcData,
			},
			wantResult: taskData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := TransformTaskData(tt.args.rpcdata); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TransformTaskData() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestTransformTaskDataRPC(t *testing.T) {
	mockTime := time.Date(2020, 10, 25, 0, 0, 0, 0, time.UTC)
	monkey.Patch(time.Now, func() time.Time {
		return mockTime
	})

	defer monkey.UnpatchAll()

	rpcData := &todolist.Task{
		Id:          1,
		Description: "Description",
		Completed:   false,
		CreatedAt:   mockTime.Unix(),
		UpdatedAt:   mockTime.Unix(),
	}

	unixTime := time.Unix(mockTime.Unix(), 0)

	taskData := types.Task{
		ID:          1,
		Description: "Description",
		Completed:   false,
		CreatedAt:   &unixTime,
		UpdatedAt:   &unixTime,
	}

	type args struct {
		data *types.Task
	}
	tests := []struct {
		name       string
		args       args
		wantResult *todolist.Task
	}{
		{
			name: "Success Transform Data to RPC",
			args: args{
				data: &taskData,
			},
			wantResult: rpcData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := TransformTaskDataRPC(tt.args.data); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TransformTaskDataRPC() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
