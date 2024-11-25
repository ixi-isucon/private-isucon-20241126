# Description: Aliases for the shell

# shortcuts
alias cd-log="cd /var/log"
alias cd-log-nginx="cd /var/log/nginx"
alias cd-log-mysql="cd /var/log/mysql"

alias cd-mysql="cd /etc/mysql"
alias cd-mysql-conf="cd /etc/mysql/conf.d"

alias cd-nginx="cd /etc/nginx"
alias cd-nginx-sites="cd /etc/nginx/sites-available"

alias cd-systemd="cd /etc/systemd/system"

# alias
alias ll="ls -l"
alias la="ls -la"

alias sys-restart="systemctl restart"
alias sys-start="systemctl start"
alias sys-status="systemctl status"
alias sys-stop="systemctl stop"

# promptの設定
if [ "$(id -u)" -eq 0 ]; then
    # Root用のプロンプト（赤で#、秒数付き）
    export PS1="\[\e[0;31m\]\t \u@\h:\w#\[\e[0m\] "
else
    # 通常ユーザ用のプロンプト（緑で$、秒数付き）
    export PS1="\[\e[0;32m\]\t \u@\h:\w\$\[\e[0m\] "
fi
