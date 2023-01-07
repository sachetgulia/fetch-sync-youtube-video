# fetch-sync-youtube-video
Clone the project
Install Go in your system as Project is in Golang
Need to install Mysql With database name 'pro_db' and User created named 'sachet' and password '1234' and grant all permissions
Install all dependencies using go mod tidy
Run go run main.go to run server


Search api curl = 
curl --location --request GET 'http://localhost:8080/search' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"TOP 10 MOST GOAL PLAYERS OF ALL TIME IN FOOTBAL #football",
    "description":"TOP 10 MOST GOAL PLAYERS OF ALL TIME IN FOOTBAL #football Here is a list of the top 10 most goal-scoring football players ...",
    "limit":4
}'


