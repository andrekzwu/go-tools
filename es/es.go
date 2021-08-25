package es

import (
	"errors"
	"github.com/olivere/elastic"
	"time"
)

var ESClient *elastic.Client

type ESEntry struct {
	URLList       []string
	Interval      time.Duration
	Gzip          bool
	TraceLogger   elastic.Logger
	ErrorLogger   elastic.Logger
}

func RegisterES(entry *ESEntry) error {
	if entry == nil {
		return errors.New("redis entry empty nil points")
	}
	client, err := elastic.NewClient(
		elastic.SetURL(entry.URLList...),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(entry.Interval),
		elastic.SetGzip(entry.Gzip),
		elastic.SetErrorLog(entry.ErrorLogger),
		elastic.SetSendGetBodyAs("GET"),
		elastic.SetTraceLog(entry.TraceLogger))
	if err != nil {
		return err
	}
	ESClient = client
	return nil
}
