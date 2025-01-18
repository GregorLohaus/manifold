package types

import (
	"gitlab.com/manifold555112/manifold/lib"
	at "gitlab.com/manifold555112/manifold/server/api/types"
)

type NodeType string

const (
	TEXT_INPUT NodeType = "TEXT_INPUT"
	PARSE_CSV  NodeType = "PARSE_CSV"
	ADD        NodeType = "ADD"
	REMOVE     NodeType = "REMOVE"
	MERGE      NodeType = "MERGE"
	SET        NodeType = "SET"
	SELECT     NodeType = "SELECT"
	IF         NodeType = "IF"
	DATAVIEW   NodeType = "DATAVIEW"
)

type Node interface {
	GetId() string
	GetType() NodeType
	GetData() map[string]interface{}
	GetSources() []SourceHandle
	GetTargets() []TargetHandle
	GetSourceHandleById(string) SourceHandle
	GetTargetHandleById(string) TargetHandle
	GetPosition() Position
	GetOrigin() []float32
	GetSelected() bool
	Execute() *at.ErrorResponse
}

func implCheck() Node {
	return &BaseNode{}
}

type BaseNode struct {
	Id       string                 `json:"id,omitempty"`
	Type     NodeType               `json:"type,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
	Position Position               `json:"position,omitempty"`
	Origin   []float32              `json:"origin,omitempty"`
	Selected bool                   `json:"selected,omitempty"`
	Sources  []SourceHandle         `json:"omit"`
	Targets  []TargetHandle         `json:"omit"`
}

func (n *BaseNode) Execute() *at.ErrorResponse {
	return &at.ErrorResponse{
		Error:          lib.Ptr("Should not execute base node."),
		ErrorCode:      lib.Ptr(at.TRIED_TO_EXECUTE_BASENODE),
		EntityId:       &n.Id,
		ExpectedSchema: nil,
	}
}

func (n *BaseNode) GetId() string {
	return n.Id
}

func (n *BaseNode) GetType() NodeType {
	return n.Type
}
func (n *BaseNode) GetData() map[string]interface{} {
	return n.Data
}

func (n *BaseNode) GetSources() []SourceHandle {
	return n.Sources
}

func (n *BaseNode) GetPosition() Position {
	return n.Position
}

func (n *BaseNode) GetOrigin() []float32 {
	return n.Origin
}

func (n *BaseNode) GetSelected() bool {
	return n.Selected
}

func (n *BaseNode) GetTargets() []TargetHandle {
	return n.Targets
}

func (n *BaseNode) GetSourceHandleById(i string) SourceHandle {
	for _, sh := range n.Sources {
		if sh != nil && sh.GetId() == i {
			return sh
		}
	}
	return nil
}

func (n *BaseNode) GetTargetHandleById(i string) TargetHandle {
	for _, th := range n.Targets {
		if th != nil && th.GetId() == i {
			return th
		}
	}
	return nil
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NodeBuilder(nodeType NodeType) Node {
	switch nodeType {
	case TEXT_INPUT:
		break
	case PARSE_CSV:
		break
	case ADD:
		break
	case REMOVE:
		break
	case MERGE:
		break
	case SET:
		break
	case SELECT:
		break
	case IF:
		break
	case DATAVIEW:
		break
	}
	return nil
}
