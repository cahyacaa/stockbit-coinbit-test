package above_threshold

import (
	"context"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/proto/proto_models"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/topic_init"
	proto "github.com/golang/protobuf/proto"
	"github.com/lovoo/goka"
	"google.golang.org/protobuf/types/known/timestamppb"

	"log"
	"time"
)

var (
	Deposits   goka.Stream   = "deposits"
	group      goka.Group    = "above_threshold"
	Table                    = goka.GroupTable(group)
	timeWindow time.Duration = time.Minute * 2
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

	//assign empty data in newWalletData struct with existing data
	newWalletData.TimeWindowBalance = existingWalletData.TimeWindowBalance + newWalletData.Amount
	newWalletData.TimeExpired = existingWalletData.TimeExpired

	//if deposit from new wallet id, assign time window expired
	if existingWalletData.WalletID == "" || existingWalletData.TimeExpired.AsTime().Local().IsZero() {
		newWalletData.TimeExpired = timestamppb.New(time.Now().Local().Add(timeWindow))
	} else {
		if existingWalletData.TimeExpired.AsTime().Local().After(time.Now()) && newWalletData.TimeWindowBalance >= 10000 {
			newWalletData.IsAboveThreshold = true
		}

		//assign new expired time if time window is over
		if existingWalletData.TimeExpired.AsTime().Local().Before(time.Now()) {
			newWalletData.TimeWindowBalance = newWalletData.Amount
			newWalletData.TimeExpired = timestamppb.New(time.Now().Local().Add(timeWindow))
		}
	}

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

		log.Println("deposit flagger is running")
		return p.Run(ctx)
	}
}
