package seeds

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)

func generateRoles() []*store.Role {
	return []*store.Role{
		{
			Name:        "admin",
			Description: "Admin role",
			Level:       1,
		},
		{
			Name:        "user",
			Description: "User role",
			Level:       2,
		},
	}
}


func generateAdminUser() (*store.User, error) {

	user := &store.User{
		Email:     "admin@org.com",
		FirstName: "Admin",
		LastName:  "User",
		IsActive:  true,
		RoleID:    1,
	}

	if err := user.Password.Set("secretpassword123"); err != nil {
		return nil, err
	}

	return user, nil
}

func generateDefaultProject() *store.Project {
	return &store.Project{
		Name:        "Default Project",
		Description: "Default project for the organisation",
		ProjectKey: uuid.NewString(),
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}


func Seed(store *store.Storage, db *sql.DB) error {
	log.Println("[+] Seeding database ...")

	ctx := context.Background()

	tx, _ := db.BeginTx(ctx, nil)

	log.Println("[+] Seeding roles ...")
	roles := generateRoles()
	for _, role := range roles {
		if err := store.Roles.Create(ctx, tx, role); err != nil {
			_ = tx.Rollback()
			log.Printf("[-] error seeding role: %v", err)
			return err
		}
	}
	log.Println("[+] Roles seeded successfully")

	log.Println("[+] Seeding admin user ...")
	adminUser, err := generateAdminUser()
	if err != nil {
		log.Printf("[-] error generating admin user: %v", err)
		return err
	}

	if err := store.Users.Create(ctx, tx, adminUser); err != nil {
		_ = tx.Rollback()
		log.Printf("[-] error seeding admin user: %v", err)
		return err
	}else{
		log.Println("[+] Admin user seeded successfully")
	}

	log.Println("[+] Seeding default project ...")
	defaultProject := generateDefaultProject()
	if err := store.Projects.Create(ctx, defaultProject); err != nil {
		_ = tx.Rollback()
		log.Printf("[-] error seeding default project: %v", err)
		return err
	}
	log.Println("[+] Default project seeded successfully")

	tx.Commit()

	log.Printf("[+] Database seeded successfully")

	return nil
}