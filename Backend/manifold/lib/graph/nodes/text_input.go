package nodes

import (
	"gitlab.com/manifold555112/manifold/lib"
	lgt "gitlab.com/manifold555112/manifold/lib/graph/types"
	p "gitlab.com/manifold555112/manifold/lib/graph/types/proto"
	at "gitlab.com/manifold555112/manifold/server/api/types"
)

type TextInput struct {
	lgt.BaseNode
}

func (t *TextInput) Execute() *at.ErrorResponse {
	if _, ok := t.Data["text"].(string); ok {

	}
	return &at.ErrorResponse{
		Error:          lib.Ptr("Text input is missing data field text"),
		ErrorCode:      lib.Ptr(at.MISSING_NODE_DATA),
		EntityId:       &t.Id,
		ExpectedSchema: nil,
	}
}

func (t *TextInput) New(n lgt.Node) *TextInput {
	t.BaseNode = lgt.BaseNode{
		Id:       n.GetId(),
		Type:     n.GetType(),
		Data:     n.GetData(),
		Position: n.GetPosition(),
		Origin:   n.GetOrigin(),
		Selected: n.GetSelected(),
	}
	t.Sources = []lgt.SourceHandle{lgt.NewSourceHandle(1, p.IOType_STRING, "default")}
	t.Targets = nil
	return t
}

func checkImpl() lgt.Node {
	return &TextInput{}
}
