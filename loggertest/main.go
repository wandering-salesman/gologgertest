package main

import (
	"context"
	logger "loggerpackage"
	"time"

	"go.uber.org/zap"
)

type UserService struct{}

func main() {

	log := logger.New(zap.DebugLevel)

	// Create context with logger attached
	ctx := log.WithContext(context.Background())

	// Set log levels for specific methods
	log.SetFunctionLevel(ctx, "UserService.Create", zap.ErrorLevel)
	log.SetFunctionLevel(ctx, "UserService.Delete", zap.ErrorLevel)
	log.SetFunctionLevel(ctx, "UserService.ProcessPayment", zap.ErrorLevel)

	// Add some metadata
	log.AddMetadata(ctx, "user_id", 12345)

	userSvc := &UserService{}

	// Simulate some service calls
	go userSvc.Create(ctx, "paagalaadmi@mail.com")
	go userSvc.Delete(ctx, "paagalaurat@mail.com")
	go userSvc.ProcessPayment(ctx, "MH12")

	// Dynamically update the logging level for Create function after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		log.SetFunctionLevel(ctx, "UserService.Create", zap.ErrorLevel)
		userSvc.Create(ctx, "paagalaurat@mail.com")
	}()

	// Allow goroutines to complete
	time.Sleep(4 * time.Second)
}

func (s *UserService) Create(ctx context.Context, email string) {
	// Use WithFunctionContext to get the logger and set the function name
	lgr, _ := logger.WithFunctionContext(ctx, "UserService.Create")

	// Log messages based on set log level
	lgr.Debugf("Starting user creation process for %s", email)
	time.Sleep(500 * time.Millisecond)
	lgr.Infof("User validation completed for %s", email)
	time.Sleep(500 * time.Millisecond)
	lgr.Warnf("Using default user preferences for %s", email)
	lgr.Errorf("Failed to send welcome email to %s", email) // This message will be filtered based on the filter we set
}

func (s *UserService) Delete(ctx context.Context, email string) {
	lgr, _ := logger.WithFunctionContext(ctx, "UserService.Delete")
	lgr.Debugf("Starting user deletion process for %s", email)
	time.Sleep(500 * time.Millisecond)
	lgr.Errorf("Failed to remove user files for %s", email)
}

func (s *UserService) ProcessPayment(ctx context.Context, txID string) {
	lgr, _ := logger.WithFunctionContext(ctx, "UserService.ProcessPayment")
	lgr.Debugf("Processing payment %s", txID)
	lgr.Infof("Payment validation passed for %s", txID)
	lgr.Warnf("Using fallback payment processor for %s", txID)
	lgr.Errorf("Payment processing failed for transaction %s", txID)
}
