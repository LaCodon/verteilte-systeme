# Args[0] = on / off for partition on or off

if ($Args[0] -eq "on")
{
    # partition node 1 and 2 from 3,4,5
    for ($i = 1; $i -le 2; $i++) {
        vagrant ssh smkvs-node-$i -c "for n in {3..5}; do sudo iptables -I INPUT -s 10.0.0.1`$n -j DROP; sudo iptables -I OUTPUT -s 10.0.0.1`$n -j DROP; done; sudo iptables -L INPUT; sudo iptables -L OUTPUT"
    }

    Write-Output "Partition activated."
}
else
{
    # unpartition node 1 and 2 from 3,4,5
    for ($i = 1; $i -le 2; $i++) {
        vagrant ssh smkvs-node-$i -c "for n in {3..5}; do sudo iptables -D INPUT -s 10.0.0.1`$n -j DROP; sudo iptables -D OUTPUT -s 10.0.0.1`$n -j DROP; done; sudo iptables -L INPUT; sudo iptables -L OUTPUT"
    }

    Write-Output "Partition deactivated."
}
