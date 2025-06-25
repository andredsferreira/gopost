# GOPOST

## Description

GoPost is a simple full-stack web application i made that allows users to make simple posts and to explore other user's posts.
The application also supports user authentication via password login. It functions by saving sensitive data in a JWT and storing it in a cookie.
This app's purpose is purely academic, i didn't focus much on UI/UX but rather on correct implementation of the backend API, and writing idiomatic go code.
It was deployed on my uni's cloud, an alpine container running NGINX. The configuration was manually performed by me which also allowed me to gain skills in NGINX and basic deployment.

## Technologies used

- net/http package from the go standard library to run the HTTP server and handle  the endpoints.

- html/template package to render HTML templates from the server.

- HTMX to communicate with the backend and exchange HTML.

- golang-jwt package for JWT authentication.

- database/sql go package for data persistance in a MySQL database (deployed in my college cloud).

## Packages

### auth
Contains the code that handles authentication logic such as generating JWT, or signing them.

### data
Contains struct types and methods for that communicate with the database for persistence.

### db
Contains database infrastructure logic.

### middleware
Middleware used in the application.

### static
Static directory for css and js.

### templates
HTML templates/views that are rendered on the server and served to the client.

## Authentication in detail

The authentication works via JWT. The application stores the JWT in a cookie along with the CSRF token. Then, to validate authentication it compares the JWT with the CSRF stored in the database.

## Deployment in detail

The deployment was made in a LXC alpine container. I configured key-based ssh access to the machine and installed and configured NGINX as a reverse proxy, pointing the exposed domain and port (vs-gate.dei.isep.ipp.pt:10245) to the internal app one (localhost:8080)
