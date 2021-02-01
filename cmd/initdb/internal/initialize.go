package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OchiengEd/afya/pkg/platform"
)

// InitDB function initializes
// database schema and roles
func InitDB() {
	fmt.Println("Initialize database schema")
	fmt.Println(platform.PasswordGen(10))

}

func initScriptsExist() bool {
	scriptsDir := "/scripts"
	scriptName := "init.sql"
	_, err := os.Stat(scriptsDir)
	if err != nil {
		return false
	}

	_, err = os.Stat(filepath.Join(scriptsDir, scriptName))
	if err != nil {
		// generate db init script
	}

	return true
}
