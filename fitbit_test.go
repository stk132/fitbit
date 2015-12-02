package fitbit

import (
	"fmt"
	"testing"
)

func Prepare() (*Client, error) {
	c := Config()
	fitbit := &Fitbit{config: c}
	err := fitbit.SetTokenFromFile("token.json")
	if err != nil {
		return nil, err
	}
	client, err := fitbit.Client()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestAuthURL(t *testing.T) {
	config := Config()
	fitbit := &Fitbit{config: config}
	url, err := fitbit.AuthURL()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(url)

}

func TestActivity(t *testing.T) {
	c := Config()
	fitbit := &Fitbit{config: c}
	err := fitbit.SetTokenFromFile("token.json")
	if err != nil {
		t.Error("can't read token file")
	}
	client, err := fitbit.Client()
	if err != nil {
		t.Error(err)
	}
	activitySummary, err := client.Activity.DailyActivitySummary("2015-11-23")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(activitySummary.Summary.CaloriesOut)
}

func TestActivityTimeSeriesByID(t *testing.T) {
	client, err := Prepare()
	if err != nil {
		t.Error(err)
	}

	logTypes := []ActivityLogType{
		CaloriesLog,
		CaloriesBMRLog,
		StepsLog,
		FloorsLog,
		ElevationLog,
		MinutesSedentaryLog,
		MinutesLightActiveLog,
		MinutesFairlyActiveLog,
		MinutesVeryActiveLog,
		ActivityCaloriesLog,
	}

	for _, logType := range logTypes {
		activitiesLog, err := client.Activity.ActivityTimeSeriesByID("-", "2015-11-20", OneWeek, logType)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("データサイズ:%d\n", len(activitiesLog.Logs))
	}

}

func TestBrowseActivityTypes(t *testing.T) {
	client, err := Prepare()
	if err != nil {
		t.Error(err)
		return
	}
	response, err := client.Activity.BrowseActivityTypes()
	if err != nil {
		t.Error(err)
		return
	}

	for _, category := range response.Categories {
		for _, activity := range category.Activities {
			fmt.Println(activity.Name)
		}
	}
}

func TestGetActivityType(t *testing.T) {
	client, err := Prepare()
	if err != nil {
		t.Error(err)
		return
	}

	response, err := client.Activity.GetActivityType("1010")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(response.Activity.Name)
}

func TestGetFrequentActivities(t *testing.T) {
	client, err := Prepare()
	if err != nil {
		t.Error(err)
		return
	}

	response, err := client.Activity.GetFrequentActivities()
	if err != nil {
		t.Error(err)
		return
	}

	for _, frequentActivity := range response {
		fmt.Println(frequentActivity.Name)
	}
}

func TestGetRecentActivities(t *testing.T) {
	client, err := Prepare()
	if err != nil {
		t.Error(err)
		return
	}

	response, err := client.Activity.GetRecentActivities()
	if err != nil {
		t.Error(err)
		return
	}

	for _, recentActivity := range response {
		fmt.Println(recentActivity.Name)
	}
}
