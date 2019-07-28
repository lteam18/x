FROM busybox:latest

ADD ./vvx/dist/vvx.linux /bin/vvx

# USER vvx
CMD [ "/bin/vvx" ]
