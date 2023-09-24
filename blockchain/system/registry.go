// Copyright 2023 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

package system

import (
	"context"
	"math/big"

	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/state"
	"github.com/klaytn/klaytn/common"
	contracts "github.com/klaytn/klaytn/contracts/system_contracts"
)

type RegistryRecord struct {
	Addr       common.Address
	Activation *big.Int
}

type AllocRegistryInit struct {
	Records map[string][]RegistryRecord
	Names   []string
	Owner   common.Address
}

// Create storage state from the given initial values.
// The storage slots are calculated according to the solidity layout rule.
// https://docs.soliditylang.org/en/v1.8.20/internals/layout_in_storage.html
func AllocRegistry(init *AllocRegistryInit) map[common.Hash]common.Hash {
	storage := make(map[common.Hash]common.Hash)

	// slot[0]: mapping(string => Record[]) records;
	// - records[x].length @ Hash(x, 0)
	// - records[x][i].addr @ Hash(Hash(x, 0)) + (2*i)
	// - records[x][i].activation @ Hash(Hash(x, 0)) + (2*i + 1)
	for name, records := range init.Records {
		arraySlot := calcMappingSlot(0, name)
		storage[arraySlot] = lpad32(len(records))

		for i, record := range records {
			addrSlot := calcArraySlot(arraySlot, 2, i, 0)
			activationSlot := calcArraySlot(arraySlot, 2, i, 1)

			storage[addrSlot] = lpad32(record.Addr)
			storage[activationSlot] = lpad32(record.Activation)
		}
	}

	// slot[1]: string[] names;
	// - names.length @ 1
	// - names[i] @ Hash(1) + i
	storage[lpad32(1)] = lpad32(len(init.Names))
	for i, name := range init.Names {
		nameSlot := calcArraySlot(1, 1, i, 0)
		storage[nameSlot] = encodeShortString(name)
	}

	// slot[2]: address _owner;
	storage[lpad32(2)] = lpad32(init.Owner)

	return storage
}

func InstallRegistry(state *state.StateDB) error {
	return state.SetCode(RegistryAddr, RegistryCode)
}

func ReadRegistryActiveAddr(backend bind.ContractCaller, name string, num *big.Int) (common.Address, error) {
	caller, err := contracts.NewRegistryCaller(RegistryAddr, backend)
	if err != nil {
		return common.Address{}, err
	}

	code, err := backend.CodeAt(context.Background(), RegistryAddr, nil)
	if err != nil {
		return common.Address{}, err
	}
	if code == nil {
		return common.Address{}, ErrRegistryNotInstalled
	}

	opts := &bind.CallOpts{BlockNumber: num}
	return caller.GetActiveAddr(opts, name)
}
