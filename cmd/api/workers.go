package main

import (
	"context"
	"time"

	"github.com/iangechuki/go-ecommerce/internal/mailer"
)

type EmailPayload struct {
	Template string
	Username string
	Email    string
	Vars     interface{} // template variables
	Debug    bool
	UserID   int64
	Ctx      context.Context
}
type Job struct {
	Type    string
	Payload interface{}
}

func (app *application) startJobWorker() {
	go func() {
		for job := range app.jobQueue {
			switch job.Type {
			case "sendEmail":
				payload, ok := job.Payload.(EmailPayload)
				if !ok {
					app.logger.Errorw("invalid payload type", "type", job.Type)
					continue
				}

				// Attempt sending an email
				status, err := app.mailer.Send(payload.Template, payload.Username, payload.Email, payload.Vars, payload.Debug)
				if err != nil {
					app.logger.Errorw("error sending welcome email", "error", err.Error())
					// rollback user creation if email fails(SAGA pattern)
					if payload.Template == mailer.UserWelcomeTemplate {
						if err := app.store.Users.Delete(payload.Ctx, payload.UserID); err != nil {
							app.logger.Errorw("error deleting user", "error", err.Error())
						}
					}
				} else {
					app.logger.Errorw("email sent", "status", status)
				}
			case "deleteExpiredSessions":
				ctx := context.Background()
				result, err := app.store.Sessions.DeleteExpired(ctx)
				if err != nil {
					app.logger.Errorw("Error deleting expired sessions", "error", err.Error())
				} else {
					rowsAffected, _ := result.RowsAffected()
					if rowsAffected > 0 {
						app.logger.Infow("Expired sessions deleted successfully", "rowsDeleted", rowsAffected)
					} else {
						app.logger.Infow("No expired sessions to delete")
					}
				}
			default:
				app.logger.Errorw("unknown job type", "type", job.Type)
			}
		}

	}()
}
func (app *application) ScheduleExpiredSessionCleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			<-ticker.C
			app.jobQueue <- Job{Type: "deleteExpiredSessions", Payload: nil}
		}
	}()
}
