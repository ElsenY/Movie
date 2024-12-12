# How to run the app
1. clone the repository
2. run `go mod download && go mod vendor`
3. create a .env file in the same directory level as main.go, fill in the value based on the .env.example file
4. run `go run main.go` to start the API server on port 8081
   
# Fill in database
for the database, I'm using postgresql, to create the table and populate the data, can run this query : 

***Create movies table***

`CREATE TABLE movies(
	id SERIAL PRIMARY KEY NOT NULL,
	title VARCHAR,
	description VARCHAR,
	duration VARCHAR,
	artists VARCHAR,
	genre VARCHAR,
	watchURL VARCHAR,
	vote int,
	viewcount int);`
 <br>
 <br>
 <br>
 ***Populate the table with random data (50 data)***
 
 `WITH random_data AS (
    SELECT
        'Movie ' || floor(random() * 10000) AS title,
        'Description ' || floor(random() * 10000) AS description,
        (floor(random() * (180 - 90 + 1)) + 90) || ' minutes' AS duration,
        'Artist ' || floor(random() * 10000) AS artists,
        CASE
            WHEN random() < 0.2 THEN 'Action'
            WHEN random() < 0.4 THEN 'Comedy'
            WHEN random() < 0.6 THEN 'Drama'
            WHEN random() < 0.8 THEN 'Horror'
            ELSE 'Romance'
        END AS genre,
        'https://example.com/watch/' || floor(random() * 10000) AS watchURL,
        floor(random() * (1000 - 1 + 1)) + 1 AS vote,
        floor(random() * (10000 - 1000 + 1)) + 1000 AS viewcount
    FROM generate_series(1, 50)
);`

# Calling the endpoint
I have made a public collection in [postman](https://www.postman.com/maintenance-architect-99534403/elsen-public/collection/vxjqpvo/movie) that include all the basic required endpoints

# Tools used
1. Gin for endpoint
2. pq lib for postgresql connection
3. Postgresql database
4. Postman for endpoint testing
