package operations

//go:generate ffjson $GOFILE

import (
	"github.com/blocktree/whitecoin-adapter/libs/types"
	"github.com/blocktree/whitecoin-adapter/libs/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeTransfer] = func() types.Operation {
		op := &TransferOperation{}
		return op
	}
}

type TransferOperation struct {
	types.OperationFee
	From       types.AccountID   `json:"from"`
	To         types.AccountID   `json:"to"`
	FromAddr   types.Address     `json:"from_addr"`
	ToAddr     types.Address     `json:"to_addr"`
	Amount     types.AssetAmount `json:"amount"`
	Memo       *types.Memo       `json:"memo,omitempty"`
	Extensions types.Extensions  `json:"extensions"`
	GuaranteeId types.ObjectID `json:"guarantee_id"`
}

func (p TransferOperation) Type() types.OperationType {
	return types.OperationTypeTransfer
}

func (p TransferOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	if ppk, ok := params["price_per_kbyte"]; ok {
		if err := enc.Encode(types.UInt32(ppk.(float64))); err != nil {
			return errors.Annotate(err, "encode PricePerKByte")
		}
	}

	return nil
}

func (p TransferOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.GuaranteeId); err != nil {
		return errors.Annotate(err, "encode guarantee_id")
	}

	if err := enc.Encode(p.From); err != nil {
		return errors.Annotate(err, "encode from")
	}

	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode to")
	}

	if err := enc.Encode(p.FromAddr); err != nil {
		return errors.Annotate(err, "encode from addr")
	}

	if err := enc.Encode(p.ToAddr); err != nil {
		return errors.Annotate(err, "encode to addr")
	}

	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}

	if err := enc.Encode(p.Memo != nil); err != nil {
		return errors.Annotate(err, "encode have Memo")
	}

	if err := enc.Encode(p.Memo); err != nil {
		return errors.Annotate(err, "encode memo")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
