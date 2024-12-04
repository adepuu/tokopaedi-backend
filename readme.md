# Tokopaedi Backend

This project is for fun. It is written in Go, using GORM for ORM, Fiber for the web framework, and PostgreSQL as the database. Docker is used to manage the PostgreSQL instance.

## Prerequisites

- Go 1.21.3 or later
- Docker
- Docker Compose

## Setup

1. Clone the repository:

    ```sh
    git clone https://github.com/adepuu/tokopaedi-backend.git
    cd tokopaedi-backend
    ```

2. Copy the `.env.example` file to `.env` and adjust the environment variables if necessary:

    ```sh
    cp .env.example .env
    ```

### .env Example

    ```env
    POSTGRES_DB=postgres
    POSTGRES_HOST=0.0.0.0
    POSTGRES_PASSWORD=admin
    POSTGRES_PORT=5433
    POSTGRES_USER=admin

    APP_PORT=8080
    ```

3. Start the PostgreSQL database using Docker Compose:

    ```sh
    docker-compose up -d
    ```

4. Install the Go dependencies:

    ```sh
    go mod tidy
    ```

5. Run the application:

    ```sh
    go run main.go
    ```

## Project Structure

- `go.mod` - Go module file
- `go.sum` - Go dependencies checksum file
- `compose.yml` - Docker Compose file for PostgreSQL
- `.env` - Environment variables file
- `.env.example` - Example environment variables file
- `.gitignore` - Git ignore file

## Dependencies

- [GORM](https://gorm.io/)
- [Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.