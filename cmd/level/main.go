package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	dockerleveldb "docker_leveldb"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Config struct {
	FilePath         string
	Port             string
	IsBloomfilter    bool
	BloomFilterCount int
}

var (
	config Config
	db     *leveldb.DB
	err    error
)

func init() {
	fs := flag.NewFlagSet("leveldb_go", flag.ExitOnError)
	fs.StringVar(&config.FilePath, "--Path", "/path/data", "leveldb data path")
	fs.StringVar(&config.Port, "--port", "8096", "http port")
	fs.BoolVar(&config.IsBloomfilter, "--use-BloomFilter", true, "select use bloom")
	fs.IntVar(&config.BloomFilterCount, "--count", 10, "BloomFilter count")
	fs.Parse(os.Args[1:])
}

func main() {
	if config.IsBloomfilter {
		db, err = newFloomFilter(config.FilePath, config.BloomFilterCount)
	}
	db, err = leveldb.OpenFile(config.FilePath, nil)
	level := dockerleveldb.NewLevelResource(db)
	restful.Add(level.WebService())

	//swagger api
	{
		rsconfig := restfulspec.Config{
			WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
			APIPath:                       "/apidocs.json",
			PostBuildSwaggerObjectHandler: enrichSwaggerObject}
		restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(rsconfig))
	}
	//http start
	{
		log.Printf("start listening on localhost:8080")
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil))
	}
}

func newFloomFilter(path string, count int) (*leveldb.DB, error) {
	o := &opt.Options{
		Filter: filter.NewBloomFilter(count),
	}

	db, err := leveldb.OpenFile("path/to/db", o)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	return db, nil
}
