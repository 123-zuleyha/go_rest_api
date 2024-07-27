package config 
import(
	"fmt"
	"os"
	"github.com/joho/godotenv" //env dosyasını yüklemek için kullanılan paket 
)

// config den env değerini almak için func ı yapılandırıyoruz
func Config(key string) string{
	 //load .env file
	 err :=godotenv.Load(".env")
	 if err !=nil{
		fmt.Print("Error loading .env file")
	 }
	 return os.Getenv(key)
}