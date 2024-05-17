# Distributed-File-Storage

A distributed file system (DFS) implementation that manages storage across multiple networked servers and provides users with a unified view of the stored data.

## Features
- **Distributed Storage**: Manages files across multiple servers.
- **P2P Architecture**: Implements peer-to-peer communication.
- **Fault Tolerance**: Ensures data integrity and availability.

## Installation

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/ValRobichaux/Distributed-File-Storage.git
cd Distributed-File-Storage
```

## Usage

Build and run the project using `Makefile`:

```sh
make
./bin/distributed-file-storage
```

## Project Structure

- **main.go**: Entry point of the application.
- **server.go**: Server-side logic.
- **store.go**: File storage logic.
- **p2p/**: Peer-to-peer communication implementation.

## Testing

Run tests with:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for suggestions or bug reports.

## License

This project is licensed under the MIT License.

---