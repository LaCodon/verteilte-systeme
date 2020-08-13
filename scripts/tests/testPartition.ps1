. (".\utils.ps1")

Write-Output "Tests partitioning of nodes 1 and 2 from nodes 3, 4 and 5"

Disconnect -nodeX 1 -nodeY "3 4 5"
Disconnect -nodeX 2 -nodeY "3 4 5"

Start-Sleep 5

Connect -nodeX 1 -nodeY "3 4 5"
Connect -nodeX 2 -nodeY "3 4 5"