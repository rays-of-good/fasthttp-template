NAME = template
HOST = localhost

ENVIRONMENT = CGO_ENABLED="0" GOARCH="amd64" GOOS="linux"

define SERVICE
[Unit]
Description=$(NAME)
After=network.target
[Service]
Type=simple
User=root
WorkingDirectory=/root/$(NAME)-data
ExecStart=/root/$(NAME)-server/main
[Install]
WantedBy=multi-user.target
endef

export SERVICE

define INSTALL
rm -rf /root/$(NAME)-server
rm -rf /root/$(NAME)-data
mkdir /root/$(NAME)-data
mkdir /root/$(NAME)-data/static
mkdir /root/$(NAME)-data/medias
mv /root/$(NAME)-update/$(NAME)-server /root/$(NAME)-server
mv -f /root/$(NAME)-update/$(NAME).service /etc/systemd/system/$(NAME).service
mv -f /root/$(NAME)-update/.env /root/$(NAME)-data/.env
rm -rf /root/$(NAME)-update
systemctl enable $(NAME)
systemctl restart $(NAME)
exit
endef

export INSTALL

define UPDATE
rm -rf /root/$(NAME)-server
mv /root/$(NAME)-update/$(NAME)-server /root/$(NAME)-server
mv -f /root/$(NAME)-update/.env /root/$(NAME)-data/.env
rm -rf /root/$(NAME)-update
systemctl restart $(NAME)
exit
endef

export UPDATE

.PHONY: i u c

.ONESHELL:

template:
	qtc -dir=templates

build:
	$(ENVIRONMENT) go build -ldflags='-s -w' -trimpath -o update/$(NAME)-server/main cmd/$(NAME)-server/main.go

service:
	echo "$$SERVICE" > update/$(NAME).service

env:
	cp .env update/.env

upload:
	scp -pr update root@$(HOST):/root/$(NAME)-update

clean:
	rm -rf update

install:
	ssh root@$(HOST) "$$INSTALL"

update:
	ssh root@$(HOST) "$$UPDATE"

i: template build service env upload clean install

u: template build env upload clean update

c: clean
