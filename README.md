- `pwctl -U admin -P admin -u localhost:8080 subcommand`
	- `-u`: username 
	- `-p`: password (optionally we can save the credentials in `.config` or environment variable, to avoid putting password in command line and make it persistent)
	- `-h`: host (optional, default to `localhost:8080`)

1. `list`: list only essential information. For detail, use `get` subcommand.
    - `database`
    - `metric`
    - `preset`
    - `stat` (same as `get`)
2. `get`: get detailed information.
   - `database`
      - `--id` or `--name` for detail, by default `list` only shows id, name
   - `metric`
      - `--id` for detail, by default `list` only shows id, name
      - limit the output number?
   - `preset`
   - `stat`
3. `add`: use yaml file
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
5. `update`: use yaml file
   - `database`
   - `metric`
   - `preset`
6. `status`: check status by webpage content keyword or add a dedicated status API
7. `sink`: need to add additional api for this
8. `pause`: need to add additional api for this 