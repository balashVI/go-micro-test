package services

import (
	"errors"
	"github.com/balashVI/go-micro-test/proto"
	"golang.org/x/net/context"
	"log"
)

// TODO: add store
type ToDoService struct {
	toDos []*proto.ToDo
}

func (_ ToDoService) Ping(ctx context.Context, _ *proto.Empty, rsp *proto.PingResponse) error {
	log.Println("Ping")

	rsp.Message = "Pong"
	return nil
}

func (s *ToDoService) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	log.Println("List", req)

	if req.Count > int32(0) {
		startIndex := req.Page * req.Count
		endIndex := startIndex + req.Count

		if endIndex > int32(len(s.toDos)) {
			return errors.New("Out of the range")
		}

		rsp.Todos = s.toDos[startIndex:endIndex]
	} else {
		rsp.Todos = s.toDos
	}
	return nil
}

func (s *ToDoService) Get(ctx context.Context, req *proto.GetRequest, todo *proto.ToDo) error {
	log.Println("Get", req)

	index := -1
	for i, toDo := range s.toDos {
		if toDo.Id == req.Id {
			index = i

			break
		}
	}

	if index == -1 {
		return errors.New("Not found")
	}

	*todo = *s.toDos[index]
	return nil
}

func (s *ToDoService) Add(ctx context.Context, newToDo *proto.ToDo, _ *proto.Empty) error {
	log.Println("Add", newToDo)

	if len(newToDo.Message) == 0 {
		return errors.New("Message is empty")
	}

	newToDo.Id = int64(len(s.toDos))
	s.toDos = append(s.toDos, newToDo)

	return nil
}

func (s *ToDoService) Update(ctx context.Context, updateToDo *proto.ToDo, _ *proto.Empty) error {
	log.Println("Update", updateToDo)

	index := -1

	for i, toDo := range s.toDos {
		if toDo.Id == updateToDo.Id {
			index = i
			*toDo = *updateToDo
			break
		}
	}

	if index == -1 {
		return errors.New("Not found")
	} else {
		*s.toDos[index] = *updateToDo
		return nil
	}
}

func (s *ToDoService) Delete(ctx context.Context, req *proto.DeleteRequest, _ *proto.Empty) error {
	log.Println("Delete", req)

	index := -1
	for i, toDo := range s.toDos {
		if toDo.Id == req.Id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("Not found")
	}

	s.toDos = append(s.toDos[:index], s.toDos[index+1:]...)

	return nil
}
