package config

import "errors"

type DensoConfig struct {
	Host string
	Auth string
}



func (c *DensoConfig) SetConf(method string)  error{
    if method == "DEV"{
    	c.Host = "dnsosoapub-de4.opc.oracleoutsourcing.com"
    	c.Auth = "Basic b2hzY3VzdGFkbWluOkhpZDhfSHlfcg=="
	}else if method == "UAT"{
		c.Host = "dnsosoapub-ts3.opc.oracleoutsourcing.com"
		c.Auth = "Basic b2hzY3VzdGFkbWluOkhpZDhfSHlfcg=="
	}else{
		return errors.New("METHOD Error, it should be DEV or UAT！")
	}

	return nil


}

