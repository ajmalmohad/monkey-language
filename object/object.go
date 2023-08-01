package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

/**
 * * Integers
 */

type Integer struct {
	Value int64
}

func (i *Integer) Token() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

/**
 * * Boolean
 */

type Boolean struct {
	Value bool
}

func (b *Boolean) Token() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

/**
 * * Null
 */

type Null struct{}

func (n *Null) Token() ObjectType {
	return NULL_OBJ
}

func (b *Null) Inspect() string {
	return "null"
}
