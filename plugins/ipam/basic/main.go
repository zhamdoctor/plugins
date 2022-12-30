package basic

import (
	"github.com/containernetworking/cni/pkg/skel"
	"log"
)

// ip分配和网桥配置
func cmdAdd(args *skel.CmdArgs) error {
	//加载cni配置
	cniConfig, err := LoadCNIConfig(args.StdinData)
	if err != nil {
		log.Fatal("加载cni配置文件失败,err:%v", err)
		return err
	}
	//存储本机ip分配列表
	store, err := NewStore(cniConfig.DataDir, cniConfig.Name)
	defer store.Close()
	//ipam分配ip
	ipam := NewIpam(cniConfig, store)

	//创建网桥,虚拟设备绑定到网桥
	//返回网络配置信息

}
