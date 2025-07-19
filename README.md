# Inforce Task - Rarible NFT API Client

## Prerequisites

- Go 1.24.2 or higher
- Docker (optional, for containerization)
- Kubernetes cluster (optional, for deployment)
- Helm 3.x (optional, for Kubernetes deployment)

## Configuration

Create a `.env` file in the root directory based on `.env.example`:

```bash
cp .env.example .env
```

Configure the following environment variables:

```bash
SERVER_PORT=8080

# Rarible API Configuration
RARIBLE_BASE_URL=https://api.rarible.org/v0.1/
RARIBLE_API_KEY=your_actual_api_key_here

# HTTP Client Configuration
HTTP_CLIENT_TIMEOUT=10s
HTTP_CLIENT_CONNECT_TIMEOUT=5s
HTTP_CLIENT_MAX_IDLE_CONNS=100
```

## Running the Application

### Method 1: Direct Go Run

1. Install dependencies:
```bash
go mod download
```

2. Build the application:
```bash
make build; make run
```

3. Run the application:
```bash
make run
```

### Method 2: Using Go Command

```bash
go run cmd/app/main.go
```

### Method 3: Docker

1. Build the Docker image:
```bash
docker build -t inforce-task .
```

2. Run the container:
```bash
docker run -p 8080:8080 --env-file .env inforce-task
```

### Method 4: Kubernetes with Helm

1. Update the Helm values in `chart/values.yaml`:
```yaml
secret:
  RARIBLE_API_KEY: "your_actual_api_key_here"
```

2. Deploy to Kubernetes:
```bash
helm install inforce-task ./chart
```

## API Endpoints

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 1. Get NFT Ownership
**GET** `/ownership`

Returns NFT ownership information for a given wallet address.

**Query Parameters:**
- `id` (string, required): Wallet address

**Example:**
```bash
curl "http://localhost:8080/ownership?id=0x1234567890abcdef1234567890abcdef12345678"
```

#### 2. Get Collection Traits
**POST** `/traits`

Returns trait rarity information for items in an NFT collection.

**Request Body:**
```json
{
  "collectionId": "ETHEREUM:0x1234567890abcdef1234567890abcdef12345678",
  "properties": [
    {
      "key": "Background",
      "value": "Blue"
    },
    {
      "key": "Eyes",
      "value": "Laser"
    }
  ]
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/traits \
  -H "Content-Type: application/json" \
  -d '{
    "collectionId": "ETHEREUM:0x1234567890abcdef1234567890abcdef12345678",
    "properties": [
      {"key": "Background", "value": "Blue"}
    ]
  }'
```

## API Documentation

Interactive API documentation is available via Swagger UI:

```
http://localhost:8080/swagger/
```

## Testing

### Run All Tests
```bash
make test
```

### Run Tests with Coverage
```bash
go test ./internal/client/... ./internal/service/... -cover -v
```

## Development

### Generate Swagger Documentation
```bash
make swagger
```

### Generate Mocks
The project uses mockery for generating mocks. Mocks are already generated in the `mocks/` directory.

### Project Structure Guidelines

- **Controllers**: Handle HTTP requests and responses
- **Services**: Contain business logic
- **Clients**: Handle external API communication
- **DTOs**: Define data transfer objects for API contracts
- **Config**: Manage application configuration

## Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `SERVER_PORT` | Server port number | `8080` | Yes |
| `RARIBLE_BASE_URL` | Rarible API base URL | `https://api.rarible.org/v0.1/` | Yes |
| `RARIBLE_API_KEY` | Rarible API key | - | Yes |
| `HTTP_CLIENT_TIMEOUT` | HTTP client timeout | `10s` | No |
| `HTTP_CLIENT_CONNECT_TIMEOUT` | HTTP client connection timeout | `5s` | No |
| `HTTP_CLIENT_MAX_IDLE_CONNS` | Maximum idle connections | `100` | No |

## Error Handling

The API uses structured error responses:

```json
{
  "error": "error message here"
}
```

HTTP status codes used:
- `200` - Success
- `400` - Bad Request (validation errors, missing parameters)
- `401` - Unauthorized (invalid API key)
- `403` - Forbidden
- `404` - Not Found
- `500` - Internal Server Error

## Troubleshooting

### Common Issues

1. **"Failed to load config" error**
   - Ensure `.env` file exists and contains all required variables
   - Check that `RARIBLE_API_KEY` is set correctly

2. **Connection timeout errors**
   - Verify internet connection and Rarible API availability
   - Check if API key is valid and has proper permissions

3. **Port already in use**
   - Change `SERVER_PORT` in `.env` to an available port
   - Kill any existing processes using the port
