#!/bin/sh

for f in web/css/*; do
    minify -qo "$f" "$f"
done

for f in web/js/*; do
    minify -qo "$f" "$f"
done

for f in internal/pkg/i18n/translations/*; do
    minify -qo "$f" "$f"
done
