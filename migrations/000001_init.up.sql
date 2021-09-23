CREATE TABLE users
(
    id serial NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE images (id serial NOT NULL UNIQUE, img bytea);

CREATE TABLE directions
(
    id serial NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL
);
CREATE TABLE communication
(
    id serial NOT NULL UNIQUE,
    VK VARCHAR(255) NOT NULL,
    WhatsApp VARCHAR(255) NOT NULL
);

CREATE TABLE student
(
    id serial NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    sur_name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    ressult TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    image_id int DEFAULT 0 REFERENCES images (id) on delete CASCADE NOT NULL,
    date_register date NOT NULL,
    complete_direction_id int[] DEFAULT 0 REFERENCES directions (id) on delete CASCADE NOT NULL,
    student_representative_id int DEFAULT 0 REFERENCES student_representative (id) on delete CASCADE NOT NULL,
    communication_id int DEFAULT 0 REFERENCES communication (id) on delete CASCADE NOT NULL
);
CREATE TABLE group
(
    id serial NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    direction int DEFAULT 0 REFERENCES directions (id) on delete CASCADE NOT NULL,
    timetable time NOT NULL,
    students_id int DEFAULT 0 REFERENCES student (id) on delete CASCADE NOT NULL
);

CREATE TABLE student_representative 
(
    id serial NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    sur_name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    place_of_work TEXT NOT NULL,
    communication_id int DEFAULT 0 REFERENCES communication (id) on delete CASCADE NOT NULL
);


CREATE TABLE teacher 
(
    id serial NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    sur_name VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    salary int NOT NULL,
    group_id int[] DEFAULT 0 REFERENCES group (id) on delete CASCADE NOT NULL,
    communication_id int DEFAULT 0 REFERENCES communication (id) on delete CASCADE NOT NULL
);
