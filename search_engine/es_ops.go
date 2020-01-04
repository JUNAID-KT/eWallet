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
			Error("error occurred while saving transactions")
		return err
	}

	return nil
}

// Search in DB ; matching a given term, user
func (es *esEngine) GetTransactions(user string) (error, []models.Transaction) {
	// Search with a term query
	var transactions []models.Transaction
	termQuery := elastic.NewTermQuery("from.keyword", user)
	searchResult, err := es.Client.Search().
		Index(util.TransactionIndexName).
		Query(termQuery).
		Do(es.Ctx)
	if err != nil {
		return err, transactions
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index
			var transaction models.Transaction
			// Deserialize hit.Source into a Transaction
			err := json.Unmarshal(*hit.Source, &transaction)
			if err != nil {
				// Deserialization failed
				return err, transactions
			}
			transactions = append(transactions, transaction)
		}
	} else {
		// No hits
		return nil, transactions
	}
	return nil, transactions
}
