# Gintest

## Develop tools
- [Golang 1.15](https://golang.org/)
- [Make](https://www.gnu.org/software/make/)
- [Docker](https://www.docker.com/)
- [Editorconfig](https://editorconfig.org/)

## Quick Start
> requirentment
> - golang 1.15
> - gcc
> - make

- Config configuration
    - Run `cp .env.example .env`
    - Edit `.env` file
- Generate executable program 
    - Run `make clean all`
- Execute program 
    - Run `./app`

## Development

### File Structure
- `.github\`
    - `workflows\` : github action job descriptions
- `docker\` : docker image descriptions
- `docs\` : api documents
- `factories\` : fake entities
- `GithooksExample\` : git hook scripts
- `mocks\` : mock interfaces
- `src\` : main application components
- `tests\` : application tests
- `.editorconfig` : editor configuration
- `.env.example` : env example
- `.gitignore` : git ignore rules
- `docker-compose.yml` : docker containers description
- `go.mod` : application used modules
- `go.sum` : dependent modules of application used modules
- `main.go` : application entrypoint
- `Makefile` : make script
- `readme.md` : application documents


### Main Application Components
- `controllers` : services provider
- `entities` : application entities
- `handlers` : handle services
- `managers` : other applications connector
- `middlewares` : features before/after controllers execution
- `repositories` : data operations
- `routes` : application routes
- `services` : application features
- `utilities`: helpful tools

### Setup 
- Run `docker-compose up -d`

### Execute development application shell
- Run `docker-compose exec application /bin/sh`

### Generate executable program
- Run `make clena all`

### Execute program
- Run `./app`

### Test
- Run `go test ./tests/... -v`

### Add missing and remove unused modules
- Run `go mod tidy`

### Generate api documents
- Run `swag init`
