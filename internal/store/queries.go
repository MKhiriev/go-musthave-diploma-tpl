package store

import "go-musthave-diploma-tpl/models"

const (
	createNewOrderOrReturnExisting = `
		WITH s AS (
			SELECT status_id FROM statuses WHERE name = @status_name
		),
		ins AS (
			INSERT INTO orders (number, status_id, user_id, accrual)
			SELECT @number, s.status_id, @user_id, @accrual FROM s
			ON CONFLICT (number) DO NOTHING
			RETURNING *, TRUE AS is_new
		)
		SELECT * FROM ins
		UNION ALL
		SELECT *, FALSE AS is_new
		FROM orders 
		WHERE number = @number 
		AND NOT EXISTS (SELECT 1 FROM ins);
	`
	withdrawSumWithBalanceCheck = `
		WITH updated_balance AS (
			UPDATE balance b
				SET
					current = b.current - @sum,
					withdrawn = b.withdrawn + @sum
				WHERE b.user_id = @user_id
					AND
					  b.current >= @sum
						  AND
					 NOT EXISTS (SELECT 1 FROM orders WHERE number = @order)
				RETURNING *
		),
			 inserted_withdrawal AS (
				 INSERT INTO withdrawals (user_id, order_num, sum)
					 SELECT
						 @user_id,
						 @order,
						 @sum
					 WHERE EXISTS (SELECT 1 FROM updated_balance)
					 RETURNING *
			 )
		SELECT * FROM inserted_withdrawal;
	`
	updateOrderAccrualAndBalance = `
		WITH status_data AS (
			SELECT status_id FROM statuses WHERE name = @status_name
		),
		updated_order AS (
			UPDATE orders o
			SET 
				status_id = (SELECT status_id FROM status_data),
				accrual = @accrual
			WHERE o.number = @order_number
			RETURNING o.user_id
		),
		updated_balance AS (
			UPDATE balance b
			SET current = current + @accrual
			WHERE b.user_id = (SELECT user_id FROM updated_order)
			RETURNING b.user_id
		)
		SELECT 
			(SELECT COUNT(*) FROM updated_order) AS orders_updated,
			(SELECT COUNT(*) FROM updated_balance) AS balance_updated;
	`
	createTablesIfNotExist = `
		-- User table
		CREATE TABLE IF NOT EXISTS users (
			user_id BIGSERIAL PRIMARY KEY,
			login TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL
		);
		
		-- status table
		CREATE TABLE IF NOT EXISTS statuses (
			status_id SERIAL PRIMARY KEY,
			name TEXT UNIQUE NOT NULL
		);
		
		-- Order table
		CREATE TABLE IF NOT EXISTS orders (
			order_id BIGSERIAL PRIMARY KEY,
			number TEXT NOT NULL UNIQUE,
			status_id INT references statuses,
			user_id BIGINT references users NOT NULL,
			accrual NUMERIC(12, 5),
			uploaded_at TIMESTAMPTZ DEFAULT current_timestamp
		);
		
		-- Withdrawal table
		CREATE TABLE IF NOT EXISTS withdrawals (
			withdrawal_id BIGSERIAL PRIMARY KEY,
			user_id INT references users NOT NULL,
			order_num TEXT NOT NULL UNIQUE,
			sum NUMERIC(12, 5),
			processed_at TIMESTAMPTZ DEFAULT current_timestamp
		);
		
		-- Balance table
		CREATE TABLE IF NOT EXISTS balance (
			balance_id BIGSERIAL PRIMARY KEY,
			user_id BIGINT references users UNIQUE NOT NULL,
			current NUMERIC(12, 5) NOT NULL DEFAULT 0,
			withdrawn NUMERIC(12, 5) DEFAULT 0
			CONSTRAINT not_negative_balance CHECK (balance.current >= 0 AND withdrawn >= 0)
		);
		
		-- ############################ DATA INSERTION ###############################
		INSERT INTO statuses (name) VALUES
		('NEW'), ('PROCESSING'), ('INVALID'), ('PROCESSED')
		ON CONFLICT DO NOTHING;
	`
)

type OrderWithFlag struct {
	models.Order
	IsNew bool `gorm:"column:is_new"`
}
