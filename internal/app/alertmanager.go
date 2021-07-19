package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Group struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	Alerts            []struct {
		Status       string            `json:"status"`
		Labels       map[string]string `json:"labels"`
		Annotations  map[string]string `json:"annotations"`
		StartsAt     time.Time         `json:"startsAt"`
		EndsAt       time.Time         `json:"endsAt"`
		GeneratorURL string            `json:"generatorURL"`
		Fingerprint  string            `json:"fingerprint"`
	} `json:"alerts"`
}

type Alert struct {
	Fingerprint string
	Status      string
	StartsAt    time.Time
	EndsAt      time.Time
}

func AlertmanagerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/alertmanager" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		var group Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			log.Println(err.Error())
			return
		}

		for _, a := range group.Alerts {

			// TODO: need to make key:value pair uniq across lables map
			labels := make(map[string]string)

			for key, val := range group.GroupLabels {
				labels[key] = val
			}

			for key, val := range group.CommonLabels {
				labels[key] = val
			}

			for key, val := range a.Labels {
				labels[key] = val
			}

			alert := Alert{
				Status:      a.Status,
				StartsAt:    a.StartsAt,
				EndsAt:      a.EndsAt,
				Fingerprint: a.Fingerprint,
			}

			switch alert.Status {
			case "firing":
				alertId := UpdateAlerts(DatabaseURL, alert)
				labelsIds := UpdateLabels(DatabaseURL, labels)

				log.Println(alertId)
				log.Println(labelsIds)

			case "resolved":
				UpdateAlerts(DatabaseURL, alert)
			default:
				log.Println("Unexpected alert status")
			}

		}
	default:
		log.Printf("Received %s request", r.Method)
	}
}
