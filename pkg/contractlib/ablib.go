package contractlib

import (
	"context"
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/caller-contract/caller"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/rarity"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

type AbLib struct {
	*RarityLib
	cAbi   abi.ABI
	rAbi   abi.ABI
	caller *caller.Caller
}

func NewAbLib(rarityLib *RarityLib) (*AbLib, error) {
	cabi, err := abi.JSON(strings.NewReader(caller.CallerABI))
	if err != nil {
		return nil, err
	}

	rabi, err := abi.JSON(strings.NewReader(rarity.RarityABI))
	if err != nil {
		return nil, err
	}

	callerContract, err := caller.NewCaller(common.HexToAddress(CallerContractAddress), rarityLib.conn)
	if err != nil {
		return nil, err
	}

	return &AbLib{RarityLib: rarityLib, cAbi: cabi, rAbi: rabi, caller: callerContract}, nil
}

func (l *AbLib) IsOperator(address string) (bool, error) {
	result, err := l.rarity.IsApprovedForAll(nil, common.HexToAddress(address), common.HexToAddress(CallerContractAddress))
	if err != nil {
		log.Logger.Error("req rarity contract is approved for all failed", zap.Error(err))
		return false, err
	}

	return result, nil
}

func (l *AbLib) SetOperator() (*types.RawTxParam, error) {
	input, err := l.rAbi.Pack("setApprovalForAll", common.HexToAddress(CallerContractAddress), true)
	if err != nil {
		return nil, err
	}

	return &types.RawTxParam{
		To:    RarityContractAddress,
		Value: "0",
		Input: hexutil.Encode(input),
	}, nil
}

func (l *AbLib) Adventure(from string, ids []uint64) (*types.RawTxParam, error) {
	return l.genRarityInput(from, "adventure", ids)
}

func (l *AbLib) Levelup(from string, ids []uint64) (*types.RawTxParam, error) {
	return l.genRarityInput(from, "level_up", ids)
}

func (l *AbLib) ClaimGold(from string, ids []uint64) (*types.RawTxParam, error) {
	return l.genNeedApproveInput(from, "claim_gold", ids)
}

func (l *AbLib) Dungeon(from string, ids []uint64) (*types.RawTxParam, error) {
	return l.genRarityInput(from, "dungeon", ids)
}

func (l *AbLib) genNeedApproveInput(from string, method string, ids []uint64) (*types.RawTxParam, error) {
	_ids := l.convertUint2BigInt(ids)

	needApprove, err := l.needApproved(_ids)

	if err != nil {
		return nil, err
	}

	cost, err := l.estimateGasCost(common.HexToAddress(from), method, _ids, needApprove)
	if err != nil {
		return nil, err
	}

	input, err := l.rAbi.Pack(method, _ids, needApprove)
	if err != nil {
		return nil, err
	}

	return &types.RawTxParam{
		To:    CallerContractAddress,
		Value: cost.String(),
		Input: hexutil.Encode(input),
	}, nil
}

func (l *AbLib) genRarityInput(from string, method string, ids []uint64) (*types.RawTxParam, error) {
	_ids := l.convertUint2BigInt(ids)

	cost, err := l.estimateGasCost(common.HexToAddress(from), method, _ids)
	if err != nil {
		return nil, err
	}

	input, err := l.rAbi.Pack(method, _ids)
	if err != nil {
		return nil, err
	}

	return &types.RawTxParam{
		To:    CallerContractAddress,
		Value: cost.String(),
		Input: hexutil.Encode(input),
	}, nil
}

func (l *AbLib) needApproved(_ids []*big.Int) ([]*big.Int, error) {
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

func (l *AbLib) convertUint2BigInt(ids []uint64) []*big.Int {
	var _ids []*big.Int

	for i := range ids {
		id := big.NewInt(int64(ids[i]))
		_ids = append(_ids, id)
	}

	return _ids
}

func (l *AbLib) estimateGasCost(from common.Address, method string, param ...interface{}) (*big.Int, error) {
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

func (l *AbLib) unsafeFloatToBigInt(val float64, gasCost *big.Int) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := new(big.Float)
	coin.SetInt(gasCost)

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
}
