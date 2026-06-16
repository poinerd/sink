func SigninHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AuthRequest

		// 1. Decode incoming credentials
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// 2. Look up the user in Postgres by email

		var storedUser User
		
		query := `SELECT id, password_hash FROM users WHERE email = $1;`
		err = db.QueryRow(query, req.Email).Scan(&storedUser.ID, &storedUser.PasswordHash)
		
		if err != nil {
			// If email doesn't exist, sql.ErrNoRows is returned
			// Security Tip: Use a generic error message so hackers don't know which part failed
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// 3. Compare plain text password against the stored database hash
		err = bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(req.Password))
		if err != nil {
			// Passwords do not match
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// 4. Password is valid! Let's build the JWT wristband (Claims)
		expirationTime := time.Now().Add(24 * time.Hour) // Valid for 1 day
		claims := &CustomClaims{
			UserID: storedUser.ID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		// 5. Securely stamp the token with your secret key
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			fmt.Println("JWT signature error:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// 6. Return the wristband token to the client frontend
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
	}
}