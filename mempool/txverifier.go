package mempool

import (
	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo/message"
	"github.com/aergoio/aergo/types"
	"github.com/golang/protobuf/proto"
)

type TxVerifier struct {
	mp *MemPool
}

func NewTxVerifier(p *MemPool) *TxVerifier {
	return &TxVerifier{mp: p}
}

//Receive actor message
func (s *TxVerifier) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *message.MemPoolPut:
		tx := msg.Tx
		var err error
		if proto.Size(tx) > txMaxSize {
			err = types.ErrTxSizeExceedLimit
		} else if s.mp.exist(tx.GetHash()) != nil {
			err = types.ErrTxAlreadyInMempool
		} else {
			tx := types.NewTransaction(tx)
			err = s.mp.verifyTx(tx)
			if err == nil {
				err = s.mp.put(tx)
			}
		}
		context.Respond(&message.MemPoolPutRsp{Err: err, Sender:msg.Sender})
	}
}
