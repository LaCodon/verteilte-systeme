# $Args[0] = nodeId

. (".\utils.ps1")

Write-Output "Tests reconnecting a node to the network (plug off, then plug on)"

$node = $Args[0]

Unplug -nodeX $node
Start-Sleep 5
Plug -nodeX $node