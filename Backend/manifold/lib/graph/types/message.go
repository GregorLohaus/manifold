package types

import p "gitlab.com/manifold555112/manifold/lib/graph/types/proto"

type Message interface {
	Value() *p.Value
	Done() bool
}

type BaseMessage struct {
	value *p.Value
	done  bool
}

func (m *BaseMessage) New(value *p.Value, done bool) *BaseMessage {
	m.value = value
	m.done = done
	return m
}

func (m *BaseMessage) Value() *p.Value {
	return m.value
}

func (m *BaseMessage) Done() bool {
	return m.done
}
