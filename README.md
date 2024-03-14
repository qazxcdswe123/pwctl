- `pwctl -U admin -P admin -u localhost:8080 subcommand`
	- `-U`: username (or `-u`, capital U to be consistent with `psql`)
	- `-P`: password (should we pass it via command line arguments or enter interactively?) (or `-p`?)
	- `-h`: host (optional, default to `localhost:8080`)

1. `get`
   - `database`
      - `--id` or `--name` for detail, by default it only shows id, name and connection string
   - `metric`
      - `--id` for detail, by default it only shows id, name
      - limit the output number?
   - `preset`
   - `stat`
2. `add`: probably too many parameters? Maybe use yaml file if needed.
   - `database`
   - `metric`
   - `preset`
3. `delete`
   - `database`
      - `--name`
   - `metric`
      - `--id`
      - `--name`
   - `preset`
      - `--name`
4. `update`: probably too many parameters? Maybe use yaml file if needed.
   - `database`
   - `metric`
   - `preset`
5. `status`: check status by webpage content keyword or add a dedicated status API
6. `sink`: need to add additional api for this (if we are going to do it via REST API)
7. `pause`: need to add additional api for this 
