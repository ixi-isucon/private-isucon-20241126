# promptの設定
if [ "$(id -u)" -eq 0 ]; then
    # Root用のプロンプト（赤で#、秒数付き）
    export PS1="\[\e[0;31m\]\t \u@\h:\w#\[\e[0m\] "
else
    # 通常ユーザ用のプロンプト（緑で$、秒数付き）
    export PS1="\[\e[0;32m\]\t \u@\h:\w\$\[\e[0m\] "
fi
