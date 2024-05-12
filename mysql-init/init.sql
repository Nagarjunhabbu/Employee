CREATE DATABASE IF NOT EXISTS employeedb;

USE employeedb;

CREATE TABLE IF NOT EXISTS employee (
                                        id INT AUTO_INCREMENT PRIMARY KEY,
                                        name VARCHAR(255),
    designation VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS employee_insurance (
                                                  id INT AUTO_INCREMENT PRIMARY KEY,
                                                  employee_id INT NOT NULL,
                                                  insurance_id VARCHAR(255) NOT NULL,
    insurance_no VARCHAR(255) NOT NULL,
    insurance_exp DATE NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employee(id)
    );

CREATE TABLE IF NOT EXISTS employee_salary (
                                               id INT AUTO_INCREMENT PRIMARY KEY,
                                               employee_id INT NOT NULL,
                                               salary DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    FOREIGN KEY (employee_id) REFERENCES employee(id)
    );
