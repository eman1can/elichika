# Declaration Style

## File Size and Splitting

Large files are hard to navigate and create unnecessary merge conflicts. To keep files manageable:

- Split large packages into smaller sub-packages. There is almost no cost to doing this — Go's package system makes
  cross-package access straightforward.
- Put each function or type in its own file when it helps readability.

## Bottom-Up Registration

Splitting into many packages can lead to large import lists. In particular, a router package might need to import one
package per endpoint. To avoid this, handlers use bottom-up registration via `init()`:

- Each handler file implements its endpoint handler **and** registers it using `func init()`.
- All handler files are imported into a single master file, which is imported by `main`. This import list can be
  generated from the file system automatically, so nothing gets accidentally omitted.
- A feature can be temporarily disabled by removing its entry from the master import file.

This pattern should be used for all types of handlers, not just endpoint handlers.

## Naming Conventions

- File names use `snake_case`.
- Type and function names use `camelCase` (unexported) or `PascalCase` (exported).
- A file dedicated to a single type or function should be named after that type or function, converted to `snake_case`.
