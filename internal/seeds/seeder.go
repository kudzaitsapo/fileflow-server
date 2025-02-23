package seeds

import (
	"context"
	"database/sql"
	"log"

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
	log.Println("[V] Roles seeded successfully")

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
		log.Println("[V] Admin user seeded successfully")
	}

	tx.Commit()

	log.Printf("[V] Database seeded successfully")

	return nil
}