# GOncurrently-Calculator

Распределенный вычислитель арифметических выражений

## Getting Started

### Prerequisites

Go 1.22 или выше

### Installation and Setup

1. Clone the repository:
      ```bash
      git clone https://github.com/Aspikk/GOncurrently-Calculator.git
      ```
2. Navigate to the project directory:
      ```bash
      cd GOncurrently-Calculator/cmd/app
      ```
3. Start the server:
      ```bash
      go run main.go
      ```

## Usage Examples

### Add an Expression ( 1 )

```bash
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
      "expression": "2+2*2^2/2-2"
}'
```

### Add an Expression ( 2 )

```bash
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
      "expression": "((2+2*2) / 2)^2 + (20*2-10)^2 - 30"
}'
```

### List Expressions

```bash
curl --location 'localhost/api/v1/expressions'
```

### Expression by id

```bash
curl --location 'localhost/api/v1/expressions/1'
```
