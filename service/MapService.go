package service

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"github.com/z0890412/JoinPlus/model"
)

func GetPlaceID(palceName string, Log *logrus.Entry) string {
	var response model.FindPlaceResponse
	apiUrl := viper.GetString("MAP.Url") + "findplacefromtext/json"

	header := map[string]string{}
	queryParam := map[string]string{
		"input":     palceName,
		"inputtype": "textquery",
		"key":       viper.GetString("Map.Token"),
	}

	resp, err := HttpRequestWithQuery("GET", apiUrl, header, queryParam, Log)
	defer resp.Body.Close()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"apiUrl": apiUrl,
			"error":  err,
		}).Error("GetPlaceID error")
		return "-99"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"apiUrl": apiUrl,
			"error":  err,
		}).Error("GetPlaceID error")
		return "-99"
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"apiUrl": apiUrl,
			"error":  err,
		}).Error("GetPlaceID error")
		return "-99"
	}
	if response.Status == "OK" {
		return response.Candidates[0].Place_id
	}
	return "-99"
}

func GetStore(lat string, lng string, Log *logrus.Entry) string {
	var response model.FindPlaceResponse
	apiUrl := viper.GetString("MAP.Url") + "nearbysearch/json"

	header := map[string]string{}
	queryParam := map[string]string{
		"type":     "restaurant",
		"location": lat + "," + lng,
		"key":      viper.GetString("Map.Token"),
		"radius":   "1500",
	}

	resp, err := HttpRequestWithQuery("GET", apiUrl, header, queryParam, Log)
	defer resp.Body.Close()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"apiUrl": apiUrl,
			"error":  err,
		}).Error("GetPlaceID error")
		return "-99"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"apiUrl": apiUrl,
			"error":  err,
		}).Error("GetPlaceID error")
		return "-99"
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"apiUrl": apiUrl,
			"error":  err,
		}).Error("GetPlaceID error")
		return "-99"
	}
	if response.Status == "status" {
		return response.Candidates[0].Place_id
	}
	return "-99"
}

func GetPlaceDetail(palceID string, Log *logrus.Entry) (lat string, lng string) {
	apiUrl := viper.GetString("MAP.Url") + "details/json"

	header := map[string]string{}
	queryParam := map[string]string{
		"placeid": palceID,
		"key":     viper.GetString("Map.Token"),
	}

	resp, err := HttpRequestWithQuery("GET", apiUrl, header, queryParam, Log)
	defer resp.Body.Close()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"apiUrl": apiUrl,
			"error":  err,
		}).Error("GetPlaceDetail error")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"apiUrl": apiUrl,
			"error":  err,
		}).Error("GetPlaceDetail error")
		return
	}

	lat = gjson.GetBytes(body, "result.geometry.location.lat").Raw
	lng = gjson.GetBytes(body, "result.geometry.location.lng").Raw
	Log.Info(lat, lng)

	return "", "-99"
}
