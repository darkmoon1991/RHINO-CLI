#!/bin/bash
#generate rhino autocompletion script and install it, or uninstall it
option=$1
completion_dir="$HOME/.bash_completion.d"
rhino_completion_file="$completion_dir/rhino.bash"
case $option in
"install")
    # 创建自动补全脚本目录
    mkdir -p "$completion_dir"
    # 生成自动补全脚本
    rhino completion bash > "$rhino_completion_file"
    # 加载自动补全脚本
    echo "source $rhino_completion_file" >> ~/.bashrc
    source ~/.bashrc
    ;;
"uninstall")
    # 删除自动补全脚本目录
    rm -rf "$completion_dir"
    # 删除自动补全脚本目录配置
    sed -i '' -e "/source ~\/.bash_completion.d\/rhino.bash/d" ~/.bashrc
    ;;
*)
    echo "Usage: completion [install|uninstall]"
    ;;
esac

