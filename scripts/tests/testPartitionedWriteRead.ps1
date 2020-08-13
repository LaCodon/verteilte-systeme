. (".\utils.ps1")

Write-Output "Tests writing and reading a key"

SetKey -key x -value 1 -target 2
Disconnect -nodeX 1 -nodeY "2 3 4 5"
Start-Sleep 5

SetKey -key x -value 2 -target 2
Connect -nodeX 1 -nodeY "2 3 4 5"
Start-Sleep 5

SetKey -key x -value 3 -target 2
Disconnect -nodeX 2 -nodeY "1 3 4 5"

pause

SetKey -key x -value 4 -target 2

pause

Connect -nodeX 2 -nodeY "1 3 4 5"

pause

SetKey -key x -value 5 -target 2

GetAllValues