package metaInfoUtil

import (
	"fmt"
	"log"
	"net/http"

	m "github.com/keighl/metabolize"
)

type MetaData struct {
	Title       string `meta:"og:title,title"`
	Description string `meta:"og:description,description"`
	Type        string `meta:"og:type"`
	Image       string `meta:"og:image"`
}

func GetMetaInfo(url string) (*MetaData, error) {
	res, err := http.Get(url)

	if err != nil {
		log.Printf("Network error: %v", err)
		return nil, err
	}

	data := new(MetaData)

	err = m.Metabolize(res.Body, data)
	fmt.Print(err)
	if err != nil {
		log.Printf("Cannot get meta info of this url: %v", err)
		return nil, err
	}
	return data, nil
}
