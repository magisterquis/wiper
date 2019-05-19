Wiper
=====
Bare-bones, crude wiper for use in demos and testing.  I sincerely hope this is
never used for anything else.

All it does is remove regular files, something like `rm -rf /*`, as opposed to
the fancypants wipers which zero out chunks of the drive and so on.
Directories, symlinks, device files and anything else are left untouched. The
removed files will be printed to stdout.  No error reporting happens; if a file
isn't removable, it silently skipped.

For legal use only.

Usage
-----
Compile, put it on target, and run it.  There's no configuration whatsoever.

Windows
-------
Files in lettered drives (e.g. `C:`) will be removed.  Network shares not
mounted with a drive letter won't be touched.
