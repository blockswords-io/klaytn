package governance

import (
	"errors"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/consensus/istanbul"
	"github.com/klaytn/klaytn/params"
	"github.com/klaytn/klaytn/storage/database"
)

type MixedEngine struct {
	initialConfig *params.ChainConfig
	initialParams *params.GovParamSet
	currentParams *params.GovParamSet

	// Subordinate engines
	defaultGov     HeaderEngine
	contractEngine ReaderEngine
}

// newMixedEngine instantiate a new MixedEngine struct.
// Only if doInit is true, subordinate engines will be initialized.
func newMixedEngine(config *params.ChainConfig, db database.DBManager, doInit bool) *MixedEngine {
	e := &MixedEngine{}

	if govParams, err := params.NewGovParamSetChainConfig(config); err == nil {
		e.initialConfig = config
		e.initialParams = govParams
		e.currentParams = govParams
	} else {
		logger.Crit("Error parsing initial ChainConfig", "err", err)
	}

	// Setup subordinate engines
	if doInit {
		e.defaultGov = NewGovernanceInitialize(config, db)
	} else {
		e.defaultGov = NewGovernance(config, db)
	}

	e.contractEngine = NewContractEngine(config)

	// Load last state
	e.UpdateParams()
	return e
}

// Developers are encouraged to call this constructor in most cases.
func NewMixedEngine(config *params.ChainConfig, db database.DBManager) *MixedEngine {
	return newMixedEngine(config, db, true)
}

// Does not load initial data for test purposes
func NewMixedEngineNoInit(config *params.ChainConfig, db database.DBManager) *MixedEngine {
	return newMixedEngine(config, db, false)
}

func (e *MixedEngine) Params() *params.GovParamSet {
	return e.currentParams
}

func (e *MixedEngine) ParamsAt(num uint64) (*params.GovParamSet, error) {
	_, strMap, err := e.defaultGov.ReadGovernance(num)
	if err != nil {
		return nil, err
	}

	headerParams, err := params.NewGovParamSetStrMap(strMap)
	if err != nil {
		return nil, err
	}

	// If governancegen is "header" in either headerParams or initialParams,
	// use contractgen.
	isContractGen := false
	if v, ok := headerParams.Get(params.GovernanceGen); ok {
		if v.(string) == "contract" {
			isContractGen = true
		}
	} else if v, ok := e.initialParams.Get(params.GovernanceGen); ok {
		if v.(string) == "contract" {
			isContractGen = true
		}
	}

	if isContractGen {
		if e.contractEngine == nil {
			logger.Error("ContractEngine is not ready")
			return nil, errors.New("ContractEngine is not ready")
		}
		return e.contractEngine.ParamsAt(num)
	} else {
		// Use ReadGovernance(), fallback to initialParams.
		p := params.NewGovParamSetMerged(e.initialParams, headerParams)
		return p, nil
	}
}

func (e *MixedEngine) UpdateParams() error {
	strMap := e.defaultGov.CurrentSet()
	govParams, err := params.NewGovParamSetStrMap(strMap)
	if err != nil {
		return err
	}

	// Use CurrentSet(), fallback to initialParams.
	e.currentParams = params.NewGovParamSetMerged(e.initialParams, govParams)
	return nil
}

func (e *MixedEngine) SetBlockchain(chain blockChain) {
	e.defaultGov.SetBlockchain(chain)
	if e.contractEngine != nil {
		e.contractEngine.SetBlockchain(chain)
	}
}

// Pass-through to HeaderEngine
func (e *MixedEngine) AddVote(key string, val interface{}) bool {
	return e.defaultGov.AddVote(key, val)
}

func (e *MixedEngine) ValidateVote(vote *GovernanceVote) (*GovernanceVote, bool) {
	return e.defaultGov.ValidateVote(vote)
}

func (e *MixedEngine) CanWriteGovernanceState(num uint64) bool {
	return e.defaultGov.CanWriteGovernanceState(num)
}

func (e *MixedEngine) WriteGovernanceState(num uint64, isCheckpoint bool) error {
	return e.defaultGov.WriteGovernanceState(num, isCheckpoint)
}

func (e *MixedEngine) ReadGovernance(num uint64) (uint64, map[string]interface{}, error) {
	return e.defaultGov.ReadGovernance(num)
}

func (e *MixedEngine) WriteGovernance(num uint64, data GovernanceSet, delta GovernanceSet) error {
	return e.defaultGov.WriteGovernance(num, data, delta)
}

func (e *MixedEngine) GetEncodedVote(addr common.Address, number uint64) []byte {
	return e.defaultGov.GetEncodedVote(addr, number)
}

func (e *MixedEngine) GetGovernanceChange() map[string]interface{} {
	return e.defaultGov.GetGovernanceChange()
}

func (e *MixedEngine) VerifyGovernance(received []byte) error {
	return e.defaultGov.VerifyGovernance(received)
}

func (e *MixedEngine) ClearVotes(num uint64) {
	e.defaultGov.ClearVotes(num)
}

func (e *MixedEngine) WriteGovernanceForNextEpoch(number uint64, governance []byte) {
	e.defaultGov.WriteGovernanceForNextEpoch(number, governance)
}

func (e *MixedEngine) UpdateCurrentSet(num uint64) {
	e.defaultGov.UpdateCurrentSet(num)
}

func (e *MixedEngine) HandleGovernanceVote(
	valset istanbul.ValidatorSet, votes []GovernanceVote, tally []GovernanceTallyItem,
	header *types.Header, proposer common.Address, self common.Address,
) (
	istanbul.ValidatorSet, []GovernanceVote, []GovernanceTallyItem,
) {
	return e.defaultGov.HandleGovernanceVote(valset, votes, tally, header, proposer, self)
}

func (e *MixedEngine) ChainId() uint64 {
	return e.defaultGov.ChainId()
}

func (e *MixedEngine) InitialChainConfig() *params.ChainConfig {
	return e.defaultGov.InitialChainConfig()
}

func (e *MixedEngine) GetVoteMapCopy() map[string]VoteStatus {
	return e.defaultGov.GetVoteMapCopy()
}

func (e *MixedEngine) GetGovernanceTalliesCopy() []GovernanceTallyItem {
	return e.defaultGov.GetGovernanceTalliesCopy()
}

func (e *MixedEngine) CurrentSet() map[string]interface{} {
	return e.defaultGov.CurrentSet()
}

func (e *MixedEngine) PendingChanges() map[string]interface{} {
	return e.defaultGov.PendingChanges()
}

func (e *MixedEngine) Votes() []GovernanceVote {
	return e.defaultGov.Votes()
}

func (e *MixedEngine) IdxCache() []uint64 {
	return e.defaultGov.IdxCache()
}

func (e *MixedEngine) IdxCacheFromDb() []uint64 {
	return e.defaultGov.IdxCacheFromDb()
}

func (e *MixedEngine) NodeAddress() common.Address {
	return e.defaultGov.NodeAddress()
}

func (e *MixedEngine) TotalVotingPower() uint64 {
	return e.defaultGov.TotalVotingPower()
}

func (e *MixedEngine) MyVotingPower() uint64 {
	return e.defaultGov.MyVotingPower()
}

func (e *MixedEngine) BlockChain() blockChain {
	return e.defaultGov.BlockChain()
}

func (e *MixedEngine) DB() database.DBManager {
	return e.defaultGov.DB()
}

func (e *MixedEngine) SetNodeAddress(addr common.Address) {
	e.defaultGov.SetNodeAddress(addr)
}

func (e *MixedEngine) SetTotalVotingPower(t uint64) {
	e.defaultGov.SetTotalVotingPower(t)
}

func (e *MixedEngine) SetMyVotingPower(t uint64) {
	e.defaultGov.SetMyVotingPower(t)
}

func (e *MixedEngine) SetTxPool(txpool txPool) {
	e.defaultGov.SetTxPool(txpool)
}
