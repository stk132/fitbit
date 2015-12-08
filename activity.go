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
	ActivityURL                 string = "https://api.fitbit.com/1/user/%s/activities/date/%s.json"
	ActivityTimeSeriesURL       string = "https://api.fitbit.com/1/user/%s/%s/date/%s/%s.json"
	BrowseActivityTypesURL      string = "https://api.fitbit.com/1/activities.json"
	GetActivityTypeURL          string = "https://api.fitbit.com/1/activities/%s.json"
	GetFrequentActivitiesURL    string = "https://api.fitbit.com/1/user/-/activities/frequent.json"
	GetRecentActivitiesURL      string = "https://api.fitbit.com/1/user/-/activities/recent.json"
	GetFavoriteActivitiesURL    string = "https://api.fitbit.com/1/user/%s/activities/favorite.json"
	FavoriteActivityResourceURL string = "https://api.fitbit.com/1/user/-/activities/favorite/%s.json"
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

type ActivitiesLog struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

type ActivitiesLogIntradayDataSet struct {
	Time  string `json:"time"`
	Value string `json:"value"`
}

type ActivitiesLogIntraday struct {
	DataSetInterval uint64                          `json:"datasetInterval"`
	DataSet         []*ActivitiesLogIntradayDataSet `json:"dataset"`
}

type ActivityTimeSeriesResponse struct {
	Logs     []*ActivitiesLog
	Intraday *ActivitiesLogIntraday
}

type ActivitiesLogStepsResponse struct {
	ActivitiesLogSteps         []*ActivitiesLog       `json:"activities-steps"`
	ActivitiesLogStepsIntraday *ActivitiesLogIntraday `json:"activities-steps-intraday"`
}

type ActivitiesLogCaloriesResponse struct {
	ActivitiesLogCalories         []*ActivitiesLog       `json:"activities-calories"`
	ActivitiesLogCaloriesIntraday *ActivitiesLogIntraday `json:"activities-calories-intraday"`
}

type ActivitiesLogCaloriesBMRResponse struct {
	ActivitiesLogCaloriesBMR []*ActivitiesLog `json:"activities-caloriesBMR"`
}

type ActivitiesLogDistanceResponse struct {
	ActivitiesLogDistance         []*ActivitiesLog       `json:"activities-distance"`
	ActivitiesLogDistanceIntraday *ActivitiesLogIntraday `json:"activities-distance-intraday"`
}

type ActivitiesLogFloorsResponse struct {
	ActivitiesLogFloors         []*ActivitiesLog       `json:"activities-floors"`
	ActivitiesLogFloorsIntraday *ActivitiesLogIntraday `json:"activities-floors-intraday"`
}

type ActivitiesLogElevationResponse struct {
	ActivitiesLogElevation         []*ActivitiesLog       `json:"activities-elevation"`
	ActivitiesLogElevationIntraday *ActivitiesLogIntraday `json:"activities-elevation-intraday"`
}

type ActivitiesLogMinutesSedentaryResponse struct {
	ActivitiesLogMinutesSedentary []*ActivitiesLog `json:"activities-minutesSedentary"`
}

type ActivitiesLogMinutesLightActiveResponse struct {
	ActivitiesLogMinutesLightActive []*ActivitiesLog `json:"activities-minutesLightActive"`
}

type ActivitiesLogMinutesFairyActiveResponse struct {
	ActivitiesLogMinutesFairyActive []*ActivitiesLog `json:"activities-minutesFairyActive"`
}

type ActivitiesLogMinutesVeryActiveResponse struct {
	ActivitiesLogMinutesVeryActive []*ActivitiesLog `json:"activities-minutesVeryActive"`
}

type ActivitiesLogActivityCaloriesResponse struct {
	ActivitiesLogActivityCalories []*ActivitiesLog `json:"activities-activityCalories"`
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
func (a *Activity) ActivityTimeSeriesByID(userID string, date string, period Period, activityLogType ActivityLogType) (*ActivityTimeSeriesResponse, error) {
	url := fmt.Sprintf(ActivityTimeSeriesURL, userID, string(activityLogType), date, string(period))
	resultByteArray, err := a.c.Get(url)
	if err != nil {
		return nil, err
	}
	return activityLogConvert(resultByteArray, activityLogType)
}

func (a *Activity) ActivityTimeSeries(date string, period Period, activityLogType ActivityLogType) (*ActivityTimeSeriesResponse, error) {
	return a.ActivityTimeSeriesByID("-", date, period, activityLogType)
}

func activityLogConvert(resultByteArray []byte, activityLogType ActivityLogType) (*ActivityTimeSeriesResponse, error) {
	switch activityLogType {
	case StepsLog:
		activityLogSteps := &ActivitiesLogStepsResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogSteps); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogSteps.ActivitiesLogSteps, Intraday: activityLogSteps.ActivitiesLogStepsIntraday}, nil
	case CaloriesLog:
		activityLogCalories := &ActivitiesLogCaloriesResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogCalories); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogCalories.ActivitiesLogCalories, Intraday: activityLogCalories.ActivitiesLogCaloriesIntraday}, nil
	case CaloriesBMRLog:
		activityLogCaloriesBMR := &ActivitiesLogCaloriesBMRResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogCaloriesBMR); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogCaloriesBMR.ActivitiesLogCaloriesBMR}, nil
	case DistanceLog:
		activityLogDistance := &ActivitiesLogDistanceResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogDistance); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogDistance.ActivitiesLogDistance, Intraday: activityLogDistance.ActivitiesLogDistanceIntraday}, nil
	case FloorsLog:
		activityLogFloor := &ActivitiesLogFloorsResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogFloor); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogFloor.ActivitiesLogFloors, Intraday: activityLogFloor.ActivitiesLogFloorsIntraday}, nil
	case ElevationLog:
		activityLogElevation := &ActivitiesLogElevationResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogElevation); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogElevation.ActivitiesLogElevation, Intraday: activityLogElevation.ActivitiesLogElevationIntraday}, nil
	case MinutesSedentaryLog:
		activityLogMinutesSedentary := &ActivitiesLogMinutesSedentaryResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogMinutesSedentary); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogMinutesSedentary.ActivitiesLogMinutesSedentary}, nil
	case MinutesLightActiveLog:
		activityLogMinutesLightActive := &ActivitiesLogMinutesLightActiveResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogMinutesLightActive); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogMinutesLightActive.ActivitiesLogMinutesLightActive}, nil
	case MinutesFairlyActiveLog:
		activityLogMinutesFairyActive := &ActivitiesLogMinutesFairyActiveResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogMinutesFairyActive); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogMinutesFairyActive.ActivitiesLogMinutesFairyActive}, nil
	case MinutesVeryActiveLog:
		activityLogMinutesVeryActive := &ActivitiesLogMinutesVeryActiveResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogMinutesVeryActive); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogMinutesVeryActive.ActivitiesLogMinutesVeryActive}, nil
	case ActivityCaloriesLog:
		activityLogActivityCalories := &ActivitiesLogActivityCaloriesResponse{}
		if err := json.Unmarshal(resultByteArray, activityLogActivityCalories); err != nil {
			return nil, err
		}
		return &ActivityTimeSeriesResponse{Logs: activityLogActivityCalories.ActivitiesLogActivityCalories}, nil
	default:
		return nil, errors.New(string(activityLogType) + " not implemented")
	}

}

type ActivityLevel struct {
	ID          uint64  `json:"id"`
	MasSpeedMPH float64 `json:"maxSpeedMPH"`
	MinSpeedMPH float64 `json:"minSpeedMPH"`
	Mets        float64 `json:"mets"`
	Name        string  `json:"name"`
}

type ActivityType struct {
	AccessLevel    string           `json:"accessLevel"`
	ActivityLevels []*ActivityLevel `json:"activityLevels"`
	HasSpeed       bool             `json:"hasSpeed"`
	ID             uint64           `json:"id"`
	Name           string           `json:"name"`
}

type Category struct {
	Activities []*ActivityType `json:"activities"`
}

type BrowseActivityTypesResponse struct {
	Categories []*Category `json:"categories"`
}

type GetActivityTypeResponse struct {
	Activity *ActivityType `json:"activity"`
}

type UserActivity struct {
	ActivityID  uint64 `json:"activityId"`
	Calories    uint64 `json:"calories"`
	Description string `json:"description"`
	Distance    uint64 `json:"distance"`
	Duration    uint64 `json:"duration"`
	Name        string `json:"name"`
}

type FavoriteActivity struct {
	ActivityID  uint64  `json:"activityId"`
	Description string  `json:"description"`
	Mets        float64 `json:"mets"`
	Name        string  `json:"name"`
}

func (a *Activity) BrowseActivityTypes() (*BrowseActivityTypesResponse, error) {
	resultByteArray, err := a.c.Get(BrowseActivityTypesURL)
	if err != nil {
		return nil, err
	}

	response := &BrowseActivityTypesResponse{}
	if err = json.Unmarshal(resultByteArray, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (a *Activity) GetActivityType(activityID string) (*GetActivityTypeResponse, error) {
	resultByteArray, err := a.c.Get(fmt.Sprintf(GetActivityTypeURL, activityID))
	if err != nil {
		return nil, err
	}

	response := &GetActivityTypeResponse{}
	if err = json.Unmarshal(resultByteArray, response); err != nil {
		return nil, err
	}
	return response, nil
}

func unmarshalUserActivity(resultByteArray []byte) ([]UserActivity, error) {
	var response []UserActivity
	if err := json.Unmarshal(resultByteArray, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (a *Activity) getUserActivities(url string) ([]UserActivity, error) {
	resultByteArray, err := a.c.Get(url)
	if err != nil {
		return nil, err
	}

	return unmarshalUserActivity(resultByteArray)
}

func (a *Activity) GetFrequentActivities() ([]UserActivity, error) {
	return a.getUserActivities(GetFrequentActivitiesURL)
}

func (a *Activity) GetRecentActivities() ([]UserActivity, error) {
	return a.getUserActivities(GetRecentActivitiesURL)
}

func (a *Activity) GetFavoriteActivitiesByID(userID string) ([]FavoriteActivity, error) {
	responseByteArray, err := a.c.Get(fmt.Sprintf(GetFavoriteActivitiesURL, userID))
	if err != nil {
		return nil, err
	}

	var response []FavoriteActivity
	if err = json.Unmarshal(responseByteArray, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (a *Activity) GetFavoriteActivities() ([]FavoriteActivity, error) {
	return a.GetFavoriteActivitiesByID("-")
}

func (a *Activity) AddFavoriteActivity(activityID string) error {
	if err := a.c.Post(fmt.Sprintf(FavoriteActivityResourceURL, activityID)); err != nil {
		return err
	}
	return nil
}

func (a *Activity) DeleteFavoriteActivity(activityID string) error {
	if err := a.c.Delete(fmt.Sprintf(FavoriteActivityResourceURL, activityID)); err != nil {
		return err
	}
	return nil
}
