# Contributing to Jordan Tech Companies Directory

Thank you for your interest in contributing! This guide will walk you through everything you need to get started.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Ways to Contribute](#ways-to-contribute)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Development Workflow](#development-workflow)
- [Coding Conventions](#coding-conventions)
- [Submitting a Pull Request](#submitting-a-pull-request)
- [Reporting Issues](#reporting-issues)

---

## Code of Conduct

Be respectful, inclusive, and constructive in all interactions. We welcome contributors of all backgrounds and experience levels.

---

## Ways to Contribute

- 🐛 **Bug reports** — Found something broken? Open an issue.
- 💡 **Feature suggestions** — Have an idea? Open a discussion or issue first before writing code.
- 🏢 **Company data** — Know a Jordanian tech company that's missing? Add it via the UI or submit a data PR.
- 🌐 **Translations / localization** — Help make the directory accessible to Arabic speakers.
- 📝 **Documentation** — Improve the README, add inline comments, or fix typos.
- 🔧 **Code improvements** — Bug fixes, refactors, performance improvements, new features.

---

## Getting Started

### Prerequisites

Make sure you have the following installed before cloning:

| Tool | Version | Notes |
|------|---------|-------|
| [Go](https://go.dev/doc/install) | 1.20+ | Core language |
| [Templ CLI](https://templ.guide/quick-start/installation) | Latest | For compiling `.templ` files |
| GCC / CGO | Any | Required on Windows for SQLite (`gcc` via [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or [MSYS2](https://www.msys2.org/)) |

### 1. Fork & Clone

```bash
git clone https://github.com/<your-username>/Jordan-Tech-Companies.git
cd Jordan-Tech-Companies
```

### 2. Set Up Environment

Copy the example env file and fill in any values you need:

```bash
cp .env.example .env
```

> [!NOTE]
> The `GEMINI_API_KEY` and `JINA_API_KEY` are only required if you plan to use the AI web scraper service (`internal/services/scraper.go`). For most contributions, you can leave them as-is.

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Compile Templates

If you change or add any `.templ` files, regenerate the Go output:

```bash
templ generate
```

### 5. Run the Application

```bash
go run ./cmd/server/main.go
```

The app will be available at **http://localhost:8080**.

> [!NOTE]
> The SQLite database (`internal/database/jordanTechCompanies.db`) is included in the repo and comes pre-populated with company data. No migrations or seed scripts needed.

---

## Project Structure

```
Jordan-Tech-Companies/
├── cmd/server/
│   └── main.go               # Entry point — server setup and route registration
├── internal/
│   ├── database/
│   │   ├── db.go             # DB connection, AutoMigrate, and query functions
│   │   └── *.db              # SQLite database file (included in repo)
│   ├── handlers/             # HTTP handlers (one file per route group)
│   │   ├── getCompanies.go
│   │   ├── postAddCompany.go
│   │   └── ...
│   ├── models/               # GORM models and shared types
│   │   ├── company.go
│   │   └── helper.go
│   └── services/             # Background services (e.g., AI scraper)
│       ├── scraper.go
│       └── image.go
└── web/
    ├── static/               # CSS, JS, images, uploaded logos
    └── templates/            # Templ components and pages
        ├── pages/            # Full-page templates
        └── partials/         # Reusable UI fragments (used with HTMX)
```

### Key Conventions

- **Handlers** live in `internal/handlers/`. File names follow the pattern `<method><Resource>.go` (e.g., `getCompanies.go`, `postAddCompany.go`).
- **Database queries** live in `internal/database/db.go` — not inside handlers.
- **Templates** are written using [Templ](https://templ.guide/). Always run `templ generate` after editing `.templ` files.
- **HTMX partials** go in `web/templates/partials/` and are returned by handlers that respond to HTMX requests.

---

## Development Workflow

### Branching Strategy

Branch off `main` using descriptive names:

| Type | Pattern | Example |
|------|---------|---------|
| Feature | `feature/<short-description>` | `feature/add-events-filter` |
| Bug fix | `fix/<short-description>` | `fix/search-empty-results` |
| Docs | `docs/<short-description>` | `docs/improve-readme` |
| Data | `data/<short-description>` | `data/add-missing-companies` |

```bash
git checkout -b feature/my-new-feature
```

### Adding a New Route

1. Create the handler file in `internal/handlers/` (e.g., `getMyFeature.go`).
2. Add the query function to `internal/database/db.go` if new DB access is needed.
3. Register the route in `cmd/server/main.go`.
4. Create the Templ template in `web/templates/pages/` or `web/templates/partials/`.
5. Run `templ generate`.

### Adding a New Model Field

1. Add the field to the struct in `internal/models/company.go`.
2. GORM's `AutoMigrate` will apply the schema change on next startup — no manual SQL needed.
3. Update any relevant handlers and templates.

---

## Coding Conventions

### Go

- Follow standard Go formatting — run `gofmt` or `goimports` before committing.
- Use `go vet ./...` to catch common mistakes.
- **No debug `fmt.Println` in committed code.** Use `log.Println` for intentional logging.
- Error handling: always check and handle errors; don't silently ignore them.
- Keep handler functions thin — move business/query logic to `internal/database/db.go` or a service.

### Templ

- Components should be in their own `.templ` file.
- Full pages go in `web/templates/pages/`, reusable fragments in `web/templates/partials/`.
- Always run `templ generate` before committing — never commit stale generated `*_templ.go` files without the updated `.templ` source.

### CSS / Frontend

- Custom styles go in `web/static/css/`.
- The project uses Bootstrap for layout; avoid overriding Bootstrap internals where possible.
- HTMX attributes (`hx-get`, `hx-post`, `hx-target`, etc.) should be defined in the Templ templates.

### Commit Messages

Use clear, imperative-style commit messages:

```
Add filter by city to companies search
Fix superfluous WriteHeader call in Contact handler
Update README with Docker instructions
```

Avoid vague messages like `fix stuff`, `update`, or `wip`.

---

## Submitting a Pull Request

1. **Ensure your branch is up to date** with `main`:
   ```bash
   git fetch origin
   git rebase origin/main
   ```

2. **Verify the build passes:**
   ```bash
   go build ./...
   go vet ./...
   ```

3. **Regenerate templates** if you changed any `.templ` files:
   ```bash
   templ generate
   ```

4. **Open a Pull Request** against the `main` branch.
   - Write a clear title and description.
   - Reference any related issues (e.g., `Closes #12`).
   - Add screenshots for UI changes.

5. A maintainer will review your PR. Please respond to feedback promptly.

> [!IMPORTANT]
> PRs that include debug `fmt.Println` statements, ungenerated Templ files, or a broken `go build` will be asked to revise before merging.

---

## Reporting Issues

When opening a bug report, please include:

- **Go version** (`go version`)
- **OS and architecture**
- **Steps to reproduce**
- **Expected vs. actual behavior**
- **Relevant logs or error messages**

For feature requests, describe the use case and why it would benefit the directory.

---

*Thank you for helping make the Jordan Tech Companies Directory better for everyone!* 🇯🇴
