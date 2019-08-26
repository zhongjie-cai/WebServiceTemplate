#!/bin/sh

# -e  causes Exit immediately when a command exits with a non-zero status.
# When a test fails in Go, it returns a non-zero status and causes the 
# build pipeline to fail and exit when unit tests fails.
set -e

# Install dependency libraries
./init/dependency.sh

# These values are coming from script parameters
AppVersion=${1:-0.0.0}
AppPort=${2:-18605}
AppName=${3:-MyApp}
AppPath=${4:-.}

# Read version values from input parameter
IFS=. read AppMajorVer AppMinorVer AppPatchVer AppBuildVer <<EOF
${AppVersion}
EOF

# Name of the binary
BINARY=./../bin/${AppName}

# Builds the project
# -a: forces rebuild
# -o: defines the name of the output binary file
AppVersionFlag="-X main.appVersion=${AppVersion}"
AppPortFlag="-X main.appPort=${AppPort}"
AppNameFlag="-X main.appName=${AppName}"
AppPathFlag="-X main.appPath=${AppPath}"
LDFlags="${AppVersionFlag} ${AppPortFlag} ${AppNameFlag} ${AppPathFlag}"
go build -ldflags "${LDFlags}" -a -o ${BINARY} ./..

# Copy docs so that they can be found by the WebServiceTemplate binaries.
cp -R ./../docs/ ./../bin/

# Copy favicon for the application
cp ./../favicon.ico ./../bin/

# Replace dynamic variables for Swagger UI
sed -i "s/\${APP_NAME}/${AppName}/g" ./../bin/docs/openapi.json
sed -i "s/\${APP_VERSION}/${AppVersion}/g" ./../bin/docs/openapi.json
sed -i "s/\${APP_PORT}/${AppPort}/g" ./../bin/docs/openapi.json
sed -i "s/\${APP_MAJOR_VER}/${AppMajorVer}/g" ./../bin/docs/openapi.json
sed -i "s/\${APP_MINOR_VER}/${AppMinorVer}/g" ./../bin/docs/openapi.json
sed -i "s/\${APP_PATCH_VER}/${AppPatchVer}/g" ./../bin/docs/openapi.json
sed -i "s/\${APP_BUILD_VER}/${AppBuildVer}/g" ./../bin/docs/openapi.json
