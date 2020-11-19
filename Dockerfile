FROM golang

LABEL maintainer="alpaca5541@foxmail.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /var/www/main

# 添加应用可执行文件，并设置执行权限
ADD ./gf-music   $WORKDIR/main
#暂时添加log文件
ADD log $WORKDIR/log
RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD i18n     $WORKDIR/i18n
# 添加静态文件、配置文件、模板文件
ADD public   $WORKDIR/public
ADD config   $WORKDIR/config
ADD template $WORKDIR/template

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
#ENTRYPOINT mkdir $WORKDIR/log
CMD $WORKDIR/main