# TO DO Lists done here

## TODO LIST v1.0212344109103

- [x] Implement fetching all projects on API
- [x] Implement fetching all files uploaded to a project on the backend
- [x] Implement file list on the front end
- [x] Implement fetching file info based on file id
- [x] Instead of using int64 id, implement uuids for stored_file primary key
- [x] Implement project settings on the front-end - includes
  - [x] File based icons
  - [x] Allowed file types / file validation
  - [x] Maximum / minimum file size allowed
- [x] Implement file resource / mime type matching => Define mime types and file extensions
- [x] User Management - front-end template
- [x] Implement User management on the backend
- [x] Auth Login / Logout on front-end + backend (integrated)
- [x] Front end - Initialise front end
- [x] Project management - create project on the front end
- [x] Fix seeding of items so that if something exists, it's not seeded instead of breaking the backend initialisation
- [x] Update project creation API to accept file size limits
- [x] Update file upload API to validate settings from project
- [x] Seed mime types
- [x] Update project settings - backend
- [x] Update project settings - front-end
- [x] Implement API key re-generation
- [x] Update backend API for project creation to allow mime types at project creation / settings update
- [x] Implement API to fetch available mime types
- [x] Implement front-end fetching of available mime types
- [x] Refactor API to use wrapped responses
- [x] Refactor API to return pagination information
- [x] Implement pagination on the front-end for files list

## TODO LIST some other version

- [ ] Refactor the front-end approach to use project assignments and project based viewing
- [ ] Refactor backend API to accept project id in the headers to filter data
- [x] Implement project based file viewing & user management
- [ ] Implement Project deletion & file deletion
- [ ] Implement File download on the front-end
- [ ] Implement File deletion on the front-end
- [x] Implement Security template (IP Address white listing / Firewall) on the front-end
- [ ] Implement User project allocation on the backend on creation
- [x] Implement User creation on backend
- [ ] Implement User creation on front-end
- [ ] Refactor project settings to use tabs, and split security details into a new tab
- [ ] Implement redis caching (figure out what to cache)
- [ ] Implement email notifications on the backend
- [ ] Implement email settings on the front-end
- [ ] Find a way to implement working hot reload for the backend
- [ ] Implement system notifications on the backend
- [ ] Implement system notifications on the front-end
- [x] Refactor authorization to send JSON instead of a string message
- [x] Re-design side bar for project settings
- [ ] Implement user permissions and roles
- [ ] Implement audit logs for access (login, file uploads, CRUD operations etc)
- [ ] Implement file versioning ? (Not sure if this is necessary to be honest)
- [ ] Implement usage & analytics reporting template on the front-end
- [ ] Implement user profile editing, adding of avatars, etc
- [ ] Vault ? To think about
