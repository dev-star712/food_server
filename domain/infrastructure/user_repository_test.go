package infrastructure

import (
	"fmt"
	"food-app/database/rdbms"
	"food-app/domain/entity"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func init() {
	if err := godotenv.Load(os.ExpandEnv("./../../.env")); err != nil {
		log.Println("no env gotten")
	}
}

func Database() (*gorm.DB, error) {
	dbdriver := os.Getenv("TEST_DB_DRIVER")
	host := os.Getenv("TEST_DB_HOST")
	password := os.Getenv("TEST_DB_PASSWORD")
	user := os.Getenv("TEST_DB_USER")
	dbname := os.Getenv("TEST_DB_NAME")
	port := os.Getenv("TEST_DB_PORT")

	conn, err := rdbms.NewDBConnection(dbdriver, user, password, port, host, dbname)
	if err != nil {
		return nil, err
	}
	err = conn.DropTableIfExists(&entity.User{}).Error
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(
		entity.User{},
	).Error
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func TestUserRepo_SaveUser(t *testing.T) {
	conn, err := Database()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = entity.User{}
	user.ID = 1
	user.Email = "manaan@gmail.com"
	user.FirstName = "Kedu"
	user.LastName = "Manner"
	user.Password = "password"

	repo := NewUserRepository(conn)

	u, saveErr := repo.SaveUser(&user)
	if saveErr != nil {
		t.Fatalf("want non error, got %#v", saveErr)
	}
	fmt.Println(u)
}
