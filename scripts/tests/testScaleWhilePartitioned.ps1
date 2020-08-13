. (".\utils.ps1")

Write-Output "Tests adding a new node to the cluster while one follower has no connection for a short time"

# first stop all nodes
StopNode -nodeX 1
StopNode -nodeX 2
StopNode -nodeX 3
StopNode -nodeX 4
StopNode -nodeX 5

# start 4 nodes
StartNode -nodeX 1
StartNode -nodeX 2
StartNode -nodeX 3
StartNode -nodeX 4

Start-Sleep 5
Write-Output "Unplugging node 4"

# disconnect node 4
Disconnect -nodeX 4 -nodeY "1 2 3 5"
Start-Sleep 5

Write-Output "Start new node"

# add new node
StartNode -nodeX 5

Start-Sleep 10

Write-Output "Replug node 4"

# reconnect node 4
Connect -nodeX 4 -nodeY "1 2 3 5"