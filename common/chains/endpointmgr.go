package chains

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/web3eye-io/Web3Eye/common/ctredis"
)

type endpointIntervalMGR struct {
	RedisExpireTime time.Duration
}

type EndpointInterval struct {
	Address         string
	MinInterval     time.Duration
	BackoffIndex    int
	MaxBackoffIndex int
	MaxInterval     time.Duration
}

var _eIMGR *endpointIntervalMGR

const (
	keyExpireTime        = time.Minute * 5
	lockEndpointWaitTime = time.Millisecond * 100
	eIMGRPrefix          = "eIMGR"
	goaheadLockTime      = time.Minute
)

func GetEndpintIntervalMGR() *endpointIntervalMGR {
	if _eIMGR == nil {
		_eIMGR = &endpointIntervalMGR{RedisExpireTime: keyExpireTime}
	}
	return _eIMGR
}

func (eIMGR *endpointIntervalMGR) putEndpoint(item *EndpointInterval, autoResetBackoffIndex bool) error {
	if autoResetBackoffIndex {
		item.BackoffIndex = 0
		_maxBackoffIndex := math.Log2(float64(item.MaxInterval) / float64(item.MinInterval))
		item.MaxBackoffIndex = int(_maxBackoffIndex)
	}

	return ctredis.Set(eIMGR.getKey(item.Address), item, eIMGR.RedisExpireTime)
}

func (eIMGR *endpointIntervalMGR) GoAheadEndpoint(item *EndpointInterval) error {
	locked, err := ctredis.TryPubLock(eIMGR.getLockKey(item.Address), goaheadLockTime)
	if !locked || err != nil {
		return nil
	}

	_item := &EndpointInterval{}
	err = ctredis.Get(eIMGR.getKey(item.Address), _item)
	if err != nil {
		return eIMGR.putEndpoint(item, true)
	}

	if _item.BackoffIndex > 0 {
		_item.BackoffIndex--
	}

	return eIMGR.putEndpoint(_item, false)
}

func (eIMGR *endpointIntervalMGR) BackoffEndpoint(address string) error {
	item := &EndpointInterval{}
	err := ctredis.Get(eIMGR.getKey(address), item)
	if err != nil {
		return err
	}

	if item.BackoffIndex < item.MaxBackoffIndex {
		item.BackoffIndex++
	}

	return eIMGR.putEndpoint(item, false)
}

func (eIMGR *endpointIntervalMGR) GetEndpointInterval(address string) (time.Duration, error) {
	item := &EndpointInterval{}
	err := ctredis.Get(eIMGR.getKey(address), item)
	if err != nil {
		return 0, err
	}
	interval := item.MinInterval << item.BackoffIndex
	if interval > item.MaxInterval {
		return item.MaxInterval, nil
	}
	return interval, nil
}

func (eIMGR *endpointIntervalMGR) getKey(address string) string {
	return fmt.Sprintf("%v-%v", eIMGRPrefix, address)
}

func (eIMGR *endpointIntervalMGR) getLockKey(address string) string {
	return fmt.Sprintf("%v-lock-%v", eIMGRPrefix, address)
}

func (e *EndpointInterval) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(e)
	return data, err
}

func (e *EndpointInterval) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}

func LockEndpoint(ctx context.Context, keys []string, lockTimes uint16) (string, error) {
	for {
		select {
		case <-time.NewTicker(lockEndpointWaitTime).C:
			_randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(keys))))
			if err != nil {
				return "", err
			}
			randIndex := int(_randIndex.Int64())
			for j := 0; j < len(keys); j++ {
				lockKey := keys[(randIndex+j)%len(keys)]
				interval, err := GetEndpintIntervalMGR().GetEndpointInterval(lockKey)
				if err != nil {
					fmt.Println(err)
					continue
				}
				locked, _ := ctredis.TryPubLock(lockKey, interval*time.Duration(lockTimes))
				if locked {
					return lockKey, nil
				}
			}
		case <-ctx.Done():
			return "", nil
		}
	}
}

// func test() {
// 	if err := logger.Init(logger.DebugLevel, "./a.log"); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	if err := db.Init(); err != nil {
// 		logger.Sugar().Fatalln(err)
// 	}
// 	keys := []string{"https://mainnet.infura.io/v3/00000000000", "s2", "s3"}
// 	eIMGR := NewNFTConllectionMGR()
// 	for _, key := range keys {
// 		eIMGR.PutEndpoint(&EndpointInterval{
// 			Address:     key,
// 			MinInterval: time.Second,
// 			MaxInterval: time.Minute * 5,
// 		}, true)
// 	}

// 	for i := 0; i < 3; i++ {
// 		go func(i int) {
// 			for j := 0; j < 100; j++ {
// 				ret, err := LockEndpoint(keys)
// 				fmt.Printf("%v,%v,%v,%v\n", i, j, ret, err)
// 				if j%5 == 0 {
// 					err = eIMGR.BackoffEndpoint(ret)
// 					fmt.Printf("backoff,%v,%v\n", ret, err)
// 				}
// 			}
// 		}(i)
// 	}
// 	time.Sleep(time.Minute)
// }
