package graph

import (
	"fmt"
	"strconv"

	"gitlab.com/manifold555112/manifold/lib"
	l "gitlab.com/manifold555112/manifold/lib"
	gn "gitlab.com/manifold555112/manifold/lib/graph/nodes"
	lgt "gitlab.com/manifold555112/manifold/lib/graph/types"
	at "gitlab.com/manifold555112/manifold/server/api/types"
)

type Graph struct {
	Id       string     `json:"id,omitempty"`
	Nodes    []lgt.Node `json:"nodes,omitempty"`
	Edges    []lgt.Edge `json:"edges,omitempty"`
	Channels []chan lgt.Message
}

func (g *Graph) GetNodeById(id string) lgt.Node {
	for _, n := range g.Nodes {
		fmt.Println(id)
		fmt.Printf("%v", n)
		if n.GetId() == id {
			return n
		}
	}
	return nil
}

func (g *Graph) PushNode(n lgt.Node) {
	g.Nodes = append(g.Nodes, n)
}
func (g *Graph) PushEdge(n lgt.Edge) {
	g.Edges = append(g.Edges, n)
}

func (g *Graph) Validate() *at.ErrorResponse {
	for _, e := range g.Edges {
		target := g.GetNodeById(e.GetTargetNodeId())
		if target == nil {
			return &at.ErrorResponse{
				Error:          l.Ptr("Target node: " + e.GetTargetNodeId() + " not found."),
				ErrorCode:      l.Ptr(at.GRAPH_NODE_NOT_FOUND),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
		source := g.GetNodeById(e.GetSourceNodeId())
		if source == nil {
			return &at.ErrorResponse{
				Error:          lib.Ptr("Source node: " + e.GetSourceNodeId() + " not found."),
				ErrorCode:      lib.Ptr(at.GRAPH_NODE_NOT_FOUND),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
		targetHandel := target.GetTargetHandleById(e.GetTargetHandleId())
		if targetHandel == nil {
			return &at.ErrorResponse{
				Error:          lib.Ptr(fmt.Sprintf("Target handle: " + e.GetTargetHandleId() + " not found.")),
				ErrorCode:      lib.Ptr(at.GRAPH_HANDLE_NOT_FOUND),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
		sourceHandle := source.GetSourceHandleById(e.GetSourceHandleId())
		if sourceHandle == nil {
			return &at.ErrorResponse{
				Error:          lib.Ptr("Source handle: " + e.GetSourceHandleId() + " not found."),
				ErrorCode:      lib.Ptr(at.GRAPH_HANDLE_NOT_FOUND),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
		if !targetHandel.HandlesInput(sourceHandle.CreatesOutput()) {
			tagetHandleIdString := targetHandel.GetId()
			sourceHandleIdString := sourceHandle.GetId()
			sourceHandleTypeString := sourceHandle.CreatesOutput()
			return &at.ErrorResponse{
				Error:          lib.Ptr("Target handle: " + tagetHandleIdString + " doesnt handle output: " + strconv.Itoa(int(sourceHandleTypeString)) + " of source handle: " + sourceHandleIdString),
				ErrorCode:      lib.Ptr(at.INCOMPATIBLE_HANDLE_IO_TYPES),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
	}
	return nil
}

func (g *Graph) Build() *at.ErrorResponse {
	for i, n := range g.Nodes {
		g.Nodes[i] = gn.RebuildNode(n)
	}
	return nil
}

func (g *Graph) Execute() *at.ErrorResponse {
	for _, e := range g.Edges {
		sourceNode := g.GetNodeById(e.GetSourceNodeId())
		targetNode := g.GetNodeById(e.GetTargetHandleId())
		if sourceNode == nil {
			return &at.ErrorResponse{
				Error:          lib.Ptr("Source node: " + e.GetSourceNodeId() + " not found."),
				ErrorCode:      lib.Ptr(at.GRAPH_NODE_NOT_FOUND),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
		if targetNode == nil {
			return &at.ErrorResponse{
				Error:          lib.Ptr("Target node: " + e.GetTargetNodeId() + " not found."),
				ErrorCode:      lib.Ptr(at.GRAPH_NODE_NOT_FOUND),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
		targetHandle := targetNode.GetTargetHandleById(e.GetTargetHandleId())
		sourceHandle := sourceNode.GetSourceHandleById(e.GetSourceHandleId())
		if targetHandle == nil {
			return &at.ErrorResponse{
				Error:          lib.Ptr("Target handle: " + e.GetTargetHandleId() + " not found."),
				ErrorCode:      lib.Ptr(at.GRAPH_NODE_NOT_FOUND),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
		if sourceHandle == nil {
			return &at.ErrorResponse{
				Error:          lib.Ptr("Source Handle: " + e.GetSourceHandleId() + " not found."),
				ErrorCode:      lib.Ptr(at.GRAPH_NODE_NOT_FOUND),
				EntityId:       lib.Ptr(e.GetId()),
				ExpectedSchema: nil,
			}
		}
		c := make(chan lgt.Message)
		g.Channels = append(g.Channels, c)
		sourceHandle.PushChannel(c)
		targetHandle.PushChannel(c)
	}
	return nil
}
