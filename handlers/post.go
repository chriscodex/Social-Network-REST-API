package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChrisCodeX/REST-API-Go/models"
	"github.com/ChrisCodeX/REST-API-Go/repository"
	"github.com/ChrisCodeX/REST-API-Go/server"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

// Receive from the client (Insert and Update)
type UpsertPostRequest struct {
	PostContent string `json:"post_content"`
}

// Response to the client to Insert a Post
type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

// Response to the client to Update a Post
type PostUpdateResponse struct {
	Message string `json:"message"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// // Get the token from Authorization header
		token, err := GetTokenAuthorizationHeader(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Validation of the user with token
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			// Decode json into into the struct
			var postRequest = UpsertPostRequest{}

			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Generate the id to the post
			id, err := ksuid.NewRandom()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			post := models.Post{
				Id:          id.String(),
				PostContent: postRequest.PostContent,
				UserId:      claims.UserId,
			}

			// Insert the post into the database
			err = repository.InsertPost(r.Context(), &post)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Response to the client
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				Id:          post.Id,
				PostContent: post.PostContent,
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the path parameter {id}
		params := mux.Vars(r)

		post, err := repository.GetPostById(r.Context(), params["id"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Response to the client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// // Get the token from Authorization header
		token, err := GetTokenAuthorizationHeader(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Validation of the user with token
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			// Decode json into into the UpsertPostRequest struct
			var postRequest = UpsertPostRequest{}

			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Get id
			params := mux.Vars(r)

			post := models.Post{
				Id:          params["id"],
				PostContent: postRequest.PostContent,
				UserId:      claims.UserId,
			}

			// Insert the post into the database
			err = repository.UpdatePost(r.Context(), &post)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Response to the client
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostUpdateResponse{
				Message: "Post Updated",
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeletePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from Authorization header
		token, err := GetTokenAuthorizationHeader(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Validation of the user with token
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			params := mux.Vars(r)

			// Delete the post from the database
			err = repository.DeletePost(r.Context(), params["id"], claims.UserId)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Response to the client
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostUpdateResponse{
				Message: "Post deleted",
			})

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ListPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		// Read Query Parameter page
		pageStr := r.URL.Query().Get("page")

		var page = uint64(0)

		if pageStr != "" {
			page, err = strconv.ParseUint(pageStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		// Read Query Parameter size
		sizeStr := r.URL.Query().Get("size")

		var size = uint64(2)

		if sizeStr != "" {
			size, err = strconv.ParseUint(sizeStr, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		// Get the posts from database
		posts, err := repository.ListPost(r.Context(), page, size)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Response to the client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}
