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
- `version`:Print the version of RhinoClient and kubernetes installed on the local machine,and the version of RhinoServer


To get more information about a specific command, use the following command:

```bash
rhino [command] --help
```
## AutoCompletion

To enable command autocompletion for RHINO-CLI, run the following command for your shell:

- For bash
```bash
bash completion.bash [install|uninstall] ##generate rhino autocompletion script and install it, or uninstall it
```
- For zsh
```bash
zsh completion.zsh [install|uninstall] ##generate rhino autocompletion script and install it, or uninstall it
```
## Demo
[RHINO-CLI demo](https://user-images.githubusercontent.com/20229719/220574704-eb67afd6-ce2c-408d-b708-b660ccfeabc2.mp4)



