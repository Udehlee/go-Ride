CREATE TABLE drivers(
	driver_id SERIAL PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	pass_word VARCHAR(255) NOT NULL,
	ratings INT,
	);