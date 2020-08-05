# Dumka Go
Before trying to run the application, please rename `.env.example` to `.env` and config environmental variables

# DB structure

Check out `assets/db.sql` for the full SQL code (tested on `MariaDB`).
Visualisation of the database is located [here](assets/db.png).
# API Endpoints

These examples were taken from projects mainly using [Django Rest
Framework](https://github.com/tomchristie/django-rest-framework) and so the
JSON responses are often similar to the way in which DRF makes responses.

Where full URLs are provided in responses they will be rendered as if service
is running on 'http://testserver/'.

## Open Endpoints

Open endpoints require no Authentication.

* `POST /v1/status` - _API Status_ ([Docs](assets/api/status.http) | [Code](src/api/status.go))
* `POST /v1/schools` - _Schools_  ([Docs](assets/api/schools.http) | [Code](src/api/schools.go))
* `POST /v1/report_categories` - _Report Categories_ ([Docs](assets/api/report_categories.http) | [Code](src/api/report_categories.go))
* `POST /v1/auth` - _Authentication_ ([Docs](assets/api/report_categories.http) | [Code](src/api/report_categories.go))

