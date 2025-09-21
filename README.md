# MyWant

![Want Dashboard](docs/img/want_dashboard.png)

**Declarative chain programming with YAML configuration.** Express what you want to achieve, not how to do it.

📚 **Documentation:** [Want System](docs/want-system.md) | [Agent System](docs/agent-system.md) | [Examples](docs/agent-examples.md)

## Features

- **YAML-Only Configuration**: No programming required
- **Recipe-Based Templates**: Reusable component configurations
- **Independent & Dependent Wants**: Parallel processing or sequential pipelines
- **Web Dashboard**: Modern React interface for monitoring
- **Dynamic Want Addition**: Runtime system modification
- **Memory Reconciliation**: Persistent state across executions

## Quick Start

### Server Mode (with Web Dashboard)

```bash
make server  # Start server with web interface
```

### Example: Queue Processing Pipeline

**Create config** (`config/config-qnet.yaml`):
```yaml
wants:
  - metadata:
      name: generator
      type: numbers
      labels: {role: source}
    spec:
      params: {count: 1000, rate: 10.0}

  - metadata:
      name: processor
      type: queue
      labels: {role: processor}
    spec:
      params: {service_time: 0.05}
      using: [{role: source}]  # Connect to generator

  - metadata:
      name: collector
      type: sink
      labels: {role: collector}
    spec:
      using: [{role: processor}]  # Connect to processor
```

**Run:**
```bash
make run-qnet
```

## API Examples

```bash
# Start server
make server

# Create wants via API
curl -X POST http://localhost:8080/api/v1/wants \
  -H "Content-Type: application/yaml" \
  --data-binary @config/config-qnet.yaml

# Monitor status
curl http://localhost:8080/api/v1/wants/{id}/status
```

## More Examples

```bash
make run-travel-recipe    # Travel planning
make run-fibonacci       # Fibonacci sequence
make run-qnet-using-recipe # Multi-stream processing
```

## Usage

```go
import "github.com/onelittlenightmusic/MyWant"

config, _ := mywant.LoadConfigFromYAML("config.yaml")
builder := mywant.NewChainBuilder(config)
builder.RegisterWantType("your-type", yourConstructor)
builder.Execute()
```

## Requirements

- Go 1.21+
- `gopkg.in/yaml.v3`