// Copyright 2018 The klaytn Authors
// Copyright 2014 The go-ethereum Authors
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
//
// This file is derived from core/blockchain/types/block.go (2018/06/04).
// Modified and improved for the klaytn development.

/*
Package types contains data types related to Klatyn consensus.

Klaytn nodes achieve a consensus of the global state by processing the same transactions in the same order.

Transaction is an atomic operation originated from an Externally Owned Account (EOA) transiting the global state of Klaytn.
Klaytn supports multiple transaction types to efficiently trigger KLAY transfer or contract execution.
LegacyTransaction, the Ethereum compatible transaction, could be sent to both of EOA and Smart Contract (SC).
However, it cannot support advanced features of Klaytn such as multi-sig or fee-delegation.
All transaction types except LegacyTransaction are enforced to be sent to either of EOA or SC.
With a slice of transaction signatures signed by multiple accounts, they can support a multi-sig account to transact.
Some transaction types support full or partial fee-delegation feature as well as basic transaction features.
Fee-delegation transactions have additional field(s) to charge some amount of transaction fee to a fee payer's account.

Block consists of transactions and a header which includes additional information for consensus and transaction support.
A block guarantees the sequential execution of transactions inside, and a block header enables Klaytn nodes to order the blocks.
The block header also provides additional data generated by the transaction execution such as transaction receipts, and logs.

This package implements Klaytn Block, Transaction and related functions and variables to support them.

# Source Files

Functions and variables related to Block and Transaction are defined in the files listed below.
  - block.go : defines block and block header
  - bloom.go : defines bloom and its functions for bloom filter which helps user to filter transaction receipts efficiently
  - contract_ref.go : interfaces ContractRef to deliver a fee payer address to the contract execution environment
  - derive_sha.go : implements keccak256 hash functions in various ways
  - gen_header_json.go : auto-generated code for JSON marshalling/un-marshalling of block header
  - gen_log_json.go : auto-generated code for JSON marshalling/un-marshalling of contract execution log
  - gen_receipt_json.go : auto-generated code for JSON marshalling/un-marshalling of transaction receipt
  - istanbul.go : provides a block header modified for istanbul consensus
  - log.go : implements contract execution log
  - receipt.go : implements transaction receipt
  - transaction.go : defines transaction
  - transaction_signing.go : interfaces signer and implements transaction signing/verification functions
  - tx_internal_data.go : defines internal data of transaction supporting various transaction types
  - tx_internal_data_account_creation.go : implements the transaction creating an EOA account
  - tx_internal_data_account_update.go : implements the transaction updating account key of an account
  - tx_internal_data_cancel.go : implements the transaction canceling a transaction in the txpool
  - tx_internal_data_chain_data_anchoring.go : implements the transaction transferring data to service chain
  - tx_internal_data_fee_delegated_account_update.go : implements the fee-delegated version of account update transaction
  - tx_internal_data_fee_delegated_account_update_with_ratio.go : implements the partially fee-delegated version of account update transaction
  - tx_internal_data_fee_delegated_cancel.go: implements the fee-delegated version of cancel transaction
  - tx_internal_data_fee_delegated_cancel_with_ratio.go: implements the partially fee-delegated version of cancel transaction
  - tx_internal_data_fee_delegated_smart_contract_deploy.go : implements the fee-delegated version of contract deploy transaction
  - tx_internal_data_fee_delegated_smart_contract_deploy_with_ratio.go : implements the partially fee-delegated version of contract deploy transaction
  - tx_internal_data_fee_delegated_smart_contract_execution.go : implements the fee-delegated version of contract execution transaction
  - tx_internal_data_fee_delegated_smart_contract_execution_with_ratio.go : implements the partially fee-delegated version of contract execution transaction
  - tx_internal_data_fee_delegated_value_transfer.go: implements the fee-delegated version of value transfer transaction
  - tx_internal_data_fee_delegated_value_transfer_memo.go: implements the fee-delegated version of value transfer with memo transaction
  - tx_internal_data_fee_delegated_value_transfer_memo_with_ratio.go: implements the partially fee-delegated version of value transfer with memo transaction
  - tx_internal_data_fee_delegated_value_transfer_with_ratio.go :implements the partially fee-delegated version of value transfer transaction
  - tx_internal_data_legacy.go: implements the legacy transaction compatible with Ethereum
  - tx_internal_data_serializer.go: implements serialization functions of transaction internal data
  - tx_internal_data_smart_contract_deploy.go: implements the transaction deploying a smart contract
  - tx_internal_data_smart_contract_execution.go: implements the transaction executing a smart contract
  - tx_internal_data_value_transfer.go: implements the transaction sending KLAY to an EOA
  - tx_internal_data_value_transfer_memo.go: implements the transaction sending KLAY to an EOA with data
  - tx_signature.go : implements transaction signature (V, R, S)
  - tx_signatures.go : implements a slice of transaction signature to support multi-sig accounts
*/
package types
