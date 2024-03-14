- `pwctl -U admin -P admin -h localhost:8080 subcommand`
	- `-U`: username (or `-u`, capital U to be consistent with `psql`)
	- `-P`: password (should we pass it via command line arguments or enter interactively?) (or `-p`?)
	- `-h`: host (optional, default to `localhost:8080`)

1. `get` (by id)
   - `databases`
   - `metrics`
   - `presets`
   - `stats`
2. `add`: probably too many parameters? Maybe use yaml file if needed.
   - `database`
   - `metric`
   - `preset`
3. `delete` (by id)
   - `database`
   - `metric`
   - `preset`
4. `update`: probably too many parameters? Maybe use yaml file if needed.
   - `database`
   - `metric`
   - `preset`
5. `status`: check status by ???
6. `sink`: need to add additional api for this (if we are going to do it via REST API)
7. `pause`: need to add additional api for this 
