# RHINO-CLI


Serverless模式在互联网领域和数字化应用中已经得到了生产级别的使用，但对于并行计算类应用支持目前仍然不足。大量并行计算类应用仍需要以传统方式进行开发、部署和运行，工作效率和硬件利用效率均有较大提升空间。为解决该问题，我们于2021年提出了RHINO方案（seRverless HIgh performaNce cOmputing），包括：平台、框架和开发工具三个部分，在一些实际应用中得到反馈后，我们在2022年12月将部分代码重构并逐步开源，形成OpenRHINO。

OpenRHINO主要包括两个部分:
1. 开发者工具，目前主要是RHINO-CLI，即本项目，是一个基于Cobra的命令行客户端；
2. 运行在服务端的RHINO-Operator，是一个基于Kubebuilder开发的K8s Operator。
<img width="1222" alt="image" src="https://user-images.githubusercontent.com/20229719/236880254-1461d62a-bd1f-4fd1-8851-41a2811eae40.png">

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

First, make sure bash-completion is installed. You can install it with `apt-get install bash-completion` or `yum install bash-completion`, etc. 
The above commands create `/usr/share/bash-completion/bash_completion`, which is the main script of bash-completion. Depending on your package manager, you have to manually source this file in your `~/.bashrc` file.

To find out, reload your shell and run `type _init_completion`. If the command succeeds, you're already set, otherwise add the following to your `~/.bashrc` file:
```bash
source /usr/share/bash-completion/bash_completion
```
Reload your shell and verify that bash-completion is correctly installed by typing `type _init_completion`.

Then enable rhino autocompletion. You now need to ensure that the rhino completion script gets sourced in your shell sessions.

```bash
echo 'source <(rhino completion bash)' >>~/.bashrc
```
- For zsh

Add the following to your `~/.zshrc` file:
```bash
echo "compdef _rhino rhino" | cat - <(rhino completion zsh) | source /dev/stdin
```

After reloading your shell, rhino autocompletion should be working.
If you get an error like : `command not found: compdef`, then add the following to the beginning of your ~/.zshrc file:

```bash
autoload -Uz compinit && compinit
```

## Demo
[RHINO-CLI demo](https://user-images.githubusercontent.com/20229719/220574704-eb67afd6-ce2c-408d-b708-b660ccfeabc2.mp4)



