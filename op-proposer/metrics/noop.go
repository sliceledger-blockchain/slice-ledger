package metrics

import (
	"github.com/sliceledger-blockchain/slice-ledger/op-node/eth"
	opmetrics "github.com/sliceledger-blockchain/slice-ledger/op-service/metrics"
	txmetrics "github.com/sliceledger-blockchain/slice-ledger/op-service/txmgr/metrics"
)

type noopMetrics struct {
	opmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
}

var NoopMetrics Metricer = new(noopMetrics)

func (*noopMetrics) RecordInfo(version string) {}
func (*noopMetrics) RecordUp()                 {}

func (*noopMetrics) RecordL2BlocksProposed(l2ref eth.L2BlockRef) {}
