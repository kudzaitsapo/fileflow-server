# PROJECT API SERVER

## APIs

Authentication:

```http
POST /api/v1/auth/login
Request Body: JSON
    email: string
    password: string
```

Upload files to the server:

```http
POST /api/v1/files/upload
Request Body: form-data
    file: file
```

Get all files:

```http
GET /api/v1/files
headers:
    PROJECT_KEY: string
```

Get a single file:

```http
GET /api/v1/files/:id
```

Delete a file:

```http
DELETE /api/v1/files/:id
```

Create Project:

```http
POST /api/v1/projects
Request Body: JSON
    name: string
    description: string
```

Get all projects:

```http
GET /api/v1/projects
```
