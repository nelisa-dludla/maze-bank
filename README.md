# Maze Bank

Maze Bank is a simple command-line interface (CLI) program written in Go for simulating banking operations like deposits, withdrawals, and statement generation. It utilizes SQLite for database storage. The name "Maze Bank" was inspired by the bank in GTA V.

![Maze Bank Welcome Menu](https://i.imgur.com/iCBz4gV.png)

## Features

- **User Registration:** Users can register using their email and password.
- **Login:** Registered users can log in securely to access their accounts.
- **Deposits and Withdrawals:** Users can make deposits and withdrawals to/from their accounts.
- **Statement Generation:** Users can generate a statement that lists all the transactions performed by them, which is saved as a text file.

## Requirements

- Go programming language (version 1.22.1)
- SQLite

## Usage

### Using Pre-compiled Builds
1. Clone the repository:

```bash
git clone https://github.com/nelisa-dludla/maze-bank.git
```

2. Navigate to the project directory:

```bash
cd maze-bank
```

3. Choose the appropriate pre-compiled executable for your operating system:
- For Windows: `maze-bank.exe`
- For Linux: `maze-bank`

4. Run the executable:

```bash
./maze-bank # On Linux
```

```bash
.\maze-bank.exe # On Windows
```

### Building and Running from Source

1. Clone the repository:

```bash
git clone https://github.com/nelisa-dludla/maze-bank.git
```

2. Navigate to the project directory:

```bash
cd maze-bank
```

3. Build the project:

```bash
go build
```

4. Run the executable:

```bash
./maze-bank
```

5. Follow the on-screen instructions to register/login and perform banking operations.

## Database

The project uses SQLite for database storage. The database file (maze_bank.db) will be created in the project directory upon first run.
