package above_threshold

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/model"
	"github.com/cahyacaa/stockbit-coinbit-test/internal/topic_init"
	"github.com/lovoo/goka"
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

	if byteMsg, ok := value.(model.AboveThreshold); ok {
		msg, err = json.Marshal(&byteMsg)
		if err != nil {
			return msg, nil
		}
	}

	return msg, err
}

func (c *AboveThresholdCodec) Decode(data []byte) (interface{}, error) {
	//var m wallet.Wallet
	var internalData model.AboveThreshold
	//if err := proto.Unmarshal(data, &m); err != nil {
	err := json.Unmarshal(data, &internalData)
	if err != nil {
		return []byte{}, err
	}
	//}

	//output := model.AboveThreshold{
	//	WalletID: m.WalletId,
	//	Amount:   m.Amount,
	//}
	return internalData, nil
}
func Flagger(ctx goka.Context, msg interface{}) {
	var existingWalletData, newWalletData model.AboveThreshold
	if v := ctx.Value(); v != nil {
		existingWalletData = v.(model.AboveThreshold)
	}
	if messageData, ok := msg.(model.AboveThreshold); ok {
		newWalletData = messageData
	}

	newWalletData.TimeWindowBalance = existingWalletData.TimeWindowBalance + newWalletData.Amount
	newWalletData.TimeExpired = existingWalletData.TimeExpired

	if existingWalletData.WalletID == "" || existingWalletData.TimeExpired.IsZero() {
		newWalletData.TimeExpired = time.Now().Add(10 * time.Minute)
	} else {
		if existingWalletData.TimeExpired.After(time.Now()) && newWalletData.TimeWindowBalance >= 10000 {
			newWalletData.IsAboveThreshold = true
		}

		if existingWalletData.TimeExpired.Before(time.Now()) {
			newWalletData.TimeWindowBalance = newWalletData.Amount
			newWalletData.TimeExpired = existingWalletData.TimeExpired.Add(10 * time.Minute)
		}
	}

	newWalletData.UpdateVersion = existingWalletData.UpdateVersion + 1
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
