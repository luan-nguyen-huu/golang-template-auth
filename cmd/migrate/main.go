package main
import (
	"log"
	
	"github.com/luan-nguyen-huu/Adam/internal/entities"
	"github.com/luan-nguyen-huu/Adam/internal/initialize"
)
func main() {
	app := initialize.Init()
	
	err := app.DB.AutoMigrate(
		&entities.User{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	
	log.Println("Migration completed successfully")
}