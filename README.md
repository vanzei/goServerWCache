# Project Idea
This project demonstrates a simple HTTP server in Go that caches user data to avoid hitting the database repeatedly. The server supports concurrent access using a mutex to ensure thread safety.

## Implementation

The server is implemented in `main.go` and consists of the following components:

- **User**: A struct representing a user with an ID and a username.
- **Server**: A struct representing the server with a database, a cache, and a mutex for synchronization.
- **NewServer**: A function that initializes the server with a pre-populated database of 100 users.
- **tryCache**: A method that checks if a user is present in the cache.
- **handleGetUser**: A method that handles HTTP GET requests to retrieve user data. It first checks the cache and then the database if the user is not found in the cache. The user data is then cached for future requests.

### How to test locally

1. Clone the repository:
    ```
    git clone git@github.com:vanzei/goServerWCache.git
    cd goServerWCache
    ```
2. Initialize the Go module:
   ```
   go mod init goServerWCache
   ```
3. Run the tests
   ```
   go test
   ```