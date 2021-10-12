package tx

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

var (
	ExecFailedErr = errors.New("execution failed")
)

func NewFTMTXSession(pkStr string, client *ethclient.Client) (*Session, error) {
	account, err := NewAccount(pkStr)
	if err != nil {
		return nil, err
	}

	return NewSession(big.NewInt(250), account, client)
}

func NewSession(chainID *big.Int, account *account, client *ethclient.Client) (*Session, error) {
	session := &Session{chainID: chainID, account: account, client: client}

	// Get the current nonce of the address during initialization
	nonce, err := client.PendingNonceAt(context.Background(), account.address)
	if err != nil {
		return nil, err
	}

	session.nonce = nonce

	// At the same time initialize the TransactOpts of the current session
	opt, err := bind.NewKeyedTransactorWithChainID(session.PrivateKey(), session.chainID)
	if err != nil {
		return nil, err
	}

	opt.Nonce = big.NewInt(int64(nonce))

	session.sessionTxOpt = opt

	return session, nil
}

type Session struct {
	chainID *big.Int
	*account
	nonce        uint64
	client       *ethclient.Client
	sessionTxOpt *bind.TransactOpts
}

func (s *Session) incNonce() {
	s.sessionTxOpt.Nonce.Add(s.sessionTxOpt.Nonce, big.NewInt(1))
}

func (s *Session) SendTx(txExecutor func(opts *bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {
	tx, err := txExecutor(s.sessionTxOpt)
	if err != nil {
		return nil, err
	}

	s.incNonce()
	return tx, nil
}

func (s *Session) SendTxWaitConfirm(txExecutor func(opts *bind.TransactOpts) (*types.Transaction, error)) error {
	tx, err := s.SendTx(txExecutor)
	if err != nil {
		return err
	}

	return s.txConfirm(tx.Hash())
}

func (s *Session) txConfirm(txHash common.Hash) error {
	receipt, err := s.client.TransactionReceipt(context.Background(), txHash)

	switch err {
	case ethereum.NotFound:
		return s.txConfirmRetry(txHash)

	case nil:
		return s.txReceiptStatus(receipt.Status)

	default:
		return err
	}
}

func (s *Session) txReceiptStatus(status uint64) error {
	if status != types.ReceiptStatusSuccessful {
		return ExecFailedErr
	}

	return nil
}

func (s *Session) txConfirmRetry(txHash common.Hash) error {
	time.Sleep(5 * time.Second)
	for i := 0; i < 6; i++ {
		receipt, err := s.client.TransactionReceipt(context.Background(), txHash)
		switch err {
		case ethereum.NotFound:
			time.Sleep(15 * time.Second)
			continue
		case nil:
			return s.txReceiptStatus(receipt.Status)
		default:
			return err
		}
	}

	return fmt.Errorf("tx confirm overtime, tx_hash: %s", txHash.String())
}
