package pipeline

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type ETHABIEncodeTask struct {
	BaseTask `mapstructure:",squash"`
	ABI      string `json:"abi"`
	Data     string `json:"data"`
}

var _ Task = (*ETHABIEncodeTask)(nil)

func (t *ETHABIEncodeTask) Type() TaskType {
	return TaskTypeETHABIEncode
}

func (t *ETHABIEncodeTask) Run(_ context.Context, vars Vars, _ JSONSerializable, inputs []Result) (result Result) {
	_, err := CheckInputs(inputs, 0, 1, 0)
	if err != nil {
		return Result{Error: err}
	}

	var (
		inputValues MapParam
		theABI      StringParam
	)
	err = multierr.Combine(
		errors.Wrap(ResolveParam(&inputValues, From(VarExpr(t.Data, vars), JSONWithVarExprs(t.Data, vars, false), Inputs(inputs))), "data"),
		errors.Wrap(ResolveParam(&theABI, From(NonemptyString(t.ABI))), "abi"),
	)
	if err != nil {
		return Result{Error: err}
	}

	parts := strings.Split(string(theABI), ",")
	var args abi.Arguments
	var vals []interface{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		typeAndMaybeName := strings.Split(part, " ")
		for i := range typeAndMaybeName {
			typeAndMaybeName[i] = strings.TrimSpace(typeAndMaybeName[i])
		}
		var typeStr, name string
		switch len(typeAndMaybeName) {
		case 0:
			return Result{Error: errors.New("bad ABI specification, empty argument")}
		case 1:
			typeStr = typeAndMaybeName[0]
		case 2:
			typeStr = typeAndMaybeName[0]
			name = typeAndMaybeName[1]
		default:
			return Result{Error: errors.New("bad ABI specification, too many components in argument")}
		}
		typ, err := abi.NewType(typeStr, "", nil)
		if err != nil {
			return Result{Error: err}
		}

		args = append(args, abi.Argument{Type: typ, Name: name})

		val, exists := inputValues[name]
		if !exists {
			return Result{Error: errors.Errorf("ETHABIEncode: argument '%v' is missing", name)}
		}
		vals = append(vals, val)
	}

	dataBytes, err := args.Pack(vals...)
	if err != nil {
		return Result{Error: err}
	}
	return Result{Value: dataBytes}
}
