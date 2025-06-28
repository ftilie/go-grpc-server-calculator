# go-grpc-server-calculator
Go powered Calculator using a Server-Client architecture with gRPC

## Features
- Perform basic arithmetic operations: `Add`, `Subtract`, `Multiply`, `Divide`, `GetPrimes`, `GetAverage`, `GetMaximum`, `SquareRoot`.
- Built using gRPC for efficient communication between server and client.
- Written in Go for high performance and scalability.

> **Note**: This project is not intended to be a fully-featured calculator application. Instead, it serves as an experiment in building a Server-Client architecture using gRPC and Protocol Buffers. Additionally, it integrates [Evans](https://github.com/ktr0731/evans) for testability and debugging purposes.

## Prerequisites
- Go installed on your system (version 1.16 or later).
- `protoc` (Protocol Buffers compiler) installed for generating gRPC code.

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/go-grpc-server-calculator.git
    cd go-grpc-server-calculator
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```

## Usage
1. Use `Make` to generate the protobuf code.
2. Start the gRPC server:
    ```bash
    .\bin\calculator\server.exe
    ```
3. Run the client to perform calculations:
    ```bash
    .\bin\calculator\client.exe
    ```
4. Alternatively, use [Evans](https://github.com/ktr0731/evans) to test the gRPC server:
    - Start Evans in interactive mode:
        ```bash
        evans --host localhost --port 50051 -r repl
        ```
    - Select the `calculator` package and `CalculatorService` service:
        ```bash
        package calculator
        service CalculatorService
        ```
    - Call the available methods:
        ```bash
        call Add
        call Subtract
        call Multiply
        call Divide
        call GetPrimes
        call GetAverage
        call GetMaximum
        call SquareRoot
        ```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.