package store

const (
	createNewOrderOrReturnExisting = `
		WITH s AS (
			SELECT status_id FROM statuses WHERE name = @status_name
		),
		ins AS (
			INSERT INTO orders (number, status_id, user_id, accrual)
			SELECT @number, s.status_id, @user_id, @accrual FROM s
			ON CONFLICT (number) DO NOTHING
			RETURNING *
		)
		SELECT * FROM ins
		UNION ALL
		SELECT * FROM orders 
		WHERE number = @number 
		AND NOT EXISTS (SELECT 1 FROM ins);
	`
)
