package services

import (
	"fmt"
	"os"
	"time"
)

func GenerateSitemap(data []map[string]interface{}, key string) {
	var SitemapXml = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	SitemapXml += `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n"
	currentTime := time.Now().UTC()
	var urls []string
	for _, url := range data[:100] {
		urls = append(urls, url[key].(string))
	}
	for _, url := range urls {
		SitemapXml += fmt.Sprintf("<url>\n<loc>%s</loc>\n<lastmod>%s</lastmod>\n</url>\n", url, currentTime.Format("2006-01-02T15:04:05Z"))
	}
	SitemapXml += `</urlset>`
	filename := "sitemap.xml"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()
	_, err = file.WriteString(SitemapXml)
	if err != nil {
		fmt.Println("Error writing sitemap content to file:", err)
	}
	fmt.Println("Sitemap content successfully written to", filename)
}

func init() {
	var data []map[string]interface{}
	item := map[string]interface{}{
		"_id": "abc",
		"url": "www.example.com/1",
	}
	data = append(data, item)
	GenerateSitemap(data, "url")

}
