# Args[0] = Node Id
# Args[1] = on / off for network cable

Write-Output $Args

$boxes = VBoxManage list runningvms | % { $a = $_.Split(' '); Write-Output $a[0].Replace('"', '') }
VBoxManage controlvm $boxes[$Args[0] - 1] setlinkstate2 $Args[1]