This is Api service with comments and posts CRUD's and user registration with email or Google, jwt-token and oauth2 for security

Service uses echo minimalistic web framework, swagger to visualize and interact with the API’s resources.

To download the project use "git clone https://github.com/SsilenzZ/Api"

You need to have mysql or any database that could be connected with gorm drivers, if you want to change database from mysql you should go to pkg/db/connect.go and setup connection.

To set up web service you should change .env variables for yours.

Default starting page "http://localhost:8000/swagger/index.html"
