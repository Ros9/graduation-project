package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type CommentRepository interface {
	CreateComment(comment *model.Comment) (*model.Comment, error)
	FindCommentById(commentId string) (*model.Comment, error)
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	preQueries := []string{
		`create table comments (
			id text,
			challenge_id text,
			user_id text,
			description text,
			parent_comment_id text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &commentRepository{db}
}

func (cr *commentRepository) CreateComment(comment *model.Comment) (*model.Comment, error) {
	row := cr.db.QueryRow("insert into comments (id, challenge_id, user_id, description, parent_comment_id) "+
		"values ($1, $2, $3, $4, $5)", &comment.ID, &comment.ChallengeID, &comment.UserId, &comment.Description, &comment.ParentCommentID)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return comment, nil
}

func (cr *commentRepository) FindCommentById(commentId string) (*model.Comment, error) {
	comment := &model.Comment{}
	err := cr.db.QueryRow("select * from comments where id = $1", &commentId).
		Scan(&comment.ID, &comment.ChallengeID, &comment.UserId, &comment.Description, &comment.ParentCommentID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
