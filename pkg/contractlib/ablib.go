package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/caller-contract/caller"
	"github.com/Xib1uvXi/rarity-box/pkg/contractlib/rarity-contract/rarity"
	"github.com/Xib1uvXi/rarity-box/pkg/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/zap"
	"strings"
)

type AbLib struct {
	*RarityLib
	cAbi abi.ABI
	rAbi abi.ABI
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

	return &AbLib{RarityLib: rarityLib, cAbi: cabi, rAbi: rabi}, nil
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
