package main

import (
	"log"
	"net/http"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Spustenie Pocketbase servera
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Absolútna cesta k súboru pocketbase.exe
		cmd := exec.Command("C:/Users/vladi/Desktop/test admin panel/functu/backend/pocketbase.exe", "serve")
		if err := cmd.Start(); err != nil {
			log.Fatalf("Chyba pri spustení Pocketbase: %v", err)
		}
		log.Printf("Pocketbase beží s PID %d", cmd.Process.Pid)

		if err := cmd.Wait(); err != nil {
			log.Printf("Pocketbase proces skončil s chybou: %v", err)
		}
	}()

	// Spustenie Vue.js aplikácie
	wg.Add(1)
	go func() {
		defer wg.Done()

		cmd := exec.Command("npm", "run", "dev")
		cmd.Dir = "C:/Users/vladi/Desktop/test admin panel/functu/frontend/app" // Nastavte pracovný adresár

		if err := cmd.Start(); err != nil {
			log.Fatalf("Chyba pri spustení Vue.js: %v", err)
		}
		log.Printf("Vue.js beží s PID %d", cmd.Process.Pid)

		if err := cmd.Wait(); err != nil {
			log.Printf("Vue.js proces skončil s chybou: %v", err)
		}
	}()

	// Nastavenie routera pre Go aplikáciu
	mux := http.NewServeMux()
	// mux.HandleFunc("/", homeHandler) // Môžete pridať svoje vlastné handlery

	// Spustenie HTTP servera
	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("Chyba pri spustení HTTP servera: %v", err)
		}
	}()

	// Čakanie na dokončenie všetkých goroutines
	wg.Wait()
}
