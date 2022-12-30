package basic

import (
	"encoding/json"
	"github.com/containernetworking/cni/pkg/types"
	"os"
)

var (
	defaultSubnetFile = "/run/mycni/subnet.json"
)

type CNIConf struct {
	PluginConf
	SubnetConf
}

func NewCNIConf(pluginConf PluginConf, subnetConf SubnetConf) *CNIConf {
	return &CNIConf{PluginConf: pluginConf, SubnetConf: subnetConf}
}

type PluginConf struct {
	types.NetConf
	RuntimeConfig *struct {
		Config map[string]interface{} `json:"config"`
	} `json:"runtimeConfig,omitempty"`
	Args *struct {
		A map[string]interface{} `json:"cni"`
	} `json:"args"`
	DataDir string `json:"dataDir"`
}

type SubnetConf struct {
	Subnet string `json:"subnet"`
	Bridge string `json:"bridge"`
}

// LoadCNIConfig
func LoadCNIConfig(stdin []byte) (*CNIConf, error) {
	pluginConfig, err := parsePluginConfig(stdin)
	if err != nil {
		return nil, err
	}
	subnetConf, err := loadSubnetConfig()
	if err != nil {
		return nil, err
	}
	return NewCNIConf(*pluginConfig, *subnetConf), nil
}

func parsePluginConfig(stdin []byte) (*PluginConf, error) {
	var pluginConf PluginConf
	if err := json.Unmarshal(stdin, &pluginConf); err != nil {
		return nil, err
	}
	return &pluginConf, nil
}

func loadSubnetConfig() (*SubnetConf, error) {
	var subnetConf SubnetConf
	file, err := os.ReadFile(defaultSubnetFile)
	if err != nil {
		err := os.MkdirAll(defaultSubnetFile, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	if err := json.Unmarshal(file, &subnetConf); err != nil {
		return nil, err
	}
	return &subnetConf, nil
}
