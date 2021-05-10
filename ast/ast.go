package ast

import (
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementeNode()
}

type Expression interface {
	Node
	expressionNode()
}
