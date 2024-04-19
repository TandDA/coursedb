CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

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

INSERT INTO building (id, class, number_of_floors, address) VALUES
                                                                ('a0ec50e4-4f9b-467f-872c-e05605e5c40a', 1, 5, '123 Main St'),
                                                                ('34af623b-1d73-43c5-b204-664a15f3c672', 2, 10, '456 Elm St'),
                                                                ('bd7e6892-850f-4b53-9a47-cb96f6b8c0ef', 1, 7, '789 Oak St'),
                                                                ('e2f9f0c7-5889-4548-8b1e-8bc2f98a4a1b', 2, 8, '101 Maple St');


INSERT INTO floor (id, floor_number, building_id) VALUES
                                                      ('b3f05d54-35e8-4cde-a4ef-dc630e1b1ced', 1, 'a0ec50e4-4f9b-467f-872c-e05605e5c40a'),
                                                      ('ca03bb9e-22ba-442c-ad57-a9b0788dd524', 2, 'a0ec50e4-4f9b-467f-872c-e05605e5c40a'),
                                                      ('9ab26d44-47a2-423e-b0b2-72d6a5615d00', 1, '34af623b-1d73-43c5-b204-664a15f3c672'),
                                                      ('982fd7f5-d5c1-4d9a-9459-e90b14b108a7', 2, '34af623b-1d73-43c5-b204-664a15f3c672'),
                                                      ('7b57d3f4-35b0-467e-a257-a725ae4ef9d6', 3, 'bd7e6892-850f-4b53-9a47-cb96f6b8c0ef'),
                                                      ('f4e9b23d-3924-4b6d-b1e3-1b7008368e6a', 4, 'bd7e6892-850f-4b53-9a47-cb96f6b8c0ef');

INSERT INTO housekeeping_services (id, service_name, floor_id) VALUES
                                                                   ('7ef3f621-a015-4bc5-a905-638aabd5773f', 'Cleaning', 'b3f05d54-35e8-4cde-a4ef-dc630e1b1ced'),
                                                                   ('6fda9a34-4b4f-4cd9-a4e9-994d75c2e349', 'Maintenance', 'ca03bb9e-22ba-442c-ad57-a9b0788dd524'),
                                                                   ('1f7c986f-a9fa-4b3c-a6ec-2536cafbc4d2', 'Security', '7b57d3f4-35b0-467e-a257-a725ae4ef9d6');

INSERT INTO housekeeping_service_details (id, start_time, end_time, regular_price, housekeeping_services_id) VALUES
                                                                                                                 ('6b2ed96e-88be-41fe-9397-6742a13b569b', '08:00:00', '12:00:00', 100.00, '7ef3f621-a015-4bc5-a905-638aabd5773f'),
                                                                                                                 ('317a6f77-ac3f-4123-9fe8-8e3920438ec2', '13:00:00', '17:00:00', 150.00, '6fda9a34-4b4f-4cd9-a4e9-994d75c2e349'),
                                                                                                                 ('a2fe90cc-6bcf-4655-81b9-1ea1256574da', '22:00:00', '06:00:00', 180.00, '1f7c986f-a9fa-4b3c-a6ec-2536cafbc4d2');

INSERT INTO room (id, number_of_rooms, regular_price, floor_id) VALUES
                                                                    ('c3986d0d-29a3-49c8-97fe-261a0cf5c5b4', 2, 200.00, 'b3f05d54-35e8-4cde-a4ef-dc630e1b1ced'),
                                                                    ('bf26402e-68e6-4386-8e2d-247c10efa7a9', 3, 300.00, 'ca03bb9e-22ba-442c-ad57-a9b0788dd524'),
                                                                    ('9fa82515-b54f-4047-82bb-494cd8faf13c', 1, 180.00, '7b57d3f4-35b0-467e-a257-a725ae4ef9d6');

INSERT INTO guest (id, first_name, last_name, date_of_entry) VALUES
                                                                 ('a3b8d425-2b60-4ad7-8bf8-7b448b3f8bb2', 'John', 'Doe', '2023-01-10'),
                                                                 ('55f144a3-f5b8-46cb-a707-dc1de385e79d', 'Jane', 'Smith', '2023-01-12'),
                                                                 ('f4bb2658-3922-4d11-9b10-eda33fbf4f75', 'Alice', 'Johnson', '2023-01-14'),
                                                                 ('68de2f62-ee5a-4973-bf2f-9f2935b0e3ea', 'Bob', 'Brown', '2023-01-16');

INSERT INTO complain (id, complain_text, guest_id) VALUES
                                                       ('e3f7b62f-7c82-44e2-91b2-9082f5d50f14', 'Air conditioning not working', 'a3b8d425-2b60-4ad7-8bf8-7b448b3f8bb2'),
                                                       ('d147a2e0-cd3b-4b67-a3a0-a5decf2f980e', 'Noisy neighbors', '55f144a3-f5b8-46cb-a707-dc1de385e79d'),
                                                       ('6a204bd7-f86f-40a4-b3ff-65bd5cf87b72', 'Room not cleaned', 'f4bb2658-3922-4d11-9b10-eda33fbf4f75');

INSERT INTO firm (id, name) VALUES
                                ('b435f447-90a0-4a65-863c-99d0a3a68c81', 'Company XYZ'),
                                ('9255b5d9-b0ee-4c0b-9203-68e42f506dcd', 'Business Inc'),
                                ('d25c7c87-14c1-4b0d-91e6-6a691e752cb1', 'Tech Solutions');

INSERT INTO booking (id, date_of_entry, date_of_departure, guest_id, firm_id, room_id) VALUES
                                                                                           ('6cf143fc-5f10-45ef-916f-d7e5e8a4bdb3', '2023-01-10', '2023-01-15', 'a3b8d425-2b60-4ad7-8bf8-7b448b3f8bb2', 'b435f447-90a0-4a65-863c-99d0a3a68c81', 'c3986d0d-29a3-49c8-97fe-261a0cf5c5b4'),
                                                                                           ('7a2a0f2b-c861-4e90-9417-ecb23b0a59cb', '2023-01-12', '2023-01-20', '55f144a3-f5b8-46cb-a707-dc1de385e79d', '9255b5d9-b0ee-4c0b-9203-68e42f506dcd', 'bf26402e-68e6-4386-8e2d-247c10efa7a9'),
                                                                                           ('fc5e0385-f6a1-4f92-a2db-3fd07b6b7b80', '2023-01-14', '2023-01-18', 'f4bb2658-3922-4d11-9b10-eda33fbf4f75', 'd25c7c87-14c1-4b0d-91e6-6a691e752cb1', '9fa82515-b54f-4047-82bb-494cd8faf13c');
