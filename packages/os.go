package main
import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()

	host := os.Getenv("DB_HOST")
	fmt.Println("Database Host:", host)
	fmt.Println("All arguments:", os.Args)
}