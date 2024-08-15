package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
)

// Premenná pre uloženie referencie na proces PocketBase
var cmd *exec.Cmd

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[SYSTEM Online]")
		fmt.Println("Použitie: go run main.go [start|stop|restart]")
		os.Exit(1)
	}

	action := os.Args[1]

	switch action {
	case "start":
		startServer()
	case "stop":
		stopServer()
	case "restart":
		restartServer()
	default:
		fmt.Println("Neznámy príkaz. Použitie: go run main.go [start|stop|restart]")
	}
}

// Funkcia na spustenie PocketBase servera
func startServer() {
	exePath, err := filepath.Abs("./backend/pocketbase.exe")
	if err != nil {
		log.Fatalf("Chyba pri zisťovaní absolútnej cesty: %v", err)
	}

	// Príkaz pre spustenie PocketBase servera
	cmd = exec.Command(exePath, "serve")

	// Presmerovanie výstupu a chýb na konzolu
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Spustenie PocketBase servera
	err = cmd.Start()
	if err != nil {
		log.Fatalf("Chyba pri spúšťaní servera: %v", err)
	}

	fmt.Printf("Server spustený s PID %d\n", cmd.Process.Pid)

	// Kanál pre prijímanie signálov
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		stopServer()
	}()

	// Čakanie na ukončenie procesu
	err = cmd.Wait()
	if err != nil {
		log.Printf("Server bol zastavený s chybou: %v", err)
	}
}

// Funkcia na zastavenie PocketBase servera
func stopServer() {
	// Nastavenie pracovného adresára na "backend"
	cmd := exec.Command("taskkill", "/IM", "pocketbase.exe", "/F")
	cmd.Dir = "./backend" // Nastavenie pracovného adresára na "backend"

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Chyba pri zastavení servera: %v", err)
	}

	fmt.Println("Pokúsili sme sa zastaviť server.")
}

// Funkcia na reštartovanie PocketBase servera
func restartServer() {
	stopServer()
	startServer()
}
