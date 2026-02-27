# snx-cli

A CLI tool to auto-connect to Check Point SNX VPN with TOTP (Time-based One-Time Password) support.

No more manually typing `fixed_password + authenticator_code` every time.

## Install

### Using `go install`

```bash
go install github.com/nhdms/snx-cli@latest
```

Make sure `$GOPATH/bin` (usually `~/go/bin`) is in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Build from source

```bash
git clone https://github.com/nhdms/snx-cli.git
cd snx-cli
go build -o snx-cli .
sudo cp snx-cli /usr/local/bin/
```

## Prerequisites

SNX client must be installed on your system. If not:

```bash
# Download from your VPN portal
wget https://<your-vpn-server>/sslvpn/SNX/INSTALL/snx_install.sh
chmod +x snx_install.sh
sudo ./snx_install.sh
```

On 64-bit Ubuntu/Debian, you may need 32-bit libraries:

```bash
sudo dpkg --add-architecture i386
sudo apt update
sudo apt install libstdc++5:i386 libx11-6:i386 libpam0g:i386
```

## Usage

### 1. Create config

```bash
snx-cli init
```

This creates `~/.snx-cli.yaml` with `0600` permissions.

### 2. Edit config

```yaml
server: vpn.yourcompany.com
username: your-username
fixed_password: your-fixed-password
totp_secret: YOUR_BASE32_TOTP_SECRET
```

### 3. Connect

```bash
snx-cli connect
```

This will:
- Kill any existing SNX process
- Generate a TOTP code from your secret
- Connect using `fixed_password + TOTP code`

### 4. Other commands

```bash
snx-cli status       # Check connection status
snx-cli disconnect   # Disconnect from VPN
```

### Custom config path

```bash
snx-cli connect -c /path/to/config.yaml
```

## How to get your TOTP secret

The `totp_secret` is the Base32 secret key used by your authenticator app. If you no longer have the original setup QR code, you can export it from Google Authenticator:

### Export from Google Authenticator

1. Open [authenticator-export-decoder.juned.site](http://authenticator-export-decoder.juned.site) on a **camera-enabled device** different from the device running Google Authenticator.
2. On Google Authenticator (on your phone), choose **"Transfer accounts"** and then tap **"Continue"** on the "Export Accounts" screen.
3. Click/Tap the **"Scan QR Code"** button on the decoder website to scan the QR code presented by Google Authenticator.
4. The website will decode and display your TOTP secret key — copy it into your config file.

> **Security note:** The decoder runs client-side in your browser. However, always verify third-party tools before trusting them with secrets. After extracting the secret, store it securely and avoid sharing it.
