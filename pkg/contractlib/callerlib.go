package contractlib

import (
	"context"
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/caller-contract/caller"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"go.uber.org/zap"
	"math/big"
)

type CallerLib struct {
	*RarityLib
	cAbi   abi.ABI
	hasFee bool
	caller *caller.Caller
}

func (l *CallerLib) checkCallerIsOperator(address string) (bool, error) {
	result, err := l.rarity.IsApprovedForAll(nil, common.HexToAddress(address), common.HexToAddress(CallerContractAddress))
	if err != nil {
		log.Logger.Error("req rarity contract is approved for all failed", zap.Error(err))
		return false, err
	}

	return result, nil
}

func (l *CallerLib) callerSetOperator() error {
	txExecutor := func(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
		return l.rarity.SetApprovalForAll(opts, common.HexToAddress(CallerContractAddress), true)
	}

	err := l.txSender.SendTxWaitConfirm(txExecutor)
	if err != nil {
		log.Logger.Error("approve for all send tx failed", zap.Error(err))
		return err
	}

	return nil
}

func (l *CallerLib) Adventure(ids []uint64) error {
	_ids := l.convertUint2BigInt(ids)

	txExecutor := func(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {

		if l.hasFee {
			cost, err := l.estimateGasCost(opts.From, "adventure", _ids...)
			if err != nil {
				return nil, err
			}

			opts.Value = cost
		}

		return l.caller.Adventure(opts, _ids)
	}

	err := l.txSender.SendTxWaitConfirm(txExecutor)
	if err != nil {
		log.Logger.Error("caller adventure send tx failed", zap.Error(err), zap.Uint64s("token id", ids))
		return err
	}

	return nil
}

func (l *CallerLib) Levelup(ids []uint64) error {
	_ids := l.convertUint2BigInt(ids)

	txExecutor := func(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
		if l.hasFee {
			cost, err := l.estimateGasCost(opts.From, "level_up", _ids...)
			if err != nil {
				return nil, err
			}

			opts.Value = cost
		}

		return l.caller.LevelUp(opts, _ids)
	}

	err := l.txSender.SendTxWaitConfirm(txExecutor)
	if err != nil {
		log.Logger.Error("caller levelup send tx failed", zap.Error(err), zap.Uint64s("token id", ids))
		return err
	}

	return nil
}

func (l *CallerLib) ClaimGold(ids []uint64) error {
	_ids := l.convertUint2BigInt(ids)
	needApprove, err := l.needApproved(_ids)

	if err != nil {
		return err
	}

	txExecutor := func(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
		if l.hasFee {
			cost, err := l.estimateGasCost(opts.From, "level_up", _ids...)
			if err != nil {
				return nil, err
			}

			opts.Value = cost
		}

		return l.caller.ClaimGold(opts, _ids, needApprove)
	}

	err = l.txSender.SendTxWaitConfirm(txExecutor)
	if err != nil {
		log.Logger.Error("caller claim gold send tx failed", zap.Error(err), zap.Uint64s("token id", ids))
		return err
	}

	return nil
}

func (l *CallerLib) needApproved(_ids []*big.Int) ([]*big.Int, error) {
	approved, err := l.caller.IsApproved(nil, _ids)

	if err != nil {
		return nil, err
	}

	var needApproved []*big.Int

	for i := range approved {
		if !approved[i] {
			needApproved = append(needApproved, _ids[i])
		}
	}

	return needApproved, nil
}

func (l *CallerLib) convertUint2BigInt(ids []uint64) []*big.Int {
	var _ids []*big.Int

	for i := range ids {
		id := big.NewInt(int64(ids[i]))
		_ids = append(_ids, id)
	}

	return _ids
}

func (l *CallerLib) estimateGasCost(from common.Address, method string, param ...*big.Int) (*big.Int, error) {
	to := common.HexToAddress(CallerContractAddress)
	callData, err := l.cAbi.Pack(method, param)
	if err != nil {
		return nil, err
	}

	gas, err := l.conn.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  from,
		To:    &to,
		Value: new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether)),
		Data:  callData})

	if err != nil {
		return nil, err
	}

	price, err := l.conn.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	gasCost := new(big.Int).Mul(big.NewInt(int64(gas)), price)

	cost := l.unsafeFloatToBigInt(0.25, gasCost)

	return cost, nil
}

func (l *CallerLib) unsafeFloatToBigInt(val float64, gasCost *big.Int) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := new(big.Float)
	coin.SetInt(gasCost)

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
}
