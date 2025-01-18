package nodes

import lgt "gitlab.com/manifold555112/manifold/lib/graph/types"
import at "gitlab.com/manifold555112/manifold/server/api/types"
import pr "gitlab.com/manifold555112/manifold/lib/graph/types/proto"

type ParseCsv struct {
	lgt.BaseNode
}

func (p *ParseCsv) Execute() *at.ErrorResponse {
	return nil
}

func (p *ParseCsv) New(n lgt.Node) *ParseCsv {
	p.BaseNode = lgt.BaseNode{
		Id:       n.GetId(),
		Type:     n.GetType(),
		Data:     n.GetData(),
		Position: n.GetPosition(),
		Origin:   n.GetOrigin(),
		Selected: n.GetSelected(),
	}
	p.Sources = []lgt.SourceHandle{lgt.NewSourceHandle(1, pr.IOType_MAP, "default")}
	p.Targets = []lgt.TargetHandle{lgt.NewTargetHandle(1, []pr.IOType{pr.IOType_STRING}, "default")}
	return p
}
