package types

type Edge interface {
	GetId() string
	GetSourceNodeId() string
	GetTargetNodeId() string
	GetSourceHandleId() string
	GetTargetHandleId() string
}

func checkImpl() Edge {
	return &BaseEdge{}
}

type BaseEdge struct {
	Id             string `json:"id,omitempty"`
	SourceNodeId   string `json:"source,omitempty"`
	TargetNodeId   string `json:"target,omitempty"`
	SourceHandleId string `json:"sourceHandle,omitempty"`
	TargetHandleId string `json:"targetHandle,omitempty"`
}

func (e *BaseEdge) GetId() string {
	return e.Id
}

func (e *BaseEdge) GetSourceNodeId() string {
	return e.SourceNodeId
}

func (e *BaseEdge) GetTargetNodeId() string {
	return e.TargetNodeId
}

func (e *BaseEdge) GetSourceHandleId() string {
	return e.SourceHandleId
}

func (e *BaseEdge) GetTargetHandleId() string {
	return e.TargetHandleId
}
