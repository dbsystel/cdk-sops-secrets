curl -LO https://github.com/timo-reymann/deterministic-zip/releases/download/$(curl -Lso /dev/null -w %{url_effective} https://github.com/timo-reymann/deterministic-zip/releases/latest | grep -o '[^/]*$')/deterministic-zip_linux-amd64 && \
chmod +x deterministic-zip_linux-amd64 && \
mv deterministic-zip_linux-amd64 /bin/dzip