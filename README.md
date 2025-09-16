# fdecono.com - Personal Website

A modern personal website built with Go and htmx, featuring a clean design and progressive enhancement. This website showcases projects, provides contact functionality, and demonstrates modern web development practices.

## Features

- **Homepage**: Personal bio and featured projects showcase
- **Projects Section**: Dynamic project listings with individual detail pages
- **Responsive Design**: Mobile-first design with custom CSS
- **Progressive Enhancement**: Works without JavaScript, enhanced with htmx
- **Health Check**: Built-in health check endpoint for monitoring
- **Static File Serving**: Efficient serving of CSS, images, and JavaScript assets
- **Docker Deployment**: Containerized deployment with nginx proxy

## Tech Stack

- **Backend**: Go 1.21+ with net/http
- **Frontend**: htmx for progressive enhancement
- **Styling**: Custom CSS with 8-bit/pixel art theme
- **Templates**: Go html/template
- **Deployment**: Docker with multi-stage builds
- **Web Server**: nginx with SSL/TLS
- **Hosting**: AWS Lightsail

## Project Structure

```
fdecono.com/
├── cmd/
│   └── server/
│       └── main.go              # Main server entry point
├── internal/
│   ├── handlers/                # HTTP handlers
│   │   ├── home.go             # Homepage handler
│   │   ├── projects.go         # Projects listing and detail handlers
│   │   └── contact.go          # Contact form handlers
│   ├── models/                  # Data models
│   │   └── project.go          # Project data structure and sample data
│   ├── templates/               # HTML templates
│   │   ├── base.html           # Base template with common layout
│   │   ├── home.html           # Homepage template
│   │   ├── projects.html       # Projects listing template
│   │   └── project-detail.html # Individual project template
│   └── static/                  # Static assets
│       ├── css/
│       │   └── style.css       # Custom styles
│       ├── js/                 # JavaScript files
│       └── images/
│           └── fd.png          # Profile image
├── .github/
│   └── workflows/
│       └── deploy.yml          # GitHub Actions deployment
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── Dockerfile                   # Docker configuration
├── docker-compose.yml           # Docker Compose configuration
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- Docker (for containerized deployment)
- Git

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/fdecono/fdecono.com
cd fdecono.com
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the development server:
```bash
go run cmd/server/main.go
```

4. Open your browser and visit `http://localhost:8080`

### Docker Development

1. Build the Docker image:
```bash
docker build -t fdecono.com .
```

2. Run with Docker:
```bash
docker run -p 8080:8080 fdecono.com
```

3. Or use Docker Compose:
```bash
docker-compose up -d
```

## Production Deployment

### Docker Deployment Process

This project uses Docker for production deployment with the following architecture:

- **Docker Container**: Runs the Go application
- **nginx**: Reverse proxy with SSL/TLS termination
- **AWS Lightsail**: Cloud hosting platform

#### Complete Deployment Flow

The deployment process follows these steps:

1. **Local Image Creation** ️
2. **Image Serialization**
3. **File Transfer** 🚀
4. **Server-Side Image Loading** 📥
5. **Container Orchestration** 🎯
6. **nginx Proxy Configuration** 🌐

#### Step 1: Local Image Creation

```bash
# On your local machine
docker build --platform linux/amd64 -t fdecono.com .
```

**What happens here:**
- Docker reads the `Dockerfile`
- Creates a multi-stage build:
  - **Stage 1 (Builder)**: Downloads Go dependencies, compiles your Go code
  - **Stage 2 (Runtime)**: Creates minimal Alpine Linux image with just the binary and static files
- Builds for `linux/amd64` architecture (compatible with your Lightsail server)
- Tags the final image as `fdecono.com`

**Result**: A Docker image containing your compiled Go application + static files

#### Step 2: Image Serialization

```bash
# On your local machine
docker save fdecono.com | gzip > fdecono.com.tar.gz
```

**What happens here:**
- `docker save` converts the Docker image into a tar archive
- `gzip` compresses the tar file to reduce size
- Creates `fdecono.com.tar.gz` (typically 10-50MB compressed)

**Why this approach:**
- **No registry needed**: Don't need Docker Hub or private registry
- **Self-contained**: Everything needed is in one file
- **Version control**: Each deployment has a unique file
- **Offline deployment**: Works without internet on server

#### Step 3: File Transfer

```bash
# From local machine to server
scp -i ~/.ssh/<KEY_NAME>.pem fdecono.com.tar.gz <USER>@<STATIC_IP>:<HOME_PATH>
```

**What happens here:**
- `scp` (Secure Copy Protocol) transfers the file over SSH
- Uses your private key for authentication
- Copies the compressed image to `<HOME_PATH>` on the server
- Transfer time: Usually 30 seconds to 2 minutes depending on file size

#### Step 4: Server-Side Image Loading

```bash
# On the server
docker load < fdecono.com.tar.gz
```

**What happens here:**
- `docker load` reads the tar.gz file
- Extracts the Docker image layers
- Registers the image in Docker's local image store
- Image is now available as `fdecono.com` on the server

**Docker image layers:**
```
fdecono.com
├── Layer 1: Alpine Linux base (~5MB)
├── Layer 2: ca-certificates (~1MB)
├── Layer 3: Your Go binary (~10MB)
├── Layer 4: Static files (~1MB)
└── Layer 5: User permissions (~1KB)
```

#### Step 5: Container Orchestration

```bash
# On the server
docker stop fdecono.com || true
docker rm fdecono.com || true
docker run -d --name fdecono.com -p 8080:8080 --restart unless-stopped fdecono.com
```

**What happens here:**
- **Graceful shutdown**: Stops existing container if running
- **Cleanup**: Removes old container to free resources
- **New deployment**: Starts new container with updated image
- **Port mapping**: Maps server port 8080 to container port 8080
- **Auto-restart**: Container restarts if it crashes

#### Step 6: nginx Proxy Configuration

```nginx
# nginx configuration
location / {
    proxy_pass http://127.0.0.1:8080;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
}

location /static/ {
    proxy_pass http://127.0.0.1:8080/static/;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
}
```

**What happens here:**
- **External requests** → nginx (port 443/80)
- **nginx** → Docker container (port 8080)
- **Static files** → Proxied to container's static file handler
- **SSL termination** → nginx handles HTTPS, forwards HTTP to container

### Visual Deployment Flow

```
┌─────────────────┐    ┌──────────────┐    ┌─────────────────┐
│   Local Machine │    │   Network    │    │  Lightsail      │
│                 │    │              │    │  Server         │
│ 1. Build Image  │───▶│              │───▶│                 │
│ 2. Save to tar  │    │              │    │ 3. Load Image   │
│ 3. Compress     │    │              │    │ 4. Run Container│
│ 4. Upload       │    │              │    │ 5. nginx Proxy  │
└─────────────────┘    └──────────────┘    └─────────────────┘
```

### Why This Architecture?

#### **Benefits:**

1. **Isolation**: App runs in its own environment
2. **Consistency**: Same image runs everywhere
3. **Rollbacks**: Easy to switch to previous version
4. **No dependencies**: Server doesn't need Go, just Docker
5. **Resource control**: Docker manages memory/CPU limits
6. **Security**: Containerized application is isolated

#### **File Sizes:**

```
Source code:        ~1MB
Docker image:       ~20MB (uncompressed)
Compressed tar:     ~8MB (gzipped)
Transfer time:      ~30 seconds
Deployment time:    ~2 minutes total
```

### Resume Subdomain

The project also includes a separate resume page:

- **URL**: `https://resume.fdecono.com`
- **File**: `fd.html` (static HTML)
- **Configuration**: Separate nginx virtual host

## Container Management

### Docker Commands

```bash
# View running containers
docker ps

# View logs
docker logs fdecono.com

# Restart container
docker restart fdecono.com

# Stop container
docker stop fdecono.com

# Start container
docker start fdecono.com

# Update to new version
docker stop fdecono.com
docker rm fdecono.com
docker load < fdecono.com.tar.gz
docker run -d --name fdecono.com -p 8080:8080 --restart unless-stopped fdecono.com
```

### Docker Compose Commands

```bash
# Start services
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f

# Restart services
docker-compose restart

# Update services
docker-compose up -d --build
```

## API Endpoints

- `GET /` - Homepage with featured projects
- `GET /projects` - Projects listing page
- `GET /health` - Health check endpoint
- `GET /static/*` - Static file serving
- `GET /favicon.ico` - Favicon

## Adding New Projects

To add a new project, edit `internal/models/project.go` and add a new `Project` struct to the `GetProjects()` function:

```go
{
    ID:          "your-project-id",
    Title:       "Your Project Title",
    Description: "Short description",
    LongDesc:    "Detailed description of your project...",
    GitHubURL:   "https://github.com/yourusername/your-project",
    LiveURL:     "https://your-project.com", // optional
    TechStack:   []string{"Go", "React", "PostgreSQL"},
    ImageURL:    "/static/images/your-project.png", // optional
    CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
    Featured:    true, // or false
},
```

## Environment Variables

- `PORT`: Server port (default: 8080)

## Troubleshooting

### Common Issues

1. **Static files not loading**: Ensure nginx is proxying `/static/` requests to the Docker container
2. **Architecture mismatch**: Build with `--platform linux/amd64` for Linux servers
3. **Container restarting**: Check logs with `docker logs fdecono.com`
4. **Permission issues**: Ensure proper file ownership in Docker container

### Debugging Commands

```bash
# Check container status
docker ps -a

# View container logs
docker logs fdecono.com

# Execute commands in container
docker exec -it fdecono.com sh

# Check static files in container
docker exec -it fdecono.com ls -la /app/internal/static/

# Test static file serving
curl http://localhost:8080/static/css/style.css
```

### Architecture Mismatch Error

```bash
# Error: exec format error
# Solution: Build with correct platform
docker build --platform linux/amd64 -t fdecono.com .
```

### Static Files Not Loading

```bash
# Error: CSS/JS not loading
# Solution: nginx proxy configuration
location /static/ {
    proxy_pass http://127.0.0.1:8080/static/;
}
```

### Container Won't Start

```bash
# Check logs
docker logs fdecono.com

# Check if port is available
netstat -tlnp | grep :8080
```

## Alternative Deployment Approaches

### Docker Registry Approach

```bash
# Push to registry
docker tag fdecono.com your-registry.com/fdecono.com:latest
docker push your-registry.com/fdecono.com:latest

# Pull on server
docker pull your-registry.com/fdecono.com:latest
```

### Docker Compose Approach

```bash
# Copy docker-compose.yml + source code
# Build on server
docker-compose up -d --build
```

## Development Features

- **Hot Reload**: Go applications automatically reload on code changes
- **Template Caching**: Templates are parsed once for better performance
- **Error Handling**: Comprehensive error handling throughout the application
- **Logging**: Built-in logging for server events and errors
- **Health Checks**: Built-in health check endpoint for monitoring
- **Docker Support**: Full containerization with multi-stage builds
