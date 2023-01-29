package balance

import (
	"context"
	"fmt"
	deposit "github.com/cahyacaa/stockbit-coinbit-test/internal/proto_models"
	proto "github.com/golang/protobuf/proto"
	"github.com/lovoo/goka"

	"log"
)

var (
	Deposits goka.Stream = "deposits"
	group    goka.Group  = "balance"
	Table    goka.Table  = goka.GroupTable(group)
)

type BalanceCodec struct{}

func (c *BalanceCodec) Encode(value interface{}) ([]byte, error) {
	var msg []byte
	var err error

	if byteMsg, ok := value.(deposit.Balance); ok {
		msg, err = proto.Marshal(&byteMsg)
		if err != nil {
			return msg, nil
		}
	}

	return msg, err
}

func (c *BalanceCodec) Decode(data []byte) (interface{}, error) {
	var balanceData deposit.Balance
	err := proto.Unmarshal(data, &balanceData)
	if err != nil {
		return []byte{}, err
	}

	return balanceData, nil
}

func balance(ctx goka.Context, msg interface{}) {
	existingBalance := deposit.Balance{}
	if v := ctx.Value(); v != nil {
		if existingData, ok := v.(deposit.Balance); ok {
			existingBalance = existingData
		}
	}

	newBalance, ok := msg.(deposit.Balance)
	if !ok {
		newBalance = deposit.Balance{}
	}

	newBalance.Balance = existingBalance.Balance + newBalance.Amount
	fmt.Println(newBalance, existingBalance)

	ctx.SetValue(newBalance)
}

func Run(ctx context.Context, brokers []string) func() error {
	tmc := goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	return func() error {
		g := goka.DefineGroup(group,
			goka.Input(Deposits, new(BalanceCodec), balance),
			goka.Persist(new(BalanceCodec)),
		)
		p, err := goka.NewProcessor(brokers, g, goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)))
		if err != nil {
			return err
		}

		log.Println("balance service running")
		return p.Run(ctx)
	}
}
