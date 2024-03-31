# knockknock


## install go 
$ curl -OL https://go.dev/dl/go1.22.1.linux-amd64.tar.gz

$ sudo tar -C /usr/local -xvf go1.22.1.linux-amd64.tar.gz

$ sudo vi ~/.profile

```
export path=$path:/usr/local/go/bin
```

$ source ~/.profile


$  go version


#$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
#rustup toolchain install stable



## docker
$ sudo apt update


$ sudo apt install -y apt-transport-https ca-certificates curl software-properties-common

$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

$ echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null


$ sudo apt update
$ sudo apt install docker-ce


$ sudo systemctl start docker
$ sudo systemctl enable docker


$ sudo usermod -aG docker $USER

$ docker --version




## iptable

$ sudo iptables -I INPUT 1 -p tcp --dport 80 -j ACCEPT
