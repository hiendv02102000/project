package query

import (
	"errors"
	"music-api/graph/input"
	"music-api/graph/output"
	"music-api/internal/pkg/domain/domain_model/dto"
	"music-api/internal/pkg/domain/domain_model/entity"
	"music-api/internal/pkg/domain/service"
	"music-api/pkg/share/utils"

	"github.com/graphql-go/graphql"
)

func GetSongDetailQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetSongDetailOutput(),
		Description: "get a song",
		Args: graphql.FieldConfigArgument{
			"song": &graphql.ArgumentConfig{
				Type: input.SongInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["song"] == nil {
				err = errors.New("song is required")
				return
			}
			reqSong := p.Args["song"].(map[string]interface{})
			song := entity.Song{
				ID: reqSong["id"].(int),
			}

			songRepo := containerRepo["song_repository"].(service.SongRepository)
			song, err = songRepo.FirstSong(song)
			if err != nil {
				return
			}
			userRepo := containerRepo["user_repository"].(service.UserRepository)
			user, err := userRepo.FirstUser(entity.Users{ID: song.UserID})
			result = dto.GetSongDetailResponse{
				UserSong: dto.UserResponse{
					Username:  user.Username,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					AvatarURL: user.AvatarUrl,
				},
				ID:         song.ID,
				Title:      song.Title,
				ContentURL: song.ContentURL,
				View:       song.View,
				ImageURL:   song.ImageURL,
				Decription: song.Decription,
				CreatedAt:  utils.FormatTime(song.CreatedAt),
				Singer:     song.Singer,
			}
			return
		},
	}
}
