map $http_origin $allowed_origin {
  default "";
  http://localhost:4173 "$http_origin";
}

# define connection for SSE
map $http_upgrade $connection_upgrade {
  default upgrade;
  ""      close;
}

upstream frontend {
  server frontend:10.0.100.5;
}

upstream federation_orchestrator {
  server federation_orchestrator:10.0.100.20;
}

upstream feed_collector {
  server feed_collector:10.0.100.30;
}

upstream insights_api {
  server insights_api:10.0.100.40;
}


server {
  listen 80;
  server_name localhost;

  add_header X-Content-Type-Options "nosniff" always;
  add_header X-Frame-Options "SAMEORIGIN" always;
  add_header X-XSS-Protection "1; mode=block" always;
  add_header Referrer-Policy "strict-origin-when-cross-origin" always;
  add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
  add_header Content-Security-Policy "default-src 'self';
  script-src 'self' 'unsafe-inline' 'unsafe-eval';
  style-src 'self' 'unsafe-inline';
  img-src 'self' data: https:;
  font-src 'self' data:;
  connect-src 'self' http://localhost:8080;"
  always;

  server_tokens off;
  client_body_buffer_size 10K;
  client_header_buffer_size 1k;
  client_max_body_size 8m;
  large_client_header_buffers 2 1k;

  client_body_timeout 10;
  client_header_timeout 10;
  keepalive_timeout 15;
  send_timeout 10;

  location / {
    proxy_pass http://frontend;
    proxy_http_version 1.1;
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
  }


}
