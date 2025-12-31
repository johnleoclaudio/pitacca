# Pitacca

### Directories
- `bin/` contains compiled binaries for production
- `cmd/api/` contains application-specific code. This include the code for running the server, reading and writing HTTP requests, and managing authentication.
- `internal/` contains the code for interacting with database, doing data validation, sending emails and so on. Any code that is not application-specific and can potentially be reused should go here. `cmd/api` will import this but not the other way around. Anything inside this directory can't be imported outside the project.
- `migrations/` contains SQL migration files for database
- `remote/` contains the configuration files and setup scripts for production server
- `Makefile` contains **recipes** for automating common tasks like auditing Go code, building binaries, and executing DB migrations
