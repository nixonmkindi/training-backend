package filter

import (
	"fmt"
	"sync"
	"time"
)

//Filter structure
type Filter struct {
	truncatePeriod int64            //period that determine the period within which message cannot be gossipped
	lastTruncate   int64            //records the last time the message was trancated
	msgRecord      map[string]int64 //message record that keeps the hash and time the message was received
	lock           *sync.Mutex
}

//New message filter
func New(truncatePeriod int64) *Filter {
	current := time.Now().Unix()
	return &Filter{truncatePeriod, current, make(map[string]int64), &sync.Mutex{}}
}

//Check function, which checks if the message exists
func (filter *Filter) Check(msgHash string) bool {
	current := time.Now().Unix()
	filter.lock.Lock()
	defer filter.lock.Unlock()
	if recvTime, ok := filter.msgRecord[msgHash]; ok {
		if current-recvTime < filter.truncatePeriod {
			return false
		} else {
			delete(filter.msgRecord, msgHash)
		}
	}
	filter.msgRecord[msgHash] = current
	if current-filter.lastTruncate > filter.truncatePeriod {
		for k, v := range filter.msgRecord {
			if current-v > filter.truncatePeriod {
				delete(filter.msgRecord, k)
			}
		}
		filter.lastTruncate = current
	}
	//for debuging
	//filter.print()
	return true
}

func (filter *Filter) print() {
	fmt.Println()
	for k, v := range filter.msgRecord {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
