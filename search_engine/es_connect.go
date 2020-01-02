package search_engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/JUNAID-KT/eWallet/app"
	"github.com/JUNAID-KT/eWallet/util"

	log "github.com/Sirupsen/logrus"
	"github.com/olivere/elastic"
	"gopkg.in/go-playground/validator.v9"
)

type esEngine struct {
	Client   *elastic.Client
	Ctx      context.Context
	Validate *validator.Validate
}

var (
	instance *esEngine
	once     sync.Once
)

func GetESInstance() *esEngine {
	var err error
	// Singleton instance - included retries when an error occurs in creating elastic client
	once.Do(func() {
		for i := 1; i <= util.MaxRetries; i++ {
			instance, err = createInstance()
			if err != nil && i != util.MaxRetries {
				time.Sleep(time.Minute)
			} else {
				log.Info("Connection established with Elasticsearch")
				return
			}
		}
		if err != nil {
			log.Errorln("retries failed")
		}
	})
	return instance
}

func (es *esEngine) Stop() {
	log.WithFields(log.Fields{"method": "Stop"}).Infoln("stopping elastic client")
	es.Client.Stop()
}

func createInstance() (*esEngine, error) {
	elasticAddress := app.Config.ElasticAddress
	fmt.Println(elasticAddress)
	log.WithFields(log.Fields{"method": "createInstance", "elasticURL": elasticAddress}).
		Infoln("starting elastic client")
	instance := new(esEngine)

	esClient, err := elastic.NewClient(
		// url
		elastic.SetURL(elasticAddress),
		// back off and retry
		elastic.SetRetrier(CustomRetrier()),
		// default 60s
		elastic.SetHealthcheckInterval(30*time.Second),
		//A client uses a sniffing process to find all nodes of your cluster by default, automatically. For one node, not required.
		elastic.SetSniff(false))
	if err != nil {
		log.WithFields(log.Fields{"method": "createInstance", "error": err.Error(),
			"description": "unable to connect to elastic"}).Errorln("connection error")
		return nil, err
	}
	instance.Client = esClient
	instance.Ctx = context.Background()
	instance.Validate = validator.New()

	return instance, nil
}
