#!/usr/bin/env bash
set -e

DOCUMENTATION_URL="https://hakutest.shelepugin.ru"
REPOSITORY_URL="https://github.com/shelepuginivan/hakutest"
DOWNLOAD_URL=""
ICONS_DOWNLOAD_URL=""
MAN_DOWNLOAD_URL=""

# init initializes environment variables.
init() {
    TMP_DOWNLOAD_PATH="/tmp/hakutest.tar.gz"
    TMP_MAN_DOWNLOAD_PATH="/tmp/hakutest-manual.tar.gz"

    if [ -n "$XDG_DATA_HOME" ]; then
        HAKUTEST_INSTALL="$XDG_DATA_HOME/hakutest"
        ICONS="$XDG_DATA_HOME/icons"
        APPLICATIONS="$XDG_DATA_HOME/applications"
        MANPATH="$XDG_DATA_HOME/man"
    else
        HAKUTEST_INSTALL="$HOME/.local/share/hakutest"
        ICONS="$HOME/.local/share/icons"
        APPLICATIONS="$HOME/.local/share/applications"
        MANPATH="$HOME/.local/share/man"
    fi
}

# download_files downloads main Hakutest files.
download_files() {
    echo "Downloading files..."
    curl -fsSL -o "$TMP_DOWNLOAD_PATH" "$DOWNLOAD_URL"
}

# remove_existing_installation removes the existing Hakutest installation if it
# exists.
remove_existing_installation() {
    if [ -d "$HAKUTEST_INSTALL" ]; then
        echo "Removing the existing installation..."
        rm -rf "$HAKUTEST_INSTALL"
    fi
}

# extract_files extracts main Hakutest files to the installation path.
extract_files() {
    mkdir -p "$HAKUTEST_INSTALL"

    echo "Extracting files..."
    tar -C "$HAKUTEST_INSTALL" -xzf "$TMP_DOWNLOAD_PATH" --strip-components=1
}

# create_desktop_entries creates Hakutest .desktop files.
create_desktop_entries() {
    echo "Creating desktop entries..."

    # hakutest-server
    cat << EOF > "$APPLICATIONS/hakutest-server.desktop"
[Desktop Entry]
Version=1.0
Name=Hakutest Server
Comment=Hakutest Server systray applet
Exec=$HAKUTEST_INSTALL/hakutest-server
Icon=hakutest
Terminal=false
Type=Application
Categories=Education;
Keywords=education;fyne;hakutest;
EOF

# hakutest-statistics
    cat << EOF > "$APPLICATIONS/hakutest-statistics.desktop"
[Desktop Entry]
Version=1.0
Name=Hakutest Statistics
Comment=Hakutest Statistics export graphical interface
Exec=$HAKUTEST_INSTALL/hakutest-statistics
Icon=hakutest
Terminal=false
Type=Application
Categories=Education;
Keywords=education;fyne;hakutest;
EOF
}

# download_icons downloads Hakutest icons.
download_icons() {
    echo "Downloading icons..."
    curl -o "$ICONS/hakutest.svg" -fsSL "$ICONS_DOWNLOAD_URL"
}

# download_manpages downloads Hakutest manpages.
download_manpages() {
    echo "Downloading manpages..."
    curl -o "$TMP_MAN_DOWNLOAD_PATH" -fsSL "$MAN_DOWNLOAD_URL"
    tar -C "$MANPATH" -xzkf "$TMP_MAN_DOWNLOAD_PATH" --strip-components=1
}

# finalize reports a successful installation and provides further instructions.
finalize() {
    echo "Hakutest installation complete!"
    echo
    echo "- Documentation: $DOCUMENTATION_URL"
    echo "- Repository:    $REPOSITORY_URL"
    echo
    echo "Add the following line to your shell profile:"
    echo
    echo "----------------------------------------------"
    echo
    echo "export PATH=\"\$PATH:$HAKUTEST_INSTALL\""
    echo
    echo "----------------------------------------------"
}

# cleanup cleans temporary files.
cleanup() {
    rm "$TMP_DOWNLOAD_PATH" "$TMP_MAN_DOWNLOAD_PATH"
}

echo "Installing Hakutest..."
init
download_files
remove_existing_installation
extract_files
create_desktop_entries
download_icons
download_manpages
finalize
cleanup
