// main.go

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"repos/github_search"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	screen.MoveTopLeft()
	// Meminta pengguna untuk memilih jenis pencarian
	fmt.Println("Pilih jenis pencarian:")
	fmt.Println("1. Find Repositori")
	fmt.Println("2. Find Users GitHub (by Username)")

	var choice int
	fmt.Print("Masukkan pilihan (1 atau 2): ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		log.Fatal("Hanya pilih 1 atau 2")
	}

	var query string
	fmt.Print("Masukkan kata kunci pencarian: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		query = scanner.Text()
	} else {
		log.Fatal("Gagal membaca input.")
	}

	switch choice {
	case 1:
		// Menjalankan pencarian repositori GitHub
		repoResult, err := github_search.SearchRepositories(query)
		if err != nil {
			log.Fatal(err)
		}

		// Menampilkan hasil pencarian repositori
		screen.Clear()
		screen.MoveTopLeft()
		github_search.PrintRepositories(repoResult)

	case 2:
		// Menjalankan pencarian pengguna GitHub
		userResult, err := github_search.SearchUsers(query)
		if err != nil {
			log.Fatal(err)
		}

		// Menampilkan hasil pencarian pengguna
		screen.Clear()
		screen.MoveTopLeft()
		github_search.PrintUsers(userResult)

	default:
		fmt.Println("Pilihan tidak valid.")
	}
}