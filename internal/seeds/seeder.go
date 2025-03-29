package seeds

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/kudzaitsapo/fileflow-server/internal/handlers"
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
	key, _ := handlers.GenerateRandomKey()
	return &store.Project{
		Name:        "Default Project",
		Description: "Default project for the organisation",
		ProjectKey:  key,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}
}

func generateFileTypes() []*store.FileType {
	filetypes := make([]*store.FileType, 0)

	filetypes = append(filetypes, &store.FileType{
		Name:        "PDF File",
		MimeType:    "application/pdf",
		Description: "Portable Digital File is a document format",
		Icon: `<svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="20"
                      height="20"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    >
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                      <polyline points="14 2 14 8 20 8"></polyline>
                      <path d="M9 15h6"></path>
                      <path d="M9 11h6"></path>
                    </svg>`,
	}, &store.FileType{
		Name:        "PNG Image File",
		MimeType:    "image/png",
		Description: "Portable Network Graphics image file",
		Icon: `<svg
					  xmlns="http://www.w3.org/2000/svg"
					  width="20"
					  height="20"
					  viewBox="0 0 24 24"
					  fill="none"
					  stroke="currentColor"
					  strokeWidth="2"
					  strokeLinecap="round"
					  strokeLinejoin="round"
					>
					  <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
					  <circle cx="8.5" cy="8.5" r="1.5"></circle>
					  <polyline points="21 15 16 10 5 21"></polyline>
					</svg>`,
	}, &store.FileType{
		Name:        "JPEG Image File",
		MimeType:    "image/jpeg",
		Description: "Joint Photographic Experts Group image file",
		Icon: `<svg
					  xmlns="http://www.w3.org/2000/svg"
					  width="20"
					  height="20"
					  viewBox="0 0 24 24"
					  fill="none"
					  stroke="currentColor"
					  strokeWidth="2"
					  strokeLinecap="round"
					  strokeLinejoin="round"
					>
					  <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
					  <circle cx="8.5" cy="8.5" r="1.5"></circle>
					  <polyline points="21 15 16 10 5 21"></polyline>
					</svg>`},
		&store.FileType{
			Name:        "Microsoft Excel Spread sheet",
			MimeType:    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
			Description: "Microsoft Excel Spread sheet file",
			Icon: `<svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="20"
                      height="20"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                    >
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                      <polyline points="14 2 14 8 20 8"></polyline>
                      <line x1="16" y1="13" x2="8" y2="13"></line>
                      <line x1="16" y1="17" x2="8" y2="17"></line>
                      <line x1="10" y1="9" x2="8" y2="9"></line>
                    </svg>`,
		},
		&store.FileType{
			Name:        "Microsoft Word Document",
			MimeType:    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
			Description: "Microsoft Word Document file",
			Icon: `<svg 
						xmlns="http://www.w3.org/2000/svg" 
						viewBox="0 0 24 24" 
						width="20"
                        height="20"
						fill="none" 
						stroke="currentColor" 
						strokeWidth="2" 
						strokeLinecap="round" 
						strokeLinejoin="round">
							<!-- Document outline -->
							<path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
							
							<!-- Folded corner -->
							<polyline points="14 2 14 8 20 8"></polyline>
							
							<!-- Word "W" logo -->
							<path d="M7.5 13L9 17.5L10.5 13L12 17.5L13.5 13L15 17.5L16.5 13"></path>
					</svg>`,
		})

	return filetypes
}

func generateAllowedFileType(projectId int64, fileTypeId int64) *store.ProjectAllowedFileType {
	return &store.ProjectAllowedFileType{
		ProjectID:  projectId,
		FileTypeID: fileTypeId,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func Seed(storage *store.Storage, db *sql.DB) error {
	log.Println("[+] Seeding database ...")

	ctx := context.Background()

	tx, _ := db.BeginTx(ctx, nil)

	roleCount, countErr := storage.Roles.Count(ctx)

	if countErr != nil {
		log.Printf("[-] Failed to count roles! An error occurred: %v", countErr)
	} else {
		if roleCount == 0 {
			log.Println("[+] Seeding roles ...")
			roles := generateRoles()
			for _, role := range roles {
				if err := storage.Roles.Create(ctx, tx, role); err != nil {
					_ = tx.Rollback()
					log.Printf("[-] error seeding role: %v", err)
					return err
				}
			}
			log.Println("[+] Roles seeded successfully")
		} else {
			log.Println("[!] Skip seeding roles since there's live data")
		}
	}

	userCount, userCountErr := storage.Users.Count(ctx)
	adminUserId := int64(0)

	if userCountErr != nil {
		log.Printf("[-] Failed to count users! An error occurred: %v", userCountErr)
	} else {
		if userCount == 0 {
			log.Println("[+] Seeding admin user ...")
			adminUser, err := generateAdminUser()
			if err != nil {
				log.Printf("[-] error generating admin user: %v", err)
				return err
			}

			if err := storage.Users.Create(ctx, tx, adminUser); err != nil {
				_ = tx.Rollback()
				log.Printf("[-] error seeding admin user: %v", err)
				return err
			} else {
				adminUserId = adminUser.ID
				log.Println("[+] Admin user seeded successfully")
			}
		} else {
			log.Println("[!] Skip seeding admin user since there's live data")
		}
	}

	projectCount, projectCountErr := storage.Projects.Count(ctx)
	var defaultProjectId int64

	if projectCountErr != nil {
		log.Printf("[-] Failed to count projects! An error occurred: %v", projectCountErr)
	} else {
		if projectCount == 0 {
			log.Println("[+] Seeding default project ...")
			defaultProject := generateDefaultProject()
			if err := storage.Projects.Create(ctx, defaultProject); err != nil {
				_ = tx.Rollback()
				log.Printf("[-] error seeding default project: %v", err)
				return err
			}
			defaultProjectId = defaultProject.ID
			log.Println("[+] Default project seeded successfully")
		} else {
			log.Println("[!] Skip seeding default project since there's live data")
		}
	}

	fileTypesCount, fileTypesCountErr := storage.FileTypes.Count(ctx)
	if fileTypesCountErr != nil {
		log.Printf("[-] Failed to count file types! An error occurred: %v", fileTypesCountErr)
	} else {
		if fileTypesCount == 0 {
			log.Println("[+] Seeding file types ...")
			fileTypesToSeed := generateFileTypes()

			for _, fileType := range fileTypesToSeed {
				if ftSeedErr := storage.FileTypes.Create(ctx, fileType); ftSeedErr != nil {
					_ = tx.Rollback()
					log.Printf("[-] error seeding file type: %v", ftSeedErr)
					return ftSeedErr
				}

				if defaultProjectId != 0 {
					log.Printf("[+] Seeding allowed file type for default project ...")
					defaultProjectAllowedType := generateAllowedFileType(defaultProjectId, fileType.ID)
					if patCreateErr := storage.ProjectAllowedFileTypes.Create(ctx, defaultProjectAllowedType); patCreateErr != nil {
						_ = tx.Rollback()
						log.Printf("[-] error seeding project allowed file type: %v", patCreateErr)
						return patCreateErr
					}
				}
			}
			log.Println("[+] File types seeded successfully")
		} else {
			log.Println("[!] Skip seeding file types since there's live data")
		}
	}

	assignmentsCount, assignCountErr := storage.UserAssignedProjects.CountByUserId(ctx, adminUserId)
	if assignCountErr != nil {
		log.Printf("[-] Failed to count user assignments! An error occurred: %v", assignCountErr)
	} else if assignmentsCount == 0 {
		if (adminUserId != 0) && (defaultProjectId != 0) {
			log.Printf("[+] Assigning default project to admin user with Id: %d", adminUserId)
			adminUserProject := &store.UserAssignedProject{
				ProjectID: defaultProjectId,
				UserID:    adminUserId,
			}

			if assignmentErr := storage.UserAssignedProjects.Create(ctx, tx, adminUserProject); assignmentErr != nil {
				_ = tx.Rollback()
				log.Printf("[-] error assigning default project to admin user: %v", assignmentErr)
				return assignmentErr
			}
			log.Println("[+] Default project assigned to admin user successfully")

		}
	}

	tx.Commit()

	log.Printf("[+] Database seeded successfully")

	return nil
}
