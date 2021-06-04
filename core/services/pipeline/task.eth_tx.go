package pipeline

import (
    "context"

    "github.com/pkg/errors"
    "go.uber.org/multierr"
    "gorm.io/gorm"
)

type ETHCallTask struct {
    BaseTask `mapstructure:",squash"`
    Contract string `json:"contract"`
    Data     string `json:"data"`

    db *gorm.DB
}

var _ Task = (*ETHCallTask)(nil)

func (t *ETHCallTask) Type() TaskType {
    return TaskTypeETHCall
}

func (t *ETHCallTask) Run(_ context.Context, vars Vars, _ JSONSerializable, inputs []Result) (result Result) {
    _, err := CheckInputs(inputs, -1, -1, 0)
    if err != nil {
        return Result{Error: err}
    }

    var (
        contractAddr AddressParam
        data         BytesParam
    )
    err = multierr.Combine(
        errors.Wrap(ResolveParam(&contractAddr, From(NonemptyString(t.Contract))), "contract"),
        errors.Wrap(ResolveParam(&data, From(VarExpr(t.Data, vars), JSONWithVarExprs(t.Data, vars, false), nil)), "data"),
    )
    if err != nil {
        return Result{Error: err}
    }

    bulletprooftxmanager.CreateEthTransaction(t.db, fromAddress, contractAddr, , gasLimit, maxUnconfirmedTransactions uint64) (etx models.EthTx, err error) {
}
