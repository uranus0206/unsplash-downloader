package main

import (
	"flag"
	"math/rand"
	"sync"
	"time"
	"unsplash-downloader/pkg/dbmanager"
	log "unsplash-downloader/pkg/qlogger"
)

var AccessKey string
var wg sync.WaitGroup
var downloadTokens = make(chan struct{}, 10)

var topics dbmanager.DbTopics
var editorial dbmanager.DbEditorial

var downloadByTopics bool

var topicsKeys = []string{
	"Mr49Evh5Sks", "bo8jQKTaE0Y", "CDwuwXJAbEw",
	"6sMVjTLSkeQ", "Fzo3zuOHN6w", "M8jVbLbTRws",
	"xHxYTMHLgOc", "iUIsnVtjB0Y", "hmenvQhUmxM",
	"qPYsDzvJOYc", "Jpg6Kidl-Hk", "S4MKLAsBB74",
	"aeu6rL-j6ew", "xjPR4hlkBGA", "towJZFskpGg",
	"_8zFHuhRhyo", "Bn-DjrcBrwo", "_hb-dl4Q-4U",
	"BJJMtteDJA4", "bDo48cUhwnY",
}

type PhotoLink struct {
	name string
	url  string
}

func main() {
	flag.StringVar(&AccessKey, "c", "", "Client Access Key.")
	flag.Parse()

	if AccessKey == "" {
		log.Panicln("Missing access key.")
	}

	dbmanager.InitWithPath("unsplash.db").CreateTable()

	t, err := dbmanager.Dbm.SearchTopicsByKey()
	if err != nil {
		t := &dbmanager.DbTopics{
			PRKey:      "DbTopicKey",
			TopicsId:   topicsKeys[0],
			TopicsIndx: 0,
			PageOffset: 1,
		}
		dbmanager.Dbm.AddTopicRecord(t)
		topics = *t
	} else {
		topics = t
	}

	if int(topics.TopicsIndx) < len(topicsKeys) {
		downloadByTopics = true
	} else {
		downloadByTopics = false
	}

	e, err := dbmanager.Dbm.SearchEditorialByKey()
	if err != nil {
		e := &dbmanager.DbEditorial{
			PRKey:      dbmanager.DbEditorialKey,
			PageOffset: 1,
		}
		dbmanager.Dbm.AddEditorialRecord(e)
		editorial = *e
	} else {
		editorial = e
	}

	for {
		downloaded := make(chan error)

		current := time.Now()
		var photos Photos

		if downloadByTopics {
			// Crawl Topics photos
			photos, err = GetTopicsPhotos(AccessKey,
				topics.TopicsId,
				int(topics.PageOffset))

			if err != nil {
				time.Sleep(30 * time.Second)
				continue
			}

			if len(photos) == 0 {
				// Next topics
				topics.PageOffset = 1
				topics.TopicsIndx += 1
				// avoid out of range index.
				if int(topics.TopicsIndx) < len(topicsKeys) {
					topics.TopicsId = topicsKeys[topics.TopicsIndx]
					downloadByTopics = true
				} else {
					topics.TopicsId = "All topics crawled !!"
					downloadByTopics = false
				}
				log.Println("Topics ", topics, ", next page: ", topics.PageOffset)
				dbmanager.Dbm.AddTopicRecord(&topics)
				time.Sleep(72 * time.Second)
				continue
			}
		} else {
			// Crawl Editorial photos
			photos, err = GetEditorialPhotos(AccessKey,
				int(editorial.PageOffset))

			if err != nil {
				time.Sleep(30 * time.Second)
				continue
			}

			if len(photos) == 0 {
				// Wait new photos.
				time.Sleep(72 * time.Second)
				continue
			}
		}

		for _, photo := range photos {
			wg.Add(1)
			log.Println("id: ", *photo.ID, "url: ", *photo.Urls.Full)

			go func(p Photo) {
				defer wg.Done()
				error := DownloadFile(*p.Urls.Full, *p.ID)
				downloaded <- error
			}(photo)
		}

		go func() {
			log.Println("Will close ch.")
			wg.Wait()
			close(downloaded)
		}()

		hasErr := false
		for anyErr := range downloaded {
			if anyErr != nil {
				hasErr = true
			}
		}

		elapse := int32(time.Since(current).Seconds())

		log.Println("Time used: ", elapse, " seconds.")

		// Next page query only for all downloads success.
		if !hasErr {
			if downloadByTopics {
				topics.PageOffset += 1
				log.Println("Topics ", topics, ", next page: ", topics.PageOffset)
				dbmanager.Dbm.AddTopicRecord(&topics)
			} else {
				editorial.PageOffset += 1
				log.Println("Editorial next page: ", editorial.PageOffset)
				dbmanager.Dbm.AddEditorialRecord(&editorial)
			}
		}

		diff := 72 - elapse
		offset := int32(0)
		if diff <= 0 {
			offset = 72
		} else if diff > 72 {
			offset = 72
		} else {
			offset = diff
		}
		delay := offset + rand.Int31n(10)
		log.Println("Sleep: ", delay, " seconds.")
		time.Sleep(time.Duration(delay) * time.Second)
	}
}
