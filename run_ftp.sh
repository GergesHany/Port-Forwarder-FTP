# Check if the file name is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: ./run_ftp.sh <file_name>"
  exit 1
fi

FILE_NAME=$1

# Get the current directory path
CURRENT_DIR=$(pwd)

# Open a new terminal and run the (sender, receiver, forward) go files

gnome-terminal -- bash -c "cd $CURRENT_DIR/sender; go run sender.go $FILE_NAME; exec bash"

gnome-terminal -- bash -c "cd $CURRENT_DIR/receiver; go run receiver.go; exec bash"

gnome-terminal -- bash -c "cd $CURRENT_DIR/forward; go run forward.go; exec bash"