# Enable network traffic between nodeX and nodeY
Function Connect($nodeX, $nodeY)
{
    vagrant ssh smkvs-node-$nodeX -c "for n in $nodeY; do sudo iptables -D INPUT -s 10.0.0.1`$n -j DROP; sudo iptables -D OUTPUT -d 10.0.0.1`$n -j DROP; done; sudo iptables -L INPUT; sudo iptables -L OUTPUT"
    Write-Output "Partition deactivated."
}

# Disable network traffic between nodeX and nodeY
Function Disconnect($nodeX, $nodeY)
{
    vagrant ssh smkvs-node-$nodeX -c "for n in $nodeY; do sudo iptables -I INPUT -s 10.0.0.1`$n -j DROP; sudo iptables -I OUTPUT -d 10.0.0.1`$n -j DROP; done; sudo iptables -L INPUT; sudo iptables -L OUTPUT"
    Write-Output "Partition activated."
}

# Returns a list of all running VMs (all nodes hopefully)
Function Get-Nodes()
{
    VBoxManage list runningvms | ForEach-Object { $a = $_.Split(' '); Write-Output $a[0].Replace('"', '') }
}

# Plugs the network cable into nodeX
Function Plug($nodeX)
{
    $boxes = Get-Nodes
    Write-Output $boxes[$nodeX - 1]
    VBoxManage controlvm $boxes[$nodeX - 1] setlinkstate2 on
}

# Unplugs the network cable from nodeX
Function Unplug($nodeX)
{
    $boxes = Get-Nodes
    Write-Output $boxes[$nodeX - 1]
    VBoxManage controlvm $boxes[$nodeX - 1] setlinkstate2 off
}

# Writes a key to ne cluster
Function SetKey($key, $value, $target)
{
    $port = $target - 1

    # node6 is client node
    vagrant ssh smkvs-node-6 -c "/vagrant/bin/smkvs set -k $key -v $value -i 10.0.0.1${target}:3600$port"
}

# Gets all values from the cluster
Function GetAllValues($target)
{
    $port = $target - 1

    # node6 is client node
    vagrant ssh smkvs-node-6 -c "/vagrant/bin/smkvs get -i 10.0.0.1${target}:3600$port"
}

# Stops nodeX
Function StopNode($nodeX)
{
    vagrant ssh smkvs-node-$nodeX -c "sudo systemctl stop smkvs"
}

# Starts nodeX
Function StartNode($nodeX)
{
    vagrant ssh smkvs-node-$nodeX -c "sudo systemctl start smkvs"
}