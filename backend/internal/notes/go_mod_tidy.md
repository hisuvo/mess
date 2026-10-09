## What go mod tidy does?

It synchronizes your go.mod and go.sum with the packages your project actually imports:

- Adds missing dependencies.
- Removes unused dependencies.
- Updates go.sum checksums.
- Cleans dependency metadata.
- Makes the module dependency state reproducible.
