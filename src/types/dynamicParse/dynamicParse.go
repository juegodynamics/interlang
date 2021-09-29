package main

import (
	"encoding/json"
	"fmt"
)

type MeasurementType string

const (
	MeasurementType_Quantity MeasurementType = "quantity"
	MeasurementType_Ratio                    = "ratio"
)

type Quantity struct {
	Type   MeasurementType `json:"type"`
	Amount float64         `json:"amount"`
	Unit   string          `json:"unit"`
}

func newQuantity(amount float64, unit string) (q *Quantity) {
	q = &Quantity{Amount: amount, Unit: unit}
	q.Type = q.GetMeasurementType()
	return
}
func (q *Quantity) GetMeasurementType() MeasurementType { return MeasurementType_Quantity }
func (q *Quantity) isMeasurement()                      {}

type Ratio struct {
	Type        MeasurementType `json:"type"`
	Numerator   *Quantity       `json:"numerator"`
	Denominator *Quantity       `json:"denominator"`
}

func newRatio(numerator *Quantity, denominator *Quantity) (q *Ratio) {
	q = &Ratio{Numerator: numerator, Denominator: denominator}
	q.Type = q.GetMeasurementType()
	return
}
func (q *Ratio) GetMeasurementType() MeasurementType { return MeasurementType_Ratio }
func (q *Ratio) isMeasurement()                      {}

type Measurement interface {
	GetMeasurementType() MeasurementType
	isMeasurement()
}

func getAllMeasurementInterfaces() []Measurement {
	return []Measurement{&Quantity{}, &Ratio{}}
}

func getAllMeasurementInterfacesByType() (mTypeToM map[MeasurementType]Measurement) {
	mTypeToM = make(map[MeasurementType]Measurement)
	for _, m := range getAllMeasurementInterfaces() {
		mTypeToM[m.GetMeasurementType()] = m
	}
	return
}

type MeasurementList struct {
	Measurements []Measurement `json:"measurements"`
}

func (m *MeasurementList) UnmarshalJSON(data []byte) (err error) {
	measurementListGeneric := struct {
		Measurements []map[string]interface{} `json:"measurements"`
	}{}
	if err = json.Unmarshal(data, &measurementListGeneric); err != nil {
		return
	}

	for _, measurement := range measurementListGeneric.Measurements {
		if measurementTypeString, ok := measurement["type"].(string); ok {
			measurementInterface := getAllMeasurementInterfacesByType()[MeasurementType(measurementTypeString)]

			if measurementInterface == nil {
				err = fmt.Errorf("found no interface for 'type' %s", measurementTypeString)
				return
			}

			if err = json.Unmarshal(data, measurementInterface); err != nil {
				return
			}

			m.Measurements = append(m.Measurements, measurementInterface)

		} else {
			err = fmt.Errorf("missing 'type' key")
			return
		}

	}
	return
}

func main() {
	sampleData := &MeasurementList{
		Measurements: []Measurement{
			newQuantity(10, "miligrams"),
			newRatio(newQuantity(10, "miligrams"), newQuantity(20, "mililiters")),
		},
	}

	data, err := json.Marshal(sampleData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(data))

	newSampleData := &MeasurementList{}
	if err = json.Unmarshal(data, newSampleData); err != nil {
		panic(err)
	}
	fmt.Printf("Successful parsing: %v\n", *newSampleData)
}
