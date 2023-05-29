package db

const (
	QUESTION_TABLE = "CREATE TABLE IF NOT EXISTS questions (id INT(11) AUTO_INCREMENT PRIMARY KEY , question TEXT, status VARCHAR(255), created_at VARCHAR(20))"
)
