package myeval

import (
	"fmt"
	"math"
	"testing"
)

type Var string
func (v Var)Eval(env Env)float64{
	return env[v]
}

type literal float64
func (l literal)Eval(_ Env)float64{
	return float64(l)
}

type unary struct {
	op rune
	x Expr
}
func (u unary)Eval(env Env)float64{
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary:%q",u.op))
}

type binary struct {
	op rune
	x,y Expr
}
func (b binary)Eval(env Env)float64{
	switch b.op {
	case '+':
		return b.x.Eval(env)+b.y.Eval(env)
	case '-':
		return b.x.Eval(env)-b.y.Eval(env)
	case '*':
		return b.x.Eval(env)*b.y.Eval(env)
	case '/':
		return b.x.Eval(env)/b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary:%q",b.op))
}

type call struct {
	fn string
	args []Expr
}
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

type Env map[Var]float64

type Expr interface {
	Eval(evn Env)float64
	Check(vars map[Var]bool) error
}

func TestEval(t *testing.T){
	tests:=[]struct{
		expr string
		env Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	var preExpr string
	for _,test := range tests{
		if test.expr!=preExpr{
			fmt.Printf("\n%s\n",test.expr)
			preExpr=test.expr
		}
		expr,err:=Parse(test.expr)
		if err!=nil{
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{"x % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, want 1"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}
