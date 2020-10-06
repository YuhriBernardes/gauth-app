# Gauth App

A simple app with a memory database.

- [Gauth App](#gauth-app)
  - [Fake users](#fake-users)
  - [Start the app](#start-the-app)
    - [Without docker](#without-docker)
      - [Frontend](#frontend)
      - [Backend](#backend)
    - [With Docker](#with-docker)
      - [Build Images](#build-images)
      - [Start containers](#start-containers)
      - [Stop containers](#stop-containers)
  - [Routes](#routes)
    - [`/api/tokens`](#apitokens)
      - [Request Body (Format: form url encoded)](#request-body-format-form-url-encoded)
        - [Fileds](#fileds)
        - [Example](#example)
      - [Responses](#responses)
        - [Success (Status Code: 200)](#success-status-code-200)
        - [Bad request (Status Code: 400)](#bad-request-status-code-400)
        - [Unauthorized (Status: 403)](#unauthorized-status-403)
    - [Error Body](#error-body)

## Fake users

| Login       | Password | Name                     | Email                                  |
| ----------- | -------- | ------------------------ | -------------------------------------- |
| admin       | admin    | Administrator            | admin@admin.com                        |
| lucca.nunes | easyPass | Lucca Yago Matheus Nunes | luccayagomatheusnunes-75@callan.com.br |

## Start the app

### Without docker

You need to have installed `golang` and `node js`. Keep in mind that you will need one terminal for each application because they block the terminal.

#### Frontend
On the project root run the command below:

```shell
cd frontend && npm install && npm start
```

or

```shell
cd frontend && yarn install && yarn start
```
The frontend app will be running on `localhost:3000`


#### Backend
On the project root run the command below:

```shell
cd backend && go mod download && go run ./main.go
```

The backend app will be running on `localhost:3001`

### With Docker

On the project root follow the instructions below:

The frontend app will be running on `localhost:3000` and the backend app on `localhost:3001`

#### Build Images

```shell
docker-compose -f ./deploy/docker-compose.dev.yml build --force-rm --no-cache
```

#### Start containers

```shell
docker-compose -f ./deploy/docker-compose.dev.yml up -d --remove-orphans
```

#### Stop containers

```shell
docker-compose -f ./deploy/docker-compose.dev.yml down
```

## Routes

You can get the postman collection on `doc/gauth-coll.json` (See how import collections [here](https://learning.postman.com/docs/getting-started/importing-and-exporting-data/#importing-data-into-postman))

Download postman [here](https://www.postman.com/downloads/)

### `/api/tokens`

#### Request Body (Format: form url encoded)

##### Fileds

| name     | type   | required? | description   |
| -------- | ------ | --------- | ------------- |
| login    | string | Yes       | User login    |
| password | string | Yes       | User password |

##### Example

```text
login=admin&password=admin
```

#### Responses

##### Success (Status Code: 200)

```json
{
  "user": {
    "name": "some name",
    "email": "some@email.com",
    "login": "some.login"
  },
  "token": "1237846189237467891263489126348976129783aslduifhaiosfhd9178236y4 "
}
```

##### Bad request (Status Code: 400)

Returns an error structure (see [Errors](#error-body))

##### Unauthorized (Status: 403)

Returns no body. Means that the passoword or login are wrong

### Error Body
Some errors are returned in JSON format following the structure below:

```json
{
  "errMsg": "Error message"
}
```

Wher the `errMsg` field is the message with error specs.