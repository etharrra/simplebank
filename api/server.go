package api

import (
	db "github.com/etharra/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves all http requests for banking services
type Server struct {
	store  *db.Store
	router *gin.Engine
}

/**
 * NewServer creates a new server instance with the provided db.Store.
 *
 * @param store The db.Store instance to be used by the server.
 * @return A pointer to the newly created Server instance.
 */
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

/**
 * Start runs the server on the specified address.
 * 
 * @param address The address to run the server on.
 * @return An error if the server fails to run.
 */
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

/**
 * errorResponse generates a Gin H map with an error message.
 * 
 * @param err The error to be included in the response.
 * @return A Gin H map containing the error message.
 */
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
