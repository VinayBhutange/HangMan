#!/bin/bash

# Workflow validation script
echo "🔍 Validating GitHub Actions workflows..."

# Check if workflows directory exists
if [ ! -d ".github/workflows" ]; then
    echo "❌ .github/workflows directory not found"
    exit 1
fi

echo "✅ Found workflows directory"

# List all workflow files
echo "📋 Workflow files:"
find .github/workflows -name "*.yml" -o -name "*.yaml" | while read -r file; do
    echo "  - $file"
done

# Basic YAML syntax check (if yq is available)
if command -v yq &> /dev/null; then
    echo "🔧 Checking YAML syntax..."
    find .github/workflows -name "*.yml" -o -name "*.yaml" | while read -r file; do
        if yq eval '.' "$file" >/dev/null 2>&1; then
            echo "  ✅ $file - Valid YAML"
        else
            echo "  ❌ $file - Invalid YAML"
        fi
    done
else
    echo "⚠️  yq not found, skipping YAML validation"
fi

echo "🎉 Workflow validation complete!"
