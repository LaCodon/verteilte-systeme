. (".\utils.ps1")

Write-Output "Tests writing and reading a key during network failure"

SetKey -key x -value 1 -target 3
Disconnect -nodeX 1 -nodeY "2 3 4 5"

Write-Output "Node 1 disconnected"
pause

SetKey -key x -value 2 -target 3

Write-Output "Wrote x=2"
pause

Connect -nodeX 1 -nodeY "2 3 4 5"

Write-Output "Node 1 reconnected"
pause

SetKey -key x -value 3 -target 3

Write-Output "Wrote x=3"
pause

Disconnect -nodeX 2 -nodeY "1 3 4 5"

Write-Output "Node 2 disconnected"
pause

SetKey -key x -value 4 -target 3

Write-Output "Wrote x=4"
pause

Connect -nodeX 2 -nodeY "1 3 4 5"

Write-Output "Node 2 reconnected"
pause

SetKey -key x -value 5 -target 3

GetAllValues -target 3