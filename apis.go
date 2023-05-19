package main

import (
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	log "unsplash-downloader/pkg/qlogger"
)

func GetEditorialPhotos(key string, page int) (Photos, error) {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	randomApi := "https://api.unsplash.com/photos"

	req, _ := http.NewRequest(http.MethodGet, randomApi, nil)

	req.Header.Add("Authorization", "Client-ID "+key)
	q := req.URL.Query()
	q.Add("per_page", "30")
	q.Add("page", strconv.Itoa(page))
	q.Add("order_by", "oldest")
	req.URL.RawQuery = q.Encode()

	log.Println(req.URL.String(), " , query: ", req.URL.Query())

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	log.Println("Response: ", res)

	if res.StatusCode == 200 {
		// Parse links
		body, _ := ioutil.ReadAll(res.Body)
		// log.Printf("%#v", string(body))
		photos, err := UnmarshalPhotos(body)
		log.Println("photos: ", len(photos), "err: ", err)

		if err != nil {
			return nil, err
		}

		return photos, nil
	} else {
		err = errors.New(res.Status)
		log.Println("Err: ", err)
		return nil, err
	}
}

func GetTopicsPhotos(key string, topicId string, page int) (Photos, error) {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	randomApi := "https://api.unsplash.com/topics/" + topicId + "/photos"

	req, _ := http.NewRequest(http.MethodGet, randomApi, nil)

	req.Header.Add("Authorization", "Client-ID "+key)
	q := req.URL.Query()
	q.Add("per_page", "30")
	q.Add("page", strconv.Itoa(page))
	q.Add("order_by", "oldest")
	req.URL.RawQuery = q.Encode()

	log.Println(req.URL.String(), " , query: ", req.URL.Query())

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	log.Println("Response: ", res)

	if res.StatusCode == 200 {
		// Parse links
		body, _ := ioutil.ReadAll(res.Body)
		// log.Printf("%#v", string(body))
		photos, err := UnmarshalPhotos(body)
		log.Println("photos: ", len(photos), "err: ", err)

		if err != nil {
			return nil, err
		}

		return photos, nil
	} else {
		err = errors.New(res.Status)
		log.Println("Err: ", err)
		return nil, err
	}
}

func GetRandomPhoto(key string) (Photos, error) {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	randomApi := "https://api.unsplash.com/photos/random?count=30&client_id=" + AccessKey

	// req, err := http.NewRequest(http.MethodGet, randomApi, nil)

	// req.Header.Add("Authorization", "Client-ID "+key)
	// req.URL.Query().Add("count", "30")

	// log.Println(req.URL.String())

	// res, err := http.DefaultClient.Do(req)
	res, err := client.Get(randomApi)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	log.Println("Response: ", res)

	if res.StatusCode == 200 {
		// Parse links
		body, _ := ioutil.ReadAll(res.Body)
		// log.Printf("%#v", string(body))
		photos, err := UnmarshalPhotos(body)
		log.Println("photos: ", len(photos), "err: ", err)

		if err != nil {
			return nil, err
		}

		return photos, nil
		// for _, photo := range photos {
		// 	log.Println("id: ", *photo.ID, "url: ", *photo.Urls.Full)
		// 	DownloadFile(*photo.Urls.Full, *photo.ID)
		// }
	} else {
		err = errors.New(res.Status)
		log.Println("Err: ", err)
		return nil, err
	}
}

func DownloadFile(URL, fileName string) error {
	name := fileName + ".jpeg"
	_, err := os.Stat("/sharedFolder/photos/" + name)

	if err == nil {
		// File exist
		log.Println("Photo ", name, " already downloaded.")
		return nil
	}

	downloadTokens <- struct{}{}
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	//Get the response bytes from the url
	response, err := client.Get(URL)
	if err != nil {
		log.Println(fileName, " request error: ", err)
		return err
	}
	defer response.Body.Close()
	<-downloadTokens

	if response.StatusCode != 200 {
		log.Println(fileName, " Received non 200 response code: ",
			response.StatusCode)
		return errors.New("Received non 200 response code")
	}

	log.Println("Downloaded: ", fileName)
	//Create a empty file
	// name := strconv.Itoa(int(time.Now().Unix())) + "_" + fileName + ".jpeg"
	file, err := os.Create("/sharedFolder/photos/" + name)
	// file, err := os.Create("./" + name)
	if err != nil {
		log.Println("Fail create file ", fileName)
		return err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Println("Fail write file ", fileName)
		return err
	}

	return nil
}
