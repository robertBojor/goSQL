package goSQL

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type GoSQL struct {
	DB  *gorm.DB
	Err error
}

func (g *GoSQL) Connect(envPrefix string) {
	// By default we're going to use an empty env prefix
	prefixSeparator := ""
	if envPrefix != "" {
		// Update the env prefix if one is provided
		prefixSeparator = "_"
	}
	// Setup the vars based on the established prefix and env variables
	sqlUser := fmt.Sprintf("%s%sMYSQL_USER", envPrefix, prefixSeparator)
	sqlPass := fmt.Sprintf("%s%sMYSQL_PASS", envPrefix, prefixSeparator)
	sqlHost := fmt.Sprintf("%s%sMYSQL_HOST", envPrefix, prefixSeparator)
	sqlPort := fmt.Sprintf("%s%sMYSQL_PORT", envPrefix, prefixSeparator)
	sqlName := fmt.Sprintf("%s%sMYSQL_NAME", envPrefix, prefixSeparator)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=UTC",
		os.Getenv(sqlUser),
		os.Getenv(sqlPass),
		os.Getenv(sqlHost),
		os.Getenv(sqlPort),
		os.Getenv(sqlName),
		"utf8mb4, utf8")

	// Set the GORM db connection and possible error
	g.DB, g.Err = gorm.Open("mysql", connectionString)
}

