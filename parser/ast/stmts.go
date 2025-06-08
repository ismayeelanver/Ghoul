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
	Expr Expr
}

func (r ReturnStmt) stmt() {} // BRO IS AFKING

type WhileStmt struct {
	Condition Expr
	Body      BlockStmt
}

func (w WhileStmt) stmt() {}

type DoStmt struct {
	Condition Expr
	Body      BlockStmt
}

func (d DoStmt) stmt() {}

type ForStmt struct {
	Stmt1     Stmt
	Condition Expr
	Stmt2     Stmt
	Body      BlockStmt
}

func (f ForStmt) stmt() {}