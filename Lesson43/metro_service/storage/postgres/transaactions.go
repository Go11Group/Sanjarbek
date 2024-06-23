package postgres

import (
	"database/sql"
	"metro_service/models"

	"github.com/google/uuid"
)

type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{Db: db}
}

func (t *TransactionRepo) Create(transaction *models.Transaction) error {
	id := uuid.NewString()
	_, err := t.Db.Exec(`
        INSERT INTO transactions(id, card_id, amount, terminal_id, transaction_type)
        VALUES($1, $2, $3, $4, $5)`,
		id, transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType)
	return err
}

func (t *TransactionRepo) GetById(id string) (*models.Transaction, error) {
	transaction := &models.Transaction{Id: id}
	err := t.Db.QueryRow(`
        SELECT card_id, amount, terminal_id, transaction_type
        FROM transactions
        WHERE id = $1`, id).Scan(&transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *TransactionRepo) GetAll() ([]models.Transaction, error) {
	rows, err := t.Db.Query(`
        SELECT id, card_id, amount, terminal_id, transaction_type
        FROM transactions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		transaction := models.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *TransactionRepo) Update(transaction *models.Transaction) (*models.Transaction, error) {
	id := transaction.Id


	var checkTransactions models.Transaction
	err := t.Db.QueryRow(`
		SELECT id, card_id, amound, terminal_id, transaction_type
		FROM transactions
		WHERE id = $1`, id).Scan(&checkTransactions.Id, &checkTransactions.CardId, &checkTransactions.Amount, checkTransactions.TerminalId, checkTransactions.TransactionType)
	if err != nil {
		return nil, err
	}
	if transaction.Id == "" {
		transaction.Id = checkTransactions.Id
	}
	if transaction.CardId == "" {
		transaction.CardId = checkTransactions.CardId
	}
	if transaction.Amount == 0.0 {
		transaction.Amount = checkTransactions.Amount
	}
	if transaction.TerminalId == "" {
		transaction.TerminalId = checkTransactions.TerminalId
	}
	if transaction.TransactionType == "" {
		transaction.TransactionType = checkTransactions.TransactionType
	}

	_, err = t.Db.Exec(`
        UPDATE transactions
        SET card_id = $2, amount = $3, terminal_id = $4, transaction_type = $5
        WHERE id = $1`,
		transaction.Id, transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *TransactionRepo) Delete(id string) error {
	_, err := t.Db.Exec(`DELETE FROM transactions WHERE id = $1`, id)
	return err
}
