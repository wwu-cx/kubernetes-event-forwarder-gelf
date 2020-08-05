package util

import (
	"errors"

	"github.com/golang/glog"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

// GetGelfWriter tcp or udp writer
func GetGelfWriter(host, port, proto string) gelf.Writer {
	graylogEndpoint := host + ":" + port
	glog.Infof("connecting to %s://%s", proto, graylogEndpoint)

	var gelfWriter gelf.Writer
	var err error

	if proto == "udp" {
		gelfWriter, err = gelf.NewUDPWriter(graylogEndpoint)
	} else if proto == "tcp" {
		gelfWriter, err = gelf.NewTCPWriter(graylogEndpoint)
	} else {
		err = errors.New("Protocol must be either tcp or udp")
	}

	if err != nil {
		glog.Fatal(err)
	}

	return gelfWriter
}
