. (".\utils.ps1")

Write-Output "Tests writing and reading a key"

SetKey -key x -value 42
SetKey -key y -value 24
SetKey -key x -value 43
GetAllValues