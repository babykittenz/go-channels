# Check Website Status - Go Program

This Go program checks the availability of a list of websites by making HTTP requests concurrently using goroutines.

## Features

- Uses **goroutines** to check multiple websites concurrently.
- Implements **sync.WaitGroup** to ensure all goroutines complete before exiting.
- Handles potential **network errors** gracefully.

## Prerequisites

- Go 1.16+ installed ([Download Go](https://golang.org/dl/))

## Installation & Running

1. Clone or download this repository.
2. Open a terminal and navigate to the project directory.
3. Run the program using:

   ```sh
   go run main.go
   ```

## Code Explanation

### **Main Function (`main.go`):**

```go
func main() {
    var wg sync.WaitGroup

    for _, link := range links {
        wg.Add(1)
        go func(l string) {
            defer wg.Done()
            checkLink(l)
        }(link)
    }

    wg.Wait()
}
```

- Initializes a **WaitGroup** to keep track of goroutines.
- Iterates over a list of website URLs, spawning a new goroutine for each check.
- Uses a **closure** to pass `link` safely into the goroutine to avoid loop variable capture.
- Waits for all goroutines to complete before the program exits.

### **CheckLink Function (`checkLink`):**

```go
func checkLink(link string) {
    _, err := http.Get(link)
    if err != nil {
        fmt.Println(link, "might be down!")
        return
    }
    fmt.Println(link, "is up!")
}
```

- Sends an HTTP GET request to each link.
- If an error occurs, it prints that the link "might be down!".
- Otherwise, it confirms the site is "up!".

## Sample Output

```
http://google.com is up!
http://facebook.com is up!
http://stackoverflow.com is up!
http://golang.org is up!
http://amazon.com is up!
```

(Results may vary based on network conditions.)

## License

This project is licensed under the MIT License.

## Author

Michael Kidby
