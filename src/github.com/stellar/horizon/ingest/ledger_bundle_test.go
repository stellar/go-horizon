package ingest

import (
	"github.com/stellar/horizon/db"
	"github.com/stellar/horizon/test"
	"testing"
)

func TestTransactionFeeByHash(t *testing.T) {
	tt := test.Start(t).Scenario("base")
	defer tt.Finish()

	bundle := &LedgerBundle{Sequence: 2}
	err := bundle.Load(db.SqlQuery{tt.CoreDB})

	if tt.Assert.NoError(err) {
		tt.Assert.Equal(uint32(2), bundle.Header.Sequence)
		tt.Assert.Len(bundle.Transactions, 3)
		tt.Assert.Len(bundle.TransactionFees, 3)
	}
}