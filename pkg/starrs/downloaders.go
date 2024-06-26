//nolint:dupl
package starrs

import (
	"errors"
	"fmt"
	"time"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

const DownloadClients = "DownloadClients"

func (s *Starrs) Downloaders(config *AppConfig) (any, error) {
	s.log.Tracef("Call:Downloaders(%s, %s)", config.App, config.Name)

	downloaders, err := s.downloaders(config)
	if err != nil {
		msg := s.log.Translate("Getting download clients: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return downloaders, nil
}

func (s *Starrs) downloaders(config *AppConfig) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Prowlarr:
		return prowlarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Radarr:
		return radarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) DeleteDownloader(config *AppConfig, clientID int64) (any, error) {
	s.log.Tracef("Call:DeleteDownloader(%s, %s, %v)", config.App, config.Name, clientID)

	if err := s.deleteDownloader(config, clientID); err != nil {
		msg := s.log.Translate("Deleting %s download client: %d: %v", config.Name, clientID, err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return s.log.Translate("Deleted %s download client with ID %d.", config.Name, clientID), nil
}

func (s *Starrs) deleteDownloader(config *AppConfig, clientID int64) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Prowlarr:
		return prowlarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Radarr:
		return radarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Readarr:
		return readarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Sonarr:
		return sonarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Whisparr:
		return sonarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) TestLidarrDownloadClient(config *AppConfig, client *lidarr.DownloadClientInput) (string, error) {
	s.log.Tracef("Call:TestLidarrDownloadClient(%s, %s, %d)", config.App, config.Name, client.ID)
	return s.testDownloadClientReply(config.Name, client.Name, client.ID, s.testDownloadClient(config, client))
}

func (s *Starrs) TestProwlarrDownloadClient(config *AppConfig, client *prowlarr.DownloadClientInput) (string, error) {
	s.log.Tracef("Call:TestProwlarrDownloadClient(%s, %s, %d)", config.App, config.Name, client.ID)
	return s.testDownloadClientReply(config.Name, client.Name, client.ID, s.testDownloadClient(config, client))
}

func (s *Starrs) TestRadarrDownloadClient(config *AppConfig, client *radarr.DownloadClientInput) (string, error) {
	s.log.Tracef("Call:TestRadarrDownloadClient(%s, %s, %d)", config.App, config.Name, client.ID)
	return s.testDownloadClientReply(config.Name, client.Name, client.ID, s.testDownloadClient(config, client))
}

func (s *Starrs) TestReadarrDownloadClient(config *AppConfig, client *readarr.DownloadClientInput) (string, error) {
	s.log.Tracef("Call:TestReadarrDownloadClient(%s, %s, %d)", config.App, config.Name, client.ID)
	return s.testDownloadClientReply(config.Name, client.Name, client.ID, s.testDownloadClient(config, client))
}

func (s *Starrs) TestSonarrDownloadClient(config *AppConfig, client *sonarr.DownloadClientInput) (string, error) {
	s.log.Tracef("Call:TestSonarrDownloadClient(%s, %s, %d)", config.App, config.Name, client.ID)
	return s.testDownloadClientReply(config.Name, client.Name, client.ID, s.testDownloadClient(config, client))
}

func (s *Starrs) TestWhisparrDownloadClient(config *AppConfig, client *sonarr.DownloadClientInput) (string, error) {
	s.log.Tracef("Call:TestWhisparrDownloadClient(%s, %s, %d)", config.App, config.Name, client.ID)
	return s.testDownloadClientReply(config.Name, client.Name, client.ID, s.testDownloadClient(config, client))
}

func (s *Starrs) testDownloadClient(config *AppConfig, client any) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

	switch data := client.(type) {
	case *lidarr.DownloadClientInput:
		return lidarr.New(instance.Config).TestDownloadClientContext(s.ctx, data)
	case *prowlarr.DownloadClientInput:
		return prowlarr.New(instance.Config).TestDownloadClientContext(s.ctx, data)
	case *radarr.DownloadClientInput:
		return radarr.New(instance.Config).TestDownloadClientContext(s.ctx, data)
	case *readarr.DownloadClientInput:
		return readarr.New(instance.Config).TestDownloadClientContext(s.ctx, data)
	case *sonarr.DownloadClientInput:
		return sonarr.New(instance.Config).TestDownloadClientContext(s.ctx, data)
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) testDownloadClientReply(
	name, clientName string,
	clientID int64,
	err error,
) (string, error) {
	if err == nil {
		msg := s.log.Translate("Tested %s download client %s (%d).", name, clientName, clientID)
		s.log.Wails.Info(msg)

		return msg, nil
	}

	reqError := &starr.ReqError{}

	if errors.As(err, &reqError) && reqError.Msg != "" {
		err = fmt.Errorf("%s: %s", reqError.Name, reqError.Msg)
	}

	msg := s.log.Translate("Testing %s download client: %s (%d): %s", name, clientName, clientID, err.Error())
	s.log.Wails.Error(msg)

	return "", fmt.Errorf(msg)
}

func (s *Starrs) UpdateLidarrDownloadClient(
	config *AppConfig,
	downloader *lidarr.DownloadClientInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateLidarrDownloadClient(%s, %s, %d)", config.App, config.Name, downloader.ID)
	data, err := s.updateDownloadClient(config, downloader, force)

	return s.updateDownloadClientReply(config.Name, downloader.Name, downloader.ID, data, err)
}

func (s *Starrs) UpdateProwlarrDownloadClient(
	config *AppConfig,
	downloader *prowlarr.DownloadClientInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateProwlarrDownloadClient(%s, %s, %d)", config.App, config.Name, downloader.ID)
	data, err := s.updateDownloadClient(config, downloader, force)

	return s.updateDownloadClientReply(config.Name, downloader.Name, downloader.ID, data, err)
}

func (s *Starrs) UpdateRadarrDownloadClient(
	config *AppConfig,
	downloader *radarr.DownloadClientInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateRadarrDownloadClient(%s, %s, %d)", config.App, config.Name, downloader.ID)
	data, err := s.updateDownloadClient(config, downloader, force)

	return s.updateDownloadClientReply(config.Name, downloader.Name, downloader.ID, data, err)
}

func (s *Starrs) UpdateReadarrDownloadClient(
	config *AppConfig,
	downloader *readarr.DownloadClientInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateReadarrDownloadClient(%s, %s, %d)", config.App, config.Name, downloader.ID)
	data, err := s.updateDownloadClient(config, downloader, force)

	return s.updateDownloadClientReply(config.Name, downloader.Name, downloader.ID, data, err)
}

func (s *Starrs) UpdateSonarrDownloadClient(
	config *AppConfig,
	downloader *sonarr.DownloadClientInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateSonarrDownloadClient(%s, %s, %d)", config.App, config.Name, downloader.ID)
	data, err := s.updateDownloadClient(config, downloader, force)

	return s.updateDownloadClientReply(config.Name, downloader.Name, downloader.ID, data, err)
}

func (s *Starrs) UpdateWhisparrDownloadClient(
	config *AppConfig,
	downloader *sonarr.DownloadClientInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateWhisparrDownloadClient(%s, %s, %d)", config.App, config.Name, downloader.ID)
	data, err := s.updateDownloadClient(config, downloader, force)

	return s.updateDownloadClientReply(config.Name, downloader.Name, downloader.ID, data, err)
}

func (s *Starrs) updateDownloadClient(config *AppConfig, downloader any, force bool) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

	switch data := downloader.(type) {
	case *lidarr.DownloadClientInput:
		return lidarr.New(instance.Config).UpdateDownloadClientContext(s.ctx, data, force)
	case *prowlarr.DownloadClientInput:
		return prowlarr.New(instance.Config).UpdateDownloadClientContext(s.ctx, data, force)
	case *radarr.DownloadClientInput:
		return radarr.New(instance.Config).UpdateDownloadClientContext(s.ctx, data, force)
	case *readarr.DownloadClientInput:
		return readarr.New(instance.Config).UpdateDownloadClientContext(s.ctx, data, force)
	case *sonarr.DownloadClientInput:
		return sonarr.New(instance.Config).UpdateDownloadClientContext(s.ctx, data, force)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) updateDownloadClientReply(
	name, clientName string,
	clientID int64,
	data any,
	err error,
) (*DataReply, error) {
	if err == nil {
		msg := s.log.Translate("Updated %s download client %s (%d).", name, clientName, clientID)
		s.log.Wails.Info(msg)

		return &DataReply{Msg: msg, Data: data}, nil
	}

	reqError := &starr.ReqError{}

	if errors.As(err, &reqError) && reqError.Msg != "" {
		err = fmt.Errorf("%s: %s", reqError.Name, reqError.Msg)
	}

	msg := s.log.Translate("Updating %s download client: %s (%d): %s", name, clientName, clientID, err.Error())
	s.log.Wails.Error(msg)

	return nil, fmt.Errorf(msg)
}

func (s *Starrs) ExportDownloadClients(config *AppConfig, selected Selected) (string, error) {
	instance, err := s.getExportInstance(config, selected, DownloadClients)
	if err != nil {
		return "", err
	}

	switch config.App {
	case starr.Lidarr.String():
		items, err := lidarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
		return s.exportItems(DownloadClients, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Prowlarr.String():
		items, err := prowlarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
		return s.exportItems(DownloadClients, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Radarr.String():
		items, err := radarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
		return s.exportItems(DownloadClients, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Readarr.String():
		items, err := readarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
		return s.exportItems(DownloadClients, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Sonarr.String():
		items, err := sonarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
		return s.exportItems(DownloadClients, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Whisparr.String():
		items, err := sonarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
		return s.exportItems(DownloadClients, config, filterListItemsByID(items, selected), selected.Count(), err)
	}

	return "", ErrInvalidApp
}

func (s *Starrs) ImportDownloadClients(config *AppConfig) (*DataReply, error) {
	switch config.App {
	case starr.Lidarr.String():
		var input []lidarr.DownloadClientOutput
		return importItems(s, DownloadClients, config, input)
	case starr.Prowlarr.String():
		var input []prowlarr.DownloadClientOutput
		return importItems(s, DownloadClients, config, input)
	case starr.Radarr.String():
		var input []radarr.DownloadClientOutput
		return importItems(s, DownloadClients, config, input)
	case starr.Readarr.String():
		var input []readarr.DownloadClientOutput
		return importItems(s, DownloadClients, config, input)
	case starr.Sonarr.String():
		var input []sonarr.DownloadClientOutput
		return importItems(s, DownloadClients, config, input)
	case starr.Whisparr.String():
		var input []sonarr.DownloadClientOutput
		return importItems(s, DownloadClients, config, input)
	}

	return nil, ErrInvalidApp
}

func (s *Starrs) AddLidarrDownloadClient(config *AppConfig, client *lidarr.DownloadClientInput) (*DataReply, error) {
	data, err := s.addDownloadClient(config, client, client.Name)

	return &DataReply{
		Data: data,
		Msg:  fmt.Sprintf("Imported Download Client '%s (%s)' into %s", client.Name, client.Protocol, config.Name),
	}, err
}

func (s *Starrs) AddProwlarrDownloadClient(
	config *AppConfig,
	client *prowlarr.DownloadClientInput,
) (*DataReply, error) {
	data, err := s.addDownloadClient(config, client, client.Name)

	return &DataReply{
		Data: data,
		Msg:  fmt.Sprintf("Imported Download Client '%s (%s)' into %s", client.Name, client.Protocol, config.Name),
	}, err
}

func (s *Starrs) AddRadarrDownloadClient(config *AppConfig, client *radarr.DownloadClientInput) (*DataReply, error) {
	data, err := s.addDownloadClient(config, client, client.Name)

	return &DataReply{
		Data: data,
		Msg:  fmt.Sprintf("Imported Download Client '%s (%s)' into %s", client.Name, client.Protocol, config.Name),
	}, err
}

func (s *Starrs) AddReadarrDownloadClient(config *AppConfig, client *readarr.DownloadClientInput) (*DataReply, error) {
	data, err := s.addDownloadClient(config, client, client.Name)

	return &DataReply{
		Data: data,
		Msg:  fmt.Sprintf("Imported Download Client '%s (%s)' into %s", client.Name, client.Protocol, config.Name),
	}, err
}

func (s *Starrs) AddSonarrDownloadClient(config *AppConfig, client *sonarr.DownloadClientInput) (*DataReply, error) {
	data, err := s.addDownloadClient(config, client, client.Name)

	return &DataReply{
		Data: data,
		Msg:  fmt.Sprintf("Imported Download Client '%s (%s)' into %s", client.Name, client.Protocol, config.Name),
	}, err
}

func (s *Starrs) AddWhisparrDownloadClient(config *AppConfig, client *sonarr.DownloadClientInput) (*DataReply, error) {
	data, err := s.addDownloadClient(config, client, client.Name)

	return &DataReply{
		Data: data,
		Msg:  fmt.Sprintf("Imported Download Client '%s (%s)' into %s", client.Name, client.Protocol, config.Name),
	}, err
}

func (s *Starrs) addDownloadClient(config *AppConfig, downloadClient any, clientName string) (any, error) {
	s.log.Tracef("Call:Add%sDownloadClient(%s, %s)", config.App, config.Name, clientName)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

	switch data := downloadClient.(type) {
	case *lidarr.DownloadClientInput:
		return lidarr.New(instance.Config).AddDownloadClientContext(s.ctx, data)
	case *prowlarr.DownloadClientInput:
		return prowlarr.New(instance.Config).AddDownloadClientContext(s.ctx, data)
	case *radarr.DownloadClientInput:
		return radarr.New(instance.Config).AddDownloadClientContext(s.ctx, data)
	case *readarr.DownloadClientInput:
		return readarr.New(instance.Config).AddDownloadClientContext(s.ctx, data)
	case *sonarr.DownloadClientInput:
		return sonarr.New(instance.Config).AddDownloadClientContext(s.ctx, data)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
