# Ku - KubarCloud CLI Tool

`ku` is a command-line interface (CLI) tool designed to help you manage your infrastructure resources on **KubarCloud**. It allows you to define the resources you want (like virtual machines) in a configuration file and then use simple commands to create, view, or delete them.

## Table of Contents

* [Features](#features)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Configuration](#configuration)
* [Usage](#usage)
    * [Defining Resources (`kubar.json`)](#defining-resources-kubarjson)
    * [Applying Changes](#applying-changes)
    * [Checking State](#checking-state)
    * [Destroying Resources](#destroying-resources)
    * [Command Help](#command-help)
    * [Shell Autocompletion](#shell-autocompletion)
* [Example Workflow](#example-workflow)

## Features

  * Manage KubarCloud IaaS (Infrastructure as a Service) resources.
  * Define your desired infrastructure state using a simple JSON file.
  * Apply changes to create or update resources to match your desired state.
  * View the current state of managed resources.
  * Destroy managed resources easily.

## Prerequisites

  * **Make:** You need `make` installed to build the project.
  * **Go Toolchain:** (Assuming it's a Go project based on `make` common usage) Ensure you have a recent version of Go installed.
  * **KubarCloud Account:** You need access credentials for your KubarCloud account.

## Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/0xAFz/ku.git
    cd ku
    ```
2.  **Build the binary:**
    ```bash
    make
    ```
    This will create the `ku` executable file in the current directory. You might want to move this file to a directory in your system's `PATH` (like `/usr/local/bin`) for easier access:
    ```bash
    # Optional: Move ku to your PATH
    sudo mv ku /usr/local/bin/
    ```

## Configuration

Before using `ku`, you need to provide your KubarCloud credentials and potentially other settings.

1.  **Copy the example environment file:**
    ```bash
    cp .env.example .env
    ```
2.  **Edit the `.env` file:**
    Open the `.env` file with your preferred text editor:
    ```bash
    vim .env
    # or nano .env, code .env, etc.
    ```
3.  **Add your details:**
    Fill in the necessary values. The `.env.example` file should indicate what's needed (e.g., API keys). For example:
    ```dotenv
    KUBARCLOUD_APIKEY="api_key xyz"
    ```
    **Important:** Keep your `.env` file secure and **do not** commit it to version control (it should be listed in your `.gitignore` file). `ku` will automatically load these variables when it runs.

## Usage

`ku` works by comparing your desired state (defined in `kubar.json`) with the actual resources in KubarCloud and making changes as needed.

### Defining Resources (`kubar.json`)

1.  **Create the state file:**

    ```bash
    touch kubar.json
    ```

2.  **Define your resources:**
    Open `kubar.json` and define the resources you want `ku` to manage. The file should contain a list of resource objects. Here's an example structure for a virtual machine:

    ```json
    [
        {
            "name": "my-web-server",       # A unique name for your resource
            "image": "Debian-12",          # The OS image to use (e.g., "Debian-12", "Ubuntu-22.04")
            "flavor": "g1-1-1",            # The instance size/type (e.g., "g1-1-1" for 1 vCPU, 1GB RAM)
            "disk_size": 25,               # The size of the primary disk in GB
            "key_name": [],                # Optional: List of pre-uploaded SSH key names in KubarCloud
            "public_key": [                # Optional: List of public SSH keys to add directly
                "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHlvdXJwdWJsaWNrZXlrZXlzYW1wbGU= user@hostname"
            ]
        }
    ]
    ```

      * **`name`**: A unique identifier for this resource within your `kubar.json`.
      * **`image`**: The operating system image name available in KubarCloud.
      * **`flavor`**: The resource size or plan name available in KubarCloud.
      * **`disk_size`**: The root disk size in Gigabytes (GB).
      * **`key_name`**: An array of strings, each being the name of an SSH key already uploaded to your KubarCloud account. Use this if you manage keys via the KubarCloud dashboard/API.
      * **`public_key`**: An array of strings, each being a full public SSH key. `ku` will add these keys to the instance, allowing you to log in.

### Applying Changes

To create or update resources in KubarCloud to match your `kubar.json` file:

```bash
ku iaas apply
```

`ku` will figure out what needs to be created or modified and perform the necessary actions.

### Checking State

To see the resources that `ku` is currently managing (based on its understanding of the state):

```bash
ku state
```

### Destroying Resources

To remove **all** the resources defined in your `kubar.json` file from KubarCloud:

```bash
ku iaas destroy
```

**Warning:** This command is destructive and will permanently delete the resources managed by `ku` in KubarCloud. Be careful when using this command, especially in production environments. Data on these resources will be lost.

### Command Help

You can get help for `ku` itself or any specific command:

  * Show general help and available commands:
    ```bash
    ku --help
    ```
  * Show help for a specific command (e.g., `iaas`):
    ```bash
    ku iaas --help
    ```

### Shell Autocompletion

`ku` can generate autocompletion scripts for various shells (like Bash, Zsh, Fish). This makes typing commands faster and easier.

```bash
# Example for Bash:
ku completion bash --help
ku completion bash > /etc/bash_completion.d/ku

# Example for Zsh:
ku completion zsh --help
ku completion zsh > "${fpath[1]}/_ku"
# You may need to run compinit afterwards:
# autoload -U compinit && compinit
```

Follow the instructions provided by `ku completion [your-shell] --help`.

## Example Workflow

1.  **Install `ku`** (Build from source as shown above).
2.  **Configure `ku`** by creating and editing your `.env` file with KubarCloud credentials.
3.  **Define resources** by creating and editing `kubar.json` (e.g., add a virtual machine definition).
4.  **Create the resources:** `ku iaas apply`
5.  **Check the managed state:** `ku state`
6.  **(Later) Modify resources:** Edit `kubar.json` (e.g., change `disk_size`).
7.  **Apply modifications:** `ku iaas apply`
8.  **(When done) Remove resources:** `ku iaas destroy`
