# MInitializr

MInitializr is a toolkit for initializing microservice-based applications using official framework initializers (Spring Boot, Micronaut, Quarkus, Grails, Vert.x, and more). It provides both a web service and a CLI to generate ready-to-use project structures, supporting multi-service setups and custom configurations.

---

## Features

- **Multi-service project initialization**: Define multiple services with different technologies in a single config.
- **Framework support**: Spring Boot, Micronaut, Quarkus, Grails, Vert.x, and React (planned).
- **Version validation**: Ensures only supported framework versions are used.
- **Official initializer integration**: Uses the frameworks' own project generators.
- **Flexible configuration**: YAML/JSON config for services, versions, and options.
- **CLI tool**: Automate project generation from your terminal.
- **Download as ZIP or extract directly**: Choose your preferred output format.

---

## Getting Started

### Prerequisites

- Go 1.21+
- (For CLI) [minitializr-cli](./cli)
- Internet connection (downloads from official initializers)

### Quick Start (Web Service)

1. **Clone the repo**  
   ```sh
   git clone https://github.com/HamzaBenyazid/minitializr.git
   cd minitializr
   ```

2. **Run the server**  
   ```sh
   go run main.go
   ```
   The API will be available at `http://localhost:8080`.

3. **POST a config**  
   Send a POST request to `/initialize` with a config (see [init-example.json](./init-example.json)).

   Example using `curl`:
   ```sh
   curl -X POST http://localhost:8080/initialize \
     -H "Content-Type: application/json" \
     --data-binary @init-example.json \
     -o myproject.zip
   ```

### Using the CLI

See [cli/README.md](./cli/README.md) for CLI usage.

---

## Configuration

Define your project in YAML or JSON. Example ([init-example.json](./init-example.json)):

```json
{
  "apiVersion": "0.0.1",
  "metadata": { "name": "ecom" },
  "services": {
    "costumer": {
      "technology": "spring-boot",
      "version": "3.2.3",
      "config": { "language": "java", ... }
    },
    "product": {
      "technology": "micronaut",
      "version": "prev",
      "config": { ... }
    }
    // more services...
  }
}
```

Supported frameworks and versions are listed in [supported-frameworks.yaml](./supported-frameworks.yaml).

---

## API

- `POST /initialize`  
  Request body: project config (JSON)  
  Response: ZIP file containing the generated project

---

## Development

- Main entry: `main.go`
- Service logic: `service/initialize.go`
- Initializer implementations: `initializers/`
- Types: `types/`
- Utilities: `utils/`
- Supported frameworks: `supported-frameworks.yaml`
- Example config: `init-example.json`

Run tests:
```sh
go test ./...
```

---

## Roadmap

- [x] Multi-service initialization
- [x] Version validation against supported frameworks
- [x] CLI tool for local usage
- [ ] Add more framework templates (e.g., React, Angular, Vue)
- [ ] Interactive CLI prompts
- [ ] Plugin system for custom generators
- [ ] Improved error reporting and logging
- [ ] Web UI for config editing and project download
- [ ] CI/CD pipeline templates
- [ ] More test coverage and integration tests

---

## License

[MIT](LICENSE)

# MInitializr CLI

## Overview

**MInitializr CLI** is a command-line tool for quickly generating new project structures with best-practice configurations. It is designed to work standalone or in conjunction with [minitializr.org](../minitializr.org).

## Features

- Scaffold new projects from templates
- Customizable options via CLI flags
- Supports multiple frameworks and languages
- Extensible for custom templates

## Getting Started

### Prerequisites

- Node.js (>= 18)
- npm or yarn

### Installation

Clone the repository and install dependencies:

```bash
git clone <repo-url>
cd MInitializr
npm install
```

### Usage

```bash
node cli.js --template <template-name> --name <project-name>
```

Example:

```bash
node cli.js --template react --name my-app
```

## Roadmap

- [ ] Add more templates and language support
- [ ] Interactive CLI prompts
- [ ] Integration with minitializr.org for remote template fetching
- [ ] Plugin system for custom generators
- [ ] Automated tests and CI

## Contributing

Feel free to open issues or submit pull requests.

## License

MIT
