# GOPOST

## Description

Simple project to learn web development in go. It's a basic full stack application, with server side rendering.

It allows users to make simple posts with an associated category, it supports authentication via JWT's.

## Technologies used

- net/http package from the go standard library to run the HTTP server and handle  the endpoints.

- html/template package to render HTML templates from the server.

- htmx to communicate with the backend and exchange HTML.

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
