/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */
package system

import (
	"encoding/json"
	"math/big"

	"github.com/aergoio/aergo/state"
	"github.com/aergoio/aergo/types"
)

func ExecuteSystemTx(scs *state.ContractState, txBody *types.TxBody,
	sender, receiver *state.V, blockNo types.BlockNo) ([]*types.Event, error) {

	ci, err := ValidateSystemTx(sender.ID(), txBody, sender, scs, blockNo)
	if err != nil {
		return nil, err
	}
	var event *types.Event
	switch ci.Name {
	case types.Stake:
		event, err = staking(txBody, sender, receiver, scs, blockNo)
	case types.VoteBP:
		event, err = voting(txBody, sender, receiver, scs, blockNo, ci)
	case types.Unstake:
		event, err = unstaking(txBody, sender, receiver, scs, blockNo, ci)
	}
	if err != nil {
		return nil, err
	}
	var events []*types.Event
	events = append(events, event)
	return events, nil
}

func ValidateSystemTx(account []byte, txBody *types.TxBody, sender *state.V,
	scs *state.ContractState, blockNo uint64) (*types.CallInfo, error) {
	var ci types.CallInfo
	if err := json.Unmarshal(txBody.Payload, &ci); err != nil {
		return nil, types.ErrTxInvalidPayload
	}
	var err error
	switch ci.Name {
	case types.Stake:
		if sender != nil && sender.Balance().Cmp(txBody.GetAmountBigInt()) < 0 {
			return nil, types.ErrInsufficientBalance
		}
	case types.VoteBP,
		types.VoteFee,
		types.VoteNumBP:
		staked, err := getStaking(scs, account)
		if err != nil {
			return nil, err
		}
		if staked.GetAmountBigInt().Cmp(new(big.Int).SetUint64(0)) == 0 {
			return nil, types.ErrMustStakeBeforeVote
		}
		oldvote, err := GetVote(scs, account)
		if err != nil {
			return nil, err
		}
		if oldvote.Amount != nil && staked.GetWhen()+VotingDelay > blockNo {
			return nil, types.ErrLessTimeHasPassed
		}
	case types.Unstake:
		_, err = validateForUnstaking(account, txBody, scs, blockNo)
	}
	if err != nil {
		return nil, err
	}
	return &ci, nil
}

func validateForUnstaking(account []byte, txBody *types.TxBody, scs *state.ContractState, blockNo uint64) (*types.Staking, error) {
	staked, err := getStaking(scs, account)
	if err != nil {
		return nil, err
	}
	if staked.GetAmountBigInt().Cmp(new(big.Int).SetUint64(0)) == 0 {
		return nil, types.ErrMustStakeBeforeUnstake
	}
	if staked.GetWhen()+StakingDelay > blockNo {
		return nil, types.ErrLessTimeHasPassed
	}
	return staked, nil
}
