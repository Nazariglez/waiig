// Created by nazarigonzalez on 21/8/17.

package evaluator

import (
	"fmt"
	"strings"
	"waiig/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: fnLenBuiltin,
	},
	"first": &object.Builtin{
		Fn: fnFirstBuiltin,
	},
	"last": &object.Builtin{
		Fn: fnLastBuiltin,
	},
	"rest": &object.Builtin{
		Fn: fnRestBuiltin,
	},
	"push": &object.Builtin{
		Fn: fnPushBuiltin,
	},
	"puts": &object.Builtin{
		Fn: fnPutsBuiltin,
	},
}

func fnPutsBuiltin(args ...object.Object) object.Object {
	text := []string{}
	for _, arg := range args {
		text = append(text, arg.Inspect())
	}

	fmt.Println(strings.Join(text, " "))
	return NULL
}

func fnLenBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {

	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}

	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}

	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())

	}
}

func fnFirstBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY, got=%s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func fnLastBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}

func fnRestBuiltin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func fnPushBuiltin(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong numner of arguments. got=%d, want=2", args[0].Type())
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}
