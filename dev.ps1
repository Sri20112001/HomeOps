# Run from repo root: .\dev.ps1
# Starts server + agent with air live-reload, each in its own window.
# Stop: close the two windows (Ctrl+C in each).

$root = Split-Path -Parent $MyInvocation.MyCommand.Path

Start-Process powershell -ArgumentList '-NoExit','-Command',"cd '$root\server'; air"
Start-Process powershell -ArgumentList '-NoExit','-Command',"cd '$root\agent'; air"

Write-Output 'Started server and agent (air) in separate windows.'
