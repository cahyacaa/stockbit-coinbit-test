package above_threshold

import (
	"context"
	"fmt"
	deposit "github.com/cahyacaa/stockbit-coinbit-test/internal/proto_models"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/topic_init"
	proto "github.com/golang/protobuf/proto"
	"github.com/lovoo/goka"
	"google.golang.org/protobuf/types/known/timestamppb"

	"log"
	"time"
)

var (
	Deposits goka.Stream = "deposits"
	group    goka.Group  = "above_threshold"
	Table                = goka.GroupTable(group)
)

type AboveThresholdCodec struct{}

func (c *AboveThresholdCodec) Encode(value interface{}) ([]byte, error) {
	var msg []byte
	var err error

	if byteMsg, ok := value.(deposit.DepositFlagger); ok {
		msg, err = proto.Marshal(&byteMsg)
		if err != nil {
			return msg, nil
		}
	}

	return msg, err
}

func (c *AboveThresholdCodec) Decode(data []byte) (interface{}, error) {
	var m deposit.DepositFlagger

	if err := proto.Unmarshal(data, &m); err != nil {
		return m, err
	}
	return m, nil
}

func Flagger(ctx goka.Context, msg interface{}) {
	var existingWalletData, newWalletData deposit.DepositFlagger
	if v := ctx.Value(); v != nil {
		existingWalletData = v.(deposit.DepositFlagger)
	}
	if messageData, ok := msg.(deposit.DepositFlagger); ok {
		newWalletData = messageData
	}

	newWalletData.TimeWindowBalance = existingWalletData.TimeWindowBalance + newWalletData.Amount
	newWalletData.TimeExpired = existingWalletData.TimeExpired

	if existingWalletData.WalletID == "" || existingWalletData.TimeExpired.AsTime().IsZero() {
		newWalletData.TimeExpired = timestamppb.New(time.Now().Add(2 * time.Minute))
	} else {
		if existingWalletData.TimeExpired.AsTime().After(time.Now()) && newWalletData.TimeWindowBalance >= 10000 {
			newWalletData.IsAboveThreshold = true
		}

		if existingWalletData.TimeExpired.AsTime().Before(time.Now()) {
			newWalletData.TimeWindowBalance = newWalletData.Amount
			newWalletData.TimeExpired = timestamppb.New(existingWalletData.TimeExpired.AsTime().Add(2 * time.Minute))
		}
	}

	fmt.Println(newWalletData, existingWalletData)
	ctx.SetValue(newWalletData)
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
			goka.Input(Deposits, new(AboveThresholdCodec), Flagger),
			goka.Persist(new(AboveThresholdCodec)),
		)
		p, err := goka.NewProcessor(brokers, g, goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)))
		if err != nil {
			log.Println(err)
		}

		log.Println("kafka instance running")
		return p.Run(ctx)
	}
}
