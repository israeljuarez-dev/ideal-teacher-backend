package service

import "golang.org/x/crypto/bcrypt"

// HashPassword recibe una clave en texto plano y devuelve el hash encriptado.
func HashPassword(password string) (string, error) {
    // Generamos el hash con un costo por defecto (10)
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }

    return string(hashedPassword), nil
}

// CheckPasswordHash compara una clave en texto plano con un hash de la base de datos.
func ComparePassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}