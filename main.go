package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/kafka-mas/net-checker/alerts"
	"github.com/kafka-mas/net-checker/networkchecker"
	"github.com/kafka-mas/net-checker/readconf"
)

var sleepTime time.Duration = 1
var confPath string = "config.yaml"

type NetworkCheck interface {
	Ping(addresses []string) ([]string, error)
}

type Config interface {
	ConfigRead(file string) error
	ConfigGetPhone() string
	ConfigGetAddresses() []string
}

func main() {
	phone, adderesses, err := ParseConf(confPath)
	if err != nil {
		panic("Error reading config")
	}
	ticker := time.NewTicker(sleepTime * time.Minute)
	defer ticker.Stop()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	RunTask(phone, adderesses)

	for {
		select {
		case <-ticker.C:
			RunTask(phone, adderesses)
		case <-ctx.Done():
			var user alerts.User = alerts.User(phone)
			err = user.SendSMS("\nProgram stopped by SIGTERM")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Завершение работы")
			return
		}
	}
}

func ParseConf(confPath string) (string, []string, error) {
	var config Config = &readconf.YamlConf{}
	err := config.ConfigRead(confPath)
	if err != nil {
		fmt.Println(err)
	}

	phone := config.ConfigGetPhone()
	addresses := config.ConfigGetAddresses()

	return phone, addresses, err
}

func RunTask(phone string, addresses []string) {

	var checker NetworkCheck = networkchecker.ICMPChecker{}

	ipUnavailable, err := checker.Ping(addresses)
	if err != nil {
		fmt.Println(err)
	}

	var user alerts.User = alerts.User(phone)

	if len(ipUnavailable) != 0 {
		addrList := strings.Join(ipUnavailable, "; ")
		msg := fmt.Sprintf("Недоступны адреса: %v", addrList)

		err = user.SendSMS(msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
