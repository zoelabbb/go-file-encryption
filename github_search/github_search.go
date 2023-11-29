// github_search/github_search.go

package github_search

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const githubAPIURL = "https://api.github.com/search"

// Struct untuk menangkap hasil pencarian repositori
type RepositoriesResult struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		HTMLURL     string `json:"html_url"`
	} `json:"items"`
}

// Struct untuk menangkap hasil pencarian pengguna
type UsersResult struct {
	TotalCount int `json:"total_count"`
	Items      []struct {
		Login string `json:"login"`
		URL   string `json:"html_url"`
	} `json:"items"`
}

// Fungsi untuk mencari repositori GitHub
func SearchRepositories(query string) (*RepositoriesResult, error) {
	url := fmt.Sprintf("%s/repositories?q=%s", githubAPIURL, query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result RepositoriesResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Fungsi untuk mencari pengguna GitHub
func SearchUsers(query string) (*UsersResult, error) {
	url := fmt.Sprintf("%s/users?q=%s", githubAPIURL, query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result UsersResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Fungsi untuk menampilkan hasil pencarian repositori
func PrintRepositories(result *RepositoriesResult) {
	fmt.Printf("Total Repositori: %d\n", result.TotalCount)
	fmt.Println("Repositori:")
	for _, repo := range result.Items {
		fmt.Printf("- %s: %s\n", repo.Name, repo.HTMLURL)
	}
}

// Fungsi untuk menampilkan hasil pencarian pengguna
func PrintUsers(result *UsersResult) {
	fmt.Printf("Total Pengguna: %d\n", result.TotalCount)
	fmt.Println("Pengguna:")
	for _, user := range result.Items {
		fmt.Printf("- %s: %s\n", user.Login, user.URL)
	}
}