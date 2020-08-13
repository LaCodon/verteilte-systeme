. (".\utils.ps1")

Write-Output "Tests writing and reading a key during network failure"

SetKey -key x -value 1 -target 3
Disconnect -nodeX 1 -nodeY "2 3 4 5"
Start-Sleep 5

SetKey -key x -value 2 -target 3
Connect -nodeX 1 -nodeY "2 3 4 5"
Start-Sleep 5

SetKey -key x -value 3 -target 3
Disconnect -nodeX 2 -nodeY "1 3 4 5"
Start-Sleep 5

SetKey -key x -value 4 -target 3
Start-Sleep 5
Connect -nodeX 2 -nodeY "1 3 4 5"

Start-Sleep 5
SetKey -key x -value 5 -target 3

GetAllValues -target 3