DROP DATABASE IF EXISTS bezbednost;

CREATE DATABASE bezbednost
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE USER root WITH PASSWORD 'ftn';
GRANT ALL PRIVILEGES ON DATABASE bezbednost TO root;

drop table users;
DROP TABLE IF EXISTS files;

CREATE TABLE users (
    id BIGINT PRIMARY KEY,
	name VARCHAR(255),
    surname VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE files (
   id BIGINT PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   content TEXT,
   owner VARCHAR(255) NOT NULL
);

INSERT INTO users (id, name, surname, email, password) VALUES 
    (1, 'Bobi', 'Bobic', 'bobi@example.com', '$argon2id$v=19$m=65536,t=1,p=8$dKoVtVV9hcq+4/vt/x1bGQ$kFoAJhN3fmNVRDB6krE2KNCZG/Vjtsi7m5khTAaUqoU'),
	(2, 'Ana', 'Cekic', 'ana@example.com', '$argon2id$v=19$m=65536,t=1,p=8$dKoVtVV9hcq+4/vt/x1bGQ$kFoAJhN3fmNVRDB6krE2KNCZG/Vjtsi7m5khTAaUqoU');

GRANT TEMPORARY, CONNECT ON DATABASE bezbednost TO PUBLIC;
GRANT ALL ON DATABASE bezbednost TO postgres;
GRANT ALL ON DATABASE bezbednost TO root;
GRANT ALL PRIVILEGES ON TABLE users TO root;