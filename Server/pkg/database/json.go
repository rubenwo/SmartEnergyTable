package database

import (
	"encoding/json"
	"log"
	"os"
)

type job struct {
	key string
}

type jsonDb struct {
	queue     chan job
	documents map[string]interface{}
}

func createJSONDb() *jsonDb {
	db := &jsonDb{
		queue:     make(chan job),
		documents: make(map[string]interface{}),
	}
	go db.startPersistJob()
	return db
}

//Set: Implementation of the database interface
func (jdb *jsonDb) Set(key string, value interface{}) error {
	log.Println("(JSON DB): saving data:", value, " for key:", key)
	return nil
}

//Get: Implementation of the database interface
func (jdb *jsonDb) Get(key string) (interface{}, error) {
	log.Println("(JSON DB): retrieving value from key:", key)
	return nil, nil
}

//Observe: Implementation of database interface
func (jdb *jsonDb) Observe(key string) (chan interface{}, error) {
	c := make(chan interface{})

	return c, nil
}

//Delete: Implementation of the database interface
func (jdb *jsonDb) Delete(key string) error {
	log.Println("(JSON DB): deleting value from key:", key)

	return nil
}

func (jdb *jsonDb) startPersistJob() {
	for {
		j := <-jdb.queue
		var data struct {
			Data interface{} `json:"data"`
		}
		var f *os.File
		if _, err := os.Stat(j.key + ".json"); os.IsNotExist(err) {
			f, err = os.Create(j.key + ".json")
			if err != nil {
				log.Println(err)
			}
		} else {
			f, err = os.Open(j.key + ".json")
			if err != nil {
				log.Println(err)
			}
		}

		if err := json.NewEncoder(f).Encode(&data); err != nil {
			log.Println(err)
		}

		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}
}
