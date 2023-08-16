# Mod vs Pkg vs Src Levels

This repo is made explicitly to test the different scan levels of govulncheck
(Module, Package, and Symbol). The module uses golang.org/text@0.3.2 as the vulnerable module it is importing from, and a local vulndb clone.

There are three vulns:

- One that is called directly
- One that is imported at a package level but not called
- One that is imported at a module level but not called
