package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	CAR_TYPE_SMALL = "small"
	CAR_TYPE_LARGE = "large"

	FLATE_RATE_SMALL_CAR = 25
	FLATE_RATE_LARGE_CAR = 35

	PERCENTAGE_REFUEL   = 0.10
	FUEL_RATE_PER_LITRE = 1.75

	EMPLOYEE_A = "A"
	EMPLOYEE_B = "B"
)

type Vehicle struct {
	Fuel         Fuel   `json:"fuel"`
	LicencePlate string `json:"licencePlate"`
	Size         string `json:"size"`
}

type Fuel struct {
	Capacity int64   `json:"capacity"`
	Level    float64 `json:"level"`
}

type Assignment struct {
	LicencePlate string  `json:"licencePlate"`
	Employee     string  `json:"employee"`
	FuelAdded    float64 `json:"fuelAdded"`
	Price        float64 `json:"price"`
}

type Employee struct {
	Name       string
	Commission float64
	Paid       float64
}

func main() {
	vehicles, err := getVehiclesFromFile("input-data.json")
	if err != nil {
		panic(err)
	}

	empA := Employee{Name: EMPLOYEE_A, Commission: 0.11}
	empB := Employee{Name: EMPLOYEE_B, Commission: 0.15}

	assignments, err := assignTasks(empA, empB, vehicles)
	if err != nil {
		panic(err)
	}

	result, err := json.MarshalIndent(assignments, "", "\t")
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", string(result))
}

func getVehiclesFromFile(filename string) ([]Vehicle, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	vehicles := []Vehicle{}
	if err := json.Unmarshal(data, &vehicles); err != nil {
		return nil, err
	}

	return vehicles, err
}

func assignTasks(empA, empB Employee, vehicles []Vehicle) ([]Assignment, error) {
	assignments := []Assignment{}

	for _, v := range vehicles {
		asgmt, err := processNextAssignement(&empA, &empB, v)
		if err != nil {
			return nil, err
		}

		assignments = append(assignments, asgmt)
	}

	return assignments, nil
}

func processNextAssignement(empA, empB *Employee, v Vehicle) (Assignment, error) {
	var err error
	asgmt := Assignment{}
	employeeName := getEmployeeNameForNextTask(*empA, *empB)

	switch employeeName {
	case EMPLOYEE_A:
		asgmt, err = empA.completeTask(v)
		if err != nil {
			return Assignment{}, err
		}

		empA.Paid += asgmt.Price * empA.Commission
	case EMPLOYEE_B:
		asgmt, err = empB.completeTask(v)
		if err != nil {
			return Assignment{}, err
		}

		empB.Paid += asgmt.Price * empB.Commission
	}

	return asgmt, nil
}

func getEmployeeNameForNextTask(empA, empB Employee) string {
	if empA.Paid == 0 && empB.Paid == 0 {
		if empA.Commission <= empB.Commission {
			return EMPLOYEE_A
		} else {
			return EMPLOYEE_B
		}
	}

	if empA.Paid <= empB.Paid {
		return EMPLOYEE_A
	}

	return EMPLOYEE_B
}

func (e *Employee) completeTask(v Vehicle) (Assignment, error) {
	asgmt := Assignment{
		Employee:     e.Name,
		LicencePlate: v.LicencePlate,
		FuelAdded:    calculateFuelAdded(v.Fuel),
	}

	flateRate, err := getCarFlateRate(v.Size)
	if err != nil {
		return Assignment{}, err
	}

	asgmt.Price = (asgmt.FuelAdded * FUEL_RATE_PER_LITRE) + float64(flateRate)

	return asgmt, nil
}

func getCarFlateRate(size string) (int, error) {
	switch size {
	case CAR_TYPE_LARGE:
		return FLATE_RATE_LARGE_CAR, nil
	case CAR_TYPE_SMALL:
		return FLATE_RATE_SMALL_CAR, nil
	default:
		return 0, errors.New(fmt.Sprintf("unknown size: %s", size))
	}
}

func calculateFuelAdded(f Fuel) float64 {
	if f.Level > PERCENTAGE_REFUEL {
		return 0
	}

	currentFuel := float64(f.Capacity) * f.Level
	return float64(f.Capacity) - currentFuel
}
