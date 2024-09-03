package utils

func IsValid(Tetromino [][]string) bool {
	for _, r := range Tetromino {
		TetrominoConnect := 0
		HashCount := 0
		for CheckVertical, line := range r {
			for CheckHorizontal, char := range line {
				if char != '#' && char != '.' {
					return false // check is characters are valid
				} else if char == '#' {
					HashConnect := 0
					HashCount++
					if CheckVertical > 0 && r[CheckVertical-1][CheckHorizontal] == '#' {
						HashConnect++
					}
					if CheckVertical < len(r)-1 && r[CheckVertical+1][CheckHorizontal] == '#' {
						HashConnect++
					}
					if CheckHorizontal > 0 && r[CheckVertical][CheckHorizontal-1] == '#' {
						HashConnect++
					}
					if CheckHorizontal < len(line)-1 && r[CheckVertical][CheckHorizontal+1] == '#' {
						HashConnect++
					}
					if HashConnect == 0 {
						return false
					} else {
						TetrominoConnect += HashConnect
					}
				}
			}
		}
	}
	return true
}

// Étape 1 : Lecture et validation du fichier d'entrée
// Lire l'argument de ligne de commande pour obtenir le chemin du fichier texte.
// Ouvrir et lire le fichier texte pour récupérer le contenu.
// Vérifier la validité du fichier :
// Le fichier doit contenir au moins un tétrimino.
// Chaque tétrimino doit avoir exactement 4 lignes de 4 caractères (# ou .).
// Chaque tétrimino doit contenir exactement 4 #.
// Les tétriminos doivent être séparés par une ligne vide.

// Étape 2 : Parsing des tétriminos
// Diviser le contenu du fichier en blocs correspondant à chaque tétrimino.
// Vérifier le format de chaque tétrimino pour s'assurer qu'il correspond aux règles mentionnées dans l'étape 3.
// Stocker les tétriminos dans une structure de données appropriée (par exemple, une matrice 2D ou une liste de coordonnées des #).

// Étape 3 : Normalisation des tétriminos
// Normaliser chaque tétrimino :
// Transposer chaque tétrimino pour qu'il soit dans sa position la plus en haut à gauche possible (retrait des lignes et colonnes vides en haut et à gauche).
// Conserver cette forme normalisée dans une structure de données.

// Étape 4 : Calcul du carré minimal
// Déterminer la taille minimale du carré pouvant contenir tous les tétriminos.
// Commencer avec le plus petit carré possible (par exemple, taille de 2x2 ou la racine carrée du nombre total de #).
// Utiliser une technique d'essai et erreur pour augmenter la taille si nécessaire.

// Étape 5 : Placement des tétriminos
// Essayer de placer les tétriminos dans le carré de taille initiale :
// Utiliser une méthode de backtracking (retour sur trace) pour essayer de placer chaque tétrimino dans toutes les positions possibles.
// Si un placement est impossible, augmenter la taille du carré et réessayer.

// Étape 6 : Algorithme de retour sur trace
// Mettre en place l'algorithme de retour sur trace :
// Pour chaque tétrimino, essayer de le placer dans le carré.
// Si tous les tétriminos sont placés correctement, l'algorithme est terminé.
// Si un tétrimino ne peut pas être placé, revenir à l'état précédent et essayer une autre configuration.

// Étape 7 : Génération et affichage du résultat
// Construire le carré final avec les positions des tétriminos placés correctement.
// Afficher le carré final avec les tétriminos placés ou une erreur si aucun placement n'est possible avec les configurations actuelles.

// Étape 8 : Gestion des erreurs
// Gérer les erreurs tout au long des étapes :
// Afficher "error" en cas de format incorrect du fichier d'entrée.
// Traiter les cas d'erreurs possibles pendant le parsing et le placement des tétriminos.
