#!/bin/sh
# Generate JSON schema documentation using
# [JSON Schema for Humans](https://github.com/coveooss/json-schema-for-humans).

REFERENCE_DIR="reference"
SCHEMA_DIR="public/standards"

generate-schema-doc --config template_name=md "$SCHEMA_DIR/test.schema.json" "$REFERENCE_DIR/test-schema.md"

# Insert metadata and note for users.
sed -i '1i\
---\
title: Test JSON schema\
description: "Hakutest test JSON schema is available at https:\/\/hakutest.org\/standards\/test.schema.json"\
---\
\
:::info\
\
Hakutest Test JSON schema is available at https:\/\/hakutest.org\/standards\/test.schema.json\
\
:::\
\
---' "$REFERENCE_DIR/test-schema.md"

# Fix anchor links.
sed -i -E 's/(<a name=")([^"]*)("><\/a>)([^{]*)$/\4 {#\2}/g' "$REFERENCE_DIR/test-schema.md"

generate-schema-doc --config template_name=md "$SCHEMA_DIR/result.schema.json" "$REFERENCE_DIR/result-schema.md"

sed -i '1i\
---\
title: Result JSON schema\
description: "Hakutest result JSON schema is available at https:\/\/hakutest.org\/standards\/result.schema.json"\
---\
\
:::info\
\
Hakutest Result JSON schema is available at https:\/\/hakutest.org\/standards\/result.schema.json\
\
:::\
\
---' "$REFERENCE_DIR/result-schema.md"

sed -i -E 's/(<a name=")([^"]*)("><\/a>)([^{]*)$/\4 {#\2}/g' "$REFERENCE_DIR/result-schema.md"
