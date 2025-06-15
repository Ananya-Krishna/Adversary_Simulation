# Initial Access and Persistence Attack Plan

## Overview
This exercise simulates an adversary gaining initial access to the target environment and establishing persistence through a custom C2 infrastructure.

## Prerequisites
- Lab environment deployed and running
- C2 server operational
- Windows and Linux targets accessible

## Attack Steps

### 1. Initial Access
1. Deploy the PowerShell beacon to the Windows target
   ```powershell
   # From the Windows target
   .\beacon.ps1
   ```

2. Verify beacon registration with C2 server
   - Check C2 server logs for successful registration
   - Confirm beacon appears in the C2 dashboard

### 2. Command Execution
1. Issue basic reconnaissance commands
   ```json
   {
     "command": "whoami",
     "args": []
   }
   ```

2. Gather system information
   ```json
   {
     "command": "systeminfo",
     "args": []
   }
   ```

### 3. Persistence
1. Create scheduled task for beacon persistence
   ```json
   {
     "command": "schtasks",
     "args": ["/create", "/tn", "WindowsUpdate", "/tr", "powershell.exe -File C:\\scripts\\beacon.ps1", "/sc", "onstart"]
   }
   ```

2. Verify persistence mechanism
   - Restart the Windows target
   - Confirm beacon reconnects automatically

## Detection Points

### Network Detection
- Suricata rules should alert on:
  - Beacon registration traffic
  - Regular check-ins
  - Command execution results

### Host Detection
- Windows Event Logs:
  - Scheduled task creation
  - PowerShell execution
  - Process creation

### Zeek Logs
- HTTP traffic patterns
- DNS queries
- Connection statistics

## Success Criteria
1. Beacon successfully registers with C2 server
2. Commands execute and return results
3. Persistence mechanism survives system restart
4. Detection rules trigger appropriate alerts

## Cleanup
1. Remove scheduled task
   ```json
   {
     "command": "schtasks",
     "args": ["/delete", "/tn", "WindowsUpdate", "/f"]
   }
   ```

2. Stop and remove beacon process
   ```json
   {
     "command": "taskkill",
     "args": ["/F", "/IM", "powershell.exe"]
   }
   ```

## Notes
- Document any detection bypasses or improvements needed
- Record timing of detection alerts
- Note any unexpected behavior or errors 