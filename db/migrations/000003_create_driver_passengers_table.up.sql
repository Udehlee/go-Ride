CREATE TABLE driverPassenger(
		id SERIAL PRIMARY KEY,
		driverId INT NOT NULL,
		passengerId INT NOT NULL,
		FOREIGN KEY (driverId) REFERENCES drivers(drivers_id) ON DELETE SET NULL,
		FOREIGN KEY (passengerId) REFERENCES passengers(passenger_id) ON DELETE SET NULL,
	);