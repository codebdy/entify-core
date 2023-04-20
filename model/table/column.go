package table

import "github.com/codebdy/entify/model/meta"

type Column struct {
	meta.AttributeMeta
	Key bool
}
