package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main() {
	defer fmt.Println("Successfully Migrated")
	
	cmd := exec.Command("go", "run", "github.com/steebchen/prisma-client-go", "migrate", "deploy", "--schema", "./infrastructure/prisma/schema.prisma")
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run prisma migration: %v, output: %s", err, string(output))
	}
	
	fmt.Println("Prisma migration completed successfully")
}
