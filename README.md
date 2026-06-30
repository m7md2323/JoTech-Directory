# Jordan Tech Companies Directory

A web application that serves as a comprehensive directory and portfolio of tech companies in Jordan. This project allows users to browse, search, and manage profiles of Jordanian technology companies and related tech events.

🔗 **Portfolio Page:** [View Project Portfolio](https://m7md2323.github.io/Portfolio/pages/jotech_directory.html)

## ✨ Features

- **Company Directory:** View an extensive list of tech companies operating in Jordan.
- **Detailed Profiles:** Access comprehensive information about each company, including their branches, links, and contact info.
- **Search & Filter:** Easily find specific companies using the built-in search and filtering functionality.
- **Events Management:** Browse upcoming tech events and add new ones.
- **Content Management:** Add, edit, and delete company profiles directly from the UI.
- **Server-Side Rendered:** Fast, SEO-friendly rendering utilizing Go templates.

## 🛠 Tech Stack

**Backend & Database:**
- [Go](https://golang.org/) (Standard `net/http` library)
- [SQLite](https://www.sqlite.org/) (Lightweight database engine)
- [GORM](https://gorm.io/) (Object Relational Mapper)

**Frontend:**
- [HTMX](https://htmx.org/) (Dynamic interactions without writing complex JavaScript)
- [Templ](https://templ.guide/) (Type-safe HTML templating for Go)
- [Bootstrap](https://getbootstrap.com/) + Custom CSS (Styling and responsive layout)

**Hosting and Deployment:**
- Docker Image (Containerization)
- Railway Hosting (Deployment)

## 📦 Prerequisites

Before running the project locally, ensure you have the following installed:
- [Go](https://go.dev/doc/install) (1.20 or higher recommended)
- [Templ](https://templ.guide/quick-start/installation) CLI (for compiling templates)
- GCC/CGO (Required for SQLite compilation on some systems)

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/m7md2323/Jordan-Tech-Companies.git
cd Jordan-Tech-Companies
```

### 2. Environment Setup

Create a `.env` file in the root directory by copying the provided `.env.example` file (`cp .env.example .env`) and filling in your API keys if you plan to use the AI scraper. The application will use default database and upload paths if omitted.

> [!NOTE]
> The SQLite database (`.db` file) is intentionally included in the repository! It comes pre-populated with tech companies data, so you don't need to run any migrations or seed scripts to get started.

### 3. Generate Templates

If you make any changes to the `.templ` files in the `web/` directory, you need to compile them into Go code:

```bash
templ generate
```

### 4. Install Dependencies

Download the required Go modules:

```bash
go mod tidy
```

### 5. Run the Application

Start the development server:

```bash
go run ./cmd/server/main.go
```

The application will be accessible at `http://localhost:8080`.

### 6. Run with Docker

Alternatively, you can run the application fully containerized using Docker without needing Go installed locally. This ensures a consistent environment and prevents version mismatches.

```bash
# Build the Docker image
docker build -t jordan-tech-companies .

# Run the container
docker run -p 8080:8080 jordan-tech-companies
```

The application will be accessible at `http://localhost:8080`.

## 📂 Project Structure

- `cmd/server/main.go`: The main entry point for the application. It initializes the database, configures routes, and starts the server.
- `internal/`: Contains core backend logic including database initialization (`internal/database`), models, HTTP handlers (`internal/handlers`), and background services.
- `web/`: Contains all frontend code, including static assets (`web/static`) and Templ components/pages (`web/templates`).

## 📄 License
This project is open-source. Please see the [LICENSE](LICENSE) file for more information.
