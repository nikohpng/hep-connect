package env

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
	UDPListenPort         int    `env:"UDPListenPort" envDefault:"9060"`
	MaxPackgeLength       int    `env:"MaxPackgeLength" envDefault:"4096"`
	MaxReadTimeoutSeconds int    `env:"MaxReadTimeoutSecond" envDefault:"5"`
	LogLevel              string `env:"LogLevel" envDefault:"info"`
	Hostname              string `env:"HOSTNAME" envDefault:"unknow"`
	HeaderUIDName         string `env:"HeaderUIDName"`
	HeaderFSCallIDName    string `env:"HeaderFSCallIDName"`
	DiscardMethods        string `env:"DiscardMethods" envDefault:"OPTIONS"`
	MinPackgeLength       int    `env:"MinPackgeLength" envDefault:"24"`
	SqlMaxOpenConn        int    `env:"SqlMaxOpenConn" envDefault:"64"`
	SqlMaxIdleConn        int    `env:"SqlMaxIdleConn" envDefault:"64"`
	DBUserPasswd          string `env:"DBUserPasswd" envDefault:"root:123456"`
	DBAddr                string `env:"DBAddr" envDefault:"localhost"`
	DBName                string `env:"DBName" envDefault:"siphub"`
	CalleeFrom            string `env:"CalleeFrom" envDefault:"RURI"`
	MaxBatchItems         int    `env:"MaxBatchItems" envDefault:"20"`
	TickerSecondTime      int    `env:"TickerSecondTime" envDefault:"20"`
}

var Conf = config{}

func init() {
	if err := env.Parse(&Conf); err != nil {
		log.Fatalf("%+v\n", err)
	} else {
		log.Printf("%#v\n", Conf)
	}

	if Conf.TickerSecondTime < 20 {
		log.Panicf("TickerSencondTime must bigger than 20. TickerSecondTime: %d", Conf.TickerSecondTime)
	}
}
