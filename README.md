### Project: Task Management API with User Authentication & Advanced Features

**Description**: Build a RESTful API for task management with user authentication, database integration, file uploads, and real-time features. This project will help you explore:

1. **Core Language Features**:

   - Structs and interfaces
   - Error handling
   - Concurrency (goroutines/channels)
   - Pointers and memory management
   - Package management with Go Modules

2. **Web Concepts**:

   - REST API design
   - Middleware implementation
   - JWT authentication
   - Rate limiting
   - Request validation

3. **Database & Storage**:

   - PostgreSQL integration
   - Database migrations
   - Object-Relational Mapping (ORM) with GORM
   - File storage (AWS S3 or local storage)

4. **Advanced Features**:
   - Redis caching
   - WebSocket implementation
   - Background workers
   - Docker containerization
   - Unit/Integration testing

---

### **Key Features to Implement**

1. **User Authentication System**

   - JWT-based authentication
   - Password hashing (bcrypt)
   - Refresh tokens
   - Role-based access control (Admin/User)

2. **Task Management**

   - CRUD operations for tasks
   - Task filtering/sorting/pagination
   - File attachments for tasks
   - Task comments system

3. **Advanced Functionality**

   - Real-time updates using WebSocket
   - Daily summary emails (use goroutines)
   - Rate limiting middleware
   - Request validation
   - Structured logging (Zap or Logrus)
   - Prometheus metrics endpoint

4. **Infrastructure**
   - Dockerize application
   - PostgreSQL database
   - Redis for caching
   - Configuration management (Viper)
   - Migrations (Goose or GORM AutoMigrate)

---

### **Project Structure Example**

```bash
/task-manager-api
  ├── cmd/
  │   └── main.go          # Entry point
  ├── internal/
  │   ├── config/          # Configuration setup
  │   ├── handlers/        # HTTP handlers
  │   ├── middleware/      # Custom middleware
  │   ├── models/          # Database models
  │   ├── repositories/    # Database operations
  │   ├── services/        # Business logic
  │   ├── utils/           # Helper functions
  │   └── worker/          # Background jobs
  ├── migrations/          # SQL migration files
  ├── pkg/                 # Reusable packages
  ├── docker-compose.yml   # Local services
  ├── Dockerfile
  ├── Makefile             # Common tasks
  └── README.md            # Documentation
```

---

### **Learning Opportunities**

1. **Concurrency**:

   - Use goroutines for sending emails async
   - Implement a worker pool for task processing
   - Use channels for cross-service communication

2. **Error Handling**:

   - Custom error types
   - Central error handling middleware
   - Database transaction rollbacks

3. **Testing**:

   - Unit tests for business logic
   - Integration tests for API endpoints
   - Mock database implementations

4. **Performance**:
   - Connection pooling
   - Query optimization
   - Caching strategies
   - Load testing with Vegeta

---

### **Tech Stack Suggestions**

- **Web Framework**: Chi Router (lightweight) or Gin
- **Database**: PostgreSQL + GORM
- **Cache**: Redis (go-redis)
- **Authentication**: JWT (golang-jwt)
- **Config**: Viper
- **Logging**: Zap
- **Testing**: Testify, GoMock
- **Documentation**: Swagger (swaggo)

---

### **Deployment & Bonus Points**

1. Containerize with Docker
2. Add CI/CD pipeline (GitHub Actions)
3. Create API documentation with Swagger
4. Implement health check endpoints
5. Add Prometheus metrics
6. Deploy to cloud platform (AWS EC2/DigitalOcean)

---

### **Project Evolution**

1. Start with basic CRUD operations
2. Add authentication layer
3. Implement file uploads
4. Add Redis caching
5. Create background workers
6. Add WebSocket support
7. Containerize and deploy

Here's a detailed breakdown of the Task Management system components:

---

### **Core Resources & Database Structure**

**1. Users Table**  
```go
type User struct {
    ID           uuid.UUID  `gorm:"type:uuid;primary_key"`
    Email        string     `gorm:"unique;not null"`
    PasswordHash string     `gorm:"not null"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    Role         string     // "admin" or "user"
    Tasks        []Task     // One-to-Many relationship
}
```

**2. Tasks Table**  
```go
type Task struct {
    ID          uuid.UUID  `gorm:"primary_key"`
    Title       string     `gorm:"not null"`
    Description string
    Status      string     // "todo", "in_progress", "done"
    Priority    int        // 1-5
    DueDate     time.Time
    CreatedAt   time.Time
    UpdatedAt   time.Time
    UserID      uuid.UUID  // Foreign key
    Comments    []Comment  // One-to-Many
    Attachments []Attachment
}
```

**3. Comments Table**  
```go
type Comment struct {
    ID        uuid.UUID
    Content   string
    CreatedAt time.Time
    TaskID    uuid.UUID
    UserID    uuid.UUID
}
```

**4. Attachments Table**  
```go
type Attachment struct {
    ID        uuid.UUID
    FileName  string
    FilePath  string    // S3 URL or local path
    UploadedAt time.Time
    TaskID    uuid.UUID
}
```

---

### **Key API Endpoints**

#### **Authentication**
- `POST /api/v1/signup` - User registration
- `POST /api/v1/login` - JWT token generation
- `POST /api/v1/refresh` - Refresh access token
- `POST /api/v1/logout` - Token invalidation

#### **Tasks**
- `POST /api/v1/tasks` - Create new task
- `GET /api/v1/tasks` - List tasks (with filters)
- `GET /api/v1/tasks/{id}` - Get single task
- `PUT /api/v1/tasks/{id}` - Update task
- `DELETE /api/v1/tasks/{id}` - Delete task
- `GET /api/v1/tasks/search?q=...` - Full-text search

#### **Comments**
- `POST /api/v1/tasks/{id}/comments` - Add comment
- `GET /api/v1/tasks/{id}/comments` - List comments
- `DELETE /api/v1/comments/{id}` - Delete comment

#### **Attachments**
- `POST /api/v1/tasks/{id}/attachments` - Upload file
- `GET /api/v1/attachments/{id}` - Download file
- `DELETE /api/v1/attachments/{id}` - Remove file

#### **Admin**
- `GET /api/v1/admin/users` - List all users (admin only)
- `PUT /api/v1/admin/users/{id}/role` - Change user role
- `DELETE /api/v1/admin/users/{id}` - Delete user

---

### **Email Content Examples**

**1. Daily Summary Email**  
```
Subject: Your Daily Task Summary - [Date]

Hello [Username],

Here's your task update:
- Completed tasks: 3
- Overdue tasks: 2
- Upcoming deadlines (next 3 days): 4

Urgent Tasks:
1. [Task Title] - Due Tomorrow
2. [Task Title] - Overdue by 2 days

View your dashboard: [Link]
```

**2. Task Assignment Notification**  
```
Subject: New Task Assigned: "[Task Title]"

Hi [Username],

You've been assigned a new task:
- Title: [Task Title]
- Priority: High
- Due Date: [Date]
- Description: [First 50 characters...]

View task: [Task URL]
```

---

### **Key Features in Action**

1. **Search & Filtering**  
   ```http
   GET /api/v1/tasks?status=done&priority=3&sort=-due_date&page=2&limit=10
   ```

2. **File Upload Flow**  
   ```
   Client -> POST /tasks/{id}/attachments (multipart/form-data)
   Server -> Store file (S3/local) -> Save metadata to DB
   ```

3. **Real-Time Updates**  
   WebSocket endpoint `/ws/v1/tasks` broadcasts:
   - Task status changes
   - New comments
   - File uploads

4. **Background Email Worker**  
   Uses Redis queue + goroutine worker pool to:
   - Send daily summaries at 8 AM user timezone
   - Handle failed email retries
   - Process 100+ emails concurrently

---

### **Advanced Database Operations**
- Soft deletes (archive instead of delete)
- Full-text search with PostgreSQL tsvector
- Composite indexes for common queries
- Database transactions for complex operations
- Connection pooling with pgx

---

### **Tech Stack Deep Dive**
- **Router**: Chi (with middleware chain)
- **ORM**: GORM (with prepared statements)
- **Cache**: Redis for frequent queries
- **Auth**: JWT with refresh token rotation
- **Storage**: AWS S3 SDK for Go
- **Real-Time**: Gorilla WebSocket
- **Email**: SMTP with templating (html/template)

This structure gives you exposure to real-world patterns like:
- Clean architecture separation
- Dependency injection
- Context propagation
- Graceful shutdown
- Structured logging
- Distributed tracing (OpenTelemetry)

### **How to add swagger documentation**

1. Install the following packages:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```
2. Run ``` swag init cmd/main.go``` to generate the swagger documentation
