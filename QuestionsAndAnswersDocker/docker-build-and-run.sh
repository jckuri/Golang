sudo docker build -t qaserver -f Dockerfile .
sudo docker run -d -p 8080:8080 qaserver
