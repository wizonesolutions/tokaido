package synctmpl

// SyncTemplateStr systemd template for background sync service
var SyncTemplateStr = `# Generated by Tokaido. Edits will be overwritten!
[Unit]
Description=File sync service for Tokaido {{.ProjectName}} project

[Service]
Type=simple
WorkingDirectory={{.ProjectPath}}
ExecStart=/usr/bin/unison {{.ProjectName}}
Restart=on-failure`
