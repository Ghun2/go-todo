package model

import "time"

type memoryHandler struct {
	todoMap map[int]*Todo
	lstId int
}

func (m *memoryHandler) GetTodos(sessionId string) []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddTodo(sessionId, name string) *Todo {
	m.lstId += 1
	id := m.lstId
	todo := &Todo{id, name, false, time.Now()}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler) CompleteTodo(id int, complete bool) bool {
	if todo, ok := m.todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	m.lstId = 0
	return m
}

func (m *memoryHandler) Close() {

}
