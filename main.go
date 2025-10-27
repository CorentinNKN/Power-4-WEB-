package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const (
	ROWS = 6
	COLS = 7
)

// Game représente l'état du jeu
type Game struct {
	Board         [ROWS][COLS]int `json:"board"`
	CurrentPlayer int             `json:"currentPlayer"`
	Winner        int             `json:"winner"`
	GameOver      bool            `json:"gameOver"`
	Message       string          `json:"message"`
}

// GameResponse pour les réponses API
type GameResponse struct {
	Success bool   `json:"success"`
	Game    *Game  `json:"game"`
	Error   string `json:"error,omitempty"`
}

// NewGame crée une nouvelle partie
func NewGame() *Game {
	return &Game{
		CurrentPlayer: 1,
		Winner:        0,
		GameOver:      false,
		Message:       "Joueur 1 à vous de jouer",
	}
}

// Reset remet à zéro la partie
func (g *Game) Reset() {
	g.Board = [ROWS][COLS]int{}
	g.CurrentPlayer = 1
	g.Winner = 0
	g.GameOver = false
	g.Message = "Joueur 1 à vous de jouer"
}

// MakeMove effectue un coup dans la colonne spécifiée
func (g *Game) MakeMove(col int) bool {
	if g.GameOver || col < 0 || col >= COLS {
		return false
	}

	// Trouver la ligne disponible la plus basse
	row := -1
	for i := ROWS - 1; i >= 0; i-- {
		if g.Board[i][col] == 0 {
			row = i
			break
		}
	}

	if row == -1 {
		g.Message = "Colonne pleine, choisissez une autre colonne"
		return false
	}

	// Placer le jeton
	g.Board[row][col] = g.CurrentPlayer

	// Vérifier la victoire
	if g.CheckWin(row, col) {
		g.Winner = g.CurrentPlayer
		g.GameOver = true
		g.Message = fmt.Sprintf("Joueur %d remporte la partie !", g.CurrentPlayer)
		return true
	}

	// Vérifier match nul
	if g.IsBoardFull() {
		g.GameOver = true
		g.Message = "Match nul"
		return true
	}

	// Changer de joueur
	if g.CurrentPlayer == 1 {
		g.CurrentPlayer = 2
		g.Message = "Joueur 2 à vous de jouer"
	} else {
		g.CurrentPlayer = 1
		g.Message = "Joueur 1 à vous de jouer"
	}

	return true
}

// CheckWin vérifie si le dernier coup est gagnant
func (g *Game) CheckWin(row, col int) bool {
	player := g.Board[row][col]

	// Vérifier horizontal
	count := 1
	for c := col + 1; c < COLS && g.Board[row][c] == player; c++ {
		count++
	}
	for c := col - 1; c >= 0 && g.Board[row][c] == player; c-- {
		count++
	}
	if count >= 4 {
		return true
	}

	// Vérifier vertical
	count = 1
	for r := row + 1; r < ROWS && g.Board[r][col] == player; r++ {
		count++
	}
	if count >= 4 {
		return true
	}

	// Vérifier diagonal (\)
	count = 1
	for r, c := row+1, col+1; r < ROWS && c < COLS && g.Board[r][c] == player; r, c = r+1, c+1 {
		count++
	}
	for r, c := row-1, col-1; r >= 0 && c >= 0 && g.Board[r][c] == player; r, c = r-1, c-1 {
		count++
	}
	if count >= 4 {
		return true
	}

	// Vérifier diagonal (/)
	count = 1
	for r, c := row+1, col-1; r < ROWS && c >= 0 && g.Board[r][c] == player; r, c = r+1, c-1 {
		count++
	}
	for r, c := row-1, col+1; r >= 0 && c < COLS && g.Board[r][c] == player; r, c = r-1, c+1 {
		count++
	}
	if count >= 4 {
		return true
	}

	return false
}

// IsBoardFull vérifie si le plateau est plein
func (g *Game) IsBoardFull() bool {
	for col := 0; col < COLS; col++ {
		if g.Board[0][col] == 0 {
			return false
		}
	}
	return true
}

// GetColumns retourne un slice avec les numéros de colonnes
func (g *Game) GetColumns() []int {
	cols := make([]int, COLS)
	for i := range cols {
		cols[i] = i
	}
	return cols
}

var game *Game

func main() {
	game = NewGame()

	// Routes pour l'interface web
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/move", moveHandler)
	http.HandleFunc("/reset", resetHandler)

	// Routes API JSON
	http.HandleFunc("/api/game", apiGameHandler)
	http.HandleFunc("/api/move", apiMoveHandler)
	http.HandleFunc("/api/reset", apiResetHandler)

	// Fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("🎮 Serveur Puissance 4 démarré sur http://localhost:8080")
	fmt.Println("📱 Interface web : http://localhost:8080")
	fmt.Println("🔗 API JSON : http://localhost:8080/api/game")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/index.html")
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/game.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	colStr := r.FormValue("column")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	game.MakeMove(col)
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	game.Reset()
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// API Handlers
func apiGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := GameResponse{
		Success: true,
		Game:    game,
	}
	json.NewEncoder(w).Encode(response)
}

func apiMoveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		response := GameResponse{
			Success: false,
			Error:   "Méthode non autorisée",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	colStr := r.FormValue("column")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		response := GameResponse{
			Success: false,
			Error:   "Colonne invalide",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	success := game.MakeMove(col)
	response := GameResponse{
		Success: success,
		Game:    game,
	}

	if !success {
		response.Error = "Coup invalide"
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(response)
}

func apiResetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		response := GameResponse{
			Success: false,
			Error:   "Méthode non autorisée",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}

	game.Reset()
	response := GameResponse{
		Success: true,
		Game:    game,
	}
	json.NewEncoder(w).Encode(response)
}
