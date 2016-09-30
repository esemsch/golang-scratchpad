FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y wget
RUN apt-get install -y vim
RUN apt-get install -y git
RUN wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz 
RUN tar -C /usr/local -xzf go1.7.1.linux-amd64.tar.gz 
RUN mkdir go 

RUN mkdir -p ~/.vim/autoload ~/.vim/bundle;
RUN wget -O ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim 
RUN echo "execute pathogen#infect() \n\
 syntax on \n\
filetype plugin indent on \n\
Helptags" > ~/.vimrc
RUN git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go 
RUN git clone  https://github.com/esemsch/golang-scratchpad go/src/github.com/esemsch/golang-scratchpad 
RUN export PATH=$PATH:/usr/local/go/bin; export GOPATH=/go; vim -c GoInstallBinaries -c q
RUN git config --global user.email "eduard.semsch@seznam.cz"; git config --global user.name "Eduard Semsch"

ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin 
ENV GOBIN=$GOPATH/bin 
ENV PATH=$PATH:$GOBIN 

CMD /bin/bash
