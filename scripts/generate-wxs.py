#!/usr/bin/env python3

# This script is used to generate Hakutest .wxs file, required to build an .msi
# installer.
from os import DirEntry, scandir
from os.path import basename
from uuid import uuid4

DISC_PROMPT = "Hakutest Installer"
ICON_NAME = "hakutest.ico"
MANUFACTURER = "Ivan Shelepugin"
OUTPUT = "hakutest.wxs"
PRODUCT_NAME = "Hakutest"
PACKAGE_DESC = "Hakutest Installer"
VERSION = "1.0.0"


def file_entry(file: DirEntry, keypath=False) -> str:
    filepath = file.path.replace("./", "")
    keypath_value = "yes" if keypath else "no"

    entry = f'<File Id="{file.name}" Name="{file.name}" DiskId="1" Source="{filepath}" KeyPath="{keypath_value}"'

    if not file.name.endswith(".exe"):
        return entry + "/>"

    shortcut_id = "shortcut" + file.name.replace(".", "").upper()
    shortcut_name = " ".join(file.name.removesuffix(".exe").split("-")).title()

    entry += ">"
    entry += f'<Shortcut Id="{shortcut_id}" Directory="ProgramMenuDir" Name="{shortcut_name}" WorkingDirectory="INSTALLDIR" Icon="{ICON_NAME}" IconIndex="0" Advertise="yes"/>'
    entry += "</File>"

    return entry


def collect_dir(path: str) -> tuple[str, list[str]]:
    dir_name = basename(path)
    dir_id = path.replace("./", "").upper().replace("/", "")
    xml_fragment = f'<Directory Id="{dir_id}" Name="{dir_name}">'
    
    if path == ".":
        xml_fragment = f'<Directory Id="INSTALLDIR" Name="{PRODUCT_NAME}">'

    files: list[DirEntry] = []
    components: list[str] = []

    for entry in scandir(path):
        if entry == ".":
            continue

        if not entry.is_dir():
            files.append(entry)
            continue

        sub_fragment, sub_components = collect_dir(entry.path)
        components.extend(sub_components)
        xml_fragment += sub_fragment

    if len(files) > 0:
        component_name = "component" + dir_id if path != "." else "Executables"
        components.append(component_name)
        xml_fragment += f'<Component Id="{component_name}" Guid="{uuid4()}">'
   
        for index, file in enumerate(files):
            xml_fragment += file_entry(file, keypath=index==0)

        xml_fragment += "</Component>"

    return xml_fragment + "</Directory>", components


def generate_features(components: list[str]) -> str:
    xml_fragment = '<Feature Id="Complete" Level="1">'

    for component in components:
        xml_fragment += f'<ComponentRef Id="{component}"/>'

    return xml_fragment + "</Feature>"

def main():
    structure, components = collect_dir(".")
    features = generate_features(components)

    wxs_xml = f"""<?xml version="1.0" encoding="windows-1252"?>
<Wix xmlns="http://schemas.microsoft.com/wix/2006/wi">
    <Product
        Id="*" 
        UpgradeCode="{uuid4()}"
        Name="{PRODUCT_NAME}"
        Manufacturer="{MANUFACTURER}"
        Language="1033"
        Codepage="1252"
        Version="{VERSION}"
    >
        <Package
            Id="*"
            Keywords="Installer"
            Description="{PACKAGE_DESC}"
            Manufacturer="{MANUFACTURER}"
            InstallerVersion="100"
            Languages="1033"
            Compressed="yes"
            SummaryCodepage="1252"
        />

        <Icon Id="{ICON_NAME}" SourceFile="{ICON_NAME}"/>
        <Media Id="1" Cabinet="{PRODUCT_NAME}.cab" EmbedCab="yes" DiskPrompt="CD-ROM #1"/>
        <Property Id="DiskPrompt" Value="{DISC_PROMPT}"/>

        <Directory Id='TARGETDIR' Name='SourceDir'>
            <Directory Id="ProgramFilesFolder">
                {structure}
            </Directory>
        </Directory>

        {features}
    </Product>
</Wix>
"""
    
    with open(OUTPUT, "w") as out:
        out.write(wxs_xml)


if __name__ == "__main__":
    main()

