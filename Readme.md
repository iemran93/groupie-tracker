# Groupie Trackers

![My Skills](https://skillicons.dev/icons?i=go,html,css,htmx,js)

## About
Groupie Trackers is a web application that allows users to explore information about artists, locations, dates, and relations related to music events. It fetches data from an API and presents it in a user-friendly manner.

<img src="readmesrc\main page.jpg" width=550px>

## Features
- View a list of artists and their details.
- Explore locations and their associations with artists.
- Access information about specific dates and events.
- Search the database by artist/band, member, concert or date.
- Filter artists by memeber, dates and locations.
- GeoMap in the artist page showing the locations of the concerts.

## Technologies Used
- Go: Go is the programming language used for the backend server.
- HTML/CSS: HTML and CSS are used for frontend web page rendering.
- HTTP: The application interacts with the API using HTTP requests.
- Templates: HTML templates are used to render the fetched data.
- HTMX: Used to view(render) the search&filter results in the html template.
- Google maps API : used the API to get the geolocation and the map
- JS: Javascript to render the google map api and adds the markers based on the geolocation.

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/iemran93/groupie-tracker
   ```

2. Change to the project directory:
   ```bash
   cd groupie-tracker
   ```

3. Run the application:
   ```go
   go run main.go
   ```

4. Open your web browser and access the application at http://localhost:8000

5. To enable the map in the artist page, get your google api key and set it in the following parts:
- handlers.go
- artist.html (js part)
## Usage
- Navigate to the different sections of the application using the provided links or navigation menu.
- Explore the artists, locations, dates, and relations to gain insights into music events.

## Author
[emarei](https://github.com/iemran93), [Omar Abdulraheem]()

