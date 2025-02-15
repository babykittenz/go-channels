# Link Checker

This is a simple Go program that checks the availability of multiple websites using goroutines and channels for concurrent execution.

## Features

- Concurrently checks multiple websites.
- Uses goroutines to improve efficiency.
- Implements a retry mechanism that checks each link every 5 seconds if it is initially unreachable.

## Prerequisites

- Go installed on your system ([Download Go](https://go.dev/dl/))

## Installation

1. Clone this repository or copy the code into a new directory.
2. Ensure you have Go installed by running:
   ```sh
   go version
   ```
3. Navigate to the directory where `main.go` is located.

## Usage

To run the program, execute the following command:

```sh
 go run main.go
```

## Code Structure

### `main.go`

The entry point of the program, which initializes a list of links and starts the checking process.

### `checkLink(link string, c chan string)`

- Checks if a given link is accessible by making an HTTP GET request.
- If the request fails, it prints that the link might be down and sends the link back into the channel.
- If successful, it prints that the link is up and sends the link back into the channel.

### `checkMultipleLinks(links []string)`

- Initializes a channel for communication between goroutines.
- Iterates over the list of links and starts a goroutine for each link to check its status.
- Uses a loop to continuously monitor the links and recheck them every 5 seconds if they fail.

## Example Output

```
http://google.com is up!
http://facebook.com is up!
http://stackoverflow.com is up!
http://golang.org is up!
http://amazon.com is up!
http://millmountaindigital.com might be down!
```

## Improvements

- Add logging for better debugging.
- Improve error handling and categorize different types of failures.
- Implement a user-defined interval for rechecking.
- Allow dynamic input for websites instead of hardcoded links.

## License

This project is open-source and available under the MIT License.
