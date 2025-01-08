# FTP-Like File Transfer Project in Go

This project demonstrates a simple FTP-like file transfer system using Go. It consists of three components:

1. **Sender Server**: Hosts a file and sends it to the forwarder.
2. **Receiver Server**: Receives the file and saves it in a `download` folder.
3. **Forwarder**: Retrieves the file from the sender and forwards it to the receiver.

The project is designed to simulate a basic file transfer system, where files are forwarded from one server to another.

---

## Project Structure

```plaintext
Port-Forwarder/
├── sender/
│ └── sender.go # Sender server code
├── receiver/
│ └── receiver.go # Receiver server code
├── forward/
│ └── forward.go # Forwarder code
├── run_ftp.sh # Bash script to automate running the project
└── README.md # Project documentation
```

<hr>

## How to Use

1. Clone the Repository

   - Clone the project repository to your local machine:
     ```bash
        git clone https://github.com/GergesHany/Port-Forwarder.git
        cd Port-Forwarder
     ```

2. Build and Run the Project

   Option 1: Manual Execution

   1. Start the Sender Server:
      Navigate to the sender directory and run the `sender` with a file name as an argument:

      ```bash
          cd sender
          go run sender.go <file_name>
      ```

   2. Start the Receiver Server:
      Open a new terminal, navigate to the `receiver` directory, and run the receiver:

      ```bash
          cd receiver
          go run receiver.go
      ```

   3. Start the Forwarder:
      Open another terminal, navigate to the `forward` directory, and run the forwarder:
      ```bash
          cd forward
          go run forward.go
      ```

   Option 2: Automated Execution (Using run_ftp.sh)

   1. Make the Script Executable:
      Run the following command to make the script executable:

      ```bash
      chmod +x run_ftp.sh
      ```

   2. Run the Script:
      Execute the script with the file name as an argument:

      ```bash
      ./run_ftp.sh example.txt
      ```

   This will open three new terminal windows and run the sender, receiver, and forwarder automatically.

<hr>

## Example

1. Create a Sample File:
   Create a file named example.txt in the sender directory:

   ```bash
   echo "Hello, World!" > sender/example.txt
   ```

2. Run the Project:
   Execute the run_ftp.sh script with the file name as an argument:

   ```bash
   ./run_ftp.sh example.txt
   ```

3. Check the Output:

   The sender will send `example.txt`.

   The receiver will save the file in the download folder with a `unique name` (e.g., received_file_1697049600.txt).
