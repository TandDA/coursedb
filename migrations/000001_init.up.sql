CREATE TABLE IF NOT EXISTS building(
	id VARCHAR(255) PRIMARY KEY,
    class smallint,
    number_of_floors smallint,
    address VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS floor(
	id VARCHAR(255) PRIMARY KEY,
    floor_number SMALLINT,
    building_id VARCHAR(255),
    FOREIGN KEY (building_id)  REFERENCES building (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS housekeeping_services(
	id VARCHAR(255) PRIMARY KEY,
    service_name VARCHAR(255),
    floor_id VARCHAR(255),
    FOREIGN KEY (floor_id)  REFERENCES floor (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS housekeeping_service_details(
	id VARCHAR(255) PRIMARY KEY,
    start_time TIME,
	end_time TIME,
    regular_price DECIMAL,
    housekeeping_services_id VARCHAR(255),
    FOREIGN KEY (housekeeping_services_id)  REFERENCES housekeeping_services (id)
);

CREATE TABLE IF NOT EXISTS room(
	id VARCHAR(255) PRIMARY KEY,
	number_of_rooms SMALLINT,
    regular_price DECIMAL,
    floor_id VARCHAR(255),
    FOREIGN KEY (floor_id)  REFERENCES floor (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS guest(
	id VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    date_of_entry DATE
);

CREATE TABLE IF NOT EXISTS complain(
	id VARCHAR(255) PRIMARY KEY,
    complain_text TEXT,
	guest_id VARCHAR(255),
    FOREIGN KEY (guest_id)  REFERENCES guest (id) ON DELETE CASCADE
);

CREATE TABLE firm(
	id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS booking(
	id VARCHAR(255) PRIMARY KEY,
    date_of_entry DATE,
    date_of_departure DATE,
	guest_id VARCHAR(255),
    firm_id VARCHAR(255),
    room_id VARCHAR(255),
    FOREIGN KEY (guest_id)  REFERENCES guest (id) ON DELETE CASCADE,
    FOREIGN KEY (firm_id)  REFERENCES firm (id) ON DELETE CASCADE,
    FOREIGN KEY (room_id)  REFERENCES room (id) ON DELETE CASCADE
);