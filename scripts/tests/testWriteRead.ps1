. (".\utils.ps1")

Write-Output "Tests writing and reading a key"

SetKey -key x -value 42 -target 1
SetKey -key y -value 24 -target 1
SetKey -key x -value 43 -target 1
GetAllValues -target 1