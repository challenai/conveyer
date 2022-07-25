package graph

import (
	"github.com/challenai/conveyer/pkg/graph/desc"
	"github.com/challenai/conveyer/pkg/graph/meta"
	"github.com/challenai/conveyer/pkg/graph/state"
)

type Graph struct {
	meta.Meta
	desc.Description
	*state.State
}
