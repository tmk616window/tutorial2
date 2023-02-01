package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"sidecar/graph/generated"
	"sidecar/graph/model"
)

// GetChatUsers is the resolver for the getChatUsers field.
func (r *queryResolver) GetChatUsers(ctx context.Context, input int) ([]*model.ChatUsersResponse, error) {
	rows, err := r.db.Query(`
			select
				u.id,
				u.name,
				latest_messages.message
			from
			friends f
			inner join users u on
				u.id = f.friend_id
			inner join
			(
			select
			*
			from
			direct_comments
			where
			not exists (
				select
				1
				from
				direct_comments sub
				where
				direct_comments.room_id = sub.room_id
				and direct_comments.created_at < sub.created_at
			)
			) latest_messages
				on
				f.room_id = latest_messages.room_id
			and f.user_id = 1;`)
	if err != nil {
		log.Fatalf("getRows db.Query error err:%v", err)
		return nil, err
	}
	defer rows.Close()

	var userChats []*model.ChatUsersResponse

	for rows.Next() {
		u := model.ChatUsersResponse{}
		user := model.User{}
		if err := rows.Scan(&u.User.ID, &user.Name, &u.LatestMessage); err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		u.User = &user
		fmt.Println(u.User.Name)
		userChats = append(userChats, &model.ChatUsersResponse{
			User:          &user,
			LatestMessage: u.LatestMessage,
		})
	}

	return userChats, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
