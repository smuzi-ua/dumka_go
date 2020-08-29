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

* `POST` `/v1/status` - _API Status_ ([Example](assets/api/route_open/status.http) | [Code](src/route/route_open/status.go))
* `POST` `/v1/schools` - _List of schools_  ([Example](assets/api/route_open/schools.http) | [Code](src/route/route_open/schools.go))
* `POST` `/v1/report_categories` - _List of report Categories_ ([Docs](assets/api/route_open/report_categories.http) | [Code](src/route/route_open/report_categories.go))
* `POST` `/v1/auth` - _Authentication_ ([Example](assets/api/route_open/auth.http) | [Code](src/route/route_open/auth.go))

### Token Endpoints

* `POST` `/v1/u/user` - _User Information (`name`, `nickname`, `school`, `type`)_ ([Example](assets/api/route_user/user.http) | [Code](src/route/route_user/user.go))
* `POST` `/v1/u/school` - _School information (list of students & name of the school)_ ([Example](assets/api/route_user/school.http) | [Code](src/route/route_user/school.go))
* `POST` `/v1/u/proposals` - _List of proposals_ ([Example](assets/api/route_user/proposals.http) | [Code](src/route/route_user/proposals.go))
* `POST` `/v1/u/proposals_add` - _Create a new proposal_ ([Example](assets/api/route_user/proposals_add.http) | [Code](src/route/route_user/proposals_add.go))
