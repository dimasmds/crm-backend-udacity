# CRM Backend for Udacity Go Course

Final Project of Udacity Go Fundamental Course

## Prerequisites

Before running this project, make sure you have the following installed on your system:

- Go (at least Go 1.13)

## Installation

1. Clone this repository to your local machine:

   ```shell
   git clone https://github.com/dimasmds/crm-backend-udacity.git
   ```

2. Change the working directory to the project folder:

   ```shell
   cd crm-backend-udacity
   ```

3. Install project dependencies using the `go get` command:

   ```shell
   go get ./...
   ```

## Running the Project

To run the project, execute the following command:

```shell
go run main.go
```

This will start the project, and you can access it at [http://localhost:3000](http://localhost:3000) in your web
browser.

## Usage
| HTTP Method | URL Path        | Description                 |
|-------------|-----------------|-----------------------------|
| GET         | /customers      | Get a list of all customers |
| GET         | /customers/{id} | Get a single customer       |
| POST        | /customers      | Add a new customer          |
| PUT         | /customers/{id} | Update an existing customer |
| DELETE      | /customers/{id} | Delete a customer           |

## Package in this project
| Name     | Path             | Description                                                        |
|----------|------------------|--------------------------------------------------------------------|
| Data     | `./pkg/data`     | Contain function that handle all CRUD operations                   |
| Handlers | `./pkg/handlers` | Contain function that describe logic serving HTTP request-response |
| Routes   | `./pkg/routes`   | Contain function that registering all available routes in HTTP     |
| Main     | `./main.go`      | Entry point of the application                                     |

## Library that Used In This Project (Dependencies)
- 	github.com/google/uuid
-   github.com/gorilla/mux