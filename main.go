package main

import (
	"context"
	"fmt"
	"math"

	"github.com/joho/godotenv"
	db "github.com/yashvantvala/go-sitemap-generator/DB"
	"github.com/yashvantvala/go-sitemap-generator/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Url struct {
	_id primitive.ObjectID `bson:"_id,omitempty"`
	url string             `bson:"url"`
}

const maxSize = 1000
const baseURI = "https://www.example.com"
const baseDir = "sitemap"
const IndexSitemap = baseDir + "/sitemap-example-index.xml"

func main() {
	godotenv.Load()
	pageNumber := 1
	db.ConnectDB()
	utils.CreateFolder(baseDir)
	var filenames []string
	var urls []map[string]interface{}
	count, _ := db.Collections.CountDocuments(context.Background(), bson.M{})
	totalRound := int(math.Round(float64(count) / maxSize))
	for i := 1; i <= totalRound; i++ {
		fmt.Println("Batch Number ---- ", i)
		var skipValue = (pageNumber - 1) * maxSize
		opts := options.Find().SetSkip(int64(skipValue)).SetLimit(maxSize)
		cursor, err := db.Collections.Find(context.TODO(), bson.M{}, opts)
		for cursor.Next(context.Background()) {
			var urlBSON bson.M
			err = cursor.Decode(&urlBSON)
			if err != nil {
				fmt.Println(err)
			}
			urls = append(urls, urlBSON)
		}
		pageNumber++
		filename := fmt.Sprintf(`%s/sitemap%d.xml.gz`, baseDir, i)
		filenames = append(filenames, filename)
		utils.GenerateSitemap(urls, "url", filename)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Page Bumber", pageNumber)
	utils.GenerateIndexMap(filenames, baseURI, IndexSitemap)

}
