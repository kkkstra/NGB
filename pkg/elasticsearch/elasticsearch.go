package elasticsearch

import (
	"NGB/pkg/logrus"
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"net/http"
)

type ElasticsearchConfig struct {
	Addresses []string `yaml:"addresses"`
}

func InitElasticsearch(config *ElasticsearchConfig) {
	//cert, err := os.ReadFile("./configs/certs/elastic.crt")
	//if err != nil {
	//	logrus.Logger.Fatalf("Error reading the cert for elasticsearch: %s", err)
	//}
	cfg := elasticsearch.Config{
		Addresses: config.Addresses,
		Username:  "elastic",
		Password:  "lyx204406",
		//CACert:    cert,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		logrus.Logger.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		logrus.Logger.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	io.Copy(io.Discard, res.Body)
}
