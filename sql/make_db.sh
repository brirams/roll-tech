#!/usr/bin/env bash

CREATE DATABASE IF NOT EXISTS roll_tech_db;

CREATE TABLE IF NOT EXISTS alumni(
    a_id int not null auto_increment,
    name VARCHAR(50) CHARACTER SET utf8 NOT NULL,
    year SMALLINT NOT NULL,
    occupation varchar(100) CHARACTER SET utf8,
    phone varchar(20) CHARACTER SET utf8,
    email varchar(50) CHARACTER SET utf8,
    location varchar(100) CHARACTER SET utf8,
    hobbies TINYTEXT CHARACTER SET utf8,
    talents  TINYTEXT CHARACTER SET utf8,
    interests  TINYTEXT CHARACTER SET utf8,
    PRIMARY KEY(a_id)
);

