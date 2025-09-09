# fdecono.com - Personal Website

A modern personal website built with Go and htmx, featuring a clean design and progressive enhancement. This website showcases projects, provides contact functionality, and demonstrates modern web development practices.

## Features

- **Homepage**: Personal bio and featured projects showcase
- **Projects Section**: Dynamic project listings with individual detail pages
- **Project Details**: Individual project pages with full descriptions and tech stacks
- **Contact Form**: Interactive contact form with htmx validation
- **Responsive Design**: Mobile-first design with Tailwind CSS
- **Progressive Enhancement**: Works without JavaScript, enhanced with htmx
- **Health Check**: Built-in health check endpoint for monitoring
- **Static File Serving**: Efficient serving of CSS, images, and JavaScript assets

## Tech Stack

- **Backend**: Go 1.21+ with net/http
- **Frontend**: htmx for progressive enhancement
- **Styling**: Tailwind CSS
- **Templates**: Go html/template
- **Deployment**: Docker support with multi-stage builds

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
│   │   ├── project-detail.html # Individual project template
│   │   ├── contact.html        # Contact page template
│   │   └── contact-form.html   # Contact form partial template
│   └── static/                  # Static assets
│       ├── css/
│       │   └── style.css       # Custom styles
│       ├── js/                 # JavaScript files
│       └── images/
│           └── fd.png          # Profile image
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── Dockerfile                   # Docker configuration
├── deploy.sh                    # Deployment script
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git

### Installation

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

### Development

The server will automatically reload when you make changes to the Go code. For template changes, you'll need to restart the server.

## API Endpoints

- `GET /` - Homepage with featured projects
- `GET /projects` - Projects listing page
- `GET /projects/{id}` - Individual project detail page
- `GET /contact` - Contact form page
- `POST /contact/submit` - Contact form submission
- `GET /health` - Health check endpoint
- `GET /static/*` - Static file serving

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

## Deployment

### Using Docker

The project includes a multi-stage Dockerfile for efficient builds:

1. Build and run with Docker:
```bash
docker build -t fdecono-website .
docker run -p 8080:8080 fdecono-website
```

2. Or use Docker Compose (create `docker-compose.yml`):
```yaml
version: '3.8'
services:
  website:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
```

### Using the Deploy Script

The project includes a deployment script for easy deployment:

```bash
# Run locally
./deploy.sh

# Deploy to production (requires configuration)
./deploy.sh production
```

### Manual Deployment

1. Build the binary:
```bash
go build -o fdecono cmd/server/main.go
```

2. Copy necessary files to your server:
```bash
# Copy binary
scp fdecono user@your-server:/home/user/

# Copy templates and static files
scp -r internal/templates user@your-server:/home/user/
scp -r internal/static user@your-server:/home/user/
```

3. Run on your server:
```bash
./fdecono
```

### Using a Process Manager (PM2)

1. Install PM2:
```bash
npm install -g pm2
```

2. Create `ecosystem.config.js`:
```javascript
module.exports = {
  apps: [{
    name: 'fdecono-website',
    script: './fdecono',
    instances: 1,
    autorestart: true,
    watch: false,
    max_memory_restart: '1G',
    env: {
      NODE_ENV: 'production',
      PORT: 8080
    }
  }]
}
```

3. Start the application:
```bash
pm2 start ecosystem.config.js
```

## Environment Variables

- `PORT`: Server port (default: 8080)

## Development Features

- **Hot Reload**: Go applications automatically reload on code changes
- **Template Caching**: Templates are parsed once for better performance
- **Error Handling**: Comprehensive error handling throughout the application
- **Logging**: Built-in logging for server events and errors

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Test your changes locally
5. Commit your changes (`git commit -m 'Add some amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## Contact

For questions or suggestions, please open an issue or contact me through the website's contact form.
