package types

import p "gitlab.com/manifold555112/manifold/lib/graph/types/proto"

type TargetHandle interface {
	GetId() string
	SetId(string)
	SetInputTypes([]p.IOType)
	HandlesInput(p.IOType) bool
	GetChannels() []chan Message
	PushChannel(chan Message)
}

func NewTargetHandle(version int, inputs []p.IOType, id string) TargetHandle {
	switch version {
	case 1:
		return &TargetHandleVI{
			id:     id,
			inputs: inputs,
		}
	}
	return nil
}

type TargetHandleVI struct {
	id       string
	inputs   []p.IOType
	channels []chan Message
}

func (s *TargetHandleVI) GetId() string {
	return s.id
}

func (s *TargetHandleVI) SetId(i string) {
	s.id = i
}

func (s *TargetHandleVI) HandlesInput(i p.IOType) bool {
	for _, in := range s.inputs {
		if in == i {
			return true
		}
	}
	return false
}

func (s *TargetHandleVI) SetInputTypes(t []p.IOType) {
	s.inputs = t
}

func (s *TargetHandleVI) GetChannels() []chan Message {
	return s.channels
}

func (s *TargetHandleVI) PushChannel(c chan Message) {
	s.channels = append(s.channels, c)
}

type SourceHandle interface {
	GetId() string
	SetId(string)
	SetOutputType(p.IOType)
	CreatesOutput() p.IOType
	GetChannels() []chan Message
	PushChannel(chan Message)
}

func NewSourceHandle(version int, output p.IOType, id string) SourceHandle {
	switch version {
	case 1:
		return &SourceHandleVI{
			id:     id,
			output: output,
		}
	}
	return nil
}

type SourceHandleVI struct {
	id       string
	output   p.IOType
	channels []chan Message
}

func (s *SourceHandleVI) GetId() string {
	return s.id
}

func (s *SourceHandleVI) SetId(i string) {
	s.id = i
}

func (s *SourceHandleVI) CreatesOutput() p.IOType {
	return s.output
}

func (s *SourceHandleVI) SetOutputType(t p.IOType) {
	s.output = t
}

func (s *SourceHandleVI) GetChannels() []chan Message {
	return s.channels
}

func (s *SourceHandleVI) PushChannel(c chan Message) {
	s.channels = append(s.channels, c)
}
