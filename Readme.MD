# ShortCut Link API Server ✂🔗

CutLink is a lightweight and efficient URL shortener application built using Go, Redis, and Docker. This project aims to provide a simple and scalable solution for creating and managing shortened URLs.

## Features
- **Fast and Scalable**: Powered by Go and Redis for optimal performance
- **Containerized Setup**: Easily deployable using Docker and Docker Compose
- **Custom Slugs**: Generate and use custom or random slugs for shortened URLs
- **RESTful API**: Clear and consistent APIs for URL shortening, retrieval, and analytics
- **Cloud-Ready**: Designed for deployment on AWS or other cloud platforms

## Technologies

- **Backend**: Go
- **Database**: Redis (key-value store)
- **Containerization**: Docker & Docker Compose

## Getting Started

### Prerequisites

Ensure the following tools are installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Clone the Repository

```bash
git clone https://github.com/andresilvase/cutlink.git
cd cutlink
```

### Build and Run the Application

1. Build and start the services:

```bash
make up
```

2. Access the application at `http://localhost:8080`.

## API Endpoints

- **POST /cut**  
  Creates a shortened URL  
  Request Body:  
  ```json
  {
    "url": "https://example.com/something-here?maybe=here-too"
  }
  ```
  Response:  
  ```json
  {
    "data": {
      "shortened_url": "https://example.com/{slug}"
    }
  }
  ```

- **GET /{slug}**  
  Redirects to the original URL

## Deployment

The production deployment was made in AWS **EC2**.