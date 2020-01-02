package search_engine

import (
	log "github.com/Sirupsen/logrus"
)

const (
	createIndex   = "CreateIndex"
	createMapping = "CreateIndexMapping"
)

//CreateIndex: Checks if the index is already present,if not creates the index with the given metadata
func (es *esEngine) CreateIndex(indxName string, indexType string, mapping string) {
	//check if the Index exists in ElasticSearch
	exists, err := es.Client.IndexExists(indxName).Do(es.Ctx)
	if err != nil {
		log.WithFields(log.Fields{"method": createIndex, "error": err.Error(), "Index Name": indxName}).
			Errorln("error in checking if index already exists")
		return
	}
	if !exists {
		// Create a new index with the given mapping.
		_, err := es.Client.CreateIndex(indxName).BodyString(mapping).Do(es.Ctx)
		if err != nil {
			log.WithFields(log.Fields{"method": createIndex, "error": err.Error(), "Index Name": indxName}).
				Errorln("error creating index")
			return
		}
		log.WithFields(log.Fields{"method": createIndex, "Index Name": indxName}).
			Infoln("index created successfully")
	}
}

/*CreateIndexMapping : Checks if the index is present in ES and if index is already there
checks the count of total docs present and apply the mapping when the
document count in the index is 0*/
func (es *esEngine) CreateIndexMapping(indxName string, indexType string, mapping string) {
	exists, err := es.Client.IndexExists(indxName).Do(es.Ctx)
	if err != nil {
		log.WithFields(log.Fields{"method": createMapping, "error": err.Error(), "Index Name": indxName}).
			Errorln("error in checking if index already exists")
		return
	}
	docCount, countErr := es.Client.Count(indxName).Type(indexType).Do(es.Ctx)
	if countErr != nil {
		log.WithFields(log.Fields{"method": createMapping, "error": countErr.Error(), "Index Name": indxName}).
			Errorln("error finding total doc count in the index")
		return
	}
	if exists && docCount == 0 {
		_, mapErr := es.Client.PutMapping().Index(indxName).BodyString(mapping).Type(indexType).Do(es.Ctx)
		if mapErr != nil {
			log.WithFields(log.Fields{"method": createMapping, "error": mapErr.Error(), "Index Name": indxName}).
				Errorln("error in mapping creation")
			return
		}
		log.WithFields(log.Fields{"method": createMapping, "Index Name": indxName}).
			Infoln("mapping created successfully")
	}
}
