package search_engine

import (
	"encoding/json"

	"github.com/JUNAID-KT/eWallet/models"
	"github.com/JUNAID-KT/eWallet/util"
	log "github.com/Sirupsen/logrus"
	"github.com/olivere/elastic"
)

func (es *esEngine) SaveTransactions(doc models.Transaction) error {
	_, err := es.Client.
		Index().
		Index(util.TransactionIndexName).
		Type(util.TransactionTypeName).
		BodyJson(doc).
		Do(es.Ctx)

	if err != nil {
		log.WithFields(log.Fields{"method": "SaveTransactions", "Index Name": util.TransactionIndexName,
			"error": err.Error()}).
			Error("error occurred while saving AWS credentials")
		return err
	}

	return nil
}
func (es *esEngine) GetTransactions(user string) (error, []models.Transaction) {
	deleted := elastic.NewMatchQuery("From", user)
	generalQuery := elastic.NewBoolQuery().MustNot(deleted)
	var transactions []models.Transaction
	searchResult, err := es.Client.Search().
		Index(util.TransactionIndexName).
		Type(util.TransactionTypeName).
		Query(generalQuery).
		Size(util.SearchLimit).
		Do(es.Ctx)

	if err != nil {
		log.WithFields(log.Fields{"error": err.Error(),
			"Index Name": util.TransactionIndexName}).Error("error occurred in fetching AWS credentials")
		return err, transactions
	}
	if searchResult != nil && searchResult.Hits != nil && searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			data := *hit.Source
			jsonData, jsonErr := data.MarshalJSON()
			if jsonErr != nil {
				log.WithFields(log.Fields{"error": jsonErr.Error(),
					"Index Name": util.TransactionIndexName}).
					Errorln("error occurred in marshalling AWS credentials data")
				return jsonErr, transactions
			}

			var result models.Transaction
			parseError := json.Unmarshal(jsonData, &result)
			if parseError != nil {
				log.WithFields(log.Fields{"error": parseError.Error(),
					"Index Name": util.TransactionIndexName}).
					Errorln("error occurred in unmarshalling AWS credentials data")
				return parseError, transactions
			}

			transactions = append(transactions, result)
		}
		log.Infoln("AWS account information fetched successfully.")
	}
	return nil, transactions

}
