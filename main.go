package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ferris/norgannon/encrypt"
)

func showUsage() {
	fmt.Println("Norgannon - File Encryption Tool")
	fmt.Println("\nUsage:")
	fmt.Println("  norgannon [file path] [password]        Encrypt a file")
	fmt.Println("  norgannon -d [file path] [password]     Decrypt a file")
	fmt.Println("\nOptions:")
	fmt.Println("  -d, --decrypt    Decrypt mode")
	fmt.Println("  -h, --help       Show this help")
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		showUsage()
		return
	}

	decryptMode := false
	if len(os.Args) > 1 && (os.Args[1] == "-d" || os.Args[1] == "--decrypt") {
		decryptMode = true
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	if len(os.Args) != 3 {
		showUsage()
		os.Exit(1)
	}

	filePath := os.Args[1]
	password := os.Args[2]

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' does not exist\n", filePath)
		os.Exit(1)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if decryptMode {
		decryptedData, err := encrypt.DecryptData(data, password)
		if err != nil {
			fmt.Printf("Error decrypting file: %v\n", err)
			os.Exit(1)
		}

		dir, file := filepath.Split(filePath)
		outputFile := filepath.Join(dir, file)
		if strings.HasSuffix(file, ".encrypted") {
			outputFile = filepath.Join(dir, strings.TrimSuffix(file, ".encrypted"))
		} else {
			outputFile = filepath.Join(dir, file+".decrypted")
		}

		if _, err := os.Stat(outputFile); err == nil {
			fmt.Printf("Warning: Output file '%s' already exists. Overwrite? (y/n): ", outputFile)
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
				fmt.Println("Operation cancelled.")
				os.Exit(0)
			}
		}

		err = os.WriteFile(outputFile, decryptedData, 0644)
		if err != nil {
			fmt.Printf("Error writing decrypted file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("File decrypted successfully. Output: %s\n", outputFile)

	} else {
		encryptedData, err := encrypt.EncryptData(data, password)
		if err != nil {
			fmt.Printf("Error encrypting file: %v\n", err)
			os.Exit(1)
		}

		dir, file := filepath.Split(filePath)
		outputFile := filepath.Join(dir, file+".encrypted")

		if _, err := os.Stat(outputFile); err == nil {
			fmt.Printf("Warning: Output file '%s' already exists. Overwrite? (y/n): ", outputFile)
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
				fmt.Println("Operation cancelled.")
				os.Exit(0)
			}
		}

		err = os.WriteFile(outputFile, encryptedData, 0644)
		if err != nil {
			fmt.Printf("Error writing encrypted file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("File encrypted successfully. Output: %s\n", outputFile)
	}
} 