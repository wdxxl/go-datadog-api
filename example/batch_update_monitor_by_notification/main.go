package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/wdxxl/go-datadog-api"
)

const api_key = "TODO"
const app_key = "TODO"

func main() {
	//updateMonitor(11158409)

	//getSlackDelvAlerts()
	//fmt.Println("-----")
	//getSlackDelvInfraAlerts()
}

func getSlackDelvAlerts() {
	client := datadog.NewClient(api_key, app_key)
	notification := "slack-delv-alerts"
	notificationMonitors, err := client.SearchNotification(notification, 0, 300)
	if err != nil {
		log.Print(err)
	}
	for i, value := range notificationMonitors {
		fmt.Println(i+1, value.Id, value.Name)
	}
}

func getSlackDelvInfraAlerts() {
	client := datadog.NewClient(api_key, app_key)
	notification := "slack-delv-infra-alerts"
	notificationMonitors, err := client.SearchNotification(notification, 0, 100)
	if err != nil {
		log.Print(err)
	}
	for i, value := range notificationMonitors {
		fmt.Println(i+1, value.Id, value.Name)
	}
}

func updateMonitor(id int) {
	// https://app.datadoghq.com/monitors#11158409/edit
	client := datadog.NewClient(api_key, app_key)
	monitor, _ := client.GetMonitor(id)
	log.Println(monitor.GetId())   // 11158409
	log.Println(monitor.GetName()) // [Conveyor][Staging][food-sre-reports] CPU usage too high
	creator := monitor.GetCreator()
	log.Println(creator.GetName(), creator.GetEmail(), creator.GetId(), creator.GetHandle())
	log.Println(monitor.GetQuery()) // avg(last_5m):100 - avg:system.cpu.idle{env:staging,appname:food-sre-reports} by {host} > 80
	log.Println(monitor.GetMessage()) // @webhook-delv-pprof @kexue.wang@grabtaxi.com @dingtalk-delv-alerts
	log.Printf("monitor %#+v", monitor)

	oldMessage := monitor.GetMessage()

	monitor.SetMessage(strings.ReplaceAll(oldMessage, "@kexue.wang@grabtaxi.com", "@liangliang.ma@grabtaxi.com"))
	client.UpdateMonitor(monitor)
}