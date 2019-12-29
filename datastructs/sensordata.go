package datastructs

import "fmt"

type SensorData struct {
	SensorId         int               `json:"esp8266id,string"`
	SoftwareVersion  string            `json:"software_version"`
	SensorDataValues []SensorDataValue `json:"sensordatavalues"`
}

type SensorDataValue struct {
	ValueType string  `json:"value_type"`
	Value     float64 `json:"value,string"`
}

func (s SensorData) FlattenToMap() map[string]interface{} {
	dataMap := make(map[string]interface{})
	dataMap["SensorId"] = s.SensorId

	for _, e := range s.SensorDataValues {
		dataMap[e.ValueType] = e.Value
	}
	return dataMap
}

func (s SensorData) FlattenToArray() []string {
	var dataArray []string
	dataArray = append(dataArray, fmt.Sprintf("%d", s.SensorId))

	for _, e := range s.SensorDataValues {
		pd := fmt.Sprintf("%f", e.Value)
		dataArray = append(dataArray, pd)
	}

	return dataArray
}
