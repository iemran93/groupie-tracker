# Groupie Trackers

![My Skills](https://skillicons.dev/icons?i=go,html,css,htmx,js)

## About
Groupie Trackers is a web application that allows users to explore information about artists, locations, dates, and relations related to music events. It fetches data from an API and presents it in a user-friendly manner.

## Features
- View a list of artists and their details.
- Explore locations and their associations with artists.
- Access information about specific dates and events.
- Search the database by artist/band, member, concert or date.
- GeoMap in the artist page showing the locations of the concerts.

## Technologies Used
- Go: Go is the programming language used for the backend server.
- HTML/CSS: HTML and CSS are used for frontend web page rendering.
- HTTP: The application interacts with the API using HTTP requests.
- Templates: HTML templates are used to render the fetched data.
- HTMX: Used to view(render) the search results in the html template.
- Google maps API : used the API to get the geolocation and the map
- JS: Javascript to render the google map api and adds the markers based on the geolocation.

## Setup
1. Clone the repository:
   git clone https://learn.reboot01.com/git/emarei/groupie-tracker

2. Change to the project directory:
   cd groupie-trackers

3. Run the application:
   ```go
   go run main.go
   ```

4. Open your web browser and access the application at http://localhost:8000

## Usage
- Navigate to the different sections of the application using the provided links or navigation menu.
- Explore the artists, locations, dates, and relations to gain insights into music events.

## Author
Omar Abdulraheem @oabulra
Emran Marei @emarei
