# net_cat
# NetCat Clone – A Lightweight Networking Tool

### Author:
This project was created by **Youssef jaouhar** (Yossf-jaouhar).
This project is a custom implementation inspired by **Netcat**, designed to provide simple yet powerful networking capabilities. It allows users to establish **TCP connections** for chat or data transfer, making it an ideal tool for both learning networking fundamentals and performing real-world network testing.

## Features:
- **Real-time group chat** with multiple clients.
- **Notifications** for user connections and disconnections.
- **Persistent chat history**.
- **Robust error handling** for seamless communication.

## Use Cases:
- **Debugging and testing network connections**.
- **Collaborative chatting** on private servers.
- **Educational tool** for understanding networking in Go.

## Technologies Used:
- **Go** (Golang) programming language
- **net package** (for TCP connections)
- **Mutex synchronization** for concurrent connections

## Installation:
1. Clone the repository:
    ```bash
    git clone https://github.com/Yossf-jaouhar/net_cat.git
    ```
2. Navigate to the project directory:
    ```bash
    cd net_cat
    ```
3. Install any necessary dependencies (if any).

## Running the Project:
To run the project, use the following command:
```bash
go run main.go <port>
