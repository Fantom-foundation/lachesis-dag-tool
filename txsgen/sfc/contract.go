// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sfc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ChangedValidatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"}],\"name\":\"CreatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"DeactivatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedUpStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"RestakedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"UnlockedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UpdatedBaseRewardPerSec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"UpdatedOfflinePenaltyThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"UpdatedSlashingRefundRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"syncPubkey\",\"type\":\"bool\"}],\"name\":\"_syncValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRewardPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"createValidator\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedOriginatedTxsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedUptime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochReceivedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getEpochSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochValidatorIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getLockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getLockupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStashedLockupRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getUnlockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidatorPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeDriver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"isLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxDelegatedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLockupDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minLockupDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"offlinePenaltyThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"restakeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTime\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"stashRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stashedRewardsUntilEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSlashedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockedRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateBaseRewardPerSecond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"updateOfflinePenaltyThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"updateSlashingRefundRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes3\",\"name\":\"\",\"type\":\"bytes3\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawalPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawalPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_toStakerID\",\"type\":\"uint256\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakerID\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dagAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"isDelegationLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"isStakeLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersLastID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"lockedDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"lockedStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"createDelegation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"calcDelegationRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"calcValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"claimDelegationRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"claimDelegationCompoundRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimValidatorRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimValidatorCompoundRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawStakePartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawDelegationPartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"partialWithdrawByRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"}],\"name\":\"lockUpStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"lockUpDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_Contract *ContractCaller) BaseRewardPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "baseRewardPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_Contract *ContractSession) BaseRewardPerSecond() (*big.Int, error) {
	return _Contract.Contract.BaseRewardPerSecond(&_Contract.CallOpts)
}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_Contract *ContractCallerSession) BaseRewardPerSecond() (*big.Int, error) {
	return _Contract.Contract.BaseRewardPerSecond(&_Contract.CallOpts)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) CalcDelegationRewards(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int, arg2 *big.Int, arg3 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calcDelegationRewards", delegator, toStakerID, arg2, arg3)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) CalcDelegationRewards(delegator common.Address, toStakerID *big.Int, arg2 *big.Int, arg3 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcDelegationRewards(&_Contract.CallOpts, delegator, toStakerID, arg2, arg3)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) CalcDelegationRewards(delegator common.Address, toStakerID *big.Int, arg2 *big.Int, arg3 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcDelegationRewards(&_Contract.CallOpts, delegator, toStakerID, arg2, arg3)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) CalcValidatorRewards(opts *bind.CallOpts, stakerID *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calcValidatorRewards", stakerID, arg1, arg2)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) CalcValidatorRewards(stakerID *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcValidatorRewards(&_Contract.CallOpts, stakerID, arg1, arg2)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) CalcValidatorRewards(stakerID *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcValidatorRewards(&_Contract.CallOpts, stakerID, arg1, arg2)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractCaller) ContractCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "contractCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractSession) ContractCommission() (*big.Int, error) {
	return _Contract.Contract.ContractCommission(&_Contract.CallOpts)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractCallerSession) ContractCommission() (*big.Int, error) {
	return _Contract.Contract.ContractCommission(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentSealedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCaller) DelegationLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationLockPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodEpochs(&_Contract.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCallerSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodEpochs(&_Contract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCaller) DelegationLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationLockPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodTime(&_Contract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCallerSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodTime(&_Contract.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address _from, uint256 _toStakerID) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractCaller) Delegations(opts *bind.CallOpts, _from common.Address, _toStakerID *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegations", _from, _toStakerID)

	outstruct := new(struct {
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		Amount           *big.Int
		PaidUntilEpoch   *big.Int
		ToStakerID       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreatedEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PaidUntilEpoch = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ToStakerID = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address _from, uint256 _toStakerID) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractSession) Delegations(_from common.Address, _toStakerID *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Contract.Contract.Delegations(&_Contract.CallOpts, _from, _toStakerID)
}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address _from, uint256 _toStakerID) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractCallerSession) Delegations(_from common.Address, _toStakerID *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Contract.Contract.Delegations(&_Contract.CallOpts, _from, _toStakerID)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() pure returns(uint256)
func (_Contract *ContractCaller) DelegationsNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationsNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() pure returns(uint256)
func (_Contract *ContractSession) DelegationsNum() (*big.Int, error) {
	return _Contract.Contract.DelegationsNum(&_Contract.CallOpts)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() pure returns(uint256)
func (_Contract *ContractCallerSession) DelegationsNum() (*big.Int, error) {
	return _Contract.Contract.DelegationsNum(&_Contract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) DelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationsTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractSession) DelegationsTotalAmount() (*big.Int, error) {
	return _Contract.Contract.DelegationsTotalAmount(&_Contract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) DelegationsTotalAmount() (*big.Int, error) {
	return _Contract.Contract.DelegationsTotalAmount(&_Contract.CallOpts)
}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedOriginatedTxsFee(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedOriginatedTxsFee", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedRewardPerToken(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedRewardPerToken", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedRewardPerToken(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedRewardPerToken(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedUptime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedUptime", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedUptime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedUptime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochOfflineBlocks(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochOfflineBlocks", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineBlocks(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineBlocks(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochOfflineTime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochOfflineTime", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineTime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineTime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochReceivedStake(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochReceivedStake", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochReceivedStake(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochReceivedStake(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractCaller) GetEpochSnapshot(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochSnapshot", arg0)

	outstruct := new(struct {
		EndTime               *big.Int
		EpochFee              *big.Int
		TotalBaseRewardWeight *big.Int
		TotalTxRewardWeight   *big.Int
		BaseRewardPerSecond   *big.Int
		TotalStake            *big.Int
		TotalSupply           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBaseRewardWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalTxRewardWeight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BaseRewardPerSecond = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalStake = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractSession) GetEpochSnapshot(arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	return _Contract.Contract.GetEpochSnapshot(&_Contract.CallOpts, arg0)
}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractCallerSession) GetEpochSnapshot(arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	return _Contract.Contract.GetEpochSnapshot(&_Contract.CallOpts, arg0)
}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractCaller) GetEpochValidatorIDs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochValidatorIDs", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetEpochValidatorIDs(&_Contract.CallOpts, epoch)
}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractCallerSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetEpochValidatorIDs(&_Contract.CallOpts, epoch)
}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) GetLockedStake(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLockedStake", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) GetLockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetLockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetLockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetLockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) GetLockupInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLockupInfo", arg0, arg1)

	outstruct := new(struct {
		LockedStake *big.Int
		FromEpoch   *big.Int
		EndTime     *big.Int
		Duration    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FromEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) GetLockupInfo(arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	return _Contract.Contract.GetLockupInfo(&_Contract.CallOpts, arg0, arg1)
}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) GetLockupInfo(arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	return _Contract.Contract.GetLockupInfo(&_Contract.CallOpts, arg0, arg1)
}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetSelfStake(opts *bind.CallOpts, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSelfStake", validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSelfStake(&_Contract.CallOpts, validatorID)
}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSelfStake(&_Contract.CallOpts, validatorID)
}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) GetStake(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStake", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) GetStake(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetStake(&_Contract.CallOpts, arg0, arg1)
}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) GetStake(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetStake(&_Contract.CallOpts, arg0, arg1)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address _addr) view returns(uint256)
func (_Contract *ContractCaller) GetStakerID(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStakerID", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address _addr) view returns(uint256)
func (_Contract *ContractSession) GetStakerID(_addr common.Address) (*big.Int, error) {
	return _Contract.Contract.GetStakerID(&_Contract.CallOpts, _addr)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address _addr) view returns(uint256)
func (_Contract *ContractCallerSession) GetStakerID(_addr common.Address) (*big.Int, error) {
	return _Contract.Contract.GetStakerID(&_Contract.CallOpts, _addr)
}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractCaller) GetStashedLockupRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStashedLockupRewards", arg0, arg1)

	outstruct := new(struct {
		LockupExtraReward *big.Int
		LockupBaseReward  *big.Int
		UnlockedReward    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockupExtraReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LockupBaseReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnlockedReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractSession) GetStashedLockupRewards(arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	return _Contract.Contract.GetStashedLockupRewards(&_Contract.CallOpts, arg0, arg1)
}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractCallerSession) GetStashedLockupRewards(arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	return _Contract.Contract.GetStashedLockupRewards(&_Contract.CallOpts, arg0, arg1)
}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) GetUnlockedStake(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUnlockedStake", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) GetUnlockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetUnlockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetUnlockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetUnlockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractCaller) GetValidator(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidator", arg0)

	outstruct := new(struct {
		Status           *big.Int
		DeactivatedTime  *big.Int
		DeactivatedEpoch *big.Int
		ReceivedStake    *big.Int
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		Auth             common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReceivedStake = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CreatedEpoch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Auth = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractSession) GetValidator(arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	return _Contract.Contract.GetValidator(&_Contract.CallOpts, arg0)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractCallerSession) GetValidator(arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	return _Contract.Contract.GetValidator(&_Contract.CallOpts, arg0)
}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractCaller) GetValidatorID(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorID", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractSession) GetValidatorID(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.GetValidatorID(&_Contract.CallOpts, arg0)
}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractCallerSession) GetValidatorID(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.GetValidatorID(&_Contract.CallOpts, arg0)
}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractCaller) GetValidatorPubkey(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorPubkey", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractSession) GetValidatorPubkey(arg0 *big.Int) ([]byte, error) {
	return _Contract.Contract.GetValidatorPubkey(&_Contract.CallOpts, arg0)
}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractCallerSession) GetValidatorPubkey(arg0 *big.Int) ([]byte, error) {
	return _Contract.Contract.GetValidatorPubkey(&_Contract.CallOpts, arg0)
}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractCaller) GetWithdrawalRequest(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWithdrawalRequest", arg0, arg1, arg2)

	outstruct := new(struct {
		Epoch  *big.Int
		Time   *big.Int
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractSession) GetWithdrawalRequest(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	return _Contract.Contract.GetWithdrawalRequest(&_Contract.CallOpts, arg0, arg1, arg2)
}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractCallerSession) GetWithdrawalRequest(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	return _Contract.Contract.GetWithdrawalRequest(&_Contract.CallOpts, arg0, arg1, arg2)
}

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractCaller) IsDelegationLockedUp(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isDelegationLockedUp", delegator, toStakerID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractSession) IsDelegationLockedUp(delegator common.Address, toStakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsDelegationLockedUp(&_Contract.CallOpts, delegator, toStakerID)
}

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractCallerSession) IsDelegationLockedUp(delegator common.Address, toStakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsDelegationLockedUp(&_Contract.CallOpts, delegator, toStakerID)
}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractCaller) IsLockedUp(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isLockedUp", delegator, toValidatorID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractSession) IsLockedUp(delegator common.Address, toValidatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsLockedUp(&_Contract.CallOpts, delegator, toValidatorID)
}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractCallerSession) IsLockedUp(delegator common.Address, toValidatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsLockedUp(&_Contract.CallOpts, delegator, toValidatorID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCallerSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractCaller) IsSlashed(opts *bind.CallOpts, validatorID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isSlashed", validatorID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractSession) IsSlashed(validatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsSlashed(&_Contract.CallOpts, validatorID)
}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractCallerSession) IsSlashed(validatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsSlashed(&_Contract.CallOpts, validatorID)
}

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 stakerID) view returns(bool)
func (_Contract *ContractCaller) IsStakeLockedUp(opts *bind.CallOpts, stakerID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isStakeLockedUp", stakerID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 stakerID) view returns(bool)
func (_Contract *ContractSession) IsStakeLockedUp(stakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsStakeLockedUp(&_Contract.CallOpts, stakerID)
}

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 stakerID) view returns(bool)
func (_Contract *ContractCallerSession) IsStakeLockedUp(stakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsStakeLockedUp(&_Contract.CallOpts, stakerID)
}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractCaller) LastValidatorID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastValidatorID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractSession) LastValidatorID() (*big.Int, error) {
	return _Contract.Contract.LastValidatorID(&_Contract.CallOpts)
}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractCallerSession) LastValidatorID() (*big.Int, error) {
	return _Contract.Contract.LastValidatorID(&_Contract.CallOpts)
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address delegator, uint256 toStakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) LockedDelegations(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lockedDelegations", delegator, toStakerID)

	outstruct := new(struct {
		FromEpoch *big.Int
		EndTime   *big.Int
		Duration  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address delegator, uint256 toStakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) LockedDelegations(delegator common.Address, toStakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedDelegations(&_Contract.CallOpts, delegator, toStakerID)
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address delegator, uint256 toStakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) LockedDelegations(delegator common.Address, toStakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedDelegations(&_Contract.CallOpts, delegator, toStakerID)
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 stakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) LockedStakes(opts *bind.CallOpts, stakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lockedStakes", stakerID)

	outstruct := new(struct {
		FromEpoch *big.Int
		EndTime   *big.Int
		Duration  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 stakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) LockedStakes(stakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedStakes(&_Contract.CallOpts, stakerID)
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 stakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) LockedStakes(stakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedStakes(&_Contract.CallOpts, stakerID)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractCaller) MaxDelegatedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxDelegatedRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractSession) MaxDelegatedRatio() (*big.Int, error) {
	return _Contract.Contract.MaxDelegatedRatio(&_Contract.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxDelegatedRatio() (*big.Int, error) {
	return _Contract.Contract.MaxDelegatedRatio(&_Contract.CallOpts)
}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractCaller) MaxLockupDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxLockupDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractSession) MaxLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MaxLockupDuration(&_Contract.CallOpts)
}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MaxLockupDuration(&_Contract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractSession) MinDelegation() (*big.Int, error) {
	return _Contract.Contract.MinDelegation(&_Contract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegation() (*big.Int, error) {
	return _Contract.Contract.MinDelegation(&_Contract.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegationDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDelegationDecrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractSession) MinDelegationDecrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationDecrease(&_Contract.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegationDecrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationDecrease(&_Contract.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegationIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDelegationIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractSession) MinDelegationIncrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationIncrease(&_Contract.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegationIncrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationIncrease(&_Contract.CallOpts)
}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractCaller) MinLockupDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minLockupDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractSession) MinLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MinLockupDuration(&_Contract.CallOpts)
}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractCallerSession) MinLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MinLockupDuration(&_Contract.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() pure returns(uint256)
func (_Contract *ContractCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minSelfStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() pure returns(uint256)
func (_Contract *ContractSession) MinSelfStake() (*big.Int, error) {
	return _Contract.Contract.MinSelfStake(&_Contract.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() pure returns(uint256)
func (_Contract *ContractCallerSession) MinSelfStake() (*big.Int, error) {
	return _Contract.Contract.MinSelfStake(&_Contract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractSession) MinStake() (*big.Int, error) {
	return _Contract.Contract.MinStake(&_Contract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStake() (*big.Int, error) {
	return _Contract.Contract.MinStake(&_Contract.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractCaller) MinStakeDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minStakeDecrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractSession) MinStakeDecrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeDecrease(&_Contract.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStakeDecrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeDecrease(&_Contract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractCaller) MinStakeIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minStakeIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractSession) MinStakeIncrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeIncrease(&_Contract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStakeIncrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeIncrease(&_Contract.CallOpts)
}

// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
//
// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
func (_Contract *ContractCaller) OfflinePenaltyThreshold(opts *bind.CallOpts) (struct {
	BlocksNum *big.Int
	Time      *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "offlinePenaltyThreshold")

	outstruct := new(struct {
		BlocksNum *big.Int
		Time      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlocksNum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
//
// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
func (_Contract *ContractSession) OfflinePenaltyThreshold() (struct {
	BlocksNum *big.Int
	Time      *big.Int
}, error) {
	return _Contract.Contract.OfflinePenaltyThreshold(&_Contract.CallOpts)
}

// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
//
// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
func (_Contract *ContractCallerSession) OfflinePenaltyThreshold() (struct {
	BlocksNum *big.Int
	Time      *big.Int
}, error) {
	return _Contract.Contract.OfflinePenaltyThreshold(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) PendingRewards(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pendingRewards", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.PendingRewards(&_Contract.CallOpts, delegator, toValidatorID)
}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.PendingRewards(&_Contract.CallOpts, delegator, toValidatorID)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) RewardsStash(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardsStash", delegator, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, delegator, validatorID)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, delegator, validatorID)
}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) SlashingRefundRatio(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slashingRefundRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractSession) SlashingRefundRatio(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlashingRefundRatio(&_Contract.CallOpts, arg0)
}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) SlashingRefundRatio(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlashingRefundRatio(&_Contract.CallOpts, arg0)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCaller) StakeLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakeLockPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodEpochs(&_Contract.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCallerSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodEpochs(&_Contract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCaller) StakeLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakeLockPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractSession) StakeLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodTime(&_Contract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCallerSession) StakeLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodTime(&_Contract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) StakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakeTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractSession) StakeTotalAmount() (*big.Int, error) {
	return _Contract.Contract.StakeTotalAmount(&_Contract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) StakeTotalAmount() (*big.Int, error) {
	return _Contract.Contract.StakeTotalAmount(&_Contract.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 _stakerID) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractCaller) Stakers(opts *bind.CallOpts, _stakerID *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakers", _stakerID)

	outstruct := new(struct {
		Status           *big.Int
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		StakeAmount      *big.Int
		PaidUntilEpoch   *big.Int
		DelegatedMe      *big.Int
		DagAddress       common.Address
		SfcAddress       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreatedEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StakeAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PaidUntilEpoch = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DelegatedMe = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.DagAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.SfcAddress = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 _stakerID) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractSession) Stakers(_stakerID *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	return _Contract.Contract.Stakers(&_Contract.CallOpts, _stakerID)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 _stakerID) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractCallerSession) Stakers(_stakerID *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	return _Contract.Contract.Stakers(&_Contract.CallOpts, _stakerID)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractCaller) StakersLastID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakersLastID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractSession) StakersLastID() (*big.Int, error) {
	return _Contract.Contract.StakersLastID(&_Contract.CallOpts)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractCallerSession) StakersLastID() (*big.Int, error) {
	return _Contract.Contract.StakersLastID(&_Contract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractCaller) StakersNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakersNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractSession) StakersNum() (*big.Int, error) {
	return _Contract.Contract.StakersNum(&_Contract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractCallerSession) StakersNum() (*big.Int, error) {
	return _Contract.Contract.StakersNum(&_Contract.CallOpts)
}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) StashedRewardsUntilEpoch(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stashedRewardsUntilEpoch", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) StashedRewardsUntilEpoch(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.StashedRewardsUntilEpoch(&_Contract.CallOpts, arg0, arg1)
}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) StashedRewardsUntilEpoch(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.StashedRewardsUntilEpoch(&_Contract.CallOpts, arg0, arg1)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractCaller) TotalActiveStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalActiveStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractSession) TotalActiveStake() (*big.Int, error) {
	return _Contract.Contract.TotalActiveStake(&_Contract.CallOpts)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalActiveStake() (*big.Int, error) {
	return _Contract.Contract.TotalActiveStake(&_Contract.CallOpts)
}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractCaller) TotalSlashedStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSlashedStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractSession) TotalSlashedStake() (*big.Int, error) {
	return _Contract.Contract.TotalSlashedStake(&_Contract.CallOpts)
}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSlashedStake() (*big.Int, error) {
	return _Contract.Contract.TotalSlashedStake(&_Contract.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractSession) TotalStake() (*big.Int, error) {
	return _Contract.Contract.TotalStake(&_Contract.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalStake() (*big.Int, error) {
	return _Contract.Contract.TotalStake(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractCaller) UnlockedRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "unlockedRewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractSession) UnlockedRewardRatio() (*big.Int, error) {
	return _Contract.Contract.UnlockedRewardRatio(&_Contract.CallOpts)
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractCallerSession) UnlockedRewardRatio() (*big.Int, error) {
	return _Contract.Contract.UnlockedRewardRatio(&_Contract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractCaller) ValidatorCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validatorCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractSession) ValidatorCommission() (*big.Int, error) {
	return _Contract.Contract.ValidatorCommission(&_Contract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractCallerSession) ValidatorCommission() (*big.Int, error) {
	return _Contract.Contract.ValidatorCommission(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) ([3]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new([3]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([3]byte)).(*[3]byte)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractSession) Version() ([3]byte, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractCallerSession) Version() ([3]byte, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCaller) WithdrawalPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "withdrawalPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
func (_Contract *ContractSession) WithdrawalPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.WithdrawalPeriodEpochs(&_Contract.CallOpts)
}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCallerSession) WithdrawalPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.WithdrawalPeriodEpochs(&_Contract.CallOpts)
}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() pure returns(uint256)
func (_Contract *ContractCaller) WithdrawalPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "withdrawalPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() pure returns(uint256)
func (_Contract *ContractSession) WithdrawalPeriodTime() (*big.Int, error) {
	return _Contract.Contract.WithdrawalPeriodTime(&_Contract.CallOpts)
}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() pure returns(uint256)
func (_Contract *ContractCallerSession) WithdrawalPeriodTime() (*big.Int, error) {
	return _Contract.Contract.WithdrawalPeriodTime(&_Contract.CallOpts)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractTransactor) SyncValidator(opts *bind.TransactOpts, validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_syncValidator", validatorID, syncPubkey)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.Contract.SyncValidator(&_Contract.TransactOpts, validatorID, syncPubkey)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractTransactorSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.Contract.SyncValidator(&_Contract.TransactOpts, validatorID, syncPubkey)
}

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractTransactor) ClaimDelegationCompoundRewards(opts *bind.TransactOpts, arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimDelegationCompoundRewards", arg0, toStakerID)
}

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractSession) ClaimDelegationCompoundRewards(arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationCompoundRewards(&_Contract.TransactOpts, arg0, toStakerID)
}

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) ClaimDelegationCompoundRewards(arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationCompoundRewards(&_Contract.TransactOpts, arg0, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractTransactor) ClaimDelegationRewards(opts *bind.TransactOpts, arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimDelegationRewards", arg0, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractSession) ClaimDelegationRewards(arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationRewards(&_Contract.TransactOpts, arg0, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) ClaimDelegationRewards(arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationRewards(&_Contract.TransactOpts, arg0, toStakerID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) ClaimRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimRewards", toValidatorID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, toValidatorID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, toValidatorID)
}

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 ) returns()
func (_Contract *ContractTransactor) ClaimValidatorCompoundRewards(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimValidatorCompoundRewards", arg0)
}

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 ) returns()
func (_Contract *ContractSession) ClaimValidatorCompoundRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorCompoundRewards(&_Contract.TransactOpts, arg0)
}

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 ) returns()
func (_Contract *ContractTransactorSession) ClaimValidatorCompoundRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorCompoundRewards(&_Contract.TransactOpts, arg0)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 ) returns()
func (_Contract *ContractTransactor) ClaimValidatorRewards(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimValidatorRewards", arg0)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 ) returns()
func (_Contract *ContractSession) ClaimValidatorRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorRewards(&_Contract.TransactOpts, arg0)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 ) returns()
func (_Contract *ContractTransactorSession) ClaimValidatorRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorRewards(&_Contract.TransactOpts, arg0)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactor) CreateDelegation(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createDelegation", toValidatorID)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 toValidatorID) payable returns()
func (_Contract *ContractSession) CreateDelegation(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateDelegation(&_Contract.TransactOpts, toValidatorID)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactorSession) CreateDelegation(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateDelegation(&_Contract.TransactOpts, toValidatorID)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractTransactor) CreateValidator(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createValidator", pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractTransactorSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, pubkey)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactor) DeactivateValidator(opts *bind.TransactOpts, validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deactivateValidator", validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactorSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactor) Delegate(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "delegate", toValidatorID)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, toValidatorID)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactorSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, toValidatorID)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", sealedEpoch, _totalSupply, nodeDriver, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
func (_Contract *ContractSession) Initialize(sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, sealedEpoch, _totalSupply, nodeDriver, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
func (_Contract *ContractTransactorSession) Initialize(sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, sealedEpoch, _totalSupply, nodeDriver, owner)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactor) LockStake(opts *bind.TransactOpts, toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockStake", toValidatorID, lockupDuration, amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractSession) LockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactorSession) LockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractTransactor) LockUpDelegation(opts *bind.TransactOpts, lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockUpDelegation", lockDuration, toStakerID)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractSession) LockUpDelegation(lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpDelegation(&_Contract.TransactOpts, lockDuration, toStakerID)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) LockUpDelegation(lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpDelegation(&_Contract.TransactOpts, lockDuration, toStakerID)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractTransactor) LockUpStake(opts *bind.TransactOpts, lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockUpStake", lockDuration)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractSession) LockUpStake(lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpStake(&_Contract.TransactOpts, lockDuration)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractTransactorSession) LockUpStake(lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpStake(&_Contract.TransactOpts, lockDuration)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 ) returns()
func (_Contract *ContractTransactor) PartialWithdrawByRequest(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "partialWithdrawByRequest", arg0)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 ) returns()
func (_Contract *ContractSession) PartialWithdrawByRequest(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PartialWithdrawByRequest(&_Contract.TransactOpts, arg0)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 ) returns()
func (_Contract *ContractTransactorSession) PartialWithdrawByRequest(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PartialWithdrawByRequest(&_Contract.TransactOpts, arg0)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 ) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawDelegation(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawDelegation", arg0)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 ) returns()
func (_Contract *ContractSession) PrepareToWithdrawDelegation(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegation(&_Contract.TransactOpts, arg0)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 ) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawDelegation(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegation(&_Contract.TransactOpts, arg0)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawDelegationPartial(opts *bind.TransactOpts, wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawDelegationPartial", wrID, toStakerID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegationPartial(&_Contract.TransactOpts, wrID, toStakerID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegationPartial(&_Contract.TransactOpts, wrID, toStakerID, amount)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractTransactor) PrepareToWithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawStake")
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStake(&_Contract.TransactOpts)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStake(&_Contract.TransactOpts)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawStakePartial(opts *bind.TransactOpts, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawStakePartial", wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStakePartial(&_Contract.TransactOpts, wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStakePartial(&_Contract.TransactOpts, wrID, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) RestakeRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "restakeRewards", toValidatorID)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RestakeRewards(&_Contract.TransactOpts, toValidatorID)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RestakeRewards(&_Contract.TransactOpts, toValidatorID)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpoch", offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractSession) SealEpoch(offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactorSession) SealEpoch(offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractTransactor) SealEpochValidators(opts *bind.TransactOpts, nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpochValidators", nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractTransactorSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactor) SetGenesisDelegation(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisDelegation", delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactorSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisValidator", auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactorSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) StashRewards(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stashRewards", delegator, toValidatorID)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.StashRewards(&_Contract.TransactOpts, delegator, toValidatorID)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.StashRewards(&_Contract.TransactOpts, delegator, toValidatorID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactor) Undelegate(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "undelegate", toValidatorID, wrID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractSession) Undelegate(toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, wrID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Undelegate(toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, wrID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) UnlockStake(opts *bind.TransactOpts, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlockStake", toValidatorID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractSession) UnlockStake(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnlockStake(&_Contract.TransactOpts, toValidatorID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) UnlockStake(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnlockStake(&_Contract.TransactOpts, toValidatorID, amount)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
func (_Contract *ContractTransactor) UpdateBaseRewardPerSecond(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateBaseRewardPerSecond", value)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
func (_Contract *ContractSession) UpdateBaseRewardPerSecond(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBaseRewardPerSecond(&_Contract.TransactOpts, value)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
func (_Contract *ContractTransactorSession) UpdateBaseRewardPerSecond(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBaseRewardPerSecond(&_Contract.TransactOpts, value)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
//
// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
func (_Contract *ContractTransactor) UpdateOfflinePenaltyThreshold(opts *bind.TransactOpts, blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateOfflinePenaltyThreshold", blocksNum, time)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
//
// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
func (_Contract *ContractSession) UpdateOfflinePenaltyThreshold(blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOfflinePenaltyThreshold(&_Contract.TransactOpts, blocksNum, time)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
//
// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
func (_Contract *ContractTransactorSession) UpdateOfflinePenaltyThreshold(blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOfflinePenaltyThreshold(&_Contract.TransactOpts, blocksNum, time)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractTransactor) UpdateSlashingRefundRatio(opts *bind.TransactOpts, validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSlashingRefundRatio", validatorID, refundRatio)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSlashingRefundRatio(&_Contract.TransactOpts, validatorID, refundRatio)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractTransactorSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSlashingRefundRatio(&_Contract.TransactOpts, validatorID, refundRatio)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", toValidatorID, wrID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, toValidatorID, wrID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractTransactorSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, toValidatorID, wrID)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 ) returns()
func (_Contract *ContractTransactor) WithdrawDelegation(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawDelegation", arg0)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 ) returns()
func (_Contract *ContractSession) WithdrawDelegation(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawDelegation(&_Contract.TransactOpts, arg0)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 ) returns()
func (_Contract *ContractTransactorSession) WithdrawDelegation(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawDelegation(&_Contract.TransactOpts, arg0)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// ContractChangedValidatorStatusIterator is returned from FilterChangedValidatorStatus and is used to iterate over the raw logs and unpacked data for ChangedValidatorStatus events raised by the Contract contract.
type ContractChangedValidatorStatusIterator struct {
	Event *ContractChangedValidatorStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractChangedValidatorStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractChangedValidatorStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractChangedValidatorStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractChangedValidatorStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractChangedValidatorStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractChangedValidatorStatus represents a ChangedValidatorStatus event raised by the Contract contract.
type ContractChangedValidatorStatus struct {
	ValidatorID *big.Int
	Status      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChangedValidatorStatus is a free log retrieval operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) FilterChangedValidatorStatus(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractChangedValidatorStatusIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ChangedValidatorStatus", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractChangedValidatorStatusIterator{contract: _Contract.contract, event: "ChangedValidatorStatus", logs: logs, sub: sub}, nil
}

// WatchChangedValidatorStatus is a free log subscription operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) WatchChangedValidatorStatus(opts *bind.WatchOpts, sink chan<- *ContractChangedValidatorStatus, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ChangedValidatorStatus", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractChangedValidatorStatus)
				if err := _Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedValidatorStatus is a log parse operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) ParseChangedValidatorStatus(log types.Log) (*ContractChangedValidatorStatus, error) {
	event := new(ContractChangedValidatorStatus)
	if err := _Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimedRewardsIterator is returned from FilterClaimedRewards and is used to iterate over the raw logs and unpacked data for ClaimedRewards events raised by the Contract contract.
type ContractClaimedRewardsIterator struct {
	Event *ContractClaimedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractClaimedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractClaimedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractClaimedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimedRewards represents a ClaimedRewards event raised by the Contract contract.
type ContractClaimedRewards struct {
	Delegator         common.Address
	ToValidatorID     *big.Int
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimedRewards is a free log retrieval operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) FilterClaimedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractClaimedRewardsIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimedRewardsIterator{contract: _Contract.contract, event: "ClaimedRewards", logs: logs, sub: sub}, nil
}

// WatchClaimedRewards is a free log subscription operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) WatchClaimedRewards(opts *bind.WatchOpts, sink chan<- *ContractClaimedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimedRewards)
				if err := _Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedRewards is a log parse operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) ParseClaimedRewards(log types.Log) (*ContractClaimedRewards, error) {
	event := new(ContractClaimedRewards)
	if err := _Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCreatedValidatorIterator is returned from FilterCreatedValidator and is used to iterate over the raw logs and unpacked data for CreatedValidator events raised by the Contract contract.
type ContractCreatedValidatorIterator struct {
	Event *ContractCreatedValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCreatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreatedValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCreatedValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCreatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreatedValidator represents a CreatedValidator event raised by the Contract contract.
type ContractCreatedValidator struct {
	ValidatorID  *big.Int
	Auth         common.Address
	CreatedEpoch *big.Int
	CreatedTime  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreatedValidator is a free log retrieval operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) FilterCreatedValidator(opts *bind.FilterOpts, validatorID []*big.Int, auth []common.Address) (*ContractCreatedValidatorIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}
	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreatedValidator", validatorIDRule, authRule)
	if err != nil {
		return nil, err
	}
	return &ContractCreatedValidatorIterator{contract: _Contract.contract, event: "CreatedValidator", logs: logs, sub: sub}, nil
}

// WatchCreatedValidator is a free log subscription operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) WatchCreatedValidator(opts *bind.WatchOpts, sink chan<- *ContractCreatedValidator, validatorID []*big.Int, auth []common.Address) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}
	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreatedValidator", validatorIDRule, authRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreatedValidator)
				if err := _Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedValidator is a log parse operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) ParseCreatedValidator(log types.Log) (*ContractCreatedValidator, error) {
	event := new(ContractCreatedValidator)
	if err := _Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDeactivatedValidatorIterator is returned from FilterDeactivatedValidator and is used to iterate over the raw logs and unpacked data for DeactivatedValidator events raised by the Contract contract.
type ContractDeactivatedValidatorIterator struct {
	Event *ContractDeactivatedValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDeactivatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeactivatedValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDeactivatedValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDeactivatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeactivatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeactivatedValidator represents a DeactivatedValidator event raised by the Contract contract.
type ContractDeactivatedValidator struct {
	ValidatorID      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedValidator is a free log retrieval operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) FilterDeactivatedValidator(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractDeactivatedValidatorIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DeactivatedValidator", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeactivatedValidatorIterator{contract: _Contract.contract, event: "DeactivatedValidator", logs: logs, sub: sub}, nil
}

// WatchDeactivatedValidator is a free log subscription operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) WatchDeactivatedValidator(opts *bind.WatchOpts, sink chan<- *ContractDeactivatedValidator, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DeactivatedValidator", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeactivatedValidator)
				if err := _Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivatedValidator is a log parse operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) ParseDeactivatedValidator(log types.Log) (*ContractDeactivatedValidator, error) {
	event := new(ContractDeactivatedValidator)
	if err := _Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the Contract contract.
type ContractDelegatedIterator struct {
	Event *ContractDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegated represents a Delegated event raised by the Contract contract.
type ContractDelegated struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegatedIterator{contract: _Contract.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegated, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegated)
				if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegated is a log parse operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) ParseDelegated(log types.Log) (*ContractDelegated, error) {
	event := new(ContractDelegated)
	if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLockedUpStakeIterator is returned from FilterLockedUpStake and is used to iterate over the raw logs and unpacked data for LockedUpStake events raised by the Contract contract.
type ContractLockedUpStakeIterator struct {
	Event *ContractLockedUpStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractLockedUpStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLockedUpStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractLockedUpStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractLockedUpStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLockedUpStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLockedUpStake represents a LockedUpStake event raised by the Contract contract.
type ContractLockedUpStake struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Duration    *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLockedUpStake is a free log retrieval operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) FilterLockedUpStake(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractLockedUpStakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LockedUpStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractLockedUpStakeIterator{contract: _Contract.contract, event: "LockedUpStake", logs: logs, sub: sub}, nil
}

// WatchLockedUpStake is a free log subscription operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) WatchLockedUpStake(opts *bind.WatchOpts, sink chan<- *ContractLockedUpStake, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LockedUpStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLockedUpStake)
				if err := _Contract.contract.UnpackLog(event, "LockedUpStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockedUpStake is a log parse operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) ParseLockedUpStake(log types.Log) (*ContractLockedUpStake, error) {
	event := new(ContractLockedUpStake)
	if err := _Contract.contract.UnpackLog(event, "LockedUpStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRestakedRewardsIterator is returned from FilterRestakedRewards and is used to iterate over the raw logs and unpacked data for RestakedRewards events raised by the Contract contract.
type ContractRestakedRewardsIterator struct {
	Event *ContractRestakedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRestakedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRestakedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRestakedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRestakedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRestakedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRestakedRewards represents a RestakedRewards event raised by the Contract contract.
type ContractRestakedRewards struct {
	Delegator         common.Address
	ToValidatorID     *big.Int
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRestakedRewards is a free log retrieval operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) FilterRestakedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractRestakedRewardsIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractRestakedRewardsIterator{contract: _Contract.contract, event: "RestakedRewards", logs: logs, sub: sub}, nil
}

// WatchRestakedRewards is a free log subscription operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) WatchRestakedRewards(opts *bind.WatchOpts, sink chan<- *ContractRestakedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRestakedRewards)
				if err := _Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRestakedRewards is a log parse operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) ParseRestakedRewards(log types.Log) (*ContractRestakedRewards, error) {
	event := new(ContractRestakedRewards)
	if err := _Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Contract contract.
type ContractUndelegatedIterator struct {
	Event *ContractUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUndelegated represents a Undelegated event raised by the Contract contract.
type ContractUndelegated struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	WrID          *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*ContractUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUndelegatedIterator{contract: _Contract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *ContractUndelegated, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUndelegated)
				if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegated is a log parse operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) ParseUndelegated(log types.Log) (*ContractUndelegated, error) {
	event := new(ContractUndelegated)
	if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnlockedStakeIterator is returned from FilterUnlockedStake and is used to iterate over the raw logs and unpacked data for UnlockedStake events raised by the Contract contract.
type ContractUnlockedStakeIterator struct {
	Event *ContractUnlockedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUnlockedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnlockedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUnlockedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUnlockedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnlockedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnlockedStake represents a UnlockedStake event raised by the Contract contract.
type ContractUnlockedStake struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Amount      *big.Int
	Penalty     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlockedStake is a free log retrieval operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) FilterUnlockedStake(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractUnlockedStakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UnlockedStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUnlockedStakeIterator{contract: _Contract.contract, event: "UnlockedStake", logs: logs, sub: sub}, nil
}

// WatchUnlockedStake is a free log subscription operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) WatchUnlockedStake(opts *bind.WatchOpts, sink chan<- *ContractUnlockedStake, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UnlockedStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnlockedStake)
				if err := _Contract.contract.UnpackLog(event, "UnlockedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlockedStake is a log parse operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) ParseUnlockedStake(log types.Log) (*ContractUnlockedStake, error) {
	event := new(ContractUnlockedStake)
	if err := _Contract.contract.UnpackLog(event, "UnlockedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdatedBaseRewardPerSecIterator is returned from FilterUpdatedBaseRewardPerSec and is used to iterate over the raw logs and unpacked data for UpdatedBaseRewardPerSec events raised by the Contract contract.
type ContractUpdatedBaseRewardPerSecIterator struct {
	Event *ContractUpdatedBaseRewardPerSec // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedBaseRewardPerSecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedBaseRewardPerSec)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedBaseRewardPerSec)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedBaseRewardPerSecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedBaseRewardPerSecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedBaseRewardPerSec represents a UpdatedBaseRewardPerSec event raised by the Contract contract.
type ContractUpdatedBaseRewardPerSec struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseRewardPerSec is a free log retrieval operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) FilterUpdatedBaseRewardPerSec(opts *bind.FilterOpts) (*ContractUpdatedBaseRewardPerSecIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedBaseRewardPerSecIterator{contract: _Contract.contract, event: "UpdatedBaseRewardPerSec", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseRewardPerSec is a free log subscription operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) WatchUpdatedBaseRewardPerSec(opts *bind.WatchOpts, sink chan<- *ContractUpdatedBaseRewardPerSec) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedBaseRewardPerSec)
				if err := _Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedBaseRewardPerSec is a log parse operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) ParseUpdatedBaseRewardPerSec(log types.Log) (*ContractUpdatedBaseRewardPerSec, error) {
	event := new(ContractUpdatedBaseRewardPerSec)
	if err := _Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdatedOfflinePenaltyThresholdIterator is returned from FilterUpdatedOfflinePenaltyThreshold and is used to iterate over the raw logs and unpacked data for UpdatedOfflinePenaltyThreshold events raised by the Contract contract.
type ContractUpdatedOfflinePenaltyThresholdIterator struct {
	Event *ContractUpdatedOfflinePenaltyThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedOfflinePenaltyThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedOfflinePenaltyThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedOfflinePenaltyThreshold represents a UpdatedOfflinePenaltyThreshold event raised by the Contract contract.
type ContractUpdatedOfflinePenaltyThreshold struct {
	BlocksNum *big.Int
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOfflinePenaltyThreshold is a free log retrieval operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) FilterUpdatedOfflinePenaltyThreshold(opts *bind.FilterOpts) (*ContractUpdatedOfflinePenaltyThresholdIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedOfflinePenaltyThreshold")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedOfflinePenaltyThresholdIterator{contract: _Contract.contract, event: "UpdatedOfflinePenaltyThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedOfflinePenaltyThreshold is a free log subscription operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) WatchUpdatedOfflinePenaltyThreshold(opts *bind.WatchOpts, sink chan<- *ContractUpdatedOfflinePenaltyThreshold) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedOfflinePenaltyThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedOfflinePenaltyThreshold)
				if err := _Contract.contract.UnpackLog(event, "UpdatedOfflinePenaltyThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedOfflinePenaltyThreshold is a log parse operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) ParseUpdatedOfflinePenaltyThreshold(log types.Log) (*ContractUpdatedOfflinePenaltyThreshold, error) {
	event := new(ContractUpdatedOfflinePenaltyThreshold)
	if err := _Contract.contract.UnpackLog(event, "UpdatedOfflinePenaltyThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdatedSlashingRefundRatioIterator is returned from FilterUpdatedSlashingRefundRatio and is used to iterate over the raw logs and unpacked data for UpdatedSlashingRefundRatio events raised by the Contract contract.
type ContractUpdatedSlashingRefundRatioIterator struct {
	Event *ContractUpdatedSlashingRefundRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedSlashingRefundRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedSlashingRefundRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedSlashingRefundRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedSlashingRefundRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedSlashingRefundRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedSlashingRefundRatio represents a UpdatedSlashingRefundRatio event raised by the Contract contract.
type ContractUpdatedSlashingRefundRatio struct {
	ValidatorID *big.Int
	RefundRatio *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSlashingRefundRatio is a free log retrieval operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) FilterUpdatedSlashingRefundRatio(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractUpdatedSlashingRefundRatioIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedSlashingRefundRatioIterator{contract: _Contract.contract, event: "UpdatedSlashingRefundRatio", logs: logs, sub: sub}, nil
}

// WatchUpdatedSlashingRefundRatio is a free log subscription operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) WatchUpdatedSlashingRefundRatio(opts *bind.WatchOpts, sink chan<- *ContractUpdatedSlashingRefundRatio, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedSlashingRefundRatio)
				if err := _Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedSlashingRefundRatio is a log parse operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) ParseUpdatedSlashingRefundRatio(log types.Log) (*ContractUpdatedSlashingRefundRatio, error) {
	event := new(ContractUpdatedSlashingRefundRatio)
	if err := _Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Contract contract.
type ContractWithdrawnIterator struct {
	Event *ContractWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdrawn represents a Withdrawn event raised by the Contract contract.
type ContractWithdrawn struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	WrID          *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) FilterWithdrawn(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*ContractWithdrawnIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawnIterator{contract: _Contract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractWithdrawn, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) ParseWithdrawn(log types.Log) (*ContractWithdrawn, error) {
	event := new(ContractWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
