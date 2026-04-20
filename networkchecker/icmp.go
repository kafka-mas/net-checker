package networkchecker

import (
	"fmt"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

var maxWorkers int8 = 5
var pingCount int = 4
var pingTimeout int8 = 5

type ICMPChecker struct{}

func (ICMPChecker) Ping(address []string) ([]string, error) {
	sem := make(chan struct{}, maxWorkers)

	result := []string{}

	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, i := range address {
		wg.Add(1)
		go func(addr string) error {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			pinger, err := probing.NewPinger(addr)
			if err != nil {
				fmt.Println("Error parse IPs", addr)
				return err
			}
			pinger.Count = pingCount
			pinger.Timeout = time.Duration(pingTimeout) * time.Second
			pinger.Run() // блокирует до завершения

			stats := pinger.Statistics()
			if stats.PacketLoss > 70 {
				mu.Lock()
				result = append(result, pinger.Addr())
				mu.Unlock()
			}
			return nil
		}(i)
	}

	wg.Wait()
	return result, nil
}
