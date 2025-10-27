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
					 FROM orders o
					 WHERE EXISTS (SELECT 1 FROM updated_balance)
					 RETURNING *
			 )
		SELECT * FROM inserted_withdrawal;
`
)

type OrderWithFlag struct {
	models.Order
	IsNew bool `gorm:"column:is_new"`
}
