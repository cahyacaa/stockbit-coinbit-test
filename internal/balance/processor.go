package above_threshold

import (
	"context"
	"fmt"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/topic_init"
	wallet "github.com/cahyacaa/stockbit-coinbit-test/model"
	"github.com/golang/protobuf/proto"
	"github.com/lovoo/goka"
	"log"
)

var (
	Deposits goka.Stream = "deposits"
	group    goka.Group  = "balance"
	Table    goka.Table  = goka.GroupTable(group)
)

type WalletCodec struct{}

func (c *WalletCodec) Encode(value interface{}) ([]byte, error) {
	byteMsg := value.(wallet.Wallet)
	msg, err := proto.Marshal(&byteMsg)
	return msg, err
}

func (c *WalletCodec) Decode(data []byte) (interface{}, error) {
	var m wallet.Wallet
	err := proto.Unmarshal(data, &m)
	return m, err
}

func collect(ctx goka.Context, msg interface{}) {
	wl := wallet.Wallet{}
	if v := ctx.Value(); v != nil {
		wl = v.(wallet.Wallet)
	}

	m := msg.(wallet.Wallet)
	//m.Balance += wl.Balance
	fmt.Println(wl)
	ctx.SetValue(m)
}

func PrepareTopics(brokers []string) {
	topic_init.EnsureStreamExists(string(Deposits), brokers)
}

func Run(ctx context.Context, brokers []string) func() error {
	tmc := goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	return func() error {
		g := goka.DefineGroup(group,
			goka.Input(Deposits, new(WalletCodec), collect),
			goka.Persist(new(WalletCodec)),
		)
		p, err := goka.NewProcessor(brokers, g, goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)))
		if err != nil {
			return err
		}

		log.Println("kafka instance running")
		return p.Run(ctx)
	}
}
