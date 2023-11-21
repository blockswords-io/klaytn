// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"fmt"

	"github.com/holiman/uint256"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/params"
)

// EnableEIP enables the given EIP on the config.
// This operation writes in-place, and callers need to ensure that the globally
// defined jump tables are not polluted.
func EnableEIP(eipNum int, jt *JumpTable) error {
	switch eipNum {
	case 4844:
		enable4844(jt)
	case 7516:
		enable7516(jt)
	case 6780:
		enable6780(jt)
	case 5656:
		enable5656(jt)
	case 3860:
		enable3860(jt)
	case 3855:
		enable3855(jt)
	case 4399:
		enable4399(jt)
	case 3529:
		enable3529(jt)
	case 2929:
		enable2929(jt)
	case 2200:
		enable2200(jt)
	case 1884:
		enable1884(jt)
	case 1344:
		enable1344(jt)
	case 1153:
		enable1153(jt)
	default:
		return fmt.Errorf("undefined eip %d", eipNum)
	}
	return nil
}

// enable1884 applies EIP-1884 to the given jump table:
// - Increase cost of BALANCE to 700
// - Increase cost of EXTCODEHASH to 700
// - Increase cost of SLOAD to 800
// - Define SELFBALANCE, with cost GasFastStep (5)
func enable1884(jt *JumpTable) {
	// Gas cost changes
	jt[SLOAD].constantGas = params.SloadGasEIP1884
	jt[BALANCE].constantGas = params.BalanceGasEIP1884
	jt[EXTCODEHASH].constantGas = params.ExtcodeHashGasEIP1884

	// New opcode
	jt[SELFBALANCE] = &operation{
		execute:         opSelfBalance,
		constantGas:     GasFastStep,
		minStack:        minStack(0, 1),
		maxStack:        maxStack(0, 1),
		computationCost: params.SelfBalanceComputationCost,
	}
}

func opSelfBalance(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	balance, _ := uint256.FromBig(interpreter.evm.StateDB.GetBalance(scope.Contract.Address()))
	scope.Stack.push(balance)
	return nil, nil
}

// enable1344 applies EIP-1344 (ChainID Opcode)
// - Adds an opcode that returns the current chain’s EIP-155 unique identifier
func enable1344(jt *JumpTable) {
	// New opcode
	jt[CHAINID] = &operation{
		execute:         opChainID,
		constantGas:     GasQuickStep,
		minStack:        minStack(0, 1),
		maxStack:        maxStack(0, 1),
		computationCost: params.ChainIDComputationCost,
	}
}

// opChainID implements CHAINID opcode
func opChainID(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	chainId, _ := uint256.FromBig(interpreter.evm.chainConfig.ChainID)
	scope.Stack.push(chainId)
	return nil, nil
}

// enable2200 applies EIP-2200 (Rebalance net-metered SSTORE)
func enable2200(jt *JumpTable) {
	jt[SLOAD].constantGas = params.SloadGasEIP2200
	jt[SSTORE].dynamicGas = gasSStoreEIP2200
}

// enableIstanbulComputationCostModification modifies ADDMOD, MULMOD, NOT, XOR, SHL, SHR, SAR computation cost
// The modification is activated with istanbulCompatible change activation.
func enableIstanbulComputationCostModification(jt *JumpTable) {
	jt[ADDMOD].computationCost = params.AddmodComputationCostIstanbul
	jt[MULMOD].computationCost = params.MulmodComputationCostIstanbul
	jt[NOT].computationCost = params.NotComputationCostIstanbul
	jt[XOR].computationCost = params.XorComputationCostIstanbul
	jt[SHL].computationCost = params.ShlComputationCostIstanbul
	jt[SHR].computationCost = params.ShrComputationCostIstanbul
	jt[SAR].computationCost = params.SarComputationCostIstanbul
}

// enable1153 applies EIP-1153 "Transient Storage"
// - Adds TLOAD that reads from transient storage
// - Adds TSTORE that writes to transient storage
func enable1153(jt *JumpTable) {
	jt[TLOAD] = &operation{
		execute:         opTload,
		constantGas:     params.WarmStorageReadCostEIP2929,
		minStack:        minStack(1, 1),
		maxStack:        maxStack(1, 1),
		computationCost: params.TloadComputationCost,
	}

	jt[TSTORE] = &operation{
		execute:         opTstore,
		constantGas:     params.WarmStorageReadCostEIP2929,
		minStack:        minStack(2, 0),
		maxStack:        maxStack(2, 0),
		computationCost: params.TstoreComputationCost,
	}
}

// opTload implements TLOAD opcode
func opTload(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	loc := scope.Stack.peek()
	hash := common.Hash(loc.Bytes32())
	val := interpreter.evm.StateDB.GetTransientState(scope.Contract.Address(), hash)
	loc.SetBytes(val.Bytes())
	return nil, nil
}

// opTstore implements TSTORE opcode
func opTstore(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	if interpreter.readOnly {
		return nil, ErrWriteProtection
	}
	loc := scope.Stack.pop()
	val := scope.Stack.pop()
	interpreter.evm.StateDB.SetTransientState(scope.Contract.Address(), loc.Bytes32(), val.Bytes32())
	return nil, nil
}

// enable3198 applies EIP-3198 (BASEFEE Opcode)
// - Adds an opcode that returns the current block's base fee.
func enable3198(jt *JumpTable) {
	// New opcode
	jt[BASEFEE] = &operation{
		execute:         opBaseFee,
		constantGas:     GasQuickStep,
		minStack:        minStack(0, 1),
		maxStack:        maxStack(0, 1),
		computationCost: params.BaseFeeComputationCost,
	}
}

// opBaseFee implements BASEFEE opcode
func opBaseFee(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	baseFee, _ := uint256.FromBig(interpreter.evm.Context.BaseFee)
	scope.Stack.push(baseFee)
	return nil, nil
}

// enable2929 enables "EIP-2929: Gas cost increases for state access opcodes"
// https://eips.ethereum.org/EIPS/eip-2929
func enable2929(jt *JumpTable) {
	jt[SSTORE].dynamicGas = gasSStoreEIP2929

	jt[SLOAD].constantGas = 0
	jt[SLOAD].dynamicGas = gasSLoadEIP2929

	jt[EXTCODECOPY].constantGas = params.WarmStorageReadCostEIP2929
	jt[EXTCODECOPY].dynamicGas = gasExtCodeCopyEIP2929

	jt[EXTCODESIZE].constantGas = params.WarmStorageReadCostEIP2929
	jt[EXTCODESIZE].dynamicGas = gasEip2929AccountCheck

	jt[EXTCODEHASH].constantGas = params.WarmStorageReadCostEIP2929
	jt[EXTCODEHASH].dynamicGas = gasEip2929AccountCheck

	jt[BALANCE].constantGas = params.WarmStorageReadCostEIP2929
	jt[BALANCE].dynamicGas = gasEip2929AccountCheck

	jt[CALL].constantGas = params.WarmStorageReadCostEIP2929
	jt[CALL].dynamicGas = gasCallEIP2929

	jt[CALLCODE].constantGas = params.WarmStorageReadCostEIP2929
	jt[CALLCODE].dynamicGas = gasCallCodeEIP2929

	jt[STATICCALL].constantGas = params.WarmStorageReadCostEIP2929
	jt[STATICCALL].dynamicGas = gasStaticCallEIP2929

	jt[DELEGATECALL].constantGas = params.WarmStorageReadCostEIP2929
	jt[DELEGATECALL].dynamicGas = gasDelegateCallEIP2929

	// This was previously part of the dynamic cost, but we're using it as a constantGas
	// factor here
	jt[SELFDESTRUCT].constantGas = params.SelfdestructGas
	jt[SELFDESTRUCT].dynamicGas = gasSelfdestructEIP2929
}

// enable3529 enabled "EIP-3529: Reduction in refunds":
// - Removes refunds for selfdestructs
// - Reduces refunds for SSTORE
// - Reduces max refunds to 20% gas
func enable3529(jt *JumpTable) {
	jt[SSTORE].dynamicGas = gasSStoreEIP3529
	jt[SELFDESTRUCT].dynamicGas = gasSelfdestructEIP3529
}

// enable4399 applies EIP-4399 (PREVRANDAO Opcode)
// - Change the 0x44 opcode from returning difficulty value to returning prev blockhash value
func enable4399(jt *JumpTable) {
	jt[PREVRANDAO] = &operation{
		execute:         opRandom,
		constantGas:     GasQuickStep,
		minStack:        minStack(0, 1),
		maxStack:        maxStack(0, 1),
		computationCost: params.RandomComputationCost,
	}
}

// enable3855 applies EIP-3855 (PUSH0 opcode)
func enable3855(jt *JumpTable) {
	jt[PUSH0] = &operation{
		execute:         opPush0,
		constantGas:     GasQuickStep,
		minStack:        minStack(0, 1),
		maxStack:        maxStack(0, 1),
		computationCost: params.Push0ComputationCost,
	}
}

// opPush0 implements the PUSH0 opcode
func opPush0(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	scope.Stack.push(new(uint256.Int))
	return nil, nil
}

// enable3860 enables "EIP-3860: Limit and meter initcode"
// https://eips.ethereum.org/EIPS/eip-3860
func enable3860(jt *JumpTable) {
	jt[CREATE].dynamicGas = gasCreateEip3860
	jt[CREATE2].dynamicGas = gasCreate2Eip3860
}

// enable5656 enables EIP-5656 (MCOPY opcode)
// https://eips.ethereum.org/EIPS/eip-5656
func enable5656(jt *JumpTable) {
	jt[MCOPY] = &operation{
		execute:         opMcopy,
		constantGas:     GasFastestStep,
		dynamicGas:      gasMcopy,
		minStack:        minStack(3, 0),
		maxStack:        maxStack(3, 0),
		memorySize:      memoryMcopy,
		computationCost: params.McopyComputationCost,
	}
}

// opMcopy implements the MCOPY opcode (https://eips.ethereum.org/EIPS/eip-5656)
func opMcopy(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	var (
		dst    = scope.Stack.pop()
		src    = scope.Stack.pop()
		length = scope.Stack.pop()
	)
	// These values are checked for overflow during memory expansion calculation
	// (the memorySize function on the opcode).
	scope.Memory.Copy(dst.Uint64(), src.Uint64(), length.Uint64())
	return nil, nil
}

// enable6780 applies EIP-6780 (deactivate SELFDESTRUCT)
func enable6780(jt *JumpTable) {
	jt[SELFDESTRUCT] = &operation{
		execute:         opSelfdestruct6780,
		dynamicGas:      gasSelfdestructEIP3529,
		constantGas:     params.SelfdestructGas,
		minStack:        minStack(1, 0),
		maxStack:        maxStack(1, 0),
		computationCost: params.SelfDestructComputationCost,
	}
}

// opBlobHash implements the BLOBHASH opcode
// Since blob data is generated in dank sharding, opBlobHash will perform only the default action of setting the top of the stack as zero
// as long as the blob txType is not fully supported in Klaytn.
func opBlobHash(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	index := scope.Stack.peek()
	index.Clear()

	return nil, nil
}

// opBlobBaseFee implements BLOBBASEFEE opcode
// Since blob data is generated in dank sharding, opBlobBaseFee will use only the zeroBaseFee
// as long as the blob txType is not fully supported in Klaytn.
func opBlobBaseFee(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	blobBaseFee := uint256.NewInt(params.ZeroBaseFee)
	scope.Stack.push(blobBaseFee)
	return nil, nil
}

// enable4844 applies EIP-4844 (BLOBHASH opcode)
func enable4844(jt *JumpTable) {
	jt[BLOBHASH] = &operation{
		execute:         opBlobHash,
		constantGas:     GasFastestStep,
		minStack:        minStack(1, 1),
		maxStack:        maxStack(1, 1),
		computationCost: params.BlobHashComptationCost,
	}
}

// enable7516 applies EIP-7516 (BLOBBASEFEE opcode)
func enable7516(jt *JumpTable) {
	jt[BLOBBASEFEE] = &operation{
		execute:         opBlobBaseFee,
		constantGas:     GasQuickStep,
		minStack:        minStack(0, 1),
		maxStack:        maxStack(0, 1),
		computationCost: params.BlobBaseFeeComputationCost,
	}
}

func enable1052(jt *JumpTable) {
	jt[EXTCODEHASH] = &operation{
		execute:         opExtCodeHash1052,
		constantGas:     params.ExtcodeHashGasConstantinople,
		minStack:        minStack(1, 1),
		maxStack:        maxStack(1, 1),
		computationCost: params.ExtCodeHashComputationCost,
	}
}

// As the cpu performance has been improved a lot, and as the storage size has increased a lot
// recalculated the computation cost of some opcodes
func enableCancunComputationCostModification(jt *JumpTable) {
	jt[SDIV].computationCost = params.SdivComputationCostCancun
	jt[MOD].computationCost = params.ModComputationCostCancun
	jt[ADDMOD].computationCost = params.AddmodComputationCostCancun
	jt[MULMOD].computationCost = params.MulmodComputationCostCancun
	jt[EXP].computationCost = params.ExpComputationCostCancun
	jt[SHA3].computationCost = params.Sha3ComputationCostCancun
	jt[MSTORE8].computationCost = params.Mstore8ComputationCostCancun

	jt[SLOAD].computationCost = params.SloadComputationCostCancun
	jt[SSTORE].computationCost = params.SstoreComputationCostCancun
	jt[LOG1].computationCost = params.Log1ComputationCostCancun
	jt[LOG2].computationCost = params.Log2ComputationCostCancun
	jt[LOG3].computationCost = params.Log3ComputationCostCancun
	jt[LOG4].computationCost = params.Log4ComputationCostCancun
}
