
# https://www.vagrantup.com/docs/multi-machine

Vagrant.configure("2") do |config|

    (1..6).each do |i|
        config.vm.define "smkvs-node-#{i}" do |node|
            node.vm.box = "hashicorp/bionic64"
            node.vm.network "private_network", ip: "10.0.0.1#{i}"
            node.vm.provision :shell, path: "vagrant.sh", args: "#{i}"
        end
    end

end
