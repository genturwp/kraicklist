package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"challenge.haraj.com.sa/kraicklist/entities"
	"challenge.haraj.com.sa/kraicklist/repositories"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
}

func dbConn() *pgxpool.Pool {
	//postgresql://kraicklist:rahasia@localhost:5434/kraicklist
	dbConnStr := viper.GetString("DB_CONN")
	dbConfig, err := pgxpool.ParseConfig(dbConnStr)
	if err != nil {
		log.Fatalln("Cannot parse db connection")
	}
	dbConfig.MinConns = 3
	dbConfig.MaxConns = 5
	connPool, err := pgxpool.ConnectConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}

	return connPool
}

func main() {
	connPool := dbConn()
	err := connPool.Ping(context.Background())
	if err != nil {
		log.Fatalln("Database connection failed")
	}
	repository := repositories.NewRepository(connPool)

	if _, err := os.Stat("data.json"); err == nil {
		dataFile, err := os.Open("data.json")
		if err != nil {
			log.Println("error opening data.json")
		} else {
			scanner := bufio.NewScanner(dataFile)
			scanner.Split(bufio.ScanLines)
			var adsData entities.AdsData
			for scanner.Scan() {
				rawTxt := scanner.Text()
				hash := sha256.Sum256([]byte(rawTxt))
				recHash := fmt.Sprintf("%x", hash)

				json.Unmarshal([]byte(scanner.Text()), &adsData)
				adsData.RowHash = recHash
				adsDataObj, err := repository.AdsDataRepository.Save(context.Background(), &adsData)
				if err != nil {
					log.Println(err)
				} else {
					for _, tag := range adsData.Tags {
						tagObj := &entities.Tag{
							AdsDataID: adsDataObj.ID,
							TagName:   tag,
						}
						_, err = repository.TagsRepository.Save(context.Background(), tagObj)
						if err != nil {
							log.Println(err)
						}
					}

					for _, imgUrl := range adsData.ImageUrls {
						imgObj := &entities.ImageUrl{
							AdsDataID: adsDataObj.ID,
							Image:     imgUrl,
						}
						_, err = repository.ImageUrlRepository.Save(context.Background(), imgObj)
						if err != nil {
							log.Println(err)
						}
					}
				}

			}
		}
	}
	datas, err := repository.AdsDataRepository.SearchFullText(context.Background(), "sony")
	if err != nil {
		log.Println("NO RESULT")
	}

	for _, dat := range datas {
		fmt.Println("content = ", dat.Content)
		fmt.Println("title = ", dat.Title)
	}

}
