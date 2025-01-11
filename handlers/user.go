package handlers

import (
    "html/template"
    "net/http"
    "log"
    "user-auth-app/models"
    "user-auth-app/database"
    "golang.org/x/crypto/bcrypt"
)

// Handle the registration form submission
func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // Render the registration form
        tmpl, err := template.ParseFiles("templates/register.html")
        if err != nil {
            log.Println("Error parsing template:", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
    } else if r.Method == http.MethodPost {
        // Get the user data from the form
        user := models.User{
            Username: r.FormValue("username"),
            Email:    r.FormValue("email"),
            Password: r.FormValue("password"),
        }

        // Hash the password before storing it in the database
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            log.Println("Error hashing password:", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }

        // Store the hashed password in the user struct
        user.Password = string(hashedPassword)

        // Log the hashed password for debugging purposes
        log.Println("Hashed Password (Registration):", user.Password)

        // Save the user to the database
        if err := database.DB.Create(&user).Error; err != nil {
            log.Println("Error creating user:", err)
            http.Error(w, "Error creating user", http.StatusInternalServerError)
            return
        }

        // Redirect to the success page
        http.Redirect(w, r, "/success", http.StatusSeeOther)
    }
}

// Handle the login form submission
// Handle the login form submission
func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // Render the login form
        tmpl, err := template.ParseFiles("templates/login.html")
        if err != nil {
            log.Println("Error parsing template:", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
    } else if r.Method == http.MethodPost {
        // Get user data from form (email and password)
        email := r.FormValue("email")
        password := r.FormValue("password")

        // Find user in the database by email
        var user models.User
        if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
            log.Println("Error fetching user:", err)
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        // Compare the password with the stored hash
        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
            log.Println("Invalid password:", err)
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        // Redirect to the success page
        http.Redirect(w, r, "/success", http.StatusSeeOther)
    }
}


// Success page handler after registration or login
func Success(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/success.html")
    if err != nil {
        log.Println("Error parsing template:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}
