package vm

import (
	"errors"
	"fmt"
	"reflect"
	"sort"

	"github.com/tesujiro/goa/ast"
)

var (
	ErrBreak    = errors.New("unexpected break")
	ErrContinue = errors.New("unexpected continue")
	ErrReturn   = errors.New("unexpected return")
)

func runStmts(stmts []ast.Stmt, env *Env) (interface{}, error) {

	if result, err := run(stmts, env); err == ErrReturn {
		return result, nil
	} else {
		return result, err
	}
}

func run(stmts []ast.Stmt, env *Env) (interface{}, error) {
	//fmt.Println("run -> env.Dump()")
	//env.Dump()
	var result interface{}
	var err error
	for _, stmt := range stmts {
		switch stmt.(type) {
		case *ast.BreakStmt:
			return nil, ErrBreak
		case *ast.ContinueStmt:
			return nil, ErrContinue
		case *ast.ReturnStmt:
			result, err = runSingleStmt(stmt, env)
			if err != nil && err != ErrReturn {
				return nil, err
			}
			return result, ErrReturn
		default:
			result, err = runSingleStmt(stmt, env)
			if err != nil && err != ErrReturn {
				return nil, err
			}
		}
	}
	return result, err
}

func runSingleStmt(stmt ast.Stmt, env *Env) (interface{}, error) {
	switch stmt.(type) {
	case *ast.AssStmt:
		assStmt := stmt.(*ast.AssStmt)
		left, right := assStmt.Left, assStmt.Right

		// evaluate right expressions
		right_values := make([]interface{}, len(right))
		var err error
		for i, expr := range right {
			right_values[i], err = evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
		}

		// evaluate assExpr
		switch {
		case len(left) == 1 && len(right) == 1:
			return evalAssExpr(left[0], right_values[0], env)
		case len(left) > 1 && len(right) == 1:
			val := right_values[0]
			if reflect.ValueOf(val).Kind() == reflect.Interface {
				val = reflect.ValueOf(val).Elem().Interface()
			}
			if reflect.ValueOf(val).Kind() != reflect.Slice {
				return nil, errors.New("single value assign to multi values")
			} else {
				elements := reflect.ValueOf(val)
				right_values = make([]interface{}, elements.Len())
				for i := 0; i < elements.Len(); i++ {
					right_values[i] = elements.Index(i).Interface()
				}
			}
			fallthrough
		default:
			for i, expr := range left {
				if i >= len(right_values) {
					return right_values[len(right_values)-1], nil
				}
				if _, err := evalAssExpr(expr, right_values[i], env); err != nil {
					return nil, err
				}
			}
			return right_values[len(left)-1], nil
		}
	case *ast.ExprStmt:
		return evalExpr(stmt.(*ast.ExprStmt).Expr, env)
	case *ast.DelStmt:
		expr := stmt.(*ast.DelStmt).Expr
		//fmt.Println("TypeOf(Expr):", reflect.TypeOf(expr))
		var id string
		var index string
		switch expr.(type) {
		case *ast.IdentExpr:
			id = expr.(*ast.IdentExpr).Literal
		case *ast.ItemExpr:
			e := expr.(*ast.ItemExpr).Expr
			ie, ok := e.(*ast.IdentExpr)
			if !ok {
				return nil, errors.New("non variable does not support delete operation")
			}
			id = ie.Literal
			var err error
			index, err = getHashIndex(env, expr.(*ast.ItemExpr).Index)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("type %s does not support delete operation", reflect.TypeOf(expr))
		}
		val, err := env.Get(id)
		if err == ErrUnknownSymbol {
			// Set Default Map to env
			val, err = env.DefineDefaultMapValue(id, index)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		if reflect.ValueOf(val).Kind() == reflect.Map {
			m := val.(map[interface{}]interface{})
			if index == "" {
				// Delete All Map Elements
				_, err := env.DefineDefaultMapValue(id, index)
				if err != nil {
					return nil, err
				}
			} else {
				// Delete Map Element
				delete(m, index)
			}
		} else {
			// Error
			return nil, fmt.Errorf("type %s does not support delete operation", reflect.ValueOf(val).Kind())
		}
		return nil, nil
	case *ast.PrintStmt:
		printStmt := stmt.(*ast.PrintStmt)
		for i, expr := range printStmt.Exprs {
			result, err := evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
			if 0 < i && i < len(printStmt.Exprs) {
				fmt.Printf("%v", env.builtin.OFS)
			}
			//fmt.Printf("%v", result)
			switch reflect.ValueOf(result).Kind() {
			case reflect.Int, reflect.Float64, reflect.Bool, reflect.String:
				fmt.Printf("%v", result)
			case reflect.Slice:
				len := reflect.ValueOf(result).Len()
				for i := 0; i < len; i++ {
					if i > 0 {
						fmt.Printf("%v", env.builtin.OFS)
					}
					fmt.Printf("%v", reflect.ValueOf(result).Index(i))
				}
			case reflect.Invalid:
				fmt.Printf("")
			default:
				return nil, fmt.Errorf("type %s does not support print operation", reflect.ValueOf(result).Kind().String())
			}
		}
		fmt.Printf("%v", env.builtin.ORS)
	case *ast.IfStmt:
		child := env.NewEnv()
		result, err := evalExpr(stmt.(*ast.IfStmt).If, child)
		if err != nil {
			return nil, err
		}
		if result.(bool) {
			//fmt.Println("If then -> env.Dump()")
			//child.Dump()
			result, err = run(stmt.(*ast.IfStmt).Then, child)
			if err != nil {
				return result, err
			}
			return result, nil
		}
		for _, stmt := range stmt.(*ast.IfStmt).ElseIf {
			result, err := evalExpr(stmt.(*ast.IfStmt).If, child)
			if err != nil {
				return nil, err
			}
			if result.(bool) {
				result, err = run(stmt.(*ast.IfStmt).Then, child)
				if err != nil {
					return result, err
				}
				return result, nil
			}
		}

		if len(stmt.(*ast.IfStmt).Else) > 0 {
			result, err = run(stmt.(*ast.IfStmt).Else, child)
			if err != nil {
				return result, err
			}
		}
		return result, nil
	case *ast.ReturnStmt:
		returnStmt := stmt.(*ast.ReturnStmt)
		length := len(returnStmt.Exprs)

		resultExpr := make([]interface{}, length)
		var err error
		for i, expr := range returnStmt.Exprs {
			resultExpr[i], err = evalExpr(expr, env)
			if err != nil {
				return nil, err
			}
		}

		switch length {
		case 0:
			return nil, nil
		case 1:
			return resultExpr[0], nil
		default:
			return resultExpr, nil
		}
	case *ast.LoopStmt:
		newEnv := env.NewEnv()
		for {
			exp := stmt.(*ast.LoopStmt).Expr
			if exp != nil {
				if result, err := evalExpr(exp, newEnv); err != nil {
					return nil, err
				} else if b, ok := result.(bool); !ok {
					return nil, fmt.Errorf("for condition type %s cannot convert to bool", reflect.TypeOf(result).Kind())
				} else if !b {
					break
				}
			}

			ret, err := run(stmt.(*ast.LoopStmt).Stmts, newEnv)
			if err == ErrReturn {
				return ret, nil
			}
			if err == ErrBreak {
				break
			}
			if err == ErrContinue {
				continue
			}
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case *ast.HashLoopStmt:
		key := stmt.(*ast.HashLoopStmt).Key
		hash := stmt.(*ast.HashLoopStmt).Hash
		stmts := stmt.(*ast.HashLoopStmt).Stmts
		v, err := env.Get(hash)
		if err != nil {
			return nil, err
		}
		if reflect.TypeOf(v).Kind() != reflect.Map {
			return nil, fmt.Errorf("for key loop not in associated array,%s", reflect.TypeOf(v).Kind())
		}
		// sort hash keys
		m := v.(map[interface{}]interface{})
		indecies := make([]string, len(m))
		i := 0
		for k, _ := range m {
			indecies[i] = k.(string)
			i++
		}
		sort.Strings(indecies)

		newEnv := env.NewEnv()
		for _, index := range indecies {
			if err := newEnv.Set(key, index); err == ErrUnknownSymbol {
				if err := newEnv.Define(key, index); err != nil {
					return nil, err
				}
			} else if err != nil {
				return nil, err
			}
			ret, err := run(stmts, newEnv)

			if err == ErrReturn {
				return ret, nil
			}
			if err == ErrBreak {
				break
			}
			if err == ErrContinue {
				continue
			}
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}
	return nil, nil
}
