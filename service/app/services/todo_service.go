package services

import (
	"github.com/balashVI/go-micro-test/proto"
	"golang.org/x/net/context"
)

// TODO: add store
type ToDoService struct {
	toDos []*proto.ToDo
}

func (_ ToDoService) Ping(ctx context.Context, _ *proto.Empty, rsp *proto.PingResponse) error {
	rsp.Message = "Pong"
	return nil
}

func (s *ToDoService) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	startIndex := req.Page * req.Count
	endIndex := startIndex + req.Count
	rsp.Todos = s.toDos[startIndex:endIndex]
	return nil
}

func (s *ToDoService) Get(ctx context.Context, req *proto.GetRequest, todo *proto.ToDo) error {
	for _, toDo := range s.toDos {
		if toDo.Id == req.Id {
			todo = toDo
			break
		}
	}
	return nil
}

func (s *ToDoService) Add(ctx context.Context, addToDO *proto.ToDo, resToDo *proto.ToDo) error {
	s.toDos = append(s.toDos, addToDO)
	return nil
}

func (s *ToDoService) Update(ctx context.Context, updateToDo *proto.ToDo, resToDo *proto.ToDo) error {
	for _, toDo := range s.toDos {
		if toDo.Id == updateToDo.Id {
			*toDo = *updateToDo
			break
		}
	}
	return nil
}

func (s *ToDoService) Delete(ctx context.Context, req *proto.DeleteRequest, _ *proto.Empty) error {
	index := -1
	for i, toDo := range s.toDos {
		if toDo.Id == req.Id {
			index = i
			break
		}
	}

	if index != -1 {
		s.toDos = append(s.toDos[:index], s.toDos[index+1:]...)
	}

	return nil
}
