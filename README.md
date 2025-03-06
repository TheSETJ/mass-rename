# Mass Rename

A simple Go app to rename files in a directory by prepending a prefix and/or appending a suffix to their names.

## Usage

### Flags:
- `-dir` (default: `.`) : Directory to rename files in
- `-prefix` : Prefix to add to file names
- `-suffix` : Suffix to add to file names

### Example:
1. Rename files with a prefix:
   ```bash
   go run main.go -prefix "new_"
   ```

2. Rename files with a suffix:
   ```bash
   go run main.go -suffix "_backup"
   ```

3. Rename files with both prefix and suffix:
   ```bash
   go run main.go -dir "./my_files" -prefix "new_" -suffix "_backup"
   ```

## Error Handling

If an error occurs (e.g., missing directory or both prefix and suffix are empty), the program will display an error message and exit.

## License

This project is open source and available under the MIT License.