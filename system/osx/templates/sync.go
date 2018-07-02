package synctmpl

// SyncTemplateStr systemd template for background sync service
var SyncTemplateStr = `<!-- Generated by Tokaido. Edits will be overwritten! -->
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>Label</key>
        <string>tokaido.sync.{{.ProjectName}}.plist</string>
    <key>ProgramArguments</key>
    <array>
          <string>/usr/local/bin/unison</string>
      <string>{{.ProjectName}}</string>
    </array>
        <key>WorkingDirectory</key>
        <string>/Users/{{.Username}}/g/Technocrat/repos/{{.ProjectName}}</string>
    <key>StandardErrorPath</key>
    <string>/Users/{{.Username}}/Library/Logs/tokaido.sync.{{.ProjectName}}.err</string>
    <key>StandardOutPath</key>
    <string>/Users/{{.Username}}/Library/Logs/tokaido.sync.{{.ProjectName}}.out</string>
    </dict>
</plist>`
