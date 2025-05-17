// envloader/envloader.go
package envloader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Load loads key=value pairs from a .env file into the environment.
func Load(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Could not open .env file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments and blank lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading .env file: %v", err)
	}
}
