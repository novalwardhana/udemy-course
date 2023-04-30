//go:build wireinject

// + wireinject

package simple

import "github.com/google/wire"

var studentAndLecturer = wire.NewSet(NewStudentRepository, NewLecturerRepository)

func InitializedRepository() *Repository {
	wire.Build(studentAndLecturer, NewRepository)
	return nil
}

var databaseMaster = wire.NewSet(NewDatabaseMaster)
var databaseTransaction = wire.NewSet(NewDatabaseTransaction)
var databaseSummary = wire.NewSet(NewDatabaseSummary)
var transaction = wire.NewSet(NewTransactionRepository, NewTransactionService)
var summary = wire.NewSet(NewSummaryRepository, NewSummaryService)

func InitializedService(log bool) *Service {
	wire.Build(databaseMaster, databaseTransaction, databaseSummary, transaction, summary, NewService)
	return nil
}
