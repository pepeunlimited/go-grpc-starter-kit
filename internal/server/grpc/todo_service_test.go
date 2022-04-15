package grpc

import (
	"context"
	"github.com/gogo/status"
	"testing"

	"github.com/pepeunlimited/go-grpc-starter-kit/pkg/api/v1/resources"
	"github.com/pepeunlimited/go-grpc-starter-kit/pkg/api/v1/services"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestTodoService_CreateTodo(t *testing.T) {
	ctx := context.TODO()
	server := NewTodoServer()
	result, err := server.CreateTodo(ctx, &services.CreateTodoRequest{
		Title: "helloWorld",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if result.Id == 0 {
		t.FailNow()
	}
	if result.Title != "helloWorld" {
		t.FailNow()
	}
	if len(server.keysOrder) != 1 {
		t.FailNow()
	}
}

func TestTodoService_GetTodo(t *testing.T) {
	ctx := context.TODO()
	server := NewTodoServer()
	var todoid int64
	if result, err := server.CreateTodo(ctx, &services.CreateTodoRequest{
		Title: "helloWorld",
	}); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		todoid = result.Id
	}
	result, err := server.GetTodo(ctx, &services.GetTodoRequest{
		Id: todoid,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if result.Id != todoid {
		t.FailNow()
	}
}

func TestTodoService_GetTodoNotFound(t *testing.T) {
	ctx := context.TODO()
	server := NewTodoServer()
	todoid := int64(1)
	if _, err := server.GetTodo(ctx, &services.GetTodoRequest{
		Id: todoid,
	}); err == nil {
		t.FailNow()
	} else {
		grcperr, ok := status.FromError(err)
		if !ok {
			t.Error(err)
			t.FailNow()
		}
		if grcperr.Message() != "Todo entity not exist" {
			t.Error(err)
			t.FailNow()
		}
	}
}

func TestTodoService_DeleteTodo(t *testing.T) {
	ctx := context.TODO()
	server := NewTodoServer()
	var todoid int64
	if result, err := server.CreateTodo(ctx, &services.CreateTodoRequest{
		Title: "helloWorld",
	}); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		todoid = result.Id
	}
	if _, err := server.DeleteTodo(ctx, &services.DeleteTodoRequest{
		TodoId: todoid,
	}); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if _, err := server.GetTodo(ctx, &services.GetTodoRequest{
		Id: todoid,
	}); err == nil {
		t.Error("GetTodo err should not be nil")
		t.FailNow()
	} else {
		if grpcErr, ok := status.FromError(err); ok {
			if grpcErr.Message() != "Todo entity not exist" {
				t.Error(err)
				t.FailNow()
			}
		} else {
			t.Error(err)
			t.FailNow()
		}

	}
	if len(server.keysOrder) != 0 {
		t.FailNow()
	}
}

func TestTodoService_UpdateTodo(t *testing.T) {
	ctx := context.TODO()
	server := NewTodoServer()
	var todoid int64
	if result, err := server.CreateTodo(ctx, &services.CreateTodoRequest{
		Title: "helloWorld",
	}); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		todoid = result.Id
	}
	updated, err := server.UpdateTodo(ctx, &services.UpdateTodoRequest{
		Operation: &resources.Todo{
			Id:     todoid,
			Title:  "newHelloWorld",
			IsDone: true, // this value is not updated => fieldMask: "title"
		},
		FieldMask: &fieldmaskpb.FieldMask{
			Paths: []string{
				"title",
			},
		},
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	result, err := server.GetTodo(ctx, &services.GetTodoRequest{
		Id: todoid,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if result.Id != todoid {
		t.FailNow()
	}
	if updated.Title != result.Title {
		t.Errorf("update.Title: %s", updated.Title)
		t.Errorf("result.Title: %s", result.Title)
	}
	// should not been updated
	if updated.IsDone {
		t.FailNow()
	}
	// should not been updated
	if result.IsDone {
		t.FailNow()
	}
}

func TestTodoService_generateId(t *testing.T) {
	server := NewTodoServer()
	if server.generateId() == 0 {
		t.FailNow()
	}
}

func TestTodoService_ListTodo(t *testing.T) {
	ctx := context.TODO()
	server := NewTodoServer()
	if _, err := server.CreateTodo(ctx, &services.CreateTodoRequest{
		Title: "helloWorld",
	}); err != nil {
		t.Error(err)
		t.FailNow()
	}
	result, err := server.ListTodo(ctx, &services.ListTodoRequest{
		PageSize: 20,
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(result.Todos) != 1 {
		t.FailNow()
	}
}
