package nodes

import (
	lgt "gitlab.com/manifold555112/manifold/lib/graph/types"
)

func RebuildNode(n lgt.Node) lgt.Node {
	switch n.GetType() {
	case lgt.TEXT_INPUT:
		return (&TextInput{}).New(n)
	case lgt.PARSE_CSV:
		return (&ParseCsv{}).New(n)
	case lgt.ADD:
		return nil
	case lgt.REMOVE:
		return nil
	case lgt.MERGE:
		return nil
	case lgt.SET:
		return nil
	case lgt.SELECT:
		return nil
	case lgt.IF:
		return nil
	case lgt.DATAVIEW:
		return nil

	}
	return nil
}
