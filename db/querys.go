package db

const (
	QUESTION_TABLE = "CREATE TABLE IF NOT EXISTS questions (question TEXT, answer TEXT, status VARCHAR(255), created_at DATETIME DEFAULT CURRENT_TIMESTAMP)"

	// Save question
	QUESTION_SAVE = "INSERT INTO questions (question, status) VALUES (?, ?);"
	// Get question by id (answered)
	QUESTION_BY_ID = "SELECT answer, status FROM questions WHERE status = 'answered' AND rowid = ?;"
	// Get las question unanswered
	QUESTION_LAST_UNANSWERED = "SELECT rowid, question, status FROM questions WHERE status = 'unanswered' ORDER BY rowid DESC;"
	// Update question with responds
	UPDATE_QUESTION = "UPDATE questions SET answer = ?, status = ? WHERE rowid = ?;"
)
