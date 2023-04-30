package simple

/* Example 1 */

type StudentRepository struct {
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

type LecturerRepository struct {
}

func NewLecturerRepository() *LecturerRepository {
	return &LecturerRepository{}
}

type Repository struct {
	StudentRepository  *StudentRepository
	LecturerRepository *LecturerRepository
}

func NewRepository(studentRepository *StudentRepository, lecturerRepository *LecturerRepository) *Repository {
	return &Repository{
		StudentRepository:  studentRepository,
		LecturerRepository: lecturerRepository,
	}
}

/* Example 2 */

type Database struct {
	Driver string
	Uri    string
}

type DatabaseMaster Database

func NewDatabaseMaster() *DatabaseMaster {
	return &DatabaseMaster{
		Driver: "PostgreSQL",
		Uri:    "Master",
	}
}

type DatabaseTransaction Database

func NewDatabaseTransaction() *DatabaseTransaction {
	return &DatabaseTransaction{
		Driver: "PostgreSQL",
		Uri:    "Trx",
	}
}

type DatabaseSummary Database

func NewDatabaseSummary() *DatabaseSummary {
	return &DatabaseSummary{
		Driver: "PostgreSQL",
		Uri:    "Summary",
	}
}

type TransactionRepository struct {
	Log bool
}

func NewTransactionRepository(log bool) *TransactionRepository {
	return &TransactionRepository{}
}

type TransactionService struct {
	DatabaseMaster        *DatabaseMaster
	DatabaseTransaction   *DatabaseTransaction
	TransactionRepository *TransactionRepository
}

func NewTransactionService(databaseMaster *DatabaseMaster, databaseTransaction *DatabaseTransaction, transactionRepository *TransactionRepository) *TransactionService {
	return &TransactionService{
		DatabaseMaster:        databaseMaster,
		DatabaseTransaction:   databaseTransaction,
		TransactionRepository: transactionRepository,
	}
}

type SummaryRepository struct {
	Log bool
}

func NewSummaryRepository(log bool) *SummaryRepository {
	return &SummaryRepository{
		Log: log,
	}
}

type SummaryService struct {
	DatabaseMaster    *DatabaseMaster
	DatabaseSummary   *DatabaseSummary
	SummaryRepository *SummaryRepository
}

func NewSummaryService(databaseMaster *DatabaseMaster, databaseSummary *DatabaseSummary, summaryRepository *SummaryRepository) *SummaryService {
	return &SummaryService{
		DatabaseMaster:    databaseMaster,
		DatabaseSummary:   databaseSummary,
		SummaryRepository: summaryRepository,
	}
}

type Service struct {
	TransactionService *TransactionService
	SummaryService     *SummaryService
}

func NewService(transactionService *TransactionService, summaryService *SummaryService) *Service {
	return &Service{
		TransactionService: transactionService,
		SummaryService:     summaryService,
	}
}
