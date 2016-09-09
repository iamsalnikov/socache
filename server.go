package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iamsalnikov/socache/cleaners"
	"github.com/iamsalnikov/socache/cleaners/facebook"
	"github.com/iamsalnikov/socache/cleaners/vk"
)

// Server struct
type Server struct {
	cleanersMap     map[string]cleaners.CleanerInterface
	defaultCleaners []string
}

// NewServer function return server struct instance
func NewServer() *Server {
	server := new(Server)

	server.cleanersMap = make(map[string]cleaners.CleanerInterface)
	server.cleanersMap["vk"] = vk.New()
	server.cleanersMap["facebook"] = facebook.New()

	server.defaultCleaners = make([]string, 2)
	for key := range server.cleanersMap {
		server.defaultCleaners = append(server.defaultCleaners, key)
	}

	return server
}

// Run function start server
func (s Server) Run(host, port string) {
	router := gin.Default()

	router.GET("/", s.clearCache)
	router.Run(host + ":" + port)
}

func (s Server) clearCache(c *gin.Context) {
	url, _ := c.GetQuery("url")
	net, _ := c.GetQuery("net")
	networks := strings.Split(net, ",")

	if len(networks) == 0 {
		networks = s.defaultCleaners
	}

	answer := make(map[string]bool)

	for i := range networks {
		cleaner, ok := s.cleanersMap[networks[i]]
		if ok {
			answer[networks[i]], _ = cleaner.Clear(url)
		} else {
			answer[networks[i]] = false
		}
	}

	c.JSON(http.StatusOK, answer)
}
