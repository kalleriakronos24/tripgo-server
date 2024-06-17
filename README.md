how to use this server application

steps to reproduce:

1. clone this repository

2. run "cp .env.example .env" and fills your specified

- DB_HOST
- DB_PORT
- DB_USERNAME
- DB_PASSWORD
- DB_DATABASE
  ...
  and all the over env variable availables.

into your .env file

3. run docker compose build

4. after docker compose build success, run docker compose up

5. now the api url is available at http://localhost:8000/api/v1/xxxxx

if you encounter issues when u run the command "docker compose up" containing texts "go.sum", try running "go mod tidy" in your project directory and try run "docker compose up" once again

API Documentation is at http://localhost:8080/api/v1/swagger/index.html


Notes:

Standard validation for data values in this project are:

- E16.4 Phone Number formatting
- ISO 8601 for Date formatting
- UUID v4 formatting for any tables ID. currently v7 are still on experimental, not recommended to use


