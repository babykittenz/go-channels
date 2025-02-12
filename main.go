package main

func main() {

	for _, link := range links {
		wg.Add(1) // Increment the WaitGroup counter
		go func(l string) {
			defer wg.Done() // Decrement counter when goroutine finishes
			checkLink(l)
		}(link)
	}

	wg.Wait() // Wait for all goroutines to finish before exiting
}
