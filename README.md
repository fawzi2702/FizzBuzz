# FizzBuzz

This REST API provides a customizable Fizzbuzz, and a stats system.

_Features:_

- Fizzbuzz with custom parameters
- Save Fizzbuzz requests stats
- Get the top requested fizzbuzz parameters

## Prerequisites

To build and run this project you need to have installed:

- Make
- Docker
- Docker Compose

## Installation

Create an `.env` file with the following content:

```bash
API_PORT=           # The port where the API will run
API_PORT_CONTAINER= # The port where the API will be exposed in the container
REDIS_HOST=         # A string to identify the Redis instance
REDIS_PORT=         # The Redis port (by default 6379)
REDIS_DB_INDEX=     # The Redis database index (by default 0)
```

You can use the .env.example file as a template.

## Build & Run

To build and run the project, execute the following command:

**For Production:**

```bash
make
```

or

**For Development:**

_No restart always and Log rotation_

```bash
make dev
```

Once it's done, check that everything is correct by running:

```bash
make ps
```

## Usage

### Fizzbuzz

To get a Fizzbuzz, you need to make a GET request to the `/fizzbuzz` endpoint with the following query parameters:

- `int1`: The first integer to replace by str1
- `int2`: The second integer to replace by str2
- `limit`: The limit of the Fizzbuzz
- `str1`: The first string to replace by int1 multiples
- `str2`: The second string to replace by int2 multiples

Example:

```bash
curl -X GET "http://127.0.0.1:<API_PORT>/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz"
```

### Stats

To get the top requested fizzbuzz parameters, you need to make a GET request to the `/stats` endpoint.

Example:

```bash
curl -X GET "http://127.0.0.1:<API_PORT>/stats"
```

## Swagger

You can access the Swagger documentation by visiting the `/docs/index.html` endpoint.
