#/bin/zsh
#generate rhino autocompletion script and install it, or uninstall it
option=$1
completion_dir="$HOME/.zsh_completion.d"
rhino_completion_file="$completion_dir/_rhino.zsh"
case $option in
"install")
    # 创建自动补全脚本目录
    mkdir -p "$completion_dir"
    # 生成自动补全脚本
    rhino completion zsh > "$rhino_completion_file"
    # 添加自动补全脚本目录到 $fpath
    echo "fpath=($completion_dir \$fpath)" >> ~/.zshrc
    # 初始化自动补全
    echo "autoload -U compinit && compinit" >> ~/.zshrc
    # 加载自动补全脚本
    source ~/.zshrc
    source "$rhino_completion_file"
    ;;
"uninstall")
    # 删除自动补全脚本目录
    rm -rf "$completion_dir"
    # 删除自动补全脚本目录配置
    sed -i '' -e '/\.zsh_completion\.d/d' ~/.zshrc
    sed -i '' -e "/autoload -U compinit/d" ~/.zshrc

    ;;
*)
    echo "Usage: rhino_completion [install|uninstall]"
    ;;
esac
