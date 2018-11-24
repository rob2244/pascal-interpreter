package ast

type Block struct {
	Declarations      []*VarDecl
	CompoundStatement *Compound
}
