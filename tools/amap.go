package tools

import (
	"basic_framework/configs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Amap struct {
}

// 逆地理编码API服务地址
func (this Amap) Regeo(lng, lat float64) map[string]interface{} {

	url := configs.Yaml.Amap.HostRegeo + configs.Yaml.Amap.WebServerKey + "&location=%v,%v"
	url = fmt.Sprintf(url, lng, lat)

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	r, e := client.Get(url)
	defer r.Body.Close()

	m := make(map[string]interface{})

	if e == nil {
		b, er := ioutil.ReadAll(r.Body)
		if er == nil {
			er = json.Unmarshal(b, &m)
			if er == nil {
				if m["status"] == "1" {
					m = m["regeocode"].(map[string]interface{})
				}
			}

		}
	}

	return m
}

// 坐标转换
func (this Amap) Coordinate(lng, lat float64) (amapLng float64, amapLat float64) {

	url := configs.Yaml.Amap.HostCoordinate + configs.Yaml.Amap.WebServerKey + "&locations=%v,%v" + "&coordsys=gps"
	url = fmt.Sprintf(url, lng, lat)

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	r, e := client.Get(url)
	defer r.Body.Close()

	amapLng = 0
	amapLat = 0
	m := make(map[string]interface{})

	if e == nil {
		b, er := ioutil.ReadAll(r.Body)
		if er == nil {

			er = json.Unmarshal(b, &m)
			if er == nil {
				if m["status"] == "1" {
					str := strings.Split(m["locations"].(string), ",")
					amapLng, _ = strconv.ParseFloat(str[0], 64)
					amapLat, _ = strconv.ParseFloat(str[1], 64)
				}
			}

		}
	}

	return amapLng, amapLat
}
