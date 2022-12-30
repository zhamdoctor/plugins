package basic

import (
	"github.com/alexflint/go-filemutex"
	"os"
	"path/filepath"
)

var (
	defaultDataDir = "/var/lib/cni"
)

type Store struct {
	*filemutex.FileMutex
	dir      string
	data     *data
	dataFile string
}

func NewStore(dataDir, network string) (*Store, error) {
	if dataDir == "" {
		dataDir = defaultDataDir
	}
	path := filepath.Join(dataDir, network)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}
	pt := filepath.Join(path, network+".json")
	if err := os.MkdirAll(pt, os.ModePerm); err != nil {
		return nil, err
	}
	mutex := filemutex.FileMutex{}
	data := &data{
		IPs: make(map[string]containerNetInfo),
	}
	return &Store{&mutex, path, data, pt}, nil
}

type data struct {
	IPs  map[string]containerNetInfo `json:"ips"`
	Last string                      `json:"last"`
}

type containerNetInfo struct {
	ID     string `json:"id"`
	IFName string `json:"if"`
}
