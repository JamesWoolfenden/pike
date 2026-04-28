resource "azurerm_nginx_configuration" "pike_gen" {
  nginx_deployment_id = "pike"
  root_file           = "/etc/nginx/nginx.conf"

  config_file {
    content = base64encode(<<-EOT
http {
    server {
        listen 80;
        location / {
            default_type text/html;
            return 200 '<!doctype html><html lang="en"><head></head><body>
                <div>this one will be updated</div>
                <div>at 10:38 am</div>
            </body></html>';
        }
        include site/*.conf;
    }
}
EOT
    )
    virtual_path = "/etc/nginx/nginx.conf"
  }

  config_file {
    content = base64encode(<<-EOT
location /bbb {
 default_type text/html;
 return 200 '<!doctype html><html lang="en"><head></head><body>
  <div>this one will be updated</div>
  <div>at 10:38 am</div>
 </body></html>';
}
EOT
    )
    virtual_path = "/etc/nginx/site/b.conf"
  }
}
