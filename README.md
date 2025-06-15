# Purple Team Adversary Simulation Framework

This repository contains a comprehensive purple team adversary simulation framework for testing and validating security controls in a controlled lab environment. The framework includes custom command and control (C2) tooling, detection rules, lab setup, and exercise materials.

## Repository Structure

```
.
├── tools/                  # Adversary tooling
│   ├── c2_server/         # Go-based C2 server
│   └── beacon/            # PowerShell-based beacon
├── detection_rules/       # Detection rules
│   ├── sigma/            # Sigma rules for SIEM detection
│   └── suricata/         # Suricata rules for network detection
├── lab/                   # Lab environment
│   ├── docker/           # Docker Compose configurations
│   └── zeek/             # Zeek (Bro) network monitoring configs
├── exercises/            # Attack scenarios and TTPs
│   ├── attack_plans/     # Detailed attack execution plans
│   └── ttp_library/      # Tactics, Techniques, and Procedures
└── scripts/              # Orchestration and automation
    ├── deploy/           # Lab deployment scripts
    └── attack/           # Attack execution scripts
```

## Prerequisites

- Docker and Docker Compose
- Go 1.20+ (for C2 server)
- PowerShell 7+ (for Windows components)
- Zeek (Bro) Network Security Monitor
- Suricata IDS
- SIEM platform (compatible with Sigma rules)

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/yourusername/purple-team-simulation.git
cd purple-team-simulation
```

2. Deploy the lab environment:
```bash
./scripts/deploy/setup_lab.sh
```

3. Start the C2 infrastructure:
```bash
./scripts/deploy/start_c2.sh
```

4. Run an exercise:
```bash
./scripts/attack/run_exercise.sh --exercise initial-access
```

## Components

### Tools
- Custom C2 server written in Go
- PowerShell-based beacon for Windows persistence
- Additional utilities for lateral movement and privilege escalation

### Detection Rules
- Sigma rules for SIEM detection
- Suricata rules for network-based detection
- Custom Zeek scripts for enhanced visibility

### Lab Environment
- Docker Compose configurations for Windows and Linux targets
- Zeek network monitoring setup
- SIEM integration templates

### Exercises
- Detailed attack scenarios based on MITRE ATT&CK
- Step-by-step execution guides
- Expected detection points and validation steps

### Scripts
- Automated lab deployment
- Attack scenario execution
- Results collection and analysis

## Security Notice

This framework is designed for authorized security testing and training purposes only. All components should be used in controlled, isolated environments. The tools and techniques demonstrated are for educational and defensive purposes.

## Contributing

Contributions are welcome! Please read our contributing guidelines before submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 