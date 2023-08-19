package util

import (
	"time"

	"github.com/winartodev/go-grpc/types"
	"github.com/winartodev/protobuff-collections/todolist"
)

func TransformTaskData(rpcdata *todolist.Task) (result types.Task) {

	createdAt := time.Unix(rpcdata.CreatedAt, 0)
	updatedAt := time.Unix(rpcdata.UpdatedAt, 0)

	result = types.Task{
		ID:          rpcdata.Id,
		Description: rpcdata.Description,
		Completed:   rpcdata.Completed,
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
	}

	return result
}

func TransformTaskDataRPC(data *types.Task) (result *todolist.Task) {

	createdAt := data.CreatedAt.Unix()

	var updatedAt int64
	if data.UpdatedAt != nil {
		updatedAt = data.UpdatedAt.Unix()
	}

	result = &todolist.Task{
		Id:          data.ID,
		Description: data.Description,
		Completed:   data.Completed,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return result
}
