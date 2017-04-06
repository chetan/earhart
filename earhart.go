package earhart

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/mesosphere/mesos-dns/models"
)

type record struct {
	Service string `json:"service"`
	Host    string `json:"host"`
	IP      string `json:"ip"`
	Port    string `json:"port"`
}

func GetAxfr(url string) (*models.AXFR, error) {
	resp, err := http.Get(url + "/v1/axfr")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	axfr := models.AXFR{}
	if err := json.NewDecoder(resp.Body).Decode(&axfr); err != nil {
		return nil, err
	}

	return &axfr, nil
}

func GetService(url string, svc string) {
	// http://192.168.33.7:8123/v1/services/_zurvey.2717abb2.zurvey.devel.root._tcp.aurora.mesos.

	// cleanup the svc name
	if svc[0] != '_' {
		svc = "_" + svc
	}
	match, err := regexp.MatchString("\\.(devel|prod)\\.(.*?)\\._tcp\\.aurora\\.mesos\\.?", svc)
	if !match {
		svc += ".devel.root._tcp.aurora.mesos."
	}
	if !strings.HasSuffix(svc, ".") {
		svc += "."
	}

	resp, err := http.Get(url + "/v1/services/" + svc)
	if err != nil {
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	println(string(body))
}
