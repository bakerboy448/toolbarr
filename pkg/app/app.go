package app

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/Notifiarr/toolbarr/pkg/ui"
	"github.com/gorilla/schema"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/version"
)

//nolint:gochecknoglobals // this is supposed to be global.
var decoder = schema.NewDecoder()

func (a *App) GetConfig() *config.Config {
	return a.config
}

func (a *App) IsWindows() bool {
	return runtime.GOOS == "windows"
}

// SaveConfigItem saves a single item to the running config and writes the config file.
func (a *App) SaveConfigItem(name string, value string, reload bool) (string, error) {
	config := a.config.Copy()

	err := decoder.Decode(config, map[string][]string{name: {value}})
	if err != nil {
		a.config.Errorf("Writing config: decoding '%s' value '%s' failed: %w", name, value, err)
		return "", fmt.Errorf("decoding '%s' value '%s' failed: %w", name, value, err)
	}

	if err = config.Write(); err != nil {
		a.config.Error("Error writing config: " + err.Error())
		return "", fmt.Errorf("writing config: %w", err)
	}

	a.config.Update(config)

	if reload {
		_ = a.config.Logger.Close()
		a.config.Logger.Setup(a.ctx, config.LogConfig)
	}

	msg := fmt.Sprintf("Config Item '%s' saved! Value: %s", name, value)
	a.config.Print(msg)

	return msg, nil
}

// PickFolder opens the folder selector.
func (a *App) PickFolder(id string) (string, error) {
	dir, err := wailsRuntime.OpenDirectoryDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		DefaultDirectory:     a.config.Path,
		Title:                id,
		CanCreateDirectories: true,
		ShowHiddenFiles:      true,
	})
	if err != nil {
		wailsRuntime.LogError(a.ctx, err.Error())
		return "", fmt.Errorf("opening directory browser: %w", err)
	}

	return dir, nil
}

func (a *App) OpenFolder(path string) string {
	go func() {
		err := ui.OpenFolder(a.ctx, path)
		if err != nil {
			a.config.Errorf("Opening Folder: %s: %w", path, err)
		}
	}()

	return "Opening Path: " + path
}

type Version struct {
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion string
	Started   string
	Running   int
}

func (a *App) Version() Version {
	return Version{
		Version:   version.Version,
		Revision:  version.Revision,
		Branch:    version.Branch,
		BuildUser: version.BuildUser,
		BuildDate: version.BuildDate,
		GoVersion: runtime.Version(),
		Started:   version.Started.Round(time.Second).String(),
		Running:   int(time.Since(version.Started).Seconds()),
	}
}
