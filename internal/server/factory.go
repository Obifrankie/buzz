package server

import (
	"github.com/Uncensored-Developer/buzz/internal/users/data"
	"github.com/Uncensored-Developer/buzz/internal/users/features"
	"github.com/Uncensored-Developer/buzz/pkg/authentication"
	"github.com/Uncensored-Developer/buzz/pkg/config"
	"github.com/Uncensored-Developer/buzz/pkg/db"
	"github.com/Uncensored-Developer/buzz/pkg/hash"
	"github.com/Uncensored-Developer/buzz/pkg/logger"
	"github.com/pkg/errors"
)

func InitializeServer() (*Server, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, errors.Wrap(err, "load config failed")
	}
	zapLogger := logger.NewLogger(cfg)

	// Start Initializing Authentication service dependencies
	passwordHasher := hash.NewSHA1Hasher(cfg.PasswordHasherSalt)
	manager, err := authentication.NewManager(cfg.JwtKey)
	if err != nil {
		return nil, errors.Wrap(err, "token manager error")
	}

	bunDB, err := db.Connect(cfg.DatabaseURL)
	userRepo := data.NewUserRepository(bunDB)
	if err != nil {
		return nil, errors.Wrap(err, "database connection failed")
	}

	authService := features.NewAuthenticationService(passwordHasher, manager, userRepo, cfg, zapLogger)

	return NewServer(cfg, zapLogger, authService), err
}
