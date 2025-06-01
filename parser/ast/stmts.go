package ast

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type ExprStmt struct {
	Expr Expr
}

func (e ExprStmt) stmt() {}

type VarStmt struct {
	Name     string
	Expr     *Expr
	DataType *Type
	Const    bool
}

func (v VarStmt) stmt() {}

type IfStmt struct {
	Condition  Expr
	Body       BlockStmt
	Consequent *IfStmt
}

func (i IfStmt) stmt() {}

type FunStmt struct {
	Name    string
	Params  map[string]Type
	Body    BlockStmt
	Returns Type
}

func (f FunStmt) stmt() {}

type ReturnStmt struct {
	Expr *Expr
}

func (r ReturnStmt) stmt() {}