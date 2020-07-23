Set-Location ..

vagrant up

Start-Process powershell { vagrant ssh smkvs-node-1 }
Start-Process powershell { vagrant ssh smkvs-node-2 }
Start-Process powershell { vagrant ssh smkvs-node-3 }
Start-Process powershell { vagrant ssh smkvs-node-4 }
Start-Process powershell { vagrant ssh smkvs-node-5 }

Set-Location scripts
Start-Process powershell