## Ride Matching System

### Overview

A simplified ride-matching system that assigns the closest driver to a passenger based on distance. It uses a priority queue (min-heap) to prioritize the nearest drivers and a worker pool concurrency pattern to handle multiple ride requests concurrently.

### Features

- Matches passengers with the nearest available drivers.

- Uses a min-heap priority queue for sorting of drivers by distance.

- Implements a worker pool to process multiple ride-matching requests concurrently.

- saves successful matched rides to database



### Technologies Used

- Go (Gin) 

- Postgres

### Api Endpoints

```sh
POST /request-a-ride	
```
```sh
POST /add-driver
```
### Example Request

- request a ride
```sh
{
	passenger_id : 1
	"passenger_name: "Ada"
	lat:  40.7128,
	lon: -74.0060
}	
```

### Example Response
```sh
{
 driver_id : 2,
 passenger_id : 1,
 ride_status: "matched"
 created_at: 2025-03-05 14:30:15.123456789 +0000 UTC
}
```

### Installation

- Clone the repository:

```sh 
git clone https://github.com/Udehlee/go-Ride.git 
```
```sh
cd go-Ride
 ```
- Install dependencies 
```sh
go mod tidy
```
- Run the project:
```sh
go run main.go
```





