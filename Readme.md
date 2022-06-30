# Pike 

A tool to determine the minimum permissions required to run a tf/iac run

- run ci - limit external outbound connections
- runs on a path or a file
- determines permission drift
- least privilege enabler
- policy creator
- test policy against environment

## Usage

```bash
$./pike -h
NAME:
   pike - Generate IAM policy from your IAC code

USAGE:
   pike [global options] command [command options] [arguments...]

COMMANDS:
   scan, s  scan
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE       Load configuration from FILE
   --directory value, -D value  Directory to scan
   --help, -h                   show help (default: false)
   

```

## Related Tools

<https://github.com/iann0036/iamlive>