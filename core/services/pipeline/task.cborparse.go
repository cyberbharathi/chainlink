package pipeline

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink/core/store/models"
)

type CBORParseTask struct {
	BaseTask `mapstructure:",squash"`
	Data     string `json:"data"`
}

var _ Task = (*CBORParseTask)(nil)

func (t *CBORParseTask) Type() TaskType {
	return TaskTypeCBORParse
}

func (t *CBORParseTask) Run(_ context.Context, vars Vars, _ JSONSerializable, inputs []Result) (result Result) {
	_, err := CheckInputs(inputs, 0, 1, 0)
	if err != nil {
		return Result{Error: err}
	}

	var data BytesParam

	err = multierr.Combine(
		errors.Wrap(ResolveParam(&data, From(VarExpr(t.Data, vars), Input(inputs, 0))), "data"),
	)
	if err != nil {
		return Result{Error: err}
	}

	parsed, err := models.ParseCBOR([]byte(data))
	if err != nil {
		return Result{Error: errors.Wrapf(ErrBadInput, "CBORParse: %v", err)}
	}
	return Result{Value: parsed}
}
