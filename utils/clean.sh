# This script is used for cleaning up the project directory

# Remove the binary file
# This command removes the binary file located in the bin directory
rm -rf bin/${BINARY}

# Remove the test files
# This command removes all the files in the ./app/testfiles directory
rm -rf ./app/testfiles/*

# Print a message indicating that the cleanup is complete
echo "Cleanup complete!"