// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package blockchain

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/common/math"
	"github.com/klaytn/klaytn/log"
	"github.com/klaytn/klaytn/params"
)

var _ = (*genesisSpecMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (g Genesis) MarshalJSON() ([]byte, error) {
	type Genesis struct {
		Config     *params.ChainConfig                         `json:"config"`
		Timestamp  math.HexOrDecimal64                         `json:"timestamp"`
		ExtraData  hexutil.Bytes                               `json:"extraData"`
		Governance []byte                                      `json:"governanceData"`
		BlockScore *math.HexOrDecimal256                       `json:"blockScore"`
		Alloc      map[common.UnprefixedAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		Number     math.HexOrDecimal64                         `json:"number"`
		GasUsed    math.HexOrDecimal64                         `json:"gasUsed"`
		ParentHash common.Hash                                 `json:"parentHash"`
	}
	var enc Genesis
	enc.Config = g.Config
	enc.Timestamp = math.HexOrDecimal64(g.Timestamp)
	enc.ExtraData = g.ExtraData
	enc.Governance = g.Governance
	enc.BlockScore = (*math.HexOrDecimal256)(g.BlockScore)
	if g.Alloc != nil {
		enc.Alloc = make(map[common.UnprefixedAddress]GenesisAccount, len(g.Alloc))
		for k, v := range g.Alloc {
			enc.Alloc[common.UnprefixedAddress(k)] = v
		}
	}
	enc.Number = math.HexOrDecimal64(g.Number)
	enc.GasUsed = math.HexOrDecimal64(g.GasUsed)
	enc.ParentHash = g.ParentHash
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (g *Genesis) UnmarshalJSON(input []byte) error {
	type Genesis struct {
		Config     *params.ChainConfig                         `json:"config"`
		Timestamp  *math.HexOrDecimal64                        `json:"timestamp"`
		ExtraData  *hexutil.Bytes                              `json:"extraData"`
		Governance []byte                                      `json:"governanceData"`
		BlockScore *math.HexOrDecimal256                       `json:"blockScore"`
		Alloc      map[common.UnprefixedAddress]GenesisAccount `json:"alloc"      gencodec:"required"`
		Number     *math.HexOrDecimal64                        `json:"number"`
		GasUsed    *math.HexOrDecimal64                        `json:"gasUsed"`
		ParentHash *common.Hash                                `json:"parentHash"`
	}
	printUnknownFields := func(input []byte) {
		// print unknown fields in genesis configuration
		gen := new(Genesis)
		decoder := json.NewDecoder(bytes.NewReader(input))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(gen); err != nil {
			log.NewModuleLogger(log.CMDUtilsNodeCMD).Error("Unmarshal error", "error", err)
		}
	}

	var dec Genesis
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Config != nil {
		g.Config = dec.Config
	}
	if dec.Timestamp != nil {
		g.Timestamp = uint64(*dec.Timestamp)
	}
	if dec.ExtraData != nil {
		g.ExtraData = *dec.ExtraData
	}
	if dec.Governance != nil {
		g.Governance = dec.Governance
	}
	if dec.BlockScore != nil {
		g.BlockScore = (*big.Int)(dec.BlockScore)
	}
	if dec.Alloc == nil {
		return errors.New("missing required field 'alloc' for Genesis")
	}
	g.Alloc = make(GenesisAlloc, len(dec.Alloc))
	for k, v := range dec.Alloc {
		g.Alloc[common.Address(k)] = v
	}
	if dec.Number != nil {
		g.Number = uint64(*dec.Number)
	}
	if dec.GasUsed != nil {
		g.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.ParentHash != nil {
		g.ParentHash = *dec.ParentHash
	}
	printUnknownFields(input)
	return nil
}
