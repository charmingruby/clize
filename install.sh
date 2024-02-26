echo "Building the application..."
make build-cli

cp .env /usr/bin/

echo "Moving the binary to /usr/local/bin"
mv clize /usr/bin/clize

echo "Installation completed. You can now run the application with 'clize'"