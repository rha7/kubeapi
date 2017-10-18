package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rha7/kubeapi/appinfo"
)

// Home Controller //
type Home struct{}

// NewHome HomeController builder //
func NewHome() *Home {
	return &Home{}
}

// Home is the main handler for HomeController //
func (c Home) Home(w http.ResponseWriter, r *http.Request) {
	executableFile, _ := os.Executable()
	groups, _ := os.Getgroups()
	workDir, _ := os.Getwd()
	envVars := parseEnvArray(os.Environ())
	respData := map[string]interface{}{
		"_generated_at":         time.Now(),
		"_app_version":          appinfo.Version,
		"_app_build_time":       appinfo.BuildTime,
		"user_id":               os.Getuid(),
		"effective_user_id":     os.Geteuid(),
		"group_id":              os.Getgid(),
		"effective_group_id":    os.Getegid(),
		"groups":                groups,
		"process_id":            os.Getpid(),
		"work_dir":              workDir,
		"environment_variables": envVars,
		"executable_file":       executableFile,
		"application_arguments": os.Args[1:],
		"os2": os.Getuid(),
		"os3": os.Getuid(),
		"os4": os.Getuid(),
		"os5": os.Getuid(),
		"os6": os.Getuid(),
		"os7": os.Getuid(),
		"os8": os.Getuid(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	je := json.NewEncoder(w)
	_ = je.Encode(respData)
}

func parseEnvArray(env []string) map[string]string {
	out := map[string]string{}
	for _, envVar := range env {
		parsed := strings.SplitN(envVar, "=", 2)
		out[parsed[0]] = parsed[1]
	}
	return out
}
