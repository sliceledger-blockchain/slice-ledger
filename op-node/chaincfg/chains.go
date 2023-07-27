package chaincfg

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/sliceledger-blockchain/slice-ledger/op-node/eth"
	"github.com/sliceledger-blockchain/slice-ledger/op-node/rollup"
)

var SliceMainnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x438335a20d98863a4c0c97999eb2481921ccd28553eac6f913af7c12aec04108"),
			Number: 17422590,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0xdbf6a80fef073de06add9b0d14026d6e5a86c85f6d102c36d3d8e9cf89c2afd3"),
			Number: 105235063,
		},
		L2Time: 1686068903,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x6887246668a3b87f54deb3b94ba47a6f63f32985"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000a6fe0")),
			GasLimit:    30_000_000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(1),
	L2ChainID:              big.NewInt(10),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000010"),
	DepositContractAddress: common.HexToAddress("0xbEb5Fc579115071764c7423A4f12eDde41f106Ed"),
	L1SystemConfigAddress:  common.HexToAddress("0x229047fed2591dbec1eF1118d64F7aF3dB9EB290"),
	RegolithTime:           u64Ptr(0),
}

var SliceGoerli = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x010e156149eedf17dad2c6ed48152ae9edbedce66b6c8085549fe041c272af22"),
			Number: 9415776,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0xff9635794b995836947f7d7f2a60deb7225b9b4dbaf14d0b94e07fd8d0a66d96"),
			Number: 0,
		},
		L2Time: 1690441440,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x4501285e8ffe2d38828f1e69fb57a6a2faac2620"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    25_000_000,
		},
	},
	BlockTime:              2,
	MaxSequencerDrift:      600,
	SeqWindowSize:          3600,
	ChannelTimeout:         300,
	L1ChainID:              big.NewInt(5),
	L2ChainID:              big.NewInt(8900),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000008900"),
	DepositContractAddress: common.HexToAddress("0x1e016b668d7dc57bf15402136932882f3da661e6"),
	L1SystemConfigAddress:  common.HexToAddress("0xea20db4746ccac5b37fd6f09dc2520269b0a7ab5"),
	RegolithTime:           u64Ptr(0),
}

var NetworksByName = map[string]rollup.Config{
	"Slice Testnet": SliceGoerli,
	"Slice mainnet": SliceMainnet,
}

var L2ChainIDToNetworkName = func() map[string]string {
	out := make(map[string]string)
	for name, netCfg := range NetworksByName {
		out[netCfg.L2ChainID.String()] = name
	}
	return out
}()

func AvailableNetworks() []string {
	var networks []string
	for name := range NetworksByName {
		networks = append(networks, name)
	}
	return networks
}

func GetRollupConfig(name string) (rollup.Config, error) {
	network, ok := NetworksByName[name]
	if !ok {
		return rollup.Config{}, fmt.Errorf("invalid network %s", name)
	}

	return network, nil
}

func u64Ptr(v uint64) *uint64 {
	return &v
}
