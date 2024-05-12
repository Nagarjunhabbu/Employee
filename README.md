# API Service

This is a simple API service built with Go.

## Installation

### 1. Install Go

Ensure you have Go installed on your system. You can download and install it from the official Go website: [https://golang.org/](https://golang.org/).

### 2. Clone the Repository

Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/api-service.git
```
### 3.Install Docker compose

```bash
docker-compose up
```
Server will start on local host with port 8000.
Once the service is running, you can make HTTP requests to the API. Here's an example using curl:
### 4. Example API Usage
Create:
```bash
curl --location 'localhost:8000/v1/employee' \ --header 'Content-Type: application/json' \ --data '{ "designation": "CEO", "name": "ganesh", "salary": { "salary": 12.23, "currency": "USD" }, "insurance": { "insurance_id": "123", "insurance_no": "hyahjah" } }'
```
Response
```json
{
        "id": 9,
        "name": "ganesh",
        "designation": "CEO",
        "salary": {
            "id": 1,
            "salary": 12.23,
            "currency": "USD",
            "EmployeeId": 9
        },
        "insurance": {
            "id": 1,
            "insurance_id": "123",
            "insurance_no": "hyahjah"
        }
    }

```




```bash
curl --location 'localhost:8000/v1/employee?page=5&page_size=2'
```
Response
```json
[
    {
        "id": 9,
        "name": "ganesh",
        "designation": "CEO",
        "salary": {
            "id": 1,
            "salary": 12.23,
            "currency": "USD",
            "EmployeeId": 9
        },
        "insurance": {
            "id": 1,
            "insurance_id": "123",
            "insurance_no": "hyahjah"
        }
    },
    {
        "id": 10,
        "name": "ganesh",
        "designation": "CEO",
        "salary": {
            "id": 2,
            "salary": 12.23,
            "currency": "USD",
            "EmployeeId": 10
        },
        "insurance": {
            "id": 2,
            "insurance_id": "123",
            "insurance_no": "hyahjah"
        }
    }
]
```


