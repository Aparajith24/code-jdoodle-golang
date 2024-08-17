# Go Code Execution Service

This is a simple HTTP server built in Go that executes code snippets via the JDoodle API. It supports multiple programming languages and allows users to input their code, along with any required input, and receive the output or error messages.

## Features

- **Supports Multiple Languages**: Execute code in languages like Python, Java, and more.
- **Handles User Input**: Accepts user input for code execution.
- **Preserves Code Indentation**: Handles and preserves indentation, which is critical for languages like Python.
- **JSON-based API**: Uses JSON to handle requests and responses, making it easy to integrate with other services.

## Prerequisites

- Go 1.15+ installed on your machine.
- A JDoodle API account to obtain your `clientId` and `clientSecret`.

## Installation

1. Clone this repository to your local machine:
    ```bash
    git clone https://github.com/your-username/go-code-execution-service.git
    cd go-code-execution-service
    ```

2. Set up your JDoodle credentials by replacing the placeholders in the `JDoodleRequest` struct in the `main.go` file:
    ```go
    ClientID:     "your_client_id",
    ClientSecret: "your_client_secret",
    ```

3. Build and run the server:
    ```bash
    go run main.go
    ```

### Example Usage

You will be prompted to enter:

- **Code**: Your code snippet.
- **Input**: Any input your code might need during execution.
- **Programming Language**: The language in which your code is written (e.g., `python`, `java`).

```plaintext
Enter the code:
fact = 1
num = 8
for i in range(1, num+1):
    fact *= i
print(fact)

Enter the input (if any):

Enter the programming language (e.g., python, java):
python
