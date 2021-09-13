package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_getVehiclesFromFile(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput []Vehicle
		expectedError  error
	}{
		{
			name:  "test happy path - 1 vehicle",
			input: "./test/test-data-1.json",
			expectedOutput: []Vehicle{
				{
					Fuel: Fuel{
						Capacity: 57,
						Level:    0.07,
					},
					LicencePlate: "A",
					Size:         "large",
				},
			},
		},
		{
			name:  "test happy path - 2 vehicles",
			input: "./test/test-data-2.json",
			expectedOutput: []Vehicle{
				{
					Fuel: Fuel{
						Capacity: 57,
						Level:    0.07,
					},
					LicencePlate: "A",
					Size:         "large",
				},
				{
					Fuel: Fuel{
						Capacity: 63,
						Level:    0.23,
					},
					LicencePlate: "B",
					Size:         "small",
				},
			},
		},
		{
			name:           "test happy path - no vehicle",
			input:          "./test/test-data-3.json",
			expectedOutput: []Vehicle{},
		},
		{
			name:           "test with error - wrong path",
			input:          "./test/wrong-path.json",
			expectedOutput: []Vehicle{},
			expectedError:  errors.New("open ./test/wrong-path.json: no such file or directory"),
		},
	}
	for _, test := range tests {
		res, err := getVehiclesFromFile(test.input)
		if err != nil {
			if test.expectedError != nil {
				if test.expectedError.Error() != err.Error() {
					t.Error(err)
				}
			} else {
				t.Error(err)
			}
		} else {
			if test.expectedError != nil {
				t.Error(test.expectedError)
			}

			if !reflect.DeepEqual(test.expectedOutput, res) {
				t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, res)
			}
		}
	}
}

func Test_assignTasks(t *testing.T) {
	tests := []struct {
		name           string
		empA, empB     Employee
		vehicles       []Vehicle
		expectedOutput []Assignment
		expectedError  error
	}{
		{
			name: "test happy path",
			empA: Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       0,
			},
			empB: Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       0,
			},
			vehicles: []Vehicle{
				{
					Fuel: Fuel{
						Capacity: 57,
						Level:    0.37,
					},
					LicencePlate: "ABC",
					Size:         "large",
				},
				{
					Fuel: Fuel{
						Capacity: 60,
						Level:    0.05,
					},
					LicencePlate: "DEF",
					Size:         "small",
				},
			},
			expectedOutput: []Assignment{
				{
					LicencePlate: "ABC",
					Employee:     "A",
					FuelAdded:    0,
					Price:        35,
				},
				{
					LicencePlate: "DEF",
					Employee:     "B",
					FuelAdded:    57,
					Price:        124.75,
				},
			},
		},
	}
	for _, test := range tests {
		res, err := assignTasks(test.empA, test.empB, test.vehicles)
		if err != nil {
			if test.expectedError != nil {
				if test.expectedError.Error() != err.Error() {
					t.Error(err)
				}
			} else {
				t.Error(err)
			}
		} else {
			if test.expectedError != nil {
				t.Error(test.expectedError)
			}

			if !reflect.DeepEqual(test.expectedOutput, res) {
				t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, res)
			}
		}
	}
}
func Test_processNextAssignement(t *testing.T) {
	tests := []struct {
		name                       string
		empA, empB                 *Employee
		vehicle                    Vehicle
		expectedOutput             Assignment
		ExpectedEmpA, ExpectedEmpB *Employee
		expectedError              error
	}{
		{
			name: "test happy path - employee A",
			empA: &Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       0,
			},
			empB: &Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       5,
			},
			vehicle: Vehicle{
				Fuel: Fuel{
					Capacity: 57,
					Level:    0.37,
				},
				LicencePlate: "ABC",
				Size:         "large",
			},
			expectedOutput: Assignment{
				LicencePlate: "ABC",
				Employee:     "A",
				FuelAdded:    0,
				Price:        35,
			},
			ExpectedEmpA: &Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       1.75,
			},
			ExpectedEmpB: &Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       5,
			},
		},
		{
			name: "test happy path - employee B",
			empA: &Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       10,
			},
			empB: &Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       5,
			},
			vehicle: Vehicle{
				Fuel: Fuel{
					Capacity: 57,
					Level:    0.37,
				},
				LicencePlate: "ABC",
				Size:         "large",
			},
			expectedOutput: Assignment{
				LicencePlate: "ABC",
				Employee:     "B",
				FuelAdded:    0,
				Price:        35,
			},
			ExpectedEmpA: &Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       10,
			},
			ExpectedEmpB: &Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       8.5,
			},
		},
	}
	for _, test := range tests {
		res, err := processNextAssignement(test.empA, test.empB, test.vehicle)
		if err != nil {
			if test.expectedError != nil {
				if test.expectedError.Error() != err.Error() {
					t.Error(err)
				}
			} else {
				t.Error(err)
			}
		} else {
			if test.expectedError != nil {
				t.Error(test.expectedError)
			}

			if !reflect.DeepEqual(test.expectedOutput, res) {
				t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, res)
			}

			if !reflect.DeepEqual(test.ExpectedEmpA, test.empA) {
				t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.ExpectedEmpA, test.empA)
			}

			if !reflect.DeepEqual(test.ExpectedEmpB, test.empB) {
				t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.ExpectedEmpB, test.empB)
			}
		}
	}
}

func Test_getEmployeeNameForNextTask(t *testing.T) {
	tests := []struct {
		name           string
		empA, empB     Employee
		expectedOutput string
	}{
		{
			name: "test happy path - employees with paid equal 0 & empA lower commission",
			empA: Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       0,
			},
			empB: Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       0,
			},
			expectedOutput: "A",
		},
		{
			name: "test happy path - employees with paid equal 0 & empA commission equal",
			empA: Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       0,
			},
			empB: Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       0,
			},
			expectedOutput: "A",
		},
		{
			name: "test happy path - employees with paid equal 0 & empA greater commission",
			empA: Employee{
				Name:       "A",
				Commission: 0.15,
				Paid:       0,
			},
			empB: Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       0,
			},
			expectedOutput: "B",
		},
		{
			name: "test happy path - employee A paid less than B",
			empA: Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       5,
			},
			empB: Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       10,
			},
			expectedOutput: "A",
		},
		{
			name: "test happy path - employee A paid more than B",
			empA: Employee{
				Name:       "A",
				Commission: 0.05,
				Paid:       15,
			},
			empB: Employee{
				Name:       "B",
				Commission: 0.10,
				Paid:       10,
			},
			expectedOutput: "B",
		},
	}
	for _, test := range tests {
		res := getEmployeeNameForNextTask(test.empA, test.empB)
		if test.expectedOutput != res {
			t.Errorf("for %s, expected result %s, but got %s", test.name, test.expectedOutput, res)
		}
	}
}

func Test_EmployeeCompleteTask(t *testing.T) {
	tests := []struct {
		name           string
		input          Vehicle
		expectedOutput Assignment
		expectedError  error
	}{
		{
			name: "test happy path - small car no fuel added",
			input: Vehicle{
				Fuel: Fuel{
					Capacity: 50,
					Level:    0.7,
				},
				LicencePlate: "ABC",
				Size:         "small",
			},
			expectedOutput: Assignment{
				LicencePlate: "ABC",
				Employee:     "A",
				FuelAdded:    0,
				Price:        25,
			},
		},
		{
			name: "test happy path - large car with fuel added",
			input: Vehicle{
				Fuel: Fuel{
					Capacity: 50,
					Level:    0.05,
				},
				LicencePlate: "DEF",
				Size:         "large",
			},
			expectedOutput: Assignment{
				LicencePlate: "DEF",
				Employee:     "A",
				FuelAdded:    47.5,
				Price:        118.125,
			},
		},
		{
			name: "test with error - wrong size",
			input: Vehicle{
				Fuel: Fuel{
					Capacity: 50,
					Level:    0.7,
				},
				LicencePlate: "FGH",
				Size:         "truck",
			},
			expectedError: errors.New("unknown size: truck"),
		},
	}
	for _, test := range tests {
		e := Employee{
			Name:       "A",
			Commission: 0.10,
		}

		res, err := e.completeTask(test.input)

		if err != nil {
			if test.expectedError != nil {
				if test.expectedError.Error() != err.Error() {
					t.Error(err)
				}
			} else {
				t.Error(err)
			}
		} else {
			if test.expectedError != nil {
				t.Error(test.expectedError)
			}

			if test.expectedOutput != res {
				t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, res)
			}
		}

	}
}

func Test_getCarFlateRate(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
		expectedError  error
	}{
		{
			name:           "test happy path - small car",
			input:          "small",
			expectedOutput: 25,
		},
		{
			name:           "test happy path - large car",
			input:          "large",
			expectedOutput: 35,
		},
		{
			name:          "test with error - wrong type",
			input:         "truck",
			expectedError: errors.New("unknown size: truck"),
		},
	}
	for _, test := range tests {
		res, err := getCarFlateRate(test.input)

		if err != nil {
			if test.expectedError != nil {
				if test.expectedError.Error() != err.Error() {
					t.Error(err)
				}
			} else {
				t.Error(err)
			}
		} else {
			if test.expectedError != nil {
				t.Error(test.expectedError)
			}

			if test.expectedOutput != res {
				t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, res)
			}
		}

	}
}

func Test_calculateFuelAdded(t *testing.T) {
	tests := []struct {
		name           string
		input          Fuel
		expectedOutput float64
	}{
		{
			name: "test happy path - no refuel needed",
			input: Fuel{
				Capacity: 50,
				Level:    0.5,
			},
			expectedOutput: 0,
		},
		{
			name: "test happy path - less than 10%",
			input: Fuel{
				Capacity: 30,
				Level:    0.05,
			},
			expectedOutput: 28.5,
		},
		{
			name: "test happy path - 10%",
			input: Fuel{
				Capacity: 45,
				Level:    0.05,
			},
			expectedOutput: 42.75,
		},
	}
	for _, test := range tests {
		res := calculateFuelAdded(test.input)

		if test.expectedOutput != res {
			t.Errorf("for %s, expected result %+v, but got %+v", test.name, test.expectedOutput, res)
		}
	}
}
