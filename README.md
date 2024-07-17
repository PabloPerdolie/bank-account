# Bank Account API 

This project is a REST API that simulates the operation of an ATM. The application supports account creation, 
balance replenishment, withdrawal and balance verification operations. 
All operations are logged to the console. Goroutines and channels are used to process operations.

## Features

- **Creating a new account**
- **Replenishment of the balance**
- **Withdrawal of funds**
- **Checking the balance**

## Architecture

The project follows the principles of pure architecture, separating business logic, infrastructure and interfaces. 
A service provider is used for dependency management and dependency injection.

## Technologies Used

- **Golang**: Backend service implemented in Go.
- **Docker**: Containerization for easy deployment and scalability.
- **Gorilla Mux**: HTTP request router for Go.
- **Sync** : Package in Go provides primitives for managing competitiveness, such as mutexes, conditional variables, and wait groups
- **Channels** : Channels in Go allow gorutins to securely exchange data and synchronize.

## Setup Instructions

Follow these steps to set up and run the Bank Account API locally:

### Clone the repository

```bash
git clone https://github.com/PabloPerdolie/bank-account
cd bank-account
```
### Launching the application

```bash
make run
```

The service will start running on http://localhost:8080/.

#### To clean up:
```bash
make clean
```

## Endpoints

You can use Postman to test this API.

#### Create New Account

- **URL**: `/accounts`
- **Method**: `POST`
- **Response**: 
  - `201 Created` New account created.
- **Sample response**:
    ````json
    {
      "ID": 1,
      "balance": 0
    }
    ````

#### Replenishment of the balance

- **URL**: `/accounts/{id}/deposit`
- **Method**: `POST`
- **Request Body**:
    ````json
    {
      "amount": 100.0
    }
    ````
- **Response**: 
  - `200 OK`: The balance has been successfully replenished.
  - `400 Bad Request`: Invalid request.


#### Withdrawal of funds

- **URL**: `/accounts/{id}/withdraw`
- **Method**: `POST`
- **Request Body**:
    ````json
    {
      "amount": 50.0
    }
    ````
- **Response**:
  - `200 OK`: Funds have been successfully withdrawn.
  - `400 Bad Request`: Invalid request.

#### Balance check


- **URL**: `/accounts/{id}/balance`
- **Method**: `GET`
  - **Response**:
      - `200 OK`: Funds have been successfully withdrawn.
        - Sample response:
        ````json
        {
            "balance": 50.0
        }
        ````
      - `400 Bad Request`: Invalid request.




<center>Thanks for checking out my service.</center>
