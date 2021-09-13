# Challenge Immutable

Parking service system

## Requirements

go installed

## Run the application

go run main.go

## Unit test

go test .

## Improvements

- Make the code generic to handle a list of employees
- Optimise the profit

## Example of output (JSON format)

```json
[
  {
    "licencePlate": "A",
    "employee": "A",
    "fuelAdded": 53.01,
    "price": 127.7675
  },
  {
    "licencePlate": "B",
    "employee": "B",
    "fuelAdded": 0,
    "price": 35
  },
  {
    "licencePlate": "C",
    "employee": "B",
    "fuelAdded": 0,
    "price": 35
  },
  {
    "licencePlate": "D",
    "employee": "B",
    "fuelAdded": 0,
    "price": 35
  },
  {
    "licencePlate": "E",
    "employee": "A",
    "fuelAdded": 0,
    "price": 35
  },
  {
    "licencePlate": "F",
    "employee": "B",
    "fuelAdded": 51.3,
    "price": 124.77499999999999
  },
  {
    "licencePlate": "G",
    "employee": "A",
    "fuelAdded": 53.2,
    "price": 118.10000000000001
  },
  {
    "licencePlate": "H",
    "employee": "A",
    "fuelAdded": 0,
    "price": 25
  },
  {
    "licencePlate": "I",
    "employee": "A",
    "fuelAdded": 0,
    "price": 25
  },
  {
    "licencePlate": "J",
    "employee": "B",
    "fuelAdded": 62.37,
    "price": 144.14749999999998
  }
]
```
