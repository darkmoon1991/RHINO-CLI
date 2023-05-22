# RHINO-CLI

RHINO-CLI is a command-line interface for OpenRHINO.

## Usage

To use RHINO-CLI, run the following command:

```bash
rhino [command]
```

Replace `[command]` with one of the available commands listed below.

## Available Commands

RHINO-CLI provides the following commands:

- `create`: Create a new MPI function/project
- `build`: Build an MPI function/project
- `run`: Submit an MPI function/project and run it as a RHINO job
- `list`: List all RHINO jobs
- `delete`: Delete a RHINO job
- `help`: Help about any command
- `completion`: Generate the autocompletion script for the specified shell
- `version`: Print the version of RhinoClient and kubernetes installed on the local machine, and the version of RhinoServer


To get more information about a specific command, use the following command:

```bash
rhino [command] --help
```
## AutoCompletion

To enable command autocompletion for RHINO-CLI, run the following command for your shell:

- For bash

First, make sure bash-completion is installed. You can install it with `apt-get install bash-completion` or `yum install bash-completion`, etc. Then you now need to ensure that the rhino completion script gets sourced in your shell sessions.
The above commands create /usr/share/bash-completion/bash_completion, which is the main script of bash-completion. Depending on your package manager, you have to manually source this file in your ~/.bashrc file.

To find out, reload your shell and run type _init_completion. If the command succeeds, you're already set, otherwise add the following to your ~/.bashrc file:
```bash
source /usr/share/bash-completion/bash_completion
```
Reload your shell and verify that bash-completion is correctly installed by typing type `_init_completion`.

Then enable kubectl autocompletion

```bash
echo 'source <(rhino completion bash)' >>~/.bashrc
```
- For zsh

Add the following to your ~/.zshrc file:
```bash
echo "compdef _rhino rhino" | cat - <(rhino completion zsh) | source /dev/stdin
```

After reloading your shell, rhino autocompletion should be working.
If you get an error like 2: `command not found: compdef`, then add the following to the beginning of your ~/.zshrc file:

```bash
autoload -Uz compinit && compinit
```

## Demo
[RHINO-CLI demo](https://user-images.githubusercontent.com/20229719/220574704-eb67afd6-ce2c-408d-b708-b660ccfeabc2.mp4)



