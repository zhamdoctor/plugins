package basic

import (
	"errors"
	cip "github.com/containernetworking/plugins/pkg/ip"
	"net"
)

type Ipam struct {
	subnet  *net.IPNet
	gateway net.IP
	store   *Store
}

// 将cni配置文件中subnet转换成ip网段，存入subnet，第一个ip作为网关ip
func NewIpam(conf *CNIConf, store *Store) (*Ipam, error) {
	_, ipNet, err := net.ParseCIDR(conf.Subnet)
	if err != nil {
		return nil, err
	}
	im := &Ipam{
		subnet: ipNet,
		store:  store,
	}
	im.gateway, err = im.NextIP(im.subnet.IP) //分配的第一个ip为gateway
	return im, err
}

func (im *Ipam) NextIP(ip net.IP) (net.IP, error) {
	nextIP := cip.NextIP(ip)
	if !im.subnet.Contains(nextIP) {
		return nil, errors.New("下一条ip被分配")
	}
	return nextIP, nil
}

func (im *Ipam) Gateway() net.IP {
	return im.gateway
}

func (im *Ipam) AllocateIP(id, ifName string) (net.IP, error) {

}
