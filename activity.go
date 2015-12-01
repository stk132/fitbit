package fitbit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Period string
type ActivityLogType string

const (
	StepsLog               ActivityLogType = "activities/steps"
	CaloriesLog            ActivityLogType = "activities/calories"
	CaloriesBMRLog         ActivityLogType = "activities/caloriesBMR"
	DistanceLog            ActivityLogType = "activities/distance"
	FloorsLog              ActivityLogType = "activities/floors"
	ElevationLog           ActivityLogType = "activities/elevation"
	MinutesSedentaryLog    ActivityLogType = "activities/minutesSedentary"
	MinutesLightActiveLog  ActivityLogType = "activities/minutesLightlyActive"
	MinutesFairlyActiveLog ActivityLogType = "activities/minutesFairlyActive"
	MinutesVeryActiveLog   ActivityLogType = "activities/minutesVeryActive"
	ActivityCaloriesLog    ActivityLogType = "activities/activityCalories"

	OneMonth   Period = "1m"
	OneDay     Period = "1d"
	OneWeek    Period = "7d"
	ThirtyDay  Period = "30d"
	ThreeMonth Period = "3m"
	HalfYear   Period = "6m"
	OneYear    Period = "1y"
	Max        Period = "max"
	// ActivityURL fitbit activity api url
	ActivityURL           string = "https://api.fitbit.com/1/user/%s/activities/date/%s.json"
	ActivityTimeSeriesURL string = "https://api.fitbit.com/1/user/%s/%s/date/%s/%s.json"
)

// Activities
type Activity struct {
	c *Client
}

// Activity fitbit activity data
type ActivityData struct {
	ActivityID       uint64  `json:"activityId"`
	ActivityParentID uint64  `json:"activityParentId"`
	Calories         uint64  `json:"calories"`
	Description      string  `json:"description"`
	Distance         float64 `json:"distance"`
	Duration         uint64  `json:"duration"`
	HasStartTime     bool    `json:"hasStartTime"`
	IsFavorite       bool    `json:"isFavorite"`
	LogID            uint64  `json:"logId"`
	Name             string  `json:"name"`
	StartTime        string  `json:"startTime"`
	Steps            uint64  `json:"steps"`
}

// Goals fitbit Goals
type Goals struct {
	CaloriesOut uint64  `json:"caloriesOut"`
	Distance    float64 `json:"distance"`
	Floors      uint64  `json:"floors"`
	Steps       uint64  `json:"steps"`
}

// Distance fitbit Distance
type Distance struct {
	Activity string  `json:"activity"`
	Distance float64 `json:"distance"`
}

// Summary fitbit activity summary
type Summary struct {
	ActivityCalories     uint64      `json:"activityCalories"`
	CaloriesBMR          uint64      `json:"caloriesBMR"`
	CaloriesOut          uint64      `json:"caloriesOut"`
	Distances            []*Distance `json:"distances"`
	Elevation            float64     `json:"elevation"`
	FairlyActiveMinutes  uint64      `json:"fairlyActiveMinutes"`
	Floors               uint64      `json:"floors"`
	LightlyActiveMinutes uint64      `json:"lightlyActiveMinutes"`
	MarginalCalories     uint64      `json:"marginalCalories"`
	SedentaryMinutes     uint64      `json:"sedentaryMinutes"`
	Steps                uint64      `json:"steps"`
	VeryActiveMinutes    uint64      `json:"veryActiveMinutes"`
}

// ActivityResponse fitbit activity api response
type ActivityResponse struct {
	Activities []*ActivityData `json:"activities"`
	Goals      Goals           `json:"goals"`
	Summary    Summary         `json:"summary"`
}

type ActivitiesLogString struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

type ActivitiesLogStepsResponse struct {
	ActivitiesLogSteps []*ActivitiesLogString `json:"activities-steps"`
}

type ActivitiesLogCaloriesResponse struct {
	ActivitiesLogCalories []*ActivitiesLogString `json:"activities-calories"`
}

type ActivitiesLogCaloriesBMRResponse struct {
	ActivitiesLogCaloriesBMR []*ActivitiesLogString `json:"activities-caloriesBMR"`
}

type ActivitiesLogDistanceResponse struct {
	ActivitiesLogDistance []*ActivitiesLogString `json:"activities-distance"`
}

type ActivitiesLogFloorsResponse struct {
	ActivitiesLogFloors []*ActivitiesLogString `json:"activities-floors"`
}

type ActivitiesLogElevationResponse struct {
	ActivitiesLogElevation []*ActivitiesLogString `json:"activities-elevation"`
}

type ActivitiesLogMinutesSedentaryResponse struct {
	ActivitiesLogMinutesSedentary []*ActivitiesLogString `json:"activities-minutesSedentary"`
}

type ActivitiesLogMinutesLightActiveResponse struct {
	ActivitiesLogMinutesLightActive []*ActivitiesLogString `json:"activities-minutesLightActive"`
}

type ActivitiesLogMinutesFairyActiveResponse struct {
	ActivitiesLogMinutesFairyActive []*ActivitiesLogString `json:"activities-minutesFairyActive"`
}

type ActivitiesLogMinutesVeryActiveResponse struct {
	ActivitiesLogMinutesVeryActive []*ActivitiesLogString `json:"activities-minutesVeryActive"`
}

type ActivitiesLogActivityCaloriesResponse struct {
	ActivitiesLogActivityCalories []*ActivitiesLogString `json:"activities-activityCalories"`
}

// DailyActivitySummaryByID hogehoge
func (a *Activity) DailyActivitySummaryByID(userID string, date string) (*ActivityResponse, error) {
	url := fmt.Sprintf(ActivityURL, userID, date)
	result, err := a.c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()
	resultByteArray, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	activity := &ActivityResponse{}
	err = json.Unmarshal(resultByteArray, activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

// DailyActivitySummary hogehoge
func (a *Activity) DailyActivitySummary(date string) (*ActivityResponse, error) {
	return a.DailyActivitySummaryByID("-", date)
}

// ActivityTimeSeriesByID hogehoge
func (a *Activity) ActivityTimeSeriesByID(userID string, date string, period Period, activityLogType ActivityLogType) ([]*ActivitiesLogString, error) {
	url := fmt.Sprintf(ActivityTimeSeriesURL, userID, string(activityLogType), date, string(period))
	resultByteArray, err := a.c.Get(url)
	if err != nil {
		return nil, err
	}
	return activityLogConvert(resultByteArray, activityLogType)
}

func (a *Activity) ActivityTimeSeries(date string, period Period, activityLogType ActivityLogType) ([]*ActivitiesLogString, error) {
	return a.ActivityTimeSeriesByID("-", date, period, activityLogType)
}

func activityLogConvert(resultByteArray []byte, activityLogType ActivityLogType) ([]*ActivitiesLogString, error) {
	switch activityLogType {
	case StepsLog:
		activityLogSteps := &ActivitiesLogStepsResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogSteps); err != nil {
			return nil, err
		}
		return activityLogSteps.ActivitiesLogSteps, nil
	case CaloriesLog:
		activityLogCalories := &ActivitiesLogCaloriesResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogCalories); err != nil {
			return nil, err
		}
		return activityLogCalories.ActivitiesLogCalories, nil
	case CaloriesBMRLog:
		activityLogCaloriesBMR := &ActivitiesLogCaloriesBMRResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogCaloriesBMR); err != nil {
			return nil, err
		}
		return activityLogCaloriesBMR.ActivitiesLogCaloriesBMR, nil
	case DistanceLog:
		activityLogDistance := &ActivitiesLogDistanceResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogDistance); err != nil {
			return nil, err
		}
		return activityLogDistance.ActivitiesLogDistance, nil
	case FloorsLog:
		activityLogFloor := &ActivitiesLogFloorsResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogFloor); err != nil {
			return nil, err
		}
		return activityLogFloor.ActivitiesLogFloors, nil
	case ElevationLog:
		activityLogElevation := &ActivitiesLogElevationResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogElevation); err != nil {
			return nil, err
		}
		return activityLogElevation.ActivitiesLogElevation, nil
	case MinutesSedentaryLog:
		activityLogMinutesSedentary := &ActivitiesLogMinutesSedentaryResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogMinutesSedentary); err != nil {
			return nil, err
		}
		return activityLogMinutesSedentary.ActivitiesLogMinutesSedentary, nil
	case MinutesLightActiveLog:
		activityLogMinutesLightActive := &ActivitiesLogMinutesLightActiveResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogMinutesLightActive); err != nil {
			return nil, err
		}
		return activityLogMinutesLightActive.ActivitiesLogMinutesLightActive, nil
	case MinutesFairlyActiveLog:
		activityLogMinutesFairyActive := &ActivitiesLogMinutesFairyActiveResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogMinutesFairyActive); err != nil {
			return nil, err
		}
		return activityLogMinutesFairyActive.ActivitiesLogMinutesFairyActive, nil
	case MinutesVeryActiveLog:
		activityLogMinutesVeryActive := &ActivitiesLogMinutesVeryActiveResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogMinutesVeryActive); err != nil {
			return nil, err
		}
		return activityLogMinutesVeryActive.ActivitiesLogMinutesVeryActive, nil
	case ActivityCaloriesLog:
		activityLogActivityCalories := &ActivitiesLogActivityCaloriesResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogActivityCalories); err != nil {
			return nil, err
		}
		return activityLogActivityCalories.ActivitiesLogActivityCalories, nil
	default:
		return nil, errors.New(string(activityLogType) + " not implemented")
	}

}
