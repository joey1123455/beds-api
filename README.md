# Project Setup

This guide will help you set up the necessary tools to work on this project. You'll be installing the Go compiler, Docker, Docker Compose, sqlc, and the Go Migrate CLI.

## 1. Install Go Compiler

Go to the official Go website and download the installer for your operating system.

### For Windows:

1. Download the MSI installer from the [Go downloads page](https://golang.org/dl/).
2. Run the MSI installer and follow the prompts.

### For macOS:

1. Download the package file from the [Go downloads page](https://golang.org/dl/).
2. Open the package file and follow the instructions to install.

### For Linux:

Use the following commands to download and install Go (replace `VERSION` with the latest Go version, e.g., `1.20.6`):

```bash
wget https://golang.org/dl/goVERSION.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf goVERSION.linux-amd64.tar.gz
```

Add Go to your PATH by adding the following line to your `.bashrc` or `.zshrc` file:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Reload your shell configuration:

```bash
source ~/.bashrc
# or
source ~/.zshrc
```

Verify the installation:

```bash
go version
```

## 2. Install Docker

Follow the instructions on the [Docker website](https://docs.docker.com/get-docker/) to install Docker for your operating system.

### For Windows and macOS:

1. Download Docker Desktop from the [Docker downloads page](https://www.docker.com/products/docker-desktop).
2. Run the installer and follow the prompts.

### For Linux:

Use the official Docker installation script:

```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
```

Verify the installation:

```bash
docker --version
```

## 3. Install Docker Compose

Follow the instructions on the [Docker Compose website](https://docs.docker.com/compose/install/) to install Docker Compose for your operating system.

### For Windows and macOS:

Docker Compose is included with Docker Desktop.

### For Linux:

Run the following commands:

```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

Verify the installation:

```bash
docker-compose --version
```

## 4. Install sqlc

Follow the instructions on the [sqlc website](https://docs.sqlc.dev/en/latest/overview/install.html) to install sqlc.

### Using Homebrew (macOS and Linux):

```bash
brew install sqlc
```

### Using Go (all platforms):

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

Verify the installation:

```bash
sqlc version
```

## 5. Install Go Migrate CLI

Follow the instructions on the [migrate website](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to install the Go Migrate CLI.

### Using Homebrew (macOS and Linux):

```bash
brew install golang-migrate
```

### Using Go (all platforms):

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Verify the installation:

```bash
migrate -version
```

---

## Makefile Instructions

The Makefile contains various commands to manage database migrations, run tests, ensure code quality, and build the project. Below are the available targets and their descriptions:

### General Commands

- **help**: Display this help message.
  
  ```sh
  make help
  ```

### SQL Migrations

- **create-migration**: Create a new migration file for a table.

  ```sh
  make create-migration table_name=<table_name>
  ```

- **migrate-force**: Force a migration to a specific version.

  ```sh
  make migrate-force version=<version>
  ```

- **migrate-up**: Migrate the database schema up to the latest version.

  ```sh
  make migrate-up
  ```

- **migrate-down**: Rollback the database schema back

  ```sh
  make migrate-down
  ```

- **migrate-drop**: Drop all database tables.

  ```sh
  make migrate-drop
  ```

- **migrate-to**: Migrate the database schema to a specific version.

  ```sh
  make migrate-to version=<version>
  ```

### Testing

- **test**: Run all tests with coverage and race detection.

  ```sh
  make test
  ```

### Code Quality Check

- **tidy**: Format code and tidy the mod file.

  ```sh
  make tidy
  ```

- **audit**: Run `go vet` and `golangci-lint`, and verify dependencies.

  ```sh
  make audit
  ```

- **lint**: Run `golangci-lint`.

  ```sh
  make lint
  ```

- **lint-fix**: Run `golangci-lint` with fix option.

  ```sh
  make lint-fix
  ```

### SQLC ORM CLI

- **sqlc-gen**: Generate code using sqlc.

  ```sh
  make sqlc-gen
  ```

### Build

- **docker-compose-up**: Build and start Docker Compose services.

  ```sh
  make docker-compose-up
  ```

- **docker-compose-down**: Stop Docker Compose services.

  ```sh
  make docker-compose-down
  ```

- **build-artifact**: Build the project for Linux.

  ```sh
  make build-artifact
  ```

### Mock Generation

- **mock**: Generate mock implementations for testing.

  ```sh
  make mock
  ```
