package commenthandlers

import (
	"real-time-forum/backend/models"
	"real-time-forum/database"
)

func GetCommentHandler(id int) ([]models.Comment, error){
	var comments []models.Comment
	query := `SELECT id, post_id, user_id, content, created_at FROM comments WHERE user_id = ? ORDER BY created_at ASC`
	rows, err := database.DB.Query(query, id)
	if (err != nil){
		return nil, err;
	}
	for rows.Next(){
		var comment models.Comment
		err = rows.Scan(&comment.ID, &comment.Post_id , &comment.Content, &comment.Created_at)
		if err != nil{
			return nil, err;
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
