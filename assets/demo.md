# WipeOs Demo Script

This file contains the commands used to create the demo GIF for the README.

## Demo Commands

```bash
# Show the welcome banner
wipeOs

# Show help
wipeOs --help

# Create some test files
echo "Sensitive data 1" > secret1.txt
echo "Sensitive data 2" > secret2.txt
echo "Browser cache data" > cache.tmp

# Show dry run first
wipeOs wipe *.txt --dry-run

# Wipe files with confirmation
wipeOs wipe secret1.txt secret2.txt

# Wipe with pattern
wipeOs wipe *.tmp --force

# Show browser data wiping (dry run)
wipeOs wipe --browser-data --dry-run

# Show system temp cleaning (dry run)
wipeOs wipe --system-temp --dry-run
```

## Recording Instructions

To record the demo:

1. Use [asciinema](https://asciinema.org/) or [terminalizer](https://terminalizer.com/)
2. Set terminal to 120x30 for optimal viewing
3. Use a dark theme with good contrast
4. Speak commands slowly and clearly
5. Wait 2-3 seconds between commands for visibility

Example asciinema command:
```bash
asciinema rec demo.cast --overwrite -t "WipeOs Demo"
```

Convert to GIF:
```bash
agg demo.cast demo.gif
``` 