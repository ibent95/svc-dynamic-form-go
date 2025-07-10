package commands

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Make(args []string) {
	if len(args) < 1 {
		log.Fatal("⚠️  Please specify what to make (controller, service, repository, model, request, middleware, command, module)")
	}

	target := args[0]
	name := ""
	if len(args) > 1 {
		name = args[1]
	}

	switch target {
	case "controller":
		MakeFile("controller", name, "src/controllers", "framework/stubs/controller.stub")
	case "service":
		MakeFile("service", name, "src/services", "framework/stubs/service.stub")
	case "repository":
		MakeFile("repository", name, "src/repositories", "framework/stubs/repository.stub")
	case "model":
		MakeFile("model", name, "src/models", "framework/stubs/model.stub")
	case "request":
		MakeFile("request", name, "src/requests", "framework/stubs/request.stub")
	case "middleware":
		MakeFile("middleware", name, "src/middlewares", "framework/stubs/middleware.stub")
	case "command":
		MakeFile("command", name, "src/commands", "framework/stubs/command.stub")
	case "module":
		MakeFile("model", name, "src/models", "framework/stubs/model.stub")
		MakeFile("repository", name, "src/repositories", "framework/stubs/repository.stub")
		MakeFile("service", name, "src/services", "framework/stubs/service.stub")
		MakeFile("controller", name, "src/controllers", "framework/stubs/controller.stub")
	default:
		log.Fatalf("❌ Unknown make target: %s", target)
	}
}

func MakeFile(kind, name, targetDir, stubPath string) {
	if name == "" {
		log.Fatalf("❌ You must provide a name for the %s", kind)
	}

	snake := strings.ToLower(name)
	pascal := strings.Title(snake)
	filename := fmt.Sprintf("%s/%s_%s.go", targetDir, snake, kind)

	if _, err := os.Stat(filename); err == nil {
		log.Printf("⚠️  %s already exists: %s", strings.Title(kind), filename)
		return
	}

	stub, err := os.ReadFile(stubPath)
	if err != nil {
		log.Fatalf("❌ Failed to read stub file: %v", err)
	}

	template := string(stub)
	template = strings.ReplaceAll(template, "{{Name}}", pascal)
	template = strings.ReplaceAll(template, "{{name}}", snake)

	err = os.MkdirAll(targetDir, os.ModePerm)
	if err != nil {
		log.Fatalf("❌ Failed to create target directory: %v", err)
	}

	err = os.WriteFile(filename, []byte(template), 0644)
	if err != nil {
		log.Fatalf("❌ Failed to write file: %v", err)
	}

	log.Printf("✅ Created %s: %s", kind, filename)
}
