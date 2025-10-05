CREATE PROCEDURE transaction_func(IN from_uid INT,IN to_uid INT,IN amount DECIMAL(10,2))
BEGIN
    DECLARE from_balance DECIMAL(10,2);

    START TRANSACTION;
    UPDATE accounts SET balance = balance - amount WHERE id = from_uid;

    UPDATE accounts SET balance = balance + amount WHERE id = to_uid;

    INSERT INTO transactions(from_account_id, to_account_id,amount) VALUES(from_uid,to_uid,amount);
    
    SELECT balance INTO from_balance FROM accounts WHERE id = from_uid;
    IF from_balance < 0 THEN
        SELECT "余额不足";
        ROLLBACK;
    ELSE
        COMMIT;
    END IF;
END;

CALL transaction_func(1001,1002,10.5);

DROP PROCEDURE IF EXISTS transaction_func;