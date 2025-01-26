
# Golang URL Shortener

This project is a URL shortening service developed in Go (Golang). Its primary purpose is to convert long URLs into shorter links, making them easier to use and share.

## Features

- **URL Shortening**: Convert long URLs into short links.
- **Redirection**: Redirect users from short links to the original URLs.
- **Click Tracking**: Maintain statistics on the number of clicks for shortened links.
- **Expiration**: Set expiration dates for shortened links.
- **Link Deletion**: Delete shortened links upon request.

## Requirements
- Go (version specified in `go.mod`)
- Docker and Docker Compose

## Installation

1. Clone the Repository:

   ```sh
   git clone https://github.com/altynbek07/golang-url-shortener.git
   ```

2. Navigate to the project directory:
   ```sh
   cd golang-url-shortener
   ```

3. Create a `.env` file based on the `.env.example`:
   ```sh
    cp .env.example .env
    ```

4. Set your secret key in the `.env` file (you can generate secret **Encryption key 256** and copy [here](https://acte.ltd/utils/randomkeygen)):
    ```env
    SECRET="your_secret_key"
    ```

5. Set your database dsn address in the `.env` file:
    ```env
    DSN="host=localhost user=postgres password=my_pass dbname=link port=5432 sslmode=disable"
    ```

3. **Run the Application with Docker:**

   Ensure Docker and Docker Compose are installed, then execute:

   ```sh
   docker-compose up --build
   ```

   This will deploy the application along with required services like the database.


## Project Structure

- **cmd**: Contains the entry point for the application.
- **configs**: Configuration files for the application.
- **internal**: Internal packages and business logic, including request handling and database management.
- **migrations**: Database migration files.
- **pkg**: Shared packages used across the application.


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
