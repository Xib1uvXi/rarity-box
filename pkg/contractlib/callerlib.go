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
	"strings"
)

const CallerContractAddress = "0x3B8d1f9569ec2B8C3dD76847B944bB1b0c998A38"

type CallerLib struct {
	*RarityLib
	cAbi   abi.ABI
	hasFee bool
	caller *caller.Caller
}

func NewCallerLib(rarityLib *RarityLib, hasFee bool) (*CallerLib, error) {
	clib := &CallerLib{RarityLib: rarityLib, hasFee: hasFee}

	callerContract, err := caller.NewCaller(common.HexToAddress(CallerContractAddress), clib.conn)
	if err != nil {
		return nil, err
	}

	clib.caller = callerContract

	cabi, err := abi.JSON(strings.NewReader(caller.CallerABI))
	if err != nil {
		return nil, err
	}

	clib.cAbi = cabi

	callerIsOpertor, err := clib.checkCallerIsOperator(clib.TxSender.Address())
	if err != nil {
		return nil, err
	}

	if !callerIsOpertor {
		if err := clib.callerSetOperator(); err != nil {
			return nil, err
		}
	}

	return clib, nil
}

func (l *CallerLib) checkCallerIsOperator(address common.Address) (bool, error) {
	result, err := l.rarity.IsApprovedForAll(nil, address, common.HexToAddress(CallerContractAddress))
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

	err := l.TxSender.SendTxWaitConfirm(txExecutor)
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
			cost, err := l.estimateGasCost(opts.From, "adventure", _ids)
			if err != nil {
				return nil, err
			}

			opts.Value = cost
		}

		return l.caller.Adventure(opts, _ids)
	}

	err := l.TxSender.SendTxWaitConfirm(txExecutor)
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
			cost, err := l.estimateGasCost(opts.From, "level_up", _ids)
			if err != nil {
				return nil, err
			}

			opts.Value = cost
		}

		return l.caller.LevelUp(opts, _ids)
	}

	err := l.TxSender.SendTxWaitConfirm(txExecutor)
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
			cost, err := l.estimateGasCost(opts.From, "claim_gold", _ids, needApprove)
			if err != nil {
				return nil, err
			}

			opts.Value = cost
		}

		return l.caller.ClaimGold(opts, _ids, needApprove)
	}

	err = l.TxSender.SendTxWaitConfirm(txExecutor)
	if err != nil {
		log.Logger.Error("caller claim gold send tx failed", zap.Error(err), zap.Uint64s("token id", ids))
		return err
	}

	return nil
}

func (l *CallerLib) Dungeon(ids []uint64) error {
	_ids := l.convertUint2BigInt(ids)
	needApprove, err := l.needApproved(_ids)

	if err != nil {
		return err
	}

	txExecutor := func(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
		if l.hasFee {
			cost, err := l.estimateGasCost(opts.From, "dungeon", _ids, needApprove)
			if err != nil {
				return nil, err
			}

			opts.Value = cost
		}

		return l.caller.Dungeon(opts, _ids, needApprove)
	}

	err = l.TxSender.SendTxWaitConfirm(txExecutor)
	if err != nil {
		log.Logger.Error("caller dungeon send tx failed", zap.Error(err), zap.Uint64s("token id", ids))
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

func (l *CallerLib) estimateGasCost(from common.Address, method string, param ...interface{}) (*big.Int, error) {
	to := common.HexToAddress(CallerContractAddress)
	callData, err := l.cAbi.Pack(method, param...)
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
