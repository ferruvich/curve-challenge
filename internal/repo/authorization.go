package repo

import (
	context "context"
	"database/sql"

	"github.com/ferruvich/curve-prepaid-card/internal/configuration"
	"github.com/ferruvich/curve-prepaid-card/internal/model"
	"github.com/ferruvich/curve-prepaid-card/internal/psql"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination=authorization_mock.go -package=repo github.com/ferruvich/curve-prepaid-card/internal/repo AuthorizationRequest

// AuthorizationRequest is the interface that contains all DB function for cards
type AuthorizationRequest interface {
	Write(context.Context, *model.AuthorizationRequest) error
}

// AuthorizationRequestRepo handler card operations in DB
type AuthorizationRequestRepo struct {
	dbConnection *sql.DB
}

// NewAuthorizationRequestRepo initialize the db connection and
// returns the initialized structure
func NewAuthorizationRequestRepo(ctx context.Context) (Card, error) {

	cfg, ok := ctx.Value("cfg").(*configuration.Configuration)
	if !ok {
		return nil, errors.Errorf("error loading configuration")
	}

	db, err := sql.Open(cfg.Psql.DriverName, newSessionString(*cfg))
	if err != nil {
		return nil, errors.Wrap(err, "error initializing db connection")
	}

	return &CardRepo{
		dbConnection: db,
	}, nil
}

// Write writes a new card on DB
func (ar *AuthorizationRequestRepo) Write(ctx context.Context, authReq *model.AuthorizationRequest) error {

	statements := []*psql.PipelineStmt{
		psql.NewPipelineStmt(
			"INSERT INTO authorizations(ID,merchant,card,amount,reversed) VALUES ($1,$2,$3,$4,$5)",
			authReq.ID, authReq.Merchant, authReq.Card, authReq.Amount, authReq.Reversed,
		),
	}

	_, err := psql.WithTransaction(ar.dbConnection, func(tx psql.Transaction) (*sql.Rows, error) {
		_, err := psql.RunPipeline(tx, statements[0])
		return nil, err
	})
	if err != nil {
		return errors.Wrap(err, "error writing card")
	}

	return nil
}
