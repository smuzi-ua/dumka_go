# Dumka Go
Before trying to run the application, please rename `.env.example` to `.env` and config environmental variables.

If you want to run HTTP files from `assets/api`, make sure to change `http-client.env.json`

## Database

[SQL File (MariaDB)](assets/db.sql) (`assets/db.sql`)

[Database visualization](assets/db.png) (`assets/db.png`)

## API Endpoints

The json values are not sorted properly. I am very sorry for that :>

#### Authorization

After authorizing using `/v1/auth`, you will receive `token` as one of the variables.
Pass the token as a query parameter to every endpoint that needs authorization.

### Open Endpoints

Open endpoints require no Authentication.

* `POST` `/v1/status` - _API Status_ ([Example](assets/api/status.http) | [Code](src/route/status.go))
* `POST` `/v1/schools` - _Schools_  ([Example](assets/api/schools.http) | [Code](src/route/schools.go))
* `POST` `/v1/report_categories` - _Report Categories_ ([Docs](assets/api/report_categories.http) | [Code](src/route/report_categories.go))
* `POST` `/v1/auth` - _Authentication_ ([Example](assets/api/auth.http) | [Code](src/route/auth.go))

### Token Endpoints

* `POST` `/v1/u/user` - _User Information_ ([Example](assets/api/user.http) | [Code](src/route/user.go))
* `POST` `/v1/u/school` - _School information (name & users)_ ([Example](assets/api/school.http) | [Code](src/route/school.go))
* `POST` `/v1/u/proposals` - _Proposals_ ([Example](assets/api/proposals.http) | [Code](src/route/proposals.go))
