- `pwctl -U admin -P admin -u localhost:8080 subcommand`
	- `-U`: username (or `-u`, capital U to be consistent with `psql`)
	- `-P`: password (should we pass it via command line arguments or enter interactively?) (or `-p`?)
	- `-h`: host (optional, default to `http:/localhost:8080`)
	- Should also consider saving the credential to `.config` directory to make the auth status persistent. Then we can make a `auth` command and get rid of this

1. `list`
    - `database`
    - `metric`
    - `preset`
    - `stat`
2. `get`
   - `database`
      - `--id` or `--name` for detail, by default `list` only shows id, name
   - `metric`
      - `--id` for detail, by default `list` only shows id, name
      - limit the output number?
   - `preset`
3. `add`: probably too many parameters? Maybe use yaml file if needed.
   - `database`
   - `metric`
   - `preset`
4. `delete`
   - `database`
      - `--name`
   - `metric`
      - `--id`
      - `--name`
   - `preset`
      - `--name`
5. `update`: probably too many parameters? Maybe use yaml file if needed.
   - `database`
   - `metric`
   - `preset`
6. `status`: check status by webpage content keyword or add a dedicated status API
7. `sink`: need to add additional api for this (if we are going to do it via REST API)
8. `pause`: need to add additional api for this 
