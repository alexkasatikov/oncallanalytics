package webhook

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

type Data struct {
}

func AlertmanagerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Enter handler")
	if r.URL.Path != "/alertmanager" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		log.Println("Received POST request")
		var group Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			log.Println(err.Error())
			return
		}

		//fmt.Fprintf(w, group.)
		//fmt.Println("StartsAt: ", group.Alerts[0].StartsAt)
		log.Println(group)
		log.Println("Done")

		//for _, alert := range group.Alerts {
		//	fmt.Println(alert.Status)
		//}

		//postgresql.InsertAlert(DatabaseURL)
	default:
		log.Printf("Received %s request", r.Method)
	}
}
