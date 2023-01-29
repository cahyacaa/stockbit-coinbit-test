package balance

import (
	"context"
	"encoding/json"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/model"
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

	if byteMsg, ok := value.(model.Balance); ok {
		msg, err = json.Marshal(&byteMsg)
		if err != nil {
			return msg, nil
		}
	}

	return msg, err
}

func (c *BalanceCodec) Decode(data []byte) (interface{}, error) {
	var internalData model.Balance

	err := json.Unmarshal(data, &internalData)
	if err != nil {
		return []byte{}, err
	}

	return internalData, nil
}

func balance(ctx goka.Context, msg interface{}) {
	existingBalance := model.Balance{}
	if v := ctx.Value(); v != nil {
		if existingData, ok := v.(model.Balance); ok {
			existingBalance = existingData
		}
	}

	newBalance, ok := msg.(model.Balance)
	if !ok {
		newBalance = model.Balance{}
	}

	newBalance.Balance = existingBalance.Balance + newBalance.Amount
	newBalance.UpdateVersion = existingBalance.UpdateVersion + 1

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
