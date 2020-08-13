# $Args[0] = nodeId

. (".\utils.ps1")

Write-Output "Tests restarting a node"

$node = $Args[0]

StopNode -nodeX $node
Start-Sleep 5
StartNode -nodeX $node

