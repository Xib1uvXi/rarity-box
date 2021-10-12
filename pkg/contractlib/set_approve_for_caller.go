package contractlib

import (
	"github.com/Xib1uvXi/rarity-box/pkg/common/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

func (r *RarityLib) checkCallerIsOperator(address string) (bool, error) {
	result, err := r.rarity.IsApprovedForAll(nil, common.HexToAddress(address), common.HexToAddress(CallerContractAddress))
	if err != nil {
		log.Logger.Error("req rarity contract is approved for all failed", zap.Error(err))
		return false, err
	}

	return result, nil
}

func (r *RarityLib) callerSetOperator() error {
	txExecutor := func(opts *bind.TransactOpts) (*ethTypes.Transaction, error) {
		return r.rarity.SetApprovalForAll(opts, common.HexToAddress(CallerContractAddress), true)
	}

	err := r.txSender.SendTxWaitConfirm(txExecutor)
	if err != nil {
		log.Logger.Error("approve for all send tx failed", zap.Error(err))
		return err
	}

	return nil
}
