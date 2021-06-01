package pipeline_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/services/pipeline"
)

func TestETHABIEncodeTask(t *testing.T) {
	tests := []struct {
		name                  string
		abi                   string
		data                  string
		vars                  pipeline.Vars
		inputs                []pipeline.Result
		expected              string
		expectedErrorCause    error
		expectedErrorContains string
	}{
		{
			"",
			"uint256 u256, bool b, int256 i256, string s",
			`{ "u256": $(foo), "b": $(bar), "i256": $(baz), "s": $(quux) }`,
			pipeline.NewVarsFrom(map[string]interface{}{"foo": big.NewInt(123), "bar": true, "baz": big.NewInt(-321), "quux": "foo bar baz"}),
			nil,
			"0x000000000000000000000000000000000000000000000000000000000000007b0000000000000000000000000000000000000000000000000000000000000001fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffebf0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000b666f6f206261722062617a000000000000000000000000000000000000000000",
			nil,
			"",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			task := pipeline.ETHABIEncodeTask{
				BaseTask: pipeline.NewBaseTask("encode", nil, 0, 0),
				ABI:      test.abi,
				Data:     test.data,
			}

			result := task.Run(context.Background(), test.vars, pipeline.JSONSerializable{}, test.inputs)
			require.NoError(t, result.Error)

			expected := hexutil.MustDecode(test.expected)

			require.Equal(t, expected, result.Value)
		})
	}
}
