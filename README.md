To build the OpenWhisk CLI run the following command from the OpenWhisk CLI home directory:

$ gradle distDocker

WSK CLI is produced in a Docker container during the build process which is copied from the
Docker container to the local file system in the following directory: bin/wsk. This binary will be platform
specific, it will only run on the operating system, and CPU architecture that matches the build machine.

Currently the build process is only supported on Linux, and Mac operating systems running on an x86 CPU architecture.

To get CLI command help, execute the following command:

$ wsk --help

To get CLI command debug information, include the -d, or --debug flag when executing a command.
